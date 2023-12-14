package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
	}
}

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}

	var buf bytes.Buffer
	d, t := len(s)%3, len(s)/3

	if d != 0 {
		buf.WriteString(s[:d])
		buf.WriteByte(',')
	}

	for i := 0; i < t; i++ {
		buf.WriteString(s[i*3+d : d+i*3+3])
		if i != t-1 {
			buf.WriteByte(',')
		}
	}
	return buf.String()
}
