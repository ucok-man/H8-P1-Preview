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

	result := spawn(bytes.NewBuffer(content))
	if err := result.err; err != nil {
		return err
	}

	for word, count := range result.data {
		fmt.Printf("%-20s : %d\n", word, count)
	}

	return nil
}

type Result struct {
	mu   *sync.Mutex
	data map[string]int
	err  error
}

func spawn(content *bytes.Buffer) *Result {
	var wg sync.WaitGroup

	input := make(chan io.Reader, runtime.NumCPU())
	result := &Result{
		data: make(map[string]int),
		mu:   &sync.Mutex{},
	}

	scanner := bufio.NewScanner(content)
	for scanner.Scan() && result.err == nil {
		wg.Add(1)

		go func(input <-chan io.Reader) {
			defer wg.Done()
			readAndInsert(input, result)
		}(input)

		input <- strings.NewReader(scanner.Text())
	}

	wg.Wait()
	return result
}

func readAndInsert(input <-chan io.Reader, result *Result) {
	scanner := bufio.NewScanner(<-input)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		word := scanner.Text()
		result.mu.Lock()
		result.data[word]++
		result.mu.Unlock()
	}

	if err := scanner.Err(); err != nil {
		result.err = err
	}
}
