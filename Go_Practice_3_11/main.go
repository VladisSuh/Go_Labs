package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("%s\n", comma(os.Args[i]))
	}
}

func comma(s string) string {
	b := bytes.Buffer{}
	Start := 0
	if s[0] == '+' || s[0] == '-' {
		b.WriteByte(s[0])
		Start = 1
	}
	End := strings.Index(s, ".")
	if End == -1 {
		End = len(s)
	}
	mantissa := s[Start:End]
	pre := len(mantissa) % 3
	if pre > 0 {
		b.Write([]byte(mantissa[:pre]))
		if len(mantissa) > pre {
			b.WriteString(",")
		}
	}
	for i, c := range mantissa[pre:] {
		if i%3 == 0 && i != 0 {
			b.WriteString(",")
		}
		b.WriteRune(c)
	}
	b.WriteString(s[End:])
	return b.String()
}
