package merge

import (
	"reflect"
	"testing"
)

func Test_mergeSort(t *testing.T) {
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
			name: "Two elements",
			args: args{input: []int{2, 1}},
			want: []int{1, 2},
		},
		{
			name: "Sorted list",
			args: args{input: []int{1, 2, 3, 4, 5}},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "Reverse sorted list",
			args: args{input: []int{5, 4, 3, 2, 1}},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "Duplicates",
			args: args{input: []int{2, 3, 3, 1, 1}},
			want: []int{1, 1, 2, 3, 3},
		},
		{
			name: "Negative numbers",
			args: args{input: []int{-1, -3, -2, 1, 4}},
			want: []int{-3, -2, -1, 1, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeSort(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
