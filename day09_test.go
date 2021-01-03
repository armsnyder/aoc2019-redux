package main

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func TestDay09CopySelf(t *testing.T) {
	var computer computer
	computer.init(strings.NewReader("109,1,204,-1,1001,100,1,100,1008,100,16,101,1006,101,0,99"))
	go computer.run()
	output := computer.allOutputs()
	wantOutput := []int64{109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99}
	if !reflect.DeepEqual(output, wantOutput) {
		t.Fail()
	}
}

func TestDay09SixteenDigitNumber(t *testing.T) {
	var computer computer
	computer.init(strings.NewReader("1102,34915192,34915192,7,4,7,99,0"))
	go computer.run()
	output := computer.allOutputs()[0]
	if len(fmt.Sprintf("%d", output)) != 16 {
		t.Fail()
	}
}

func TestDay09LargeNumber(t *testing.T) {
	var computer computer
	computer.init(strings.NewReader("104,1125899906842624,99"))
	go computer.run()
	output := computer.allOutputs()[0]
	if output != 1125899906842624 {
		t.Fail()
	}
}
