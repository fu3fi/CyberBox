package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

func GenerateSequence(n int) []int {
	numbers := make([]int, n)
	usedNumbers := make(map[int]bool)
	newNumbers := make([]int, n)
	for i := range numbers {
		numbers[i] = i
	}
	rand.Seed(1252152112414)
	counter := 0

	startIndex := 0
	for {

		if counter == n-1 {
			newNumbers[startIndex] = 0
			break
		}

		nextNumber := numbers[rand.Intn(len(numbers))]

		if _, check := usedNumbers[nextNumber]; check == true || nextNumber == 0 {
			continue
		}

		usedNumbers[nextNumber] = true
		newNumbers[startIndex] = nextNumber
		startIndex = nextNumber
		counter++
	}

	return newNumbers
}

func Eversion(sequence *[]int, depth int) {

	next := 0
	memory := 0
	counter := 0
	subCounter := 1

	for {

		if counter == len(*sequence) {
			break
		}

		if subCounter == depth {
			memory, (*sequence)[next] = (*sequence)[next], memory
			next = memory
			subCounter = 0
		} else {
			next = (*sequence)[next]
		}

		counter++
		subCounter++
	}

}

func SubSequence(seq []int, startI int) []int {
	sS := make([]int, 0, 8)
	startV := startI
	sS = append(sS, startI)

	for {

		startI = seq[startI]
		sS = append(sS, startI)

		if startV == startI {
			break
		}
	}

	return sS[:len(sS)-1]
}

func start(w http.ResponseWriter, req *http.Request) {
	id := req.URL.Query().Get("id")

	fmt.Println(id)
	fmt.Println(SESSIONS)

	if _, check := SESSIONS[id]; check == false {
		newId := randSeq(16)
		SESSIONS[newId] = append(SubSequence(SEQ, rand.Intn(SIZE)), []int{0}...)
		// SESSIONS[randSeq(16)] = SubSequence(SEQ, rand.Intn(SIZE))
		fmt.Fprintf(w, newId)
		return
	}

	counter := SESSIONS[id][len(SESSIONS[id])-1]

	if counter == 0 {
		fmt.Fprintf(w, strconv.Itoa(SESSIONS[id][counter]))
		SESSIONS[id][len(SESSIONS[id])-1] = counter + 1
		return
	} else {
		nextValue := req.URL.Query().Get("next")

		if strconv.Itoa(SESSIONS[id][counter]) == nextValue {

			if counter+2 == len(SESSIONS[id]) {
				fmt.Fprintf(w, "SUCCESS")
				delete(SESSIONS, id)
				return
			}

			fmt.Fprintf(w, "OK")
			SESSIONS[id][len(SESSIONS[id])-1] = counter + 1
		} else {
			fmt.Fprintf(w, "ERROR")
			delete(SESSIONS, id)
		}
	}

}

func randSeq(n int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func main() {

	Eversion(&SEQ, DEPTH)
	// fmt.Println(SubSequence(SEQ, 6626))
	http.HandleFunc("/start", start)
	fmt.Println("start on " + IP + ":" + PORT)
	err := http.ListenAndServe(IP+":"+PORT, nil)

	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println("start on " + IP + ":" + PORT)
}
