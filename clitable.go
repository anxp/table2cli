package table2cli

import (
	"fmt"
	"math"
	"strconv"
)

func PrintTable(headerRow []string, content [][]string, cellContentWidth int) {
	//Do validation before start do anything
	numOfColumns := len(headerRow)
	numOfContentRows := len(content)

	for i := 0; i < numOfContentRows; i++ {
		if numOfColumns != len(content[i]) {
			fmt.Println("Data consistency error. Number of columns in header slice != number of columns in content slice.")
			return
		}
	}

	// Calculate table width (in symbols)
	// (border_sym + space + cellContentWidth + space) * numOfColumns + border_sym
	tableWidth := (1 + 1 + cellContentWidth + 1) * numOfColumns + 1

	//================ Print header row: ===============================================================================
	printHorizontalLine(tableWidth, "=")
	printRow(headerRow, numOfColumns, cellContentWidth)
	printHorizontalLine(tableWidth, "=")
	//==================================================================================================================

	//================ Print table content: ============================================================================
	for i := 0; i < len(content); i++ {
		printRow(content[i], numOfColumns, cellContentWidth)

		if i < len(content) - 1 {
			printHorizontalLine(tableWidth, "-")
		}
	}

	printHorizontalLine(tableWidth, "=")
	//==================================================================================================================
}

func printRow(row []string, numOfColumns int, cellContentWidth int) {
	//Calculate for how many lines current row will be expanded:
	numOfLines := calculateMaxNumberOfChunksForRow(row, cellContentWidth)

	//Break each cell's string into chunks to fit in column width:
	sliceOfBrokenStrings := make([][]string, numOfColumns)

	for i := 0; i < numOfColumns; i++ {
		sliceOfBrokenStrings[i] = breakStringIntoSlice(row[i], cellContentWidth, numOfLines)
	}

	for j := 0; j < numOfLines; j++ {
		fmt.Printf("| ") //Print opening (first) symbol in line

		for i := 0; i < numOfColumns; i++ {
			format := "%-" + strconv.Itoa(cellContentWidth) + "s"
			fmt.Printf(format, sliceOfBrokenStrings[i][j])

			if i < numOfColumns - 1 {
				fmt.Printf(" | ")
			}
		}

		fmt.Printf(" |\n") //Print closing (last) symbol in line
	}
}

func calculateMaxNumberOfChunksForRow(row []string, chunkLength int) int {
	maxNumberOfChunks := 0

	for i := 0; i < len(row); i++ {
		numberOfLetters := len(convertStringToRuneSlice(row[i]))
		numberOfChunksForCurrentString := int(math.Ceil(float64(numberOfLetters) / float64(chunkLength)))

		if numberOfChunksForCurrentString > maxNumberOfChunks {
			maxNumberOfChunks = numberOfChunksForCurrentString
		}
	}

	return maxNumberOfChunks
}

func breakStringIntoSlice(inputString string, chunkLength int, returnedSliceSize int) []string {
	var resultSlice []string

	inputAdapted := convertStringToRuneSlice(inputString)

	inputStringLength := len(inputAdapted)
	numberOfChunks := int(math.Ceil(float64(inputStringLength) / float64(chunkLength)))
	resultSliceSize := max(returnedSliceSize, numberOfChunks)

	for i := 0; i < inputStringLength; i += chunkLength {
		chunkEndPosition := min(i + chunkLength, inputStringLength)
		resultSlice = append(resultSlice, string(inputAdapted[i : chunkEndPosition]))
	}

	// If user ordered bigger slice than number of chunks (substrings) we have got,
	// let's append required quantity of empty elements to it:
	for ; len(resultSlice) < resultSliceSize;  {
		resultSlice = append(resultSlice, "")
	}

	return resultSlice
}

func convertStringToRuneSlice(mbString string) []rune {
	return []rune(mbString)
}

func printHorizontalLine(length int, symbol string) {
	for i := 0; i < length; i++ {
		fmt.Print(symbol)

		if i == length - 1 {
			fmt.Print("\n")
		}
	}
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}

	return y
}