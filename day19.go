package main

import (
	"io"
)

var _ = declareDay(19, func(part2 bool, inputReader io.Reader) interface{} {
	if part2 {
		return day19Part2(inputReader)
	}
	return day19Part1(inputReader)
})

func day19Part1(_ io.Reader) interface{} {
	return nil
}

func day19Part2(inputReader io.Reader) interface{} {
	var computer computer
	computer.init(inputReader)
	staticBackup := make([]int64, len(computer.staticMemory))
	copy(staticBackup, computer.staticMemory)
	go func() {
		for {
			computer.run()
			computer.ip = 0
			copy(computer.staticMemory, staticBackup)
			computer.dynamicMemory = make(map[int]int64)
			computer.done = make(chan bool)
		}
	}()
	test := func(x, y int) bool {
		computer.input <- int64(x)
		computer.input <- int64(y)
		return <-computer.output == 1
	}
	originX := 0
	for originY := 0; ; originY++ {
		for !test(originX, originY+99) {
			originX++
		}
		if test(originX+99, originY) {
			return originX*10000 + originY
		}
	}
}
