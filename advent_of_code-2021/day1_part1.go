package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	var line string
	var lines = []int{}

	f, err := os.Open("level1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line = scanner.Text()
		lineAsInt, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}
		lines = append(lines, lineAsInt)
	}
	fmt.Println(lines)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
