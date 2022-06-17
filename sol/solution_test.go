package sol

import (
	"math"
	"reflect"
	"testing"
)

func BenchmarkTest(b *testing.B) {
	rooms := [][]int{
		{math.MaxInt32, -1, 0, math.MaxInt32},
		{math.MaxInt32, math.MaxInt32, math.MaxInt32, -1},
		{math.MaxInt32, -1, math.MaxInt32, -1},
		{0, -1, math.MaxInt32, math.MaxInt32},
	}
	for idx := 0; idx < b.N; idx++ {
		WallsAndGates(rooms)
	}
}
func TestWallsAndGates(t *testing.T) {
	type args struct {
		rooms [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "Example1",
			args: args{
				rooms: [][]int{
					{math.MaxInt32, -1, 0, math.MaxInt32},
					{math.MaxInt32, math.MaxInt32, math.MaxInt32, -1},
					{math.MaxInt32, -1, math.MaxInt32, -1},
					{0, -1, math.MaxInt32, math.MaxInt32},
				},
			},
			want: [][]int{
				{3, -1, 0, 1},
				{2, 2, 1, -1},
				{1, -1, 2, -1},
				{0, -1, 3, 4},
			},
		},
		{
			name: "Example2",
			args: args{
				rooms: [][]int{
					{0, -1},
					{math.MaxInt32, math.MaxInt32},
				},
			},
			want: [][]int{
				{0, -1},
				{1, 2},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			WallsAndGates(tt.args.rooms)
			if !reflect.DeepEqual(tt.args.rooms, tt.want) {
				t.Errorf("WallsAndGates() = %v, want %v", tt.args.rooms, tt.want)
			}

		})
	}
}
