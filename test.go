package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Products struct {
	Id    string `json:"_id"`
	Price string `json:"price"`
}

func (p Products) Floatize() float64 {
	s := strings.TrimPrefix(p.Price, "$")
	f, _ := strconv.ParseFloat(s, 64)
	return f
}

type ByPrice []Products

func (p ByPrice) Len() int { return len(p) }
func (p ByPrice) Less(i, j int) bool {
	return p[i].Floatize() < p[j].Floatize()
}
func (p ByPrice) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

func main() {
	file, err := os.Open("the-scrambler.json")
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return
	}
	decoder := json.NewDecoder(bufio.NewReader(file))
	x := make([]Products, 0)
	err = decoder.Decode(&x)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	}
	sort.Sort(ByPrice(x))
	spew.Dump(x)
}
