package binaryutil

import (
	"math"
)

var DataTrue = byte(1)
var DataFalse = byte(0)

func Decode(d []byte, data interface{}) error {
	length := len(d)
	if length == 0 {
		return ErrDataLengthNotMatch
	}
	switch data := data.(type) {
	case *bool:
		if length != 1 {
			return ErrDataLengthNotMatch
		}
		*data = d[0] != 0
	case *int8:
		if length != 1 {
			return ErrDataLengthNotMatch
		}
		*data = int8(d[0])
	case *uint8:
		if length != 1 {
			return ErrDataLengthNotMatch
		}
		*data = d[0]
	case *int16:
		if length != 2 {
			return ErrDataLengthNotMatch
		}
		*data = int16(DataOrder.Uint16(d))
	case *uint16:
		if length != 2 {
			return ErrDataLengthNotMatch
		}
		*data = DataOrder.Uint16(d)
	case *int32:
		if length != 4 {
			return ErrDataLengthNotMatch
		}
		*data = int32(DataOrder.Uint32(d))
	case *uint32:
		if length != 4 {
			return ErrDataLengthNotMatch
		}
		*data = DataOrder.Uint32(d)
	case *int:
		if length != 8 {
			return ErrDataLengthNotMatch
		}
		*data = int(int64(DataOrder.Uint64(d)))
	case *uint:
		if length != 4 {
			return ErrDataLengthNotMatch
		}

		*data = uint(DataOrder.Uint32(d))
	case *int64:
		if length != 8 {
			return ErrDataLengthNotMatch
		}
		*data = int64(DataOrder.Uint64(d))
	case *uint64:
		if length != 8 {
			return ErrDataLengthNotMatch
		}
		*data = DataOrder.Uint64(d)
	case *float32:
		if length != 4 {
			return ErrDataLengthNotMatch
		}
		*data = math.Float32frombits(DataOrder.Uint32(d))
	case *float64:
		if length != 8 {
			return ErrDataLengthNotMatch
		}
		*data = math.Float64frombits(DataOrder.Uint64(d))
	case *[]byte:
		*data = d
	case *string:
		*data = string(d)
	default:
		return ErrDataTypeNotSupported
	}
	return nil
}

func Encode(data interface{}) ([]byte, error) {
	var d []byte
	switch data := data.(type) {
	case *bool:
		if *data {
			return []byte{DataTrue}, nil
		}
		return []byte{DataFalse}, nil
	case bool:
		if data {
			return []byte{DataTrue}, nil
		}
		return []byte{DataFalse}, nil
	case *int8:
		return []byte{byte(*data)}, nil
	case int8:
		return []byte{byte(data)}, nil
	case *uint8:
		return []byte{byte(*data)}, nil
	case uint8:
		return []byte{byte(data)}, nil
	case *int16:
		d = make([]byte, 2)
		DataOrder.PutUint16(d, uint16(*data))
		return d, nil
	case int16:
		d = make([]byte, 2)
		DataOrder.PutUint16(d, uint16(data))
		return d, nil
	case *uint16:
		d = make([]byte, 2)
		DataOrder.PutUint16(d, *data)
		return d, nil
	case uint16:
		d = make([]byte, 2)
		DataOrder.PutUint16(d, data)
		return d, nil
	case *int32:
		d = make([]byte, 4)
		DataOrder.PutUint32(d, uint32(*data))
		return d, nil
	case int32:
		d = make([]byte, 4)
		DataOrder.PutUint32(d, uint32(data))
		return d, nil
	case *int:
		d = make([]byte, 8)
		DataOrder.PutUint64(d, uint64(*data))
		return d, nil
	case int:
		d = make([]byte, 8)
		DataOrder.PutUint64(d, uint64(data))
		return d, nil
	case *uint32:
		d = make([]byte, 4)
		DataOrder.PutUint32(d, *data)
		return d, nil
	case uint32:
		d = make([]byte, 4)
		DataOrder.PutUint32(d, data)
		return d, nil
	case *uint:
		d = make([]byte, 4)
		DataOrder.PutUint32(d, uint32(*data))
		return d, nil
	case uint:
		d = make([]byte, 4)
		DataOrder.PutUint32(d, uint32(data))
		return d, nil
	case *int64:
		d = make([]byte, 8)
		DataOrder.PutUint64(d, uint64(*data))
		return d, nil
	case int64:
		d = make([]byte, 8)
		DataOrder.PutUint64(d, uint64(data))
		return d, nil
	case *uint64:
		d = make([]byte, 8)
		DataOrder.PutUint64(d, *data)
		return d, nil
	case uint64:
		d = make([]byte, 8)
		DataOrder.PutUint64(d, data)
		return d, nil
	case *float32:
		d = make([]byte, 4)
		DataOrder.PutUint32(d, math.Float32bits(*data))
		return d, nil
	case float32:
		d = make([]byte, 4)
		DataOrder.PutUint32(d, math.Float32bits(data))
		return d, nil
	case *float64:
		d = make([]byte, 8)
		DataOrder.PutUint64(d, math.Float64bits(*data))
		return d, nil
	case float64:
		d = make([]byte, 8)
		DataOrder.PutUint64(d, math.Float64bits(data))
		return d, nil
	case *[]byte:
		return *data, nil
	case *string:
		return []byte(*data), nil
	case []byte:
		return data, nil
	case string:
		return []byte(data), nil
	}
	return nil, ErrDataTypeNotSupported
}
