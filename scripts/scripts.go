package scripts

import (
	"bufio"
	"bytes"
	"crypto/tls"
	"log"
	"net/http"
	"os"
	"strings"
)

func init() {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
}

func ReadLists(filename string, lists []string) {
	var res []string
	for _, list := range lists {
		res = appendCategory(res, readFromUrl(list))
	}

	err := os.WriteFile(filename, []byte(strings.Join(res, "\n")), 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func appendCategory(a []string, b []string) []string {

	check := make(map[string]int)
	d := append(a, b...)
	res := make([]string, 0)
	for _, val := range d {
		check[val] = 1
	}

	for letter, _ := range check {
		res = append(res, letter)
	}

	return res
}

func readFromUrl(url string) []string {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	scanner := bufio.NewScanner(resp.Body)
	data := []string{}

	for scanner.Scan() {
		line := scanner.Bytes()
		if i := bytes.Index(line, []byte{'#'}); i >= 0 {
			line = line[0:i]
		}
		f := bytes.Fields(line)

		switch len(f) {
		case 1:
			data = append(data, string(f[0]))
		case 2:
			data = append(data, string(f[1]))
		default:
			continue
		}
	}

	return data
}
