package board

import "fmt"

const (
	GraphicalMode = iota
	FenMode
)

func (b *Board) Debug(mode int) string {
	switch mode {
	case GraphicalMode:
		return b.grafical()
	case FenMode:
		return b.fen()
	}

	return ""
}

func (b *Board) grafical() (result string) {
	for i := 0; i < sizeOfBoard; i++ {
		line := ""

		for j := 0; j < sizeOfBoard; j++ {
			if b.rows[i][j] != nil {
				line += string(b.rows[i][j].Show())
			} else {
				line += " "
			}

			if j != 7 {
				line += " "
			}
		}
		result += line + "\n"
	}

	return
}

func (b *Board) fen() (result string) {
	counterBlank := 0

	for idxRow, row := range b.rows {
		for _, column := range row {
			if column == nil {
				counterBlank++
			} else {
				result += printCounter(&counterBlank)
				result += string(column.ShowFEN())
			}
		}
		result += printCounter(&counterBlank)

		if idxRow != len(b.rows)-1 {
			result += "/"
		}
	}

	return
}

func printCounter(counter *int) string {
	if *counter > 0 {
		aux := *counter
		*counter = 0

		return fmt.Sprint(aux)
	}

	return ""
}
