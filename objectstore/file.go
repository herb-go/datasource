package objectstore

import (
	"io"
	"os"
	"path/filepath"
	"strings"
)

type FileObjectStore struct {
	NopObjectStore
	Path string
	Mode os.FileMode
}

func (s *FileObjectStore) abs(path string) (string, error) {
	filepath := filepath.Join(s.Path, path)
	if strings.HasPrefix(filepath, s.Path) {
		return filepath, nil
	}
	return "", NewErrPathInvalid(path)
}
func (s *FileObjectStore) List(path string, iter string, limit int64) ([]*Stat, string, error) {
	return nil, "", ErrFeatureNotSupported
}
func (s *FileObjectStore) Stat(path string) (stat *Stat, err error) {
	defer func() { err = convertFileObjectStoreError(path, err) }()
	var abs string
	abs, err = s.abs(path)
	if err != nil {
		return
	}
	i, err := os.Stat(abs)
	if err != nil {
		return
	}
	return convertStat(i), nil
}

func (s *FileObjectStore) Remove(path string) (err error) {
	defer func() { err = convertFileObjectStoreError(path, err) }()
	var abs string
	abs, err = s.abs(path)
	if err != nil {
		return
	}
	err = os.Remove(abs)
	return
}
func (s *FileObjectStore) Rename(from string, to string) (err error) {
	defer func() { err = convertFileObjectStoreError(from, err) }()
	var absfrom, absto string
	absfrom, err = s.abs(from)
	if err != nil {
		return
	}
	absto, err = s.abs(from)
	if err != nil {
		return
	}
	err = os.Rename(absfrom, absto)
	return
}
func (s *FileObjectStore) Copy(from string, to string) (err error) {
	defer func() { err = convertFileObjectStoreError(from, err) }()
	var absfrom, absto string
	var filefrom, fileto *os.File
	absfrom, err = s.abs(from)
	if err != nil {
		return
	}
	absto, err = s.abs(from)
	if err != nil {
		return
	}
	filefrom, err = os.Open(absfrom)
	if err != nil {
		return
	}
	defer filefrom.Close()
	fileto, err = os.Open(absto)
	if err != nil {
		return
	}
	defer fileto.Close()
	_, err = io.Copy(fileto, filefrom)
	return
}
func (s *FileObjectStore) LoadObject(path string, w io.Writer) (writen int64, err error) {
	defer func() { err = convertFileObjectStoreError(path, err) }()
	var abs string
	var f *os.File
	abs, err = s.abs(path)
	if err != nil {
		return
	}
	f, err = os.Open(abs)
	if err != nil {
		return
	}
	defer f.Close()
	writen, err = io.Copy(w, f)
	return
}
func (s *FileObjectStore) LoadObjectPart(path string, from int64, to int64, w io.Writer) (writen int64, err error) {
	defer func() { err = convertFileObjectStoreError(path, err) }()
	var abs string
	var f *os.File
	abs, err = s.abs(path)
	if err != nil {
		return
	}
	_, err = f.Seek(from, 0)
	if err != nil {
		return
	}
	f, err = os.Open(abs)
	if err != nil {
		return
	}
	defer f.Close()
	writen, err = io.CopyN(w, f, to-from)
	return
}
func (s *FileObjectStore) SaveObject(path string, r io.Reader) (writen int64, err error) {
	defer func() { err = convertFileObjectStoreError(path, err) }()
	var abs string
	var f *os.File
	abs, err = s.abs(path)
	if err != nil {
		return
	}
	f, err = os.Open(abs)
	if err != nil {
		return
	}
	defer f.Close()
	writen, err = io.Copy(f, r)
	return
}

func convertFileObjectStoreError(path string, err error) error {
	if err == nil {
		return nil
	} else if os.IsNotExist(err) {
		return NewErrObjectNotExist(path)
	} else if os.IsExist(err) {
		return NewErrObjectExist(path)
	} else if os.IsPermission(err) {
		return NewErrPermissionDenied(path)
	}
	return err
}

func convertStat(i os.FileInfo) *Stat {
	return &Stat{
		Name:         i.Name(),
		IsFolder:     i.IsDir(),
		Size:         i.Size(),
		ModifiedTime: i.ModTime(),
	}
}
