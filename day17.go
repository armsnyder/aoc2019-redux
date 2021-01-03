package main

import (
	"io"
	"strconv"
	"strings"
)

var _ = declareDay(17, func(part2 bool, inputReader io.Reader) interface{} {
	if part2 {
		return day17Part2(inputReader)
	}
	return day17Part1(inputReader)
})

func day17Part1(_ io.Reader) interface{} {
	return nil
}

func day17Part2(inputReader io.Reader) interface{} {
	//buffered, _ := ioutil.ReadAll(inputReader)
	//cpu := &computer{}
	//cpu.init(bytes.NewReader(buffered))
	//maze := day17ReadMaze(cpu)
	//path := day17GetPath(maze)

	// I gave up trying to find a compression algorithm.
	main, a, b, c := "A,B,B,C,C,A,A,B,B,C\n", "L,12,R,4,R,4\n", "R,12,R,4,L,12\n", "R,12,R,4,L,6,L,8,L,8\n"

	cpu := &computer{}
	cpu.init(inputReader)
	cpu.staticMemory[0] = 2
	go cpu.run()
	go func() {
		cpu.inputString(main)
		cpu.inputString(a)
		cpu.inputString(b)
		cpu.inputString(c)
		cpu.inputString("n\n")
	}()

	outputs := cpu.allOutputs()
	return outputs[len(outputs)-1]
}

//func day17ReadMaze(computer *computer) string {
//	sb := &strings.Builder{}
//	go computer.run()
//	for {
//		select {
//		case <-computer.done:
//			return strings.TrimSpace(sb.String())
//		case output := <-computer.output:
//			sb.WriteRune(rune(output))
//		}
//	}
//}

func day17GetPath(maze string) []string {
	rows := strings.Split(maze, "\n")

	type coord struct{ x, y int }

	findStart := func() coord {
		for y, row := range rows {
			for x, ch := range row {
				if ch == '^' {
					return coord{x, y}
				}
			}
		}
		panic("no start")
	}

	curCoord, curDir := findStart(), 0

	isLegalCoord := func(coord coord) bool {
		return coord.x >= 0 && coord.x < len(rows[0]) && coord.y >= 0 && coord.y < len(rows)
	}

	addCoords := func(a, b coord) coord {
		return coord{a.x + b.x, a.y + b.y}
	}

	absDirToRelCoord := map[int]coord{
		0: {y: -1},
		1: {x: 1},
		2: {y: 1},
		3: {x: -1},
	}

	isPathAhead := func(relDir int) bool {
		absDir := (curDir + relDir + 4) % 4
		relCoord := absDirToRelCoord[absDir]
		absCoord := addCoords(curCoord, relCoord)
		return isLegalCoord(absCoord) && rows[absCoord.y][absCoord.x] == '#'
	}

	var result []string

	reorient := func() bool {
		switch {
		case isPathAhead(1):
			curDir = (curDir + 1) % 4
			result = append(result, "R")
		case isPathAhead(-1):
			curDir = (curDir + 3) % 4
			result = append(result, "L")
		default:
			return false
		}
		return true
	}

	move := func() {
		moved := 0
		for isPathAhead(0) {
			moved++
			curCoord = addCoords(curCoord, absDirToRelCoord[curDir])
		}
		if moved > 0 {
			result = append(result, strconv.Itoa(moved))
		}
	}

	for reorient() {
		move()
	}

	return result
}

//func day17SplitPath(path []string) (main, a, b, c string) {
//accumLen := func() []int {
//	result := make([]int, len(path))
//	lenSoFar := 0
//	for i, seg := range path {
//		result[i] = lenSoFar
//		lenSoFar += len(seg)
//	}
//	return result
//}()

//var solve func(annotatedPath day17AnnotatedSlice) day17AnnotatedSlice
//solve = func(annotatedPath day17AnnotatedSlice) day17AnnotatedSlice {
//	var annotationValue string
//	switch annotatedPath.timesAnnotated {
//	case 0:
//		annotationValue = "A"
//	case 1:
//		annotationValue = "B"
//	case 2:
//		annotationValue = "C"
//	default:
//		return annotatedPath
//	}
//
//	slicesToAnnotate := annotatedPath.nonAnnotatedSlices()
//
//	pathLen := len(slicesToAnnotate[0])
//	if pathLen > 11 {
//		pathLen = 11
//	}
//
//	for ; pathLen > 3; pathLen-- {
//		annotatedPathCopy := annotatedPath.copy()
//		annotatedPathCopy.annotate(slicesToAnnotate[0][:pathLen], annotationValue)
//		resultCandidate := solve(annotatedPathCopy)
//		if resultCandidate.timesAnnotated == 3 && len(resultCandidate.nonAnnotatedSlices()) == 0 {
//			return resultCandidate
//		}
//	}
//
//	return day17AnnotatedSlice{}
//}
//
//routine := func(value []string) string {
//	return strings.Join(value, ",") + "\n"
//}
//
//annotatedPath := solve(day17AnnotatedSlice{slice: path})
//
//main = routine(annotatedPath.annotationValues())
//a = routine(annotatedPath.subSliceFromAnnotation("A"))
//b = routine(annotatedPath.subSliceFromAnnotation("B"))
//c = routine(annotatedPath.subSliceFromAnnotation("C"))
//
//return main, a, b, c

