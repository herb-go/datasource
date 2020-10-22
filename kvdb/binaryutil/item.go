package binaryutil

import "io"

type item struct {
	v interface{}
}

func (i *item) PackBinary(w io.Writer) error {
	data, err := Encode(i.v)
	if err != nil {
		return err
	}
	return PackTo(w, nil, data)
}

func (i *item) UnpackBinary(r io.Reader) error {
	data, err := UnpackFrom(r, nil)
	if err != nil {
		return err
	}
	return Decode(data, i.v)
}

type packerFunc func(w io.Writer) error

func (f packerFunc) PackBinary(w io.Writer) error {
	return f(w)
}

type unpackerFunc func(r io.Reader) error

func (f unpackerFunc) UnpackBinary(r io.Reader) error {
	return f(r)
}

type packers struct {
	items []BinaryPacker
}

func (p *packers) PackBinary(w io.Writer) error {
	var err error
	for k := range p.items {
		err = p.items[k].PackBinary(w)
		if err != nil {
			return err
		}
	}
	return nil
}

type unpackers struct {
	items []BinaryUnpacker
}

func (p *unpackers) UnpackBinary(r io.Reader) error {
	var err error
	for k := range p.items {
		err = p.items[k].UnpackBinary(r)
		if err != nil {
			return err
		}
	}
	return nil
}

func Unpackers(values ...interface{}) BinaryUnpacker {
	us := &unpackers{
		items: make([]BinaryUnpacker, len(values)),
	}
	for k, v := range values {
		u, ok := v.(BinaryUnpacker)
		if ok {
			us.items[k] = u
			continue
		}
		f, ok := v.(func(r io.Reader) error)
		if ok {
			us.items[k] = unpackerFunc(f)
			continue
		}
		us.items[k] = &item{
			v: v,
		}
	}
	return us
}

func Packers(values ...interface{}) BinaryPacker {
	ps := &packers{
		items: make([]BinaryPacker, len(values)),
	}
	for k, v := range values {
		p, ok := v.(BinaryPacker)
		if ok {
			ps.items[k] = p
			continue
		}
		f, ok := v.(func(w io.Writer) error)
		if ok {
			ps.items[k] = packerFunc(f)
			continue
		}
		ps.items[k] = &item{
			v: v,
		}
	}
	return ps
}
