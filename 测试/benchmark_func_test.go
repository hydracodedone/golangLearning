package funcTest

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_insertElementAtIndexThird(t *testing.T) {
	type args struct {
		newElement int
		slice      *[]int
		result     *[]int
	}
	var nilSlice []int
	var lengthLessThreeSlice = make([]int, 1)
	result1 := make([]int, 4)
	result1[3] = 100
	result2 := make([]int, 101)
	result2[3] = 100

	var lessLargeThanThreeSlice = make([]int, 100)
	tests := []struct {
		name string
		args args
	}{
		{
			name: "nilSlice",
			args: args{
				newElement: 100,
				slice:      &nilSlice,
				result:     &result1,
			}},
		{
			name: "lengthLessThanThree",
			args: args{
				newElement: 100,
				slice:      &lengthLessThreeSlice,
				result:     &result1,
			}},
		{
			name: "lengthLargeThanThree",
			args: args{
				newElement: 100,
				slice:      &lessLargeThanThreeSlice,
				result:     &result2,
			}},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			insertElementAtIndexThird(tt.args.newElement, tt.args.slice)
			require.ElementsMatch(t, *(tt.args.slice), *(tt.args.result))
		})
	}
}

func Benchmark_insertElementAtIndexThird(b *testing.B) {
	for i := 0; i < b.N; i++ {
		insertSlice := make([]int, 100)
		insertElement := 100
		insertElementAtIndexThird(insertElement, &insertSlice)
	}
}
