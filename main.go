package main

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/djherbis/times.v1"
)

func run(file string) error {
	t, err := times.Stat(file)
	if err != nil {
		return err
	}

	fmt.Printf("- file: %s\n", file)
	fmt.Printf("  atime: %v\n", t.AccessTime())
	fmt.Printf("  mtime: %v\n", t.ModTime())

	if t.HasChangeTime() {
		fmt.Printf("  ctime: %v\n", t.ChangeTime())
	}

	if t.HasBirthTime() {
		fmt.Printf("  btime: %v\n", t.BirthTime())
	}

	return nil
}

func main() {
	files := os.Args
	for idx, file := range files {
		if idx == 0 {
			continue
		}
		if err := run(file); err != nil {
			log.Fatal(err)
		}
	}
}
