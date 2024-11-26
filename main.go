package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	Reader();
}

func Reader() string {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter a folder path containing banners")

	path, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	fmt.Println(path)

	return path
}
