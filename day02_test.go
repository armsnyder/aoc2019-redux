package main

import (
	"testing"
)

func TestDay02Part1(t *testing.T) {
	runDayTests(t, 2, []dayTest{
		{
			input: `1,9,10,3,2,3,11,0,99,30,40,50`,
			want:  int64(3500),
		},
		{
			input: `1,0,0,0,99`,
			want:  int64(2),
		},
		{
			input: `2,3,0,3,99`,
			want:  int64(2),
		},
		{
			input: `2,4,4,5,99,0`,
			want:  int64(2),
		},
		{
			input: `1,1,1,4,99,5,6,0,99`,
			want:  int64(30),
		},
	})
}
