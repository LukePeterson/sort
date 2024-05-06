package insertion

import (
	"reflect"
	"testing"
)

func Test_insertionSort(t *testing.T) {
	type args struct {
		input []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Empty array",
			args: args{input: []int{}},
			want: []int{},
		},
		{
			name: "Single element",
			args: args{input: []int{4}},
			want: []int{4},
		},
		{
			name: "Small test case",
			args: args{input: []int{5, 9, 1, 4, 10}},
			want: []int{1, 4, 5, 9, 10},
		},
		{
			name: "Already sorted",
			args: args{input: []int{1, 2, 3, 4, 5}},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "Reverse sorted",
			args: args{input: []int{5, 4, 3, 2, 1}},
			want: []int{1, 2, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			insertionSort(tt.args.input)
			if !reflect.DeepEqual(tt.args.input, tt.want) {
				t.Errorf("\ninsertionSort() = %v, want %v", tt.args.input, tt.want)
			}
		})
	}
}
