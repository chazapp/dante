package tigbra

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
)

type Doc struct {
	Category string `json:"category"`
	Theme    string `json:"theme"`
	Quote    string `json:"quote"`
	Page     string `json:"page"`
}

func ParseDataset(path string) ([]Doc, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	err = file.Close()
	if err != nil {
		return nil, err
	}
	var cat = ""
	var theme = ""
	var output = make([]Doc, 0)
	for _, item := range lines {
		if _, err := strconv.Atoi(string(item[0])); err == nil {
			cat = item
			continue
		}
		if item[0] == '-' {
			r, _ := regexp.Compile("p.([0-9]{1,3})")
			if r.MatchString(item) {
				page := r.FindString(item)
				quote := r.ReplaceAllString(item, "")
				var out = Doc{Page: page, Quote: quote, Category: cat, Theme: theme}
				output = append(output, out)
			}
		} else {
			theme = item
			continue
		}
	}
	return output, nil
}
