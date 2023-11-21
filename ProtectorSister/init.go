package main

import (
	"math/rand"
	"time"
)

var PORT string = "8678"
var IP string = ""
var SESSIONS map[string][]int = make(map[string][]int)

var SIZE int = 100000
var DEPTH int = 8
var SEQ []int = GenerateSequence(SIZE)

func init() {
	rand.Seed(time.Now().UnixNano())
}
