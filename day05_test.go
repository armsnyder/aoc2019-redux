package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestDay05Modes(t *testing.T) {
	tests := []string{"1002,4,3,4,33", "1101,100,-1,4,0"}
	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			var computer computer
			computer.init(strings.NewReader(test))
			computer.run()
		})
	}
}

func TestDay05IO(t *testing.T) {
	var computer computer
	computer.init(strings.NewReader("3,0,4,0,99"))
	go computer.run()
	if computer.singleInOut(7) != 7 {
		t.Fail()
	}
}

func TestDay05Comparators(t *testing.T) {
	tests := []struct {
		program     string
		inputOutput map[int64]int64
	}{
		{
			program:     "3,9,8,9,10,9,4,9,99,-1,8",
			inputOutput: map[int64]int64{7: 0, 8: 1, 9: 0},
		},
		{
			program:     "3,9,7,9,10,9,4,9,99,-1,8",
			inputOutput: map[int64]int64{7: 1, 8: 0, 9: 0},
		},
		{
			program:     "3,3,1108,-1,8,3,4,3,99",
			inputOutput: map[int64]int64{7: 0, 8: 1, 9: 0},
		},
		{
			program:     "3,3,1107,-1,8,3,4,3,99",
			inputOutput: map[int64]int64{7: 1, 8: 0, 9: 0},
		},
	}
	for _, test := range tests {
		t.Run(test.program, func(t *testing.T) {
			for input, output := range test.inputOutput {
				t.Run(fmt.Sprintf("%d->%d", input, output), func(t *testing.T) {
					var computer computer
					computer.init(strings.NewReader(test.program))
					go computer.run()
					if computer.singleInOut(input) != output {
						t.Fail()
					}
				})
			}
		})
	}
}

func TestDay05Jump(t *testing.T) {
	tests := []string{
		"3,12,6,12,15,1,13,14,13,4,13,99,-1,0,1,9",
		"3,3,1105,-1,9,1101,0,0,12,4,12,99,1",
	}
	for _, test := range tests {
		t.Run(test, func(t *testing.T) {
			for input, output := range map[int64]int64{-1: 1, 0: 0, 1: 1} {
				t.Run(fmt.Sprintf("%d->%d", input, output), func(t *testing.T) {
					var computer computer
					computer.init(strings.NewReader(test))
					go computer.run()
					if computer.singleInOut(input) != output {
						t.Fail()
					}
				})
			}
		})
	}
}
