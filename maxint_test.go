package mathx

import (
	"math"
	"testing"
)

func Test_MaxInt(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "a_gt_b",
			args: args{
				a: 200,
				b: 100,
			},
			want: 200,
		},
		{
			name: "a_lt_b",
			args: args{
				a: -23,
				b: 100,
			},
			want: 100,
		},
		{
			name: "a_eq_b",
			args: args{
				a: 9808,
				b: 9808,
			},
			want: 9808,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxInt(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("MaxInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_MaxUInt(t *testing.T) {
	type args struct {
		a uint
		b uint
	}
	tests := []struct {
		name string
		args args
		want uint
	}{
		{
			name: "a_gt_b",
			args: args{
				a: 200,
				b: 100,
			},
			want: 200,
		},
		{
			name: "a_lt_b",
			args: args{
				a: 42,
				b: 100,
			},
			want: 100,
		},
		{
			name: "a_eq_b",
			args: args{
				a: 9808,
				b: 9808,
			},
			want: 9808,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxUint(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("MaxInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_MaxInt64(t *testing.T) {
	type args struct {
		a int64
		b int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "a_gt_b",
			args: args{
				a: 200,
				b: 100,
			},
			want: 200,
		},
		{
			name: "a_lt_b",
			args: args{
				a: math.MinInt64,
				b: 100,
			},
			want: 100,
		},
		{
			name: "a_eq_b",
			args: args{
				a: 9808,
				b: 9808,
			},
			want: 9808,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxInt64(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("MaxInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_MaxUInt64(t *testing.T) {
	type args struct {
		a uint64
		b uint64
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{
			name: "a_gt_b",
			args: args{
				a: 200,
				b: 100,
			},
			want: 200,
		},
		{
			name: "a_lt_b",
			args: args{
				a: 42,
				b: 100,
			},
			want: 100,
		},
		{
			name: "a_eq_b",
			args: args{
				a: 9808,
				b: 9808,
			},
			want: 9808,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxUint64(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("MaxInt() = %v, want %v", got, tt.want)
			}
		})
	}
}