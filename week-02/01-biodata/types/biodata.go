package types

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
)

type Biodata struct {
	Nama      string
	Alamat    string
	Pekerjaan string
	Motivasi  string
}

type ListBiodata []Biodata

func (lb ListBiodata) Pompt(question string, reader io.Reader) (string, error) {
	fmt.Print(question)

	input, err := lb.inputFromStdin(reader)
	if err != nil {
		return "", err
	}

	return input, nil

}

func (lb ListBiodata) inputFromStdin(reader io.Reader) (string, error) {
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		return "", nil
	}
	return scanner.Text(), nil
}

func (lb ListBiodata) ValidateInput(input string) (int, error) {
	if input == "" {
		return -1, fmt.Errorf("invalid input: field is required")
	}

	inputNum, err := strconv.Atoi(input)
	if err != nil {
		return -1, fmt.Errorf("invalid input (%s): make sure entering valid number", input)
	}

	if inputNum < 1 {
		return -1, fmt.Errorf("invalid input (%d): must be positive >= 1", inputNum)
	}

	if inputNum > len(lb) {
		return -1, fmt.Errorf("no data found")
	}
	return inputNum, nil
}

func (lb ListBiodata) Search(index int) *Biodata {
	return &lb[index-1]
}
