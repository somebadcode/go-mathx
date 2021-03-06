package mathx

import "testing"

func Test_equal(t *testing.T) {
	type args struct {
		x float64
		y float64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "equal",
			args: args{
				x: 1.0,
				y: 1.0000000001,
			},
			want: true,
		},
		{
			name: "unequal",
			args: args{
				x: 1.0,
				y: 1.000000001,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Equal(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("Equal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_equal(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		Equal(10.0, 10.0000000001)
	}
}
