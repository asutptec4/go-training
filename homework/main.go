package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"time"
)

func main() {
	jsonFilename := flag.String("json-file", "data.json", "a JSON file")

	file, err := os.Open(*jsonFilename)
	defer file.Close()
	if err != nil {
		fmt.Printf("Failed to open the JSON file: %s, err %v\n", *jsonFilename, err)
		return
	}

	d, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("Failed to read from file: %v\n", err)
		return
	}

	var p Person
	err = json.Unmarshal(d, &p)
	if err != nil {
		fmt.Printf("Fail to parse JSON: %v\n", err)
		return
	}

	p.SayMyName()
	p.PrintAchivments()
}

// Person type describe peron
type Person struct {
	Name         string            `json:"name,omitempty"`
	Age          int               `json:"age,string,omitempty"`
	Nicknames    []string          `json:"nicknames,omitempty"`
	Achievements map[string]string `json:"achievements,omitempty"`
}

// PrintAchivments prints in console all Person's achievements as increase sequence
func (p Person) PrintAchivments() {
	var ages []int
	for k := range p.Achievements {
		a, err := strconv.Atoi(k)
		if err != nil {
			continue
		}
		ages = append(ages, a)
	}
	sort.Ints(ages)
	for _, v := range ages {
		a := strconv.Itoa(v)
		fmt.Printf("Date - %s, Achivement - %s\n", a, p.Achievements[a])
	}
}

// SayMyName prints in console random Person's nickname
func (p Person) SayMyName() {
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(len(p.Nicknames))
	fmt.Printf("SayMyName - %s\n", p.Nicknames[r])
}
