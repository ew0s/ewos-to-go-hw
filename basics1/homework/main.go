package main

import (
	"fmt"
	"strings"
)

func drawCell(cTypeName cellTypeName, color color, params ...param) error {
	cType, err := getCellType(cTypeName)
	if err != nil {
		return err
	}

	sb := &strings.Builder{}
	if err = paintText(sb, color); err != nil {
		return err
	}

	drawMarkLine(sb, cType[horizontalBorder])

	data, err := getParamData(name)
	if err != nil {
		return err
	}

	drawDataLine(sb, cType[verticalBorder], cType[fieldSpace], *data)

	printedParams := map[param]struct{}{name: {}}

	for i := 0; i < len(params); i++ {
		if _, ok := printedParams[params[i]]; ok {
			continue
		}

		printedParams[params[i]] = struct{}{}

		drawMarkLine(sb, cType[horizontalBorder])

		data, err := getParamData(params[i])
		if err != nil {
			return err
		}

		drawDataLine(sb, cType[verticalBorder], cType[fieldSpace], *data)
	}

	drawMarkLine(sb, cType[horizontalBorder])

	fmt.Println(sb.String() + "\033[0m")

	return nil
}

func main() {
	if err := drawCell(line, blue, description, location, price, delivery); err != nil {
		fmt.Println(err)
	}
}
