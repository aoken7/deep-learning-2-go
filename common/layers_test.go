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
				>>> m = m - m.max(axis=1, keepdims=True)
				>>> m = np.exp(x)
				>>> m /= m.sum(axis=1, keepdims=True)
					[[0.01587624 0.86681333 0.11731043]
					 [0.01587624 0.11731043 0.86681333]]
			*/
			want: numgo.NewMx([][]float32{
				{0.01587624, 0.86681333, 0.11731042},
				{0.015876241, 0.11731043, 0.86681336}}),
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
