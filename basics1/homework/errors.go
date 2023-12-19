package main

import "errors"

var (
	errorParamDataLen     = errors.New("length of param data out of cell width")
	errorInvalidParamData = errors.New("invalid paramData")
	errorInvalidColor     = errors.New("invalid color")
	errorInvalidCellType  = errors.New("invalid cellTypeName")
)
