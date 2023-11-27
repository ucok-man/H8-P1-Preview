package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"runtime"
	"strings"
	"sync"
)

func run(fpath string) error {
	content, err := os.ReadFile(fpath)
	if err != nil {
		return err
	}

	result, err := spawn(bytes.NewBuffer(content))
	if err != nil {
		return err
	}

	for word, count := range result {
		fmt.Printf("%-20s : %d\n", word, count)
	}

	return nil
}

func spawn(content *bytes.Buffer) (map[string]int, error) {
	var wg sync.WaitGroup
	var mu = &sync.Mutex{}

	result := make(map[string]int)
	input := make(chan io.Reader, runtime.NumCPU())

	var err error
	scanner := bufio.NewScanner(content)
	
	for scanner.Scan() && err == nil {
		wg.Add(1)

		go func(input <-chan io.Reader) {
			defer wg.Done()

			reader := <-input
			errcw := countWord(reader, result, mu)
			if errcw != nil {
				err = errcw
			}
		}(input)

		input <- strings.NewReader(scanner.Text())
	}

	wg.Wait()
	return result, err
}

func countWord(reader io.Reader, result map[string]int, mu *sync.Mutex) error {
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		word := scanner.Text()
		mu.Lock()
		result[word]++
		mu.Unlock()
	}

	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}
