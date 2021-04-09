package scripts

import (
	"io/ioutil"
	"log"
	"strings"
)

//GetLists loads lists from filesystem
func GetLists(path string) []string {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	return strings.Split(string(data), "\n")
}
