package main

import (
	"fmt"
	"io"
	"os"
)

type logWriter struct{}

func main() {
	args := os.Args

	if len(args) <= 1 {
		fmt.Println("This program needs to be provided 1 file name to read. E.g., go run main.go test.txt")
		os.Exit(1)
	}

	filenames := args[1:]

	for _, filename := range filenames {
		printFileContents(filename)
	}
}

func printFileContents(filename string) {
	f, err := os.Open(filename)

	if err != nil {
		fmt.Println("Something went wrong when opening the provided file:", err)
	}

	// 1st option

	/* content := make([]byte, 99999)

	_, err = f.Read(content)

	if err != nil {
		fmt.Println("Something went wrong when reading the provided file:", err)
	}

	fmt.Println(string(content)) */

	// 2nd option (copy to std output)

	//io.Copy(os.Stdout, f)

	// or 3rd option; we could also have our custom writer
	io.Copy(logWriter{}, f)
}

func (l logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))

	return len(bs), nil
}
