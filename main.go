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

	// I can do the following though:
	resultI := get2()
	result = resultI.(int)
	log.Println("result", result)
	// but is it possible to encapsulate the line
	// result = resultI.(int)
	// inside method "get"

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

func get2() interface{} {
	aMtx.Lock()
	defer aMtx.Unlock()
	return a
}

func get3(c interface{}) {
	aMtx.Lock()
	defer aMtx.Unlock()
//	t := reflect.TypeOf(c)
//	c = a.(t) // looks like it's impossible https://stackoverflow.com/q/27971895/1024794
}

func get4(c interface{}) {
	aMtx.Lock()
	defer aMtx.Unlock()
	//json.Unmarshal() // kind of this?
}

func get5(c interface{}) {
	aMtx.Lock()
	defer aMtx.Unlock()
	//switch v := c.(type) {
	//case int:
	//*c = v
	//}
}