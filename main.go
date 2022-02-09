package main

import "log"

func main() {
	log.SetFlags(log.LstdFlags | log.Llongfile)
	set(1)
	var result int
	get(&result) // I expect that result equals "1" after calling "get" function but it's "0"
	log.Println(result) // it prints "0" here but I would like to fix code in way so it should print "1" here
}

var (
	a interface{}

)

func set(b interface{}) {
	a = b
}

func get(c interface{}) {
	c = a
}

