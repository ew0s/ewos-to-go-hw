package main

import (
	"strings"
)

type color = string

const (
	red   color = "red"
	green color = "green"
	blue  color = "blue"
)

type colorCode = string

const (
	colorCodeRed   colorCode = "31"
	colorCodeGreen colorCode = "32"
	colorCodeBlue  colorCode = "34"
)

func paintText(sb *strings.Builder, c color) error {
	cCode, ok := map[color]colorCode{
		red:   colorCodeRed,
		blue:  colorCodeBlue,
		green: colorCodeGreen,
	}[c]
	if !ok {
		return errorInvalidColor
	}

	sb.WriteString("\033[" + cCode + "m")

	return nil
}
