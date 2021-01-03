package main

import (
	"bufio"
	"bytes"
	"io"
	"math"
	"strconv"
)

type computer struct {
	staticMemory  []int64
	dynamicMemory map[int]int64
	ip            int
	bp            int
	input         chan int64
	output        chan int64
	done          chan bool
}

func (c *computer) init(inputReader io.Reader) {
	c.staticMemory = parseCommaSeparatedInts(inputReader)
	c.dynamicMemory = make(map[int]int64)
	c.input = make(chan int64)
	c.output = make(chan int64)
	c.done = make(chan bool)
}

func (c *computer) inputString(s string) {
	for _, ch := range s {
		c.input <- int64(ch)
	}
}

func (c *computer) singleInOut(input int64) (output int64) {
	c.input <- input
	return <-c.output
}

func (c *computer) allOutputs() (output []int64) {
	for {
		select {
		case thisOutput := <-c.output:
			output = append(output, thisOutput)
		case <-c.done:
			return output
		}
	}
}

func (c *computer) run() {
	defer close(c.done)

	for {
		opcode := c.readMemory(c.ip) % 100
		readParam := c.paramReader()
		writeParam := c.paramWriter()
		switch opcode {
		case 1:
			writeParam(3, readParam(1)+readParam(2))
			c.ip += 4
		case 2:
			writeParam(3, readParam(1)*readParam(2))
			c.ip += 4
		case 3:
			writeParam(1, <-c.input)
			c.ip += 2
		case 4:
			c.output <- readParam(1)
			c.ip += 2
		case 5:
			if readParam(1) != 0 {
				c.ip = int(readParam(2))
			} else {
				c.ip += 3
			}
		case 6:
			if readParam(1) == 0 {
				c.ip = int(readParam(2))
			} else {
				c.ip += 3
			}
		case 7:
			var value int64
			if readParam(1) < readParam(2) {
				value = 1
			}
			writeParam(3, value)
			c.ip += 4
		case 8:
			var value int64
			if readParam(1) == readParam(2) {
				value = 1
			}
			writeParam(3, value)
			c.ip += 4
		case 9:
			c.bp += int(readParam(1))
			c.ip += 2
		case 99:
			return
		default:
			panic(opcode)
		}
	}
}

func (c *computer) paramReader() func(int) int64 {
	modes := int(c.readMemory(c.ip)) / 100
	return func(arg int) int64 {
		mode := (modes / int(math.Pow10(arg-1))) % 10
		argv := c.readMemory(c.ip + arg)
		switch mode {
		case 0:
			return c.readMemory(int(argv))
		case 1:
			return argv
		case 2:
			return c.readMemory(c.bp + int(argv))
		default:
			panic(mode)
		}
	}
}

func (c *computer) paramWriter() func(arg int, value int64) {
	modes := int(c.readMemory(c.ip)) / 100
	return func(arg int, value int64) {
		mode := (modes / int(math.Pow10(arg-1))) % 10
		argv := int(c.readMemory(c.ip + arg))
		switch mode {
		case 0:
			c.writeMemory(argv, value)
		case 2:
			c.writeMemory(c.bp+argv, value)
		default:
			panic(mode)
		}
	}
}

func (c *computer) readMemory(i int) int64 {
	if i < len(c.staticMemory) {
		return c.staticMemory[i]
	}
	return c.dynamicMemory[i]
}

func (c *computer) writeMemory(i int, v int64) {
	if i < len(c.staticMemory) {
		c.staticMemory[i] = v
	}
	c.dynamicMemory[i] = v
}

func parseCommaSeparatedInts(inputReader io.Reader) []int64 {
	var memory []int64
	scanner := bufio.NewScanner(inputReader)
	scanner.Split(splitComma)
	for scanner.Scan() {
		v, err := strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			panic(err)
		}
		memory = append(memory, v)
	}
	return memory
}

func splitComma(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	if i := bytes.IndexByte(data, ','); i >= 0 {
		return i + 1, bytes.TrimSpace(data[:i]), nil
	}
	if atEOF {
		return len(data), bytes.TrimSpace(data), nil
	}
	return 0, nil, nil
}
