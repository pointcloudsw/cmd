package main

import (
	"fmt"
	"log"

	"github.com/pointcloudsw/util"
)

func main() {
	// Read entire file
	content, err := util.ReadFile("/etc/group")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("File Content:\n", content)

	// Read file line by line
	lines, err := util.ReadFileLines("/etc/group")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("\nFile Lines:")
	for _, line := range lines {
		fmt.Println(line)
	}
}
