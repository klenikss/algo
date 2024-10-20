package main

import (
	"bufio"
	"fmt"
	"os"
)

const initialSize = 8
const loadFactor = 3 / 4

type HashSet struct {
	table []string
	size  int
	count int
}

func NewHashSet() *HashSet {
	return &HashSet{
		table: make([]string, initialSize),
		size:  initialSize,
		count: 0,
	}
}

// Horner's method for hashing a string
func hash(s string, m int) int {
	h := 0
	for i := 0; i < len(s); i++ {
		h = (h*31 + int(s[i])) % m
	}
	return h
}

func (set *HashSet) resize() {
	newSize := set.size * 2
	newTable := make([]string, newSize)
	set.count = 0

	for _, value := range set.table {
		if value != "" {
			set.insertIntoTable(newTable, value)
		}
	}

	set.table = newTable
	set.size = newSize
}

func (set *HashSet) insertIntoTable(table []string, s string) {
	index := hash(s, len(table))
	i := 0
	for table[index] != "" {
		i++
		index = (index + i*i) % len(table)
	}
	table[index] = s
	set.count++
}

func (set *HashSet) Add(s string) string {
	if float64(set.count)/float64(set.size) >= loadFactor {
		set.resize()
	}
	index := hash(s, set.size)
	i := 0
	for {
		if set.table[index] == "" {
			set.insertIntoTable(set.table, s)
			return "OK"
		}
		if set.table[index] == s {
			return "FAIL"
		}
		i++
		index = (index + i*i) % set.size
	}
}

func (set *HashSet) Remove(s string) string {
	index := hash(s, set.size)
	i := 0
	for {
		if set.table[index] == "" {
			return "FAIL"
		}
		if set.table[index] == s {
			set.table[index] = "" // Удаляем элемент
			set.count--
			return "OK"
		}
		i++
		index = (index + i*i) % set.size
	}
}

func (set *HashSet) Contains(s string) string {
	index := hash(s, set.size)
	i := 0
	for {
		if set.table[index] == "" {
			return "FAIL"
		}
		if set.table[index] == s {
			return "OK"
		}
		i++
		index = (index + i*i) % set.size
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	set := NewHashSet()

	for scanner.Scan() {
		line := scanner.Text()
		var op string
		var value string
		fmt.Sscanf(line, "%s %s", &op, &value)

		switch op {
		case "+":
			fmt.Println(set.Add(value))
		case "-":
			fmt.Println(set.Remove(value))
		case "?":
			fmt.Println(set.Contains(value))
		}
	}
}
