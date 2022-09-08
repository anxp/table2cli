package table2cli

import (
	"errors"
	"fmt"
	"math"
	"strconv"
)

type ContentCell struct {
	Data string
	PrintBottomBorder bool
}

type TableHeader []string
type TableContent [][]ContentCell

type Table struct {
	Header TableHeader
	Content TableContent
	ColumnWidths []int
}

func NewTable(header TableHeader, content TableContent, widths []int) (*Table, error) {
	// ================================= VALIDATE BEFORE CONSTRUCT: ====================================================
	numOfColumns := len(header)
	numOfContentRows := len(content)

	if numOfColumns != len(widths) {
		return nil, errors.New("can't create table - column widths values are missed (or in excess)")
	}

	for i:= 0; i < numOfContentRows; i++ {
		if numOfColumns != len(content[i]) {
			return nil, errors.New("data consistency error - number of columns in header slice != number of columns in content slice")
		}
	}
	// =================================================================================================================

	return &Table{Header: header, Content: content, ColumnWidths: widths}, nil
}

func (t *Table) Print() {
	numberOfColumns := len(t.Header)

	// Calculate table width (in symbols)
	// Reminder: t.ColumnWidths[i] is a width of the CONTENT IN THE CELL, not the total width of column!
	tableWidth := 0
	for i := 0; i < numberOfColumns; i++ {
		tableWidth += 1 + 1 + t.ColumnWidths[i] + 1

		if i == numberOfColumns - 1 {
			tableWidth += 1
		}
	}

	//================ Print header row: ===============================================================================
	t.printHorizontalLine(tableWidth, "=", true)
	t.printRow(t.Header)
	t.printHorizontalLine(tableWidth, "=", true)
	//==================================================================================================================

	//================ Print table content: ============================================================================
	for currentRowNumber := 0; currentRowNumber < len(t.Content); currentRowNumber++ {
		currentRow := t.Content[currentRowNumber]

		currentRowContent := make([]string, numberOfColumns)
		bottomLinePrintingInfo := make([]bool, numberOfColumns)
		for i := 0; i < numberOfColumns; i++ {
			currentRowContent[i] = currentRow[i].Data
			bottomLinePrintingInfo[i] = currentRow[i].PrintBottomBorder
		}

		t.printRow(currentRowContent)

		if currentRowNumber < len(t.Content) - 1 {

			// Print bottom line of current ROW. Please note, that this line can have unfilled segments (when PrintBottomBorder == false)
			for i := 0; i < numberOfColumns; i++ {

				// CHOICE of the FIRST symbol of the bottom line of CURRENT COLUMN:
				if (bottomLinePrintingInfo[i] == true) && (i > 0 && bottomLinePrintingInfo[i - 1] == true) {
					fmt.Print("*")
				} else {
					fmt.Print("|")
				}

				// Inner column width = one space symbol + actual content width + second space symbol.
				// This is column width without left and right borders.
				innerColumnWidth := 1 + t.ColumnWidths[i] + 1

				if bottomLinePrintingInfo[i] == true {
					t.printHorizontalLine(innerColumnWidth, "-", false)
				} else {
					t.printHorizontalLine(innerColumnWidth, " ", false)
				}

				// If we at the LAST column in ROW:
				if i == numberOfColumns - 1 {
					fmt.Print("|\n")
				}
			}
		}
	}

	t.printHorizontalLine(tableWidth, "=", true)
	//==================================================================================================================
}

func (t *Table) printRow(cells []string) {

	if len(cells) != len(t.ColumnWidths) {
		// TODO: Panic
	}

	numOfColumns := len(cells)

	//Calculate for how many lines current row will be expanded:
	numOfLines := t.calculateMaxNumberOfChunksForRow(cells)

	//Break each cell's string into chunks to fit in column width:
	sliceOfBrokenStrings := make([][]string, numOfColumns)

	for i := 0; i < numOfColumns; i++ {
		sliceOfBrokenStrings[i] = t.breakStringIntoSlice(cells[i], t.ColumnWidths[i], numOfLines)
	}

	for j := 0; j < numOfLines; j++ {
		fmt.Printf("| ") //Print opening (first) symbol in line

		for i := 0; i < numOfColumns; i++ {
			format := "%-" + strconv.Itoa(t.ColumnWidths[i]) + "s"
			fmt.Printf(format, sliceOfBrokenStrings[i][j])

			if i < numOfColumns - 1 {
				fmt.Printf(" | ")
			}
		}

		fmt.Printf(" |\n") //Print closing (last) symbol in line
	}
}

func (t *Table) calculateMaxNumberOfChunksForRow(cells []string) int {
	maxNumberOfChunks := 0

	for i := 0; i < len(cells); i++ {
		numberOfLetters := len(t.convertStringToRuneSlice(cells[i]))
		numberOfChunksForCurrentString := int(math.Ceil(float64(numberOfLetters) / float64(t.ColumnWidths[i])))

		if numberOfChunksForCurrentString > maxNumberOfChunks {
			maxNumberOfChunks = numberOfChunksForCurrentString
		}
	}

	return maxNumberOfChunks
}

func (t *Table) breakStringIntoSlice(inputString string, chunkLength int, returnedSliceSize int) []string {
	var resultSlice []string

	inputAdapted := t.convertStringToRuneSlice(inputString)

	inputStringLength := len(inputAdapted)
	numberOfChunks := int(math.Ceil(float64(inputStringLength) / float64(chunkLength)))
	resultSliceSize := t.max(returnedSliceSize, numberOfChunks)

	for i := 0; i < inputStringLength; i += chunkLength {
		chunkEndPosition := t.min(i + chunkLength, inputStringLength)
		resultSlice = append(resultSlice, string(inputAdapted[i : chunkEndPosition]))
	}

	// If user ordered bigger slice than number of chunks (substrings) we have got,
	// let's append required quantity of empty elements to it:
	for ; len(resultSlice) < resultSliceSize;  {
		resultSlice = append(resultSlice, "")
	}

	return resultSlice
}

func (t *Table) convertStringToRuneSlice(mbString string) []rune {
	return []rune(mbString)
}

func (t *Table) printHorizontalLine(length int, symbol string, terminateLine bool) {
	for i := 0; i < length; i++ {
		fmt.Print(symbol)

		if terminateLine && i == length - 1 {
			fmt.Print("\n")
		}
	}
}

func (t *Table) max(x, y int) int {
	if x > y {
		return x
	}

	return y
}

func (t *Table) min(x, y int) int {
	if x < y {
		return x
	}

	return y
}