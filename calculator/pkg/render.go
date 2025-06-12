package pkg

import (
	"fmt"
	"math"
	"strings"
)

func Render(expression string, result float64) string {
	var resultStr string

	if math.Mod(result, 1) == 0 {
		resultStr = fmt.Sprintf("%d", int(result))
	} else {
		resultStr = fmt.Sprintf("%g", result)
	}

	boxWidth := max(len(expression), len(resultStr)) + 4

	var box []string
	box = append(box, "┌"+strings.Repeat("─", boxWidth)+"┐")
	box = append(box, fmt.Sprintf("│  %s%s│", expression, strings.Repeat(" ", boxWidth-len(expression)-2)))
	box = append(box, fmt.Sprintf("│%s│", strings.Repeat(" ", boxWidth)))
	box = append(box, fmt.Sprintf("│  =%s│", strings.Repeat(" ", boxWidth-3)))
	box = append(box, fmt.Sprintf("│%s│", strings.Repeat(" ", boxWidth)))
	box = append(box, fmt.Sprintf("│  %s%s│", resultStr, strings.Repeat(" ", boxWidth-len(resultStr)-2)))
	box = append(box, "└"+strings.Repeat("─", boxWidth)+"┘")

	return strings.Join(box, "\n")
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
