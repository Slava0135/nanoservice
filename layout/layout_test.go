package layout

import (
	"reflect"
	"slava0135/nanoservice/rules"
	"testing"
)

func TestLinkedCornerSquares(t *testing.T) {
	type args struct {
		x uint
		y uint
	}
	tests := []struct {
		name string
		args args
		want []Point
	}{
		{"zero corner", args{0, 0}, []Point{{1, 1}}},
		{"n corner", args{rules.N - 1, rules.N - 1}, []Point{{8, 8}}},
		{"side y", args{1, 0}, []Point{{2, 1}, {0, 1}}},
		{"side x", args{0, 1}, []Point{{1, 0}, {1, 2}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LinkedCornerSquares(tt.args.x, tt.args.y); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LinkedCornerSquares() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinkedSideSquares(t *testing.T) {
	type args struct {
		x uint
		y uint
	}
	tests := []struct {
		name string
		args args
		want []Point
	}{
		{"zero corner", args{0, 0}, []Point{{1, 0}, {0, 1}}},
		{"n corner", args{rules.N - 1, rules.N - 1}, []Point{{8, 9}, {9, 8}}},
		{"side y", args{1, 0}, []Point{{0, 0}, {2, 0}, {1, 1}}},
		{"side x", args{0, 1}, []Point{{0, 0}, {1, 1}, {0, 2}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LinkedSideSquares(tt.args.x, tt.args.y); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LinkedCornerSquares() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinkedSquares(t *testing.T) {
	type args struct {
		x uint
		y uint
	}
	tests := []struct {
		name string
		args args
		want []Point
	}{
		{"zero corner", args{0, 0}, []Point{{1, 1}, {1, 0}, {0, 1}}},
		{"n corner", args{rules.N - 1, rules.N - 1}, []Point{{8, 8}, {8, 9}, {9, 8}}},
		{"side y", args{1, 0}, []Point{{2, 1}, {0, 1}, {0, 0}, {2, 0}, {1, 1}}},
		{"side x", args{0, 1}, []Point{{1, 0}, {1, 2}, {0, 0}, {1, 1}, {0, 2}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LinkedSquares(tt.args.x, tt.args.y); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LinkedCornerSquares() = %v, want %v", got, tt.want)
			}
		})
	}
}


