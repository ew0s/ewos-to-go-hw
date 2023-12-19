package main

import "strings"

const (
	width       = 40
	borderWidth = 1
)

func drawMarkLine(sb *strings.Builder, border rune) {
	for i := 0; i < width+borderWidth*2; i++ {
		sb.WriteRune(border)
	}
	sb.WriteByte('\n')
}

func drawDataLine(sb *strings.Builder, border rune, fieldSpace rune, data paramData) {
	for i := 0; i < borderWidth; i++ {
		sb.WriteRune(border)
	}

	sb.WriteString(data)

	for i := 0; i < width-len([]rune(data)); i++ {
		sb.WriteRune(fieldSpace)
	}

	for i := 0; i < borderWidth; i++ {
		sb.WriteRune(border)
	}

	sb.WriteByte('\n')
}
