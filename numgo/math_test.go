package numgo

import (
	"reflect"
	"testing"
)

func TestMx_sigmoid(t *testing.T) {
	type fields struct {
		Vec [][]float32
		T   [][]float32
	}
	tests := []struct {
		name   string
		fields fields
		want   Mx
	}{
		// TODO: Add test cases.
		{
			name: "(1x3)の行列のシグモイド計算",
			fields: fields{
				Vec: [][]float32{{1, 2, 3}},
			},
			/*
				>>> a = np.array([[1,2,3]])
				>>> 1 / (1 + np.exp(-a))
				array([[0.73105858, 0.88079708, 0.95257413]])
			*/
			want: NewMx([][]float32{{0.73105858, 0.880797, 0.95257413}}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Mx{
				Vec: tt.fields.Vec,
				T:   tt.fields.T,
			}
			if got := m.Sigmoid(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Mx.sigmoid()\n   = %v, \nwant %v", got, tt.want)
			}
		})
	}
}
