package scripts

import (
	"log"
	"os"
	"strings"
)

//GetLists loads lists from filesystem
func GetLists(path string) []string {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	return strings.Split(string(data), "\n")
}
