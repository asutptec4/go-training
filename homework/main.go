package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"sort"
	"strconv"
	"time"
)

func main() {
	json := `{"name":"Walter White","age":"52","nicknames":["Heisenberg", "Mr. lambert", "Mr. mayhem", "Walt jackson"],"achievements":{"49":"survive cancer","51":"make a fortune for the family","50":"build a laboratory"}}`
	var p Person
	if err := p.ParseFromJSON([]byte(json)); err != nil {
		fmt.Printf("Fail to parse JSON.\n")
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

// ParseFromJSON build Person object from JSON
func (p *Person) ParseFromJSON(data []byte) error {
	err := json.Unmarshal(data, p)
	if err != nil {
		return err
	}
	return nil
}

// PrintAchivments prints in console all Person's achievements as increase sequence
func (p *Person) PrintAchivments() {
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
func (p *Person) SayMyName() {
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(len(p.Nicknames))
	fmt.Printf("SayMyName - %s\n", p.Nicknames[r])
}
