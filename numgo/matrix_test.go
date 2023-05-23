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
			name: "(1x3)*(3x1)の行列積",
			args: args{
				a: NewMx([][]float32{{1, 2, 3}}),
				b: NewMx([][]float32{{4}, {5}, {6}}),
			},
			/*
				>>> import numpy as np
				>>> a = np.array([1,2,3])
				>>> b = np.array([4,5,6])
				>>> np.dot(a,b)
				32
			*/
			want: NewMx([][]float32{{32}}),
		},
		{
			name: "(2x2)*(2x2)の行列積",
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
			name: "(2x2)*(2x3)の行列積",
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
			want: NewMx([][]float32{{1}}),
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

func TestMx_Exp(t *testing.T) {
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
			name: "(1x3)の行列のExp",
			fields: fields{
				Vec: [][]float32{{1}, {2}, {3}},
			},
			/*
				>>> a = np.array([[1],[2],[3]])
				>>> np.exp(a)
				array([[ 2.71828183],
				       [ 7.3890561 ],
				       [20.08553692]])
			*/
			want: NewMx([][]float32{{2.71828183}, {7.3890561}, {20.08553692}}),
		},
		{
			name: "(2x2)の行列のExp",
			fields: fields{
				Vec: [][]float32{{1, 2}, {3, 4}},
			},
			/*
				>>> b = np.array([[1,2],[3,4]])
				>>> np.exp(b)
				array([[ 2.71828183,  7.3890561 ],
				       [20.08553692, 54.59815003]])
			*/
			want: NewMx([][]float32{{2.71828183, 7.3890561}, {20.08553692, 54.59815003}}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Mx{
				Vec: tt.fields.Vec,
				t:   tt.fields.T,
			}
			if got := m.Exp(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Mx.Exp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAdd(t *testing.T) {
	type args struct {
		a Mx
		w Mx
	}
	tests := []struct {
		name string
		args args
		want Mx
	}{
		// TODO: Add test cases.
		{
			name: "スカラーを加算",
			args: args{
				a: NewMx([][]float32{{1, 2, 3}}),
				w: NewMx([][]float32{{10}}),
			},
			want: NewMx([][]float32{{11, 12, 13}}),
		},
		{
			name: "行列を加算",
			args: args{
				a: NewMx([][]float32{{1, 2, 3}, {4, 5, 6}}),
				w: NewMx([][]float32{{10, 10, 10}}),
			},
			want: NewMx([][]float32{{11, 12, 13}, {14, 15, 16}}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Add(tt.args.a, tt.args.w); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMx_Shape(t *testing.T) {
	type fields struct {
		Vec [][]float32
		t   [][]float32
	}
	tests := []struct {
		name   string
		fields fields
		want   []int
	}{
		// TODO: Add test cases.
		{
			name:   "(1x3)の行列の形状確認",
			fields: fields{Vec: [][]float32{{1, 2, 3}}},
			want:   []int{1, 3},
		},
		{
			name:   "(3x2)の行列の形状確認",
			fields: fields{Vec: [][]float32{{1, 2}, {3, 4}, {5, 6}}},
			want:   []int{3, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Mx{
				Vec: tt.fields.Vec,
				t:   tt.fields.t,
			}
			if got := m.Shape(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Mx.Shape() = %v, want %v", got, tt.want)
			}
		})
	}
}
