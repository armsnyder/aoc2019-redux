package main

import (
	"io"
)

var _ = declareDay(9, func(part2 bool, inputReader io.Reader) interface{} {
	if part2 {
		return day09Part2(inputReader)
	}
	return day09Part1(inputReader)
})

func day09Part1(inputReader io.Reader) interface{} {
	var computer computer
	computer.init(inputReader)
	go computer.run()
	return computer.singleInOut(1)
}

func day09Part2(inputReader io.Reader) interface{} {
	var computer computer
	computer.init(inputReader)
	go computer.run()
	return computer.singleInOut(2)
}
