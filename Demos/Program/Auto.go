package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	f, x := os.Open("Files.txt")
	if x != nil {
		log.Fatal(x)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		FileOut := scanner.Text()
		f, x = os.Create(FileOut + ".skyline")
		if x != nil {
			log.Fatal(x)
		}
		defer f.Close()
		fmt.Println("[+] Created directory -> ", FileOut)
	}
}
