package main

import (
	"fmt"
	"io"
	"os"

	"github.com/ucok-man/PHASE_1/week_2/day1/01-biodata/types"
)

func main() {
	bio := types.ListBiodata{
		{
			Nama:      "Teman 1",
			Alamat:    "Alamat teman 1",
			Pekerjaan: "Pekerjaan",
			Motivasi:  "Motivasi 1",
		},
		{
			Nama:      "Teman 2",
			Alamat:    "Alamat teman 2",
			Pekerjaan: "Pekerjaan",
			Motivasi:  "Motivasi 2",
		},
		{
			Nama:      "Teman 3",
			Alamat:    "Alamat teman 3",
			Pekerjaan: "Pekerjaan",
			Motivasi:  "Motivasi 3",
		},
	}

	if err := run(os.Stdin, bio); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(reader io.Reader, data types.ListBiodata) error {
	input, err := data.Pompt("Input absen: ", reader)
	if err != nil {
		return err
	}

	inputAbsen, err := data.ValidateInput(input)
	if err != nil {
		return err
	}

	biodata := data.Search(inputAbsen)

	show(inputAbsen, biodata)
	return nil
}

func show(idx int, data *types.Biodata) {
	fmt.Println()
	fmt.Printf("Data teman dengan absen %d\n", idx)
	fmt.Printf("%-10s: %v\n", "Nama", data.Nama)
	fmt.Printf("%-10s: %v\n", "Alamat", data.Alamat)
	fmt.Printf("%-10s: %v\n", "Pekerjaan", data.Pekerjaan)
	fmt.Printf("%-10s: %v\n", "Motivasi", data.Motivasi)
	fmt.Println()
}
