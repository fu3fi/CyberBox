package main

import "math/rand"
import "fmt"

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


func main() {
	size := 100000
	depth := 8

	seq := GenerateSequence(size)

	Eversion(&seq, depth)

	fmt.Println(SubSequence(seq, 6626))
}

