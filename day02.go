package main

import (
	"io"
)

var _ = declareDay(2, func(part2 bool, inputReader io.Reader) interface{} {
	if part2 {
		return day02Part2(inputReader)
	}
	return day02Part1(inputReader)
})

func day02Part1(inputReader io.Reader) interface{} {
	var computer computer
	computer.init(inputReader)
	computer.run()
	return computer.readMemory(0)
}

func day02Part2(_ io.Reader) interface{} {
	return nil
}
