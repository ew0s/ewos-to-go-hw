package main

type cellTypeName = string

const (
	line        cellTypeName = "line"
	transparent cellTypeName = "transparent"
)

type cellType = [3]rune

const (
	horizontalBorder = iota
	verticalBorder
	fieldSpace
)

func getCellType(cTypeName cellTypeName) (*cellType, error) {
	cType, ok := map[cellTypeName]cellType{
		line:        {'_', '|', ' '},
		transparent: {' ', ' ', ' '},
	}[cTypeName]
	if !ok {
		return nil, errorInvalidCellType
	}

	return &cType, nil
}
