package main

import (
	"log"
	"sync"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Llongfile)
	set(1)
	var result int
	get(&result) // I expect that result equals "1" after calling "get" function but it's "0"
	log.Println(result) // it prints "0" here but I would like to fix code in way so it should print "1" here
	// set function should accept arbitrary type
}

var (
	a interface{}
	aMtx sync.Mutex
)

func set(b interface{}) {
	aMtx.Lock()
	defer aMtx.Unlock()
	a = b
}

func get(c interface{}) {
	aMtx.Lock()
	defer aMtx.Unlock()
	c = a
}

