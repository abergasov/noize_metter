package substation_test

import (
	"bufio"
	"encoding/csv"
	"fmt"
	testhelpers "noize_metter/internal/test_helpers"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestService(t *testing.T) {
	container := testhelpers.GetClean(t)
	require.NoError(t, container.ServiceSubstation.WrapIteration())
}

func TestGenerateInsert(t *testing.T) {
	file, err := os.Open("/home/alejandro/go/src/noize_metter/internal/entities/modbus.go")
	require.NoError(t, err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := scanner.Text()
		if !strings.Contains(row, "parquet") {
			continue
		}
		dataRows := strings.Split(row, "json:")
		subData := strings.Split(dataRows[0], " ")
		dataRows[1] = strings.ReplaceAll(dataRows[1], "`", "")
		subData[0] = strings.Trim(subData[0], " \t")
		println(fmt.Sprintf(`%s: item.%s,`, dataRows[1], subData[0]))
	}
}

func TestItemsList(t *testing.T) {
	rows := readCsvFile(t, "/home/alejandro/Downloads/input_registers.csv")
	rowsD := readCsvFile(t, "/home/alejandro/Downloads/discrete_registers.csv")
	targetFile := "/home/alejandro/go/src/noize_metter/internal/entities/modbus.go"

	newRows := make([]string, 0, len(rows))
	newRowsDiscrete := make([]string, 0, len(rows))
	sqlRows := make([]string, 0, len(rows))
	for i, r := range rows {
		if i == 0 {
			continue
		}
		row := prepareString(r[0])
		//scale := prepareString(r[3])
		measure := prepareStringMeasure(r[4])
		source := prepareStringSource(r[5])

		if measure != "" {
			row += "_" + measure
		}
		if source != "" {
			row += "_" + source
		}
		row = strings.ReplaceAll(row, "_alarm_alarm", "_alarm")
		row = cutDigitsAtFront(row)
		newRows = append(newRows, row)
		spaces := strings.Repeat(" ", 79-len(row)-len("FLOAT NOT NULL,"))
		sqlRows = append(sqlRows, fmt.Sprintf("%s%sFLOAT NOT NULL,", row, spaces))
	}
	sqlRows = append(sqlRows, " ")
	for i, r := range rowsD {
		if i == 0 {
			continue
		}
		row := prepareString(r[0])
		trueState := prepareString(r[4])
		source := prepareStringSource(r[5])
		if trueState != "" {
			row += "_" + trueState
		}
		if source != "" {
			row += "_" + source
		}
		row = strings.ReplaceAll(row, "_alarm_alarm", "_alarm")
		row = cutDigitsAtFront(row)
		newRowsDiscrete = append(newRowsDiscrete, row)
		spaces := strings.Repeat(" ", 70-len(row)-len("BIT NOT NULL,"))
		sqlRows = append(sqlRows, fmt.Sprintf("%s%sBIT NOT NULL,", row, spaces))
	}
	println(strings.Join(sqlRows, "\n"))
	println(targetFile)
	replaceAnnotations(t, targetFile, "%%%", newRows)
	replaceAnnotations(t, targetFile, "@@@", newRowsDiscrete)
}

func replaceAnnotations(t *testing.T, targetFile, param string, newRows []string) {
	file, err := os.Open(targetFile)
	require.NoError(t, err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	requltRows := make([]string, 0, len(newRows))
	i := 0
	for scanner.Scan() {
		row := scanner.Text()
		if !strings.Contains(row, param) {
			requltRows = append(requltRows, row)
			continue
		}
		row = strings.ReplaceAll(row, param, newRows[i])
		i++
		requltRows = append(requltRows, row)
	}

	require.NoError(t, scanner.Err())
	require.NoError(t, file.Close())

	file, err = os.Create(targetFile)
	require.NoError(t, err)
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, row := range requltRows {
		_, err = writer.WriteString(row + "\n")
		require.NoError(t, err)
	}
	require.NoError(t, writer.Flush())
}

func readCsvFile(t *testing.T, filePath string) [][]string {
	f, err := os.Open(filePath)
	require.NoError(t, err)
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	require.NoError(t, err)

	return records
}

func prepareString(row string) string {
	row = strings.ToLower(row)
	return prepare(row)
}

func prepareStringMeasure(row string) string {
	return prepare(row)
}
func prepareStringSource(row string) string {
	row = strings.ToLower(row)
	row = prepare(row)
	row = strings.ReplaceAll(row, "_", "")
	return row
}

func prepare(row string) string {
	row = strings.TrimSpace(row)
	row = strings.ReplaceAll(row, " ", "_")
	row = strings.ReplaceAll(row, "__", "_")
	row = strings.ReplaceAll(row, "___", "_")
	row = strings.ReplaceAll(row, "°c", "c")
	row = strings.ReplaceAll(row, "-", "")
	row = strings.ReplaceAll(row, ".", "")
	row = strings.ReplaceAll(row, "/", "_")
	row = strings.ReplaceAll(row, "(", "")
	row = strings.ReplaceAll(row, "%", "Percent")
	row = strings.ReplaceAll(row, ")", "")
	return row
}

func cutDigitsAtFront(row string) string {
	row = strings.TrimLeft(row, "0123456789_")
	return row
}
