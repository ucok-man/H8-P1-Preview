package main

import (
	"bytes"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"runtime"
	"strings"
	"sync"
)

func run(inputFpath, outputFpath string) error {
	content, err := os.ReadFile(inputFpath)
	if err != nil {
		return err
	}

	buff, err := spawn(bytes.NewBuffer(content))
	if err != nil && !errors.Is(err, io.EOF) {
		return err
	}

	outfile, err := os.Create(outputFpath)
	if err != nil {
		return err
	}

	if _, err := outfile.WriteString(buff.String()); err != nil {
		return err
	}

	return nil
}

func spawn(content *bytes.Buffer) (*bytes.Buffer, error) {
	var wg sync.WaitGroup
	var mu *sync.Mutex = &sync.Mutex{}

	input := make(chan []string, runtime.NumCPU())
	buffer := bytes.Buffer{}

	csvreader := csv.NewReader(content)
	rows, err := csvreader.Read()
	if err != nil {
		return nil, err
	}

	err = writeCSV(rows, &buffer)
	if err != nil {
		return nil, err
	}

	for rows, err = csvreader.Read(); err == nil; rows, err = csvreader.Read() {
		wg.Add(1)

		go func(input <-chan []string) {
			defer wg.Done()

			rows := <-input
			errcsv := process(rows, &buffer, mu)
			if errcsv != nil {
				err = errcsv
			}
		}(input)

		input <- rows
	}

	wg.Wait()
	return &buffer, err
}

func process(rows []string, writer io.Writer, mu *sync.Mutex) error {
	if len(rows) != 3 {
		return fmt.Errorf("invalid csv input: expected csv row to have 3 length")
	}
	rows[0] = strings.ToUpper(rows[0])
	rows[2] = fmt.Sprintf("Mr.%s", rows[2])

	csvwritter := csv.NewWriter(writer)

	mu.Lock()
	err := csvwritter.Write(rows)
	if err != nil {
		return err
	}
	csvwritter.Flush()
	mu.Unlock()

	return nil
}

func writeCSV(rows []string, writer io.Writer) error {
	csvwritter := csv.NewWriter(writer)
	defer csvwritter.Flush()

	err := csvwritter.Write(rows)
	if err != nil {
		return err
	}

	return nil
}
