package common

import (
	"numgo/numgo"
	"reflect"
	"testing"
)

func Test_softmax(t *testing.T) {
	type args struct {
		x numgo.Mx
	}
	tests := []struct {
		name string
		args args
		want numgo.Mx
	}{
		{
			name: "(2x3)の行列のSoftmax",
			args: args{numgo.NewMx([][]float32{
				{1, 5, 3},
				{2, 4, 6}})},
			/*
				>>> a = np.array([[1,5,3],[2,4,6]])
				>>> np.exp(a - a.max(axis=1, keepdims=True)) / a.sum(axis=1, keepdims=True)
				   array([[0.00203507, 0.11111111, 0.01503725],
				          [0.0015263 , 0.01127794, 0.08333333]])
			*/
			want: numgo.NewMx([][]float32{
				{0.00203507, 0.11111111, 0.01503725},
				{0.0015263, 0.01127794, 0.08333333}}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := softmax(tt.args.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("softmax() = %v, want %v", got, tt.want)
			}
		})
	}
}
