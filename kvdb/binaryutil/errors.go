package binaryutil

import "errors"

var ErrDataLengthNotMatch = errors.New("binarystream data length not match")
var ErrDataTypeNotSupported = errors.New("binarystream data type not supported")
