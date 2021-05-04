package scripts

import (
	"bufio"
	"bytes"
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
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

	for _, hostname := range res {
		hostShort := ""
		name := strings.Split(hostname, ".")[0]
		if len(name) == 0 {
			continue
		}
		if len(name) == 1 {
			hostShort = "_" + name
		} else {
			hostShort = hostname[0:2]
		}
		addToFile(fmt.Sprintf("%s/%s.txt", filename, hostShort), fmt.Sprintf("%s\n", hostname))
	}
}

func addToFile(filename string, text string) {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	if _, err = f.WriteString(text); err != nil {
		panic(err)
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
			if checkDomain(f[0]) {
				data = append(data, string(f[0]))
			}
		case 2:
			if checkDomain(f[1]) {
				data = append(data, string(f[1]))
			}
		default:
			continue
		}
	}

	return data
}

var r = regexp.MustCompile("^(?:[a-z0-9](?:[a-z0-9-]{0,61}[a-z0-9])?.)+[a-z0-9][a-z0-9-]{0,61}[a-z0-9]$")

func checkDomain(host []byte) bool {
	return r.Match([]byte(host))
}
