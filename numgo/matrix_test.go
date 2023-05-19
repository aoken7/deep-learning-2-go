package numgo

import (
	"reflect"
	"testing"
)

func TestDot(t *testing.T) {
	type args struct {
		a Mx
		b Mx
	}
	tests := []struct {
		name string
		args args
		want Mx
	}{
		{
			name: "2*2の行列積",
			args: args{
				a: NewMx([][]float32{{1, 2}, {3, 4}}),
				b: NewMx([][]float32{{1, 2}, {3, 4}}),
			},
			/*
				>>> a = np.array([[1,2],[3,4]])
				>>> b = np.array([[1,2],[3,4]])
				>>> c = np.dot(a,b)
				>>> c
				array([[ 7, 10],
					   [15, 22]])
			*/
			want: NewMx([][]float32{{7, 10}, {15, 22}}),
		},
		{
			name: "2*3の行列積",
			args: args{
				a: NewMx([][]float32{{1, 2}, {3, 4}}),
				b: NewMx([][]float32{{1, 2, 3}, {4, 5, 6}}),
			},
			/*
				>>> a = np.array([[1,2],[3,4]])
				>>> b = np.array([[1,2,3],[4,5,6]])
				>>> c = np.dot(a,b)
				>>> c
				array([[ 9, 12, 15],
				       [19, 26, 33]])
			*/
			want: NewMx([][]float32{{9, 12, 15}, {19, 26, 33}}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Dot(tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Dot() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRandn(t *testing.T) {
	type args struct {
		l int
		r int
	}
	tests := []struct {
		name string
		args args
		want Mx
	}{
		// TODO: Add test cases.
		{
			name: "動作確認",
			args: args{2, 3},
			want: NewMx([][]float32{}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			/* if got := Randn(tt.args.l, tt.args.r); !reflect.DeepEqual(got, tt.want) {
				fmt.Println(got)
				t.Errorf("Randn() = %v, want %v", got, tt.want)
			} */
		})
	}
}
