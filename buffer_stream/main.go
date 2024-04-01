package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	m := int64(8) // buffer
	n := 3        // to read

	start := int64(0) // starting of file is 0
	file, err := os.Open("file.txt")

	if err != nil {
		fmt.Println(err)
	}

	for {
		// read 5 bytes
		size := make([]byte, n)

		read, err := file.ReadAt(size, start)
		if err == io.EOF {
			break
		}
		if read < n {
			fmt.Println("read is : ", read)
			break
		}
		start += m
		fmt.Printf("%q\n", size[:n])
	}

}
