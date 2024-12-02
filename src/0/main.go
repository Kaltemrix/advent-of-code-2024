package main

import "os"

func main() {
	file, err := os.Open("data.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
}
