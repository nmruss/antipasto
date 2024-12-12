package filemgmt

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/gorilla/css/scanner"
)

type Property struct {
	name  string
	value string
}

// NOTE: at some point maybe make this safer with our own tokenType?
type CSSToken struct {
	Type  string
	Value string
}

func ParseCSSFromFile(filepath *string) []CSSToken {
	//takes in a CSS file and returns its properties
	//as an array of tokens
	var cssTokens []CSSToken

	file, err := os.OpenFile(*filepath, os.O_RDONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()

	r := bufio.NewReader(file)
	buf := make([]byte, 1024)
	for {
		n, err := r.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if n == 0 {
			break
		}
		s := scanner.New(string(buf))
		tok := s.Next()
		for tok.Type.String() != "EOF" {
			var cssTok CSSToken
			cssTok.Type = tok.Type.String()
			cssTok.Value = tok.Value
			cssTokens = append(cssTokens, cssTok)
			tok = s.Next()
		}
	}
	return cssTokens
}

func printBuffer(buf *[]byte, asString bool) {
	//prints a byte buffer as bytes or a string
	//followed by a new line character
	buffer := *buf
	for n := 0; n < len(buffer); n++ {
		if asString {
			fmt.Print(string(buffer[n]))
		} else {
			fmt.Print(buffer[n])
		}
		if buffer[n] == 0 {
			fmt.Println()
			break
		}
	}
}

func WriteCSS(filepath *string, id string, properties []Property) {
	//Writes a set of properties to an existing CSS object under the given id
	//takes an array of type Property
	// file, err := os.OpenFile(*filepath, os.O_RDWR, 0644)
	// if err != nil {
	// 	panic(err)
	// }
	// defer func() {
	// 	if err := file.Close(); err != nil {
	// 		panic(err)
	// 	}
	// }()

	// r := bufio.NewReader(file)
	// w := bufio.NewWriter(file)
	// buf := make([]byte, 1024)
	// for {
	// 	n, err := r.Read(buf)
	// 	if err != nil && err != io.EOF {
	// 		panic(err)
	// 	}

	// 	if n == 0 {
	// 		break
	// 	}

	// 	if _, err := w.Write(buf[:n]); err != nil {
	// 		panic(err)
	// 	}
	// }

	// if err = w.Flush(); err != nil {
	// 	panic(err)
	// }
}
