package main

import (
	"io"
)

var _ = declareDay(5, func(part2 bool, inputReader io.Reader) interface{} {
	if part2 {
		return day05(inputReader, 5)
	}
	return day05(inputReader, 5)
})

func day05(inputReader io.Reader, systemID int64) int64 {
	var computer computer
	computer.init(inputReader)
	go computer.run()
	computer.input <- systemID
	for output := range computer.output {
		if output > 0 {
			return output
		}
	}
	panic("output closed")
}
