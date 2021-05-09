// 该代码通过gotests --all -w func.go 生成的模板
package funcTest

import "testing"

func Test_add(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test_a",
			args: args{1, 2},
			want: 3,
		},

		{
			name: "test_b",
			args: args{100, 200},
			want: 300,
		},
		{
			name: "test_c",
			args: args{1000, 2000},
			want: 3000,
		},
	}
	for index, tt := range tests {
		if testing.Short() && (index == 1 || index == 2) {
			t.Skipf("跳过当前用例")
		}
		t.Run(tt.name, func(t *testing.T) {
			if got := add(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("add() = %v, want %v", got, tt.want)
			}
		})
	}
}
