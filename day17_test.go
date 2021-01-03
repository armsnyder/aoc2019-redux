package main

import (
	"reflect"
	"strings"
	"testing"
)

func Test_day17GetPath(t *testing.T) {
	maze := strings.TrimSpace(`
#######...#####
#.....#...#...#
#.....#...#...#
......#...#...#
......#...###.#
......#.....#.#
^########...#.#
......#.#...#.#
......#########
........#...#..
....#########..
....#...#......
....#...#......
....#...#......
....#####......
`)
	want := strings.Split("R,8,R,8,R,4,R,4,R,8,L,6,L,2,R,4,R,4,R,8,R,8,R,8,L,6,L,2", ",")
	if got := day17GetPath(maze); !reflect.DeepEqual(got, want) {
		t.Errorf("day17GetPath() = %v, want %v", got, want)
	}
}

//func Test_day17SplitPath(t *testing.T) {
//	path := strings.Split("R,8,R,8,R,4,R,4,R,8,L,6,L,2,R,4,R,4,R,8,R,8,R,8,L,6,L,2", ",")
//
//	wantMain := "A,B,C,B,A,C\n"
//	wantA := "R,8,R,8\n"
//	wantB := "R,4,R,4,R,8\n"
//	wantC := "L,6,L,2\n"
//
//	gotMain, gotA, gotB, gotC := day17SplitPath(path)
//
//	if gotMain != wantMain {
//		t.Errorf("day17SplitPath() gotMain = %v, want %v", gotMain, wantMain)
//	}
//	if gotA != wantA {
//		t.Errorf("day17SplitPath() gotA = %v, want %v", gotA, wantA)
//	}
//	if gotB != wantB {
//		t.Errorf("day17SplitPath() gotB = %v, want %v", gotB, wantB)
//	}
//	if gotC != wantC {
//		t.Errorf("day17SplitPath() gotC = %v, want %v", gotC, wantC)
//	}
//}
