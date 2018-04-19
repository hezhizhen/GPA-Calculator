package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type Course struct {
	Name   string  `json:"name"`
	Credit float64 `json:"credit"`
	Grade  int     `json:"grade"`
	Kind   string  `json:"kind"`
	// general: 通修
	// public: 平台
	// core: 核心
}

type Term struct {
	Name    string   `json:"name"`
	Courses []Course `json:"courses"`
}

var courses = func() []Term {
	f, err := os.Open("courses.json")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	bs, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	result := []Term{}
	if err := json.Unmarshal(bs, &result); err != nil {
		panic(err)
	}
	return result
}()

var terms = []string{
	"freshman",
	"sophomore",
	"junior",
	"senior",
}

var kindsCombination = [][]string{
	{"general", "public", "core"},
	{"public", "core"},
	{"core"},
}

func inSlice(one string, ones []string) bool {
	for _, o := range ones {
		if one == o {
			return true
		}
	}
	return false
}

func njuAverage(termPrefix string, kinds []string) float64 {
	total := 0.0
	credits := 0.0
	for _, term := range courses {
		if termPrefix != "" && !strings.HasPrefix(term.Name, termPrefix) {
			continue
		}
		for _, course := range term.Courses {
			if !inSlice(course.Kind, kinds) {
				continue
			}
			credits += course.Credit
			total += course.Credit * float64(course.Grade)
		}
	}
	return total / credits
}

func main() {
	for _, kinds := range kindsCombination {
		for _, term := range terms {
			fmt.Println(term, kinds, njuAverage(term, kinds)/20)
		}
		fmt.Println("overall", kinds, njuAverage("", kinds)/20)
	}
}
