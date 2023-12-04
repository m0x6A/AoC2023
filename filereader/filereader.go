package filereader

import (
	"bufio"
	"os"
)

func LoadData(day string) ([]string, error) {
	lines, err := readFile("../day" + day + "/day" + day + ".txt")
	return lines, err
}
func readFile(filepath string) ([]string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}
