package bubble

import (
	"reflect"
	"testing"
)

func Test_bubbleSort(t *testing.T) {
	type args struct {
		input []int
	}
	tests := []struct {
		name  string
		input []int
		args  args
		want  []int
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
			if got := bubbleSort(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("bubbleSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
