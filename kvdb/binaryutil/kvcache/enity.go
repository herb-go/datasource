package kvcache

import (
	"bytes"
	"encoding/binary"
	"io"
)

var DataOrder = binary.BigEndian

type enity struct {
	typecode byte
	version  []byte
	data     []byte
}

func (e *enity) WriteTo(w io.Writer) error {
	var err error
	_, err = w.Write([]byte{e.typecode})
	if err != nil {
		return err
	}
	if e.typecode == enityTypecodeRevocable {
		vbytes := make([]byte, 8)
		DataOrder.PutUint64(vbytes, uint64(len(e.version)))
		_, err = w.Write(vbytes)
		if err != nil {
			return err
		}
		_, err = w.Write(e.version)
		if err != nil {
			return err
		}
	}
	_, err = w.Write(e.data)
	return err
}

const enityTypecodeRevocable = byte(0)
const enityTypecodeIrrevocable = byte(1)

func newEnity() *enity {
	return &enity{}
}
func createEnity(revocable bool, version []byte, data []byte) *enity {
	e := newEnity()
	if revocable {
		e.typecode = enityTypecodeRevocable
		e.version = version
	} else {
		e.typecode = enityTypecodeIrrevocable
	}
	e.data = data
	return e
}

func loadEnity(data []byte, revocable bool, version []byte) (*enity, error) {
	datalength := len(data)
	if datalength == 0 {
		return nil, ErrUnresolvedCacheEnity
	}
	switch data[0] {
	case enityTypecodeIrrevocable:
		if revocable {
			return nil, ErrEnityTypecodeNotMatch
		}
		return createEnity(false, nil, data[1:]), nil
	case enityTypecodeRevocable:
		if datalength < 5 {
			return nil, ErrUnresolvedCacheEnity
		}
		if !revocable {
			return nil, ErrEnityTypecodeNotMatch
		}
		versionend := 5 + int(DataOrder.Uint64(data[1:5]))
		if datalength < versionend {
			return nil, ErrUnresolvedCacheEnity
		}
		versiondata := data[5:versionend]
		if bytes.Compare(version, versiondata) != 0 {
			return nil, ErrEnityVersionNotMatch
		}
		return createEnity(false, versiondata, data[versionend:]), nil
	}
	return nil, ErrUnresolvedCacheEnity
}