//for aLen := 11; aLen > 0; aLen-- {
//	if accumLen[aLen] > 11 || aLen > len(path) {
//		continue
//	}
//	aCandidate := path[:aLen]
//	for bLen := 11; bLen > 0; bLen-- {
//		if accumLen[aLen+bLen]-accumLen[aLen] > 11 || aLen+bLen > len(path) {
//			continue
//		}
//		bCandidate := path[aLen : aLen+bLen]
//		for cLen := 11; cLen > 0; cLen-- {
//			if accumLen[aLen+bLen+cLen]-accumLen[aLen+bLen] > 11 || aLen+bLen+cLen > len(path) {
//				continue
//			}
//			cCandidate := path[aLen+bLen : aLen+bLen+cLen]
//		}
//	}
//}
//}

//type day17Annotation struct {
//	start int
//	end   int
//	value string
//}
//
//type day17Annotations []day17Annotation
//
//func (a day17Annotations) Len() int {
//	return len(a)
//}
//
//func (a day17Annotations) Less(i, j int) bool {
//	return a[i].start < a[j].start
//}
//
//func (a day17Annotations) Swap(i, j int) {
//	a[i], a[j] = a[j], a[i]
//}

// day17AnnotatedSlice wraps a slice of strings with some metadata that can annotate ranges of the
// slice. Ranges that carry the same annotation value are identical. It does not assume anything
// else about the slice contents or meaning of annotations.
//type day17AnnotatedSlice struct {
//	slice          []string
//	annotations    day17Annotations // always sorted
//	timesAnnotated int
//}

// copy returns an identical day17AnnotatedSlice with copied annotations. The slice is assumed to be
// immutable and is shared by the copy.
//func (a day17AnnotatedSlice) copy() day17AnnotatedSlice {
//	result := day17AnnotatedSlice{
//		slice:          a.slice,
//		annotations:    make(day17Annotations, len(a.annotations)),
//		timesAnnotated: a.timesAnnotated,
//	}
//	copy(result.annotations, a.annotations)
//	return result
//}

// annotate applies the annotation value anywhere in the original slice that the specified pattern
// is found. It will not annotate segments of the original slice that are already annotated.
//func (a *day17AnnotatedSlice) annotate(pattern []string, value string) {
//	var sliceI, annotationI int
//	originalAnnotationsLen := len(a.annotations)
//
//	for sliceI < len(a.slice)-len(pattern) {
//		// Skip over segments of the slice that are already annotated.
//		if annotationI < originalAnnotationsLen && sliceI+len(pattern) > a.annotations[annotationI].start {
//			sliceI = a.annotations[annotationI].end
//			annotationI++
//			continue
//		}
//
//		if func() bool {
//			for patternI := 0; patternI < len(pattern); patternI++ {
//				if a.slice[sliceI+patternI] != pattern[patternI] {
//					return false
//				}
//			}
//			return true
//		}() {
//			a.annotations = append(a.annotations, day17Annotation{
//				start: sliceI,
//				end:   sliceI + len(pattern),
//				value: value,
//			})
//			sliceI += len(pattern)
//		} else {
//			sliceI++
//		}
//	}
//
//	sort.Sort(a.annotations)
//	a.timesAnnotated++
//}

//func (a day17AnnotatedSlice) nonAnnotatedSlices() [][]string {
//	var result [][]string
//	start := 0
//	for _, annotation := range a.annotations {
//		if annotation.start != start {
//			result = append(result, a.slice[start:annotation.start])
//		}
//		start = annotation.end
//	}
//	if start < len(a.slice) {
//		result = append(result, a.slice[start:])
//	}
//	return result
//}

//func (a day17AnnotatedSlice) annotationValues() []string {
//	result := make([]string, len(a.annotations))
//	for i, annotation := range a.annotations {
//		result[i] = annotation.value
//	}
//	return result
//}

//func (a day17AnnotatedSlice) subSliceFromAnnotation(annotationValue string) []string {
//	for _, annotation := range a.annotations {
//		if annotation.value == annotationValue {
//			return a.slice[annotation.start:annotation.end]
//		}
//	}
//	return nil
//}
