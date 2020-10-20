package binaryutil

import (
	"encoding/binary"
	"math"
)

var DataByteOrder = binary.BigEndian

type Data []byte

var True = byte(1)
var False = byte(0)

func (d Data) Read(data interface{}) error {
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
		*data = int16(DataByteOrder.Uint16(d))
	case *uint16:
		if length != 2 {
			return ErrDataLengthNotMatch
		}
		*data = DataByteOrder.Uint16(d)
	case *int32:
		if length != 4 {
			return ErrDataLengthNotMatch
		}
		*data = int32(DataByteOrder.Uint32(d))
	case *uint32:
		if length != 4 {
			return ErrDataLengthNotMatch
		}
		*data = DataByteOrder.Uint32(d)
	case *int:
		if length != 4 {
			return ErrDataLengthNotMatch
		}
		*data = int(int32(DataByteOrder.Uint32(d)))
	case *uint:
		if length != 4 {
			return ErrDataLengthNotMatch
		}

		*data = uint(DataByteOrder.Uint32(d))
	case *int64:
		if length != 8 {
			return ErrDataLengthNotMatch
		}
		*data = int64(DataByteOrder.Uint64(d))
	case *uint64:
		if length != 8 {
			return ErrDataLengthNotMatch
		}
		*data = DataByteOrder.Uint64(d)
	case *float32:
		if length != 4 {
			return ErrDataLengthNotMatch
		}
		*data = math.Float32frombits(DataByteOrder.Uint32(d))
	case *float64:
		if length != 8 {
			return ErrDataLengthNotMatch
		}
		*data = math.Float64frombits(DataByteOrder.Uint64(d))
	case *[]byte:
		*data = d
	case *string:
		*data = string(d)
	default:
		return ErrDataTypeNotSupported
	}
	return nil
}

func CreateData(data interface{}) (Data, error) {
	var d Data
	switch data := data.(type) {
	case *bool:
		if *data {
			return Data{True}, nil
		}
		return Data{False}, nil
	case bool:
		if data {
			return Data{True}, nil
		}
		return Data{False}, nil
	case *int8:
		return Data{byte(*data)}, nil
	case int8:
		return Data{byte(data)}, nil
	case *uint8:
		return Data{byte(*data)}, nil
	case uint8:
		return Data{byte(data)}, nil
	case *int16:
		d = make(Data, 2)
		DataByteOrder.PutUint16(d, uint16(*data))
		return d, nil
	case int16:
		d = make(Data, 2)
		DataByteOrder.PutUint16(d, uint16(data))
		return d, nil
	case *uint16:
		d = make(Data, 2)
		DataByteOrder.PutUint16(d, *data)
		return d, nil
	case uint16:
		d = make(Data, 2)
		DataByteOrder.PutUint16(d, data)
		return d, nil
	case *int32:
		d = make(Data, 4)
		DataByteOrder.PutUint32(d, uint32(*data))
		return d, nil
	case int32:
		d = make(Data, 4)
		DataByteOrder.PutUint32(d, uint32(data))
		return d, nil
	case *int:
		d = make(Data, 4)
		DataByteOrder.PutUint32(d, uint32(*data))
		return d, nil
	case int:
		d = make(Data, 4)
		DataByteOrder.PutUint32(d, uint32(data))
		return d, nil
	case *uint32:
		d = make(Data, 4)
		DataByteOrder.PutUint32(d, *data)
		return d, nil
	case uint32:
		d = make(Data, 4)
		DataByteOrder.PutUint32(d, data)
		return d, nil
	case *uint:
		d = make(Data, 4)
		DataByteOrder.PutUint32(d, uint32(*data))
		return d, nil
	case uint:
		d = make(Data, 4)
		DataByteOrder.PutUint32(d, uint32(data))
		return d, nil
	case *int64:
		d = make(Data, 8)
		DataByteOrder.PutUint64(d, uint64(*data))
		return d, nil
	case int64:
		d = make(Data, 8)
		DataByteOrder.PutUint64(d, uint64(data))
		return d, nil
	case *uint64:
		d = make(Data, 8)
		DataByteOrder.PutUint64(d, *data)
		return d, nil
	case uint64:
		d = make(Data, 8)
		DataByteOrder.PutUint64(d, data)
		return d, nil
	case *float32:
		d = make(Data, 4)
		DataByteOrder.PutUint32(d, math.Float32bits(*data))
		return d, nil
	case float32:
		d = make(Data, 4)
		DataByteOrder.PutUint32(d, math.Float32bits(data))
		return d, nil
	case *float64:
		d = make(Data, 8)
		DataByteOrder.PutUint64(d, math.Float64bits(*data))
		return d, nil
	case float64:
		d = make(Data, 8)
		DataByteOrder.PutUint64(d, math.Float64bits(data))
		return d, nil
	case *[]byte:
		return *data, nil
	case *string:
		return Data(*data), nil

	case []byte:
		return data, nil
	case string:
		return Data(data), nil

	}
	return nil, ErrDataTypeNotSupported
}
