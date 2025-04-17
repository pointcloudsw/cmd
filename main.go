package main

import (
	"fmt"
	"log"

	"github.com/pointcloudsw/util"
)

func main() {
	// Read entire file
	content, err := util.ReadFile("/home/dev/repos/wce/certs/certdocs-main/src/content/docs/en/resources/commands/loadCdb/.env")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("File Content:\n", content)

	// Read file line by line
	lines, err := util.ReadFileLines("/home/dev/repos/wce/certs/certdocs-main/src/content/docs/en/resources/commands/loadCdb/.env")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("\nFile Lines:")
	for _, line := range lines {
		fmt.Println(line)
	}
}
