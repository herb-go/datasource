package kvcache

import (
	"io"
)

func AppendBytes(key []byte, newkey []byte) []byte {
	var k []byte
	if len(newkey) < 255 {
		k = make([]byte, len(key)+1+len(newkey))
	} else {
		k = make([]byte, len(key)+9+len(newkey))
	}
	copy(k[0:len(key)], key)
	if len(newkey) < 255 {
		k[len(key)] = byte(len(newkey))
		copy(k[len(key)+1:], newkey)
	} else {
		k[len(key)] = 255
		DataOrder.PutUint64(k[len(key)+1:len(key)+9], uint64(len(newkey)))
		copy(k[len(key)+9:], newkey)
	}
	return k
}

func SplitBytes(data []byte) ([]byte, []byte, error) {
	var length int
	var start int
	if len(data) == 0 {
		return nil, nil, io.EOF
	}
	code := data[0]
	if code == 255 {
		if len(data) < 9 {
			return nil, nil, io.EOF
		}
		start = 9
		length = int(DataOrder.Uint64(data[1:9]))
	} else {
		start = 1
		length = int(code)
	}
	return data[start : start+length], data[start+length:], nil
}
