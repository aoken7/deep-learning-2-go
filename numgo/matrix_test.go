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

func TestHad(t *testing.T) {
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
			name: "(2x2)と(2x2)の行列のアダマール積",
			args: args{
				a: NewMx([][]float32{{1, 2}, {3, 4}}),
				b: NewMx([][]float32{{5, 6}, {7, 8}}),
			},
			/*
				>>> a = np.array([[1,2],[3,4]])
				>>> b = np.array([[5,6],[7,8]])
				>>> a * b
				array([[ 5, 12],
				       [21, 32]])
			*/
			want: NewMx([][]float32{{5, 12}, {21, 32}}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Had(tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Had() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSum(t *testing.T) {
	type args struct {
		a    Mx
		axis int
	}
	tests := []struct {
		name string
		args args
		want Mx
	}{
		{
			name: "(2x3)の行列を行にSum",
			args: args{
				a:    NewMx([][]float32{{1, 2, 3}, {2, 3, 4}}),
				axis: 0,
			},
			/*
				>>> a = np.array([[1,2],[3,4]])
				>>> np.sum(a, axis=0)
				array([4, 6])
			*/
			want: NewMx([][]float32{{3, 5, 7}}),
		},
		{
			name: "(2x3)の行列を列にSum",
			args: args{
				a:    NewMx([][]float32{{1, 2, 3}, {2, 3, 4}}),
				axis: 1,
			},
			/*
				>>> b = np.array([[1,2,3],[2,3,4]])
				>>> np.sum(b, axis=1)
				array([6, 9])
			*/
			want: NewMx([][]float32{{6, 9}}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sum(tt.args.a, tt.args.axis); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMx_Max(t *testing.T) {
	type fields struct {
		Vec [][]float32
		t   [][]float32
	}
	type args struct {
		axis int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Mx
	}{
		// TODO: Add test cases.
		{
			name:   "(2x3)の行列の列方向へのMax",
			fields: fields{Vec: [][]float32{{1, 5, 3}, {2, 4, 6}}},
			args:   args{axis: 0},
			/*
				>>> a = np.array([[1,5,3],[2,4,6]])
				>>> a.max(axis=0,keepdims=True)
				array([[2, 5, 6]])
			*/
			want: NewMx([][]float32{{2, 5, 6}}),
		},
		{
			name:   "(2x3)の行列の行方向へのMax",
			fields: fields{Vec: [][]float32{{1, 5, 3}, {2, 4, 6}}},
			args:   args{axis: 1},
			/*
				>>> a = np.array([[1,5,3],[2,4,6]])
				>>> a.max(axis=1,keepdims=True)
				array([[5],
				       [6]])
			*/
			want: NewMx([][]float32{{5}, {6}}),
		},
		{
			name:   "(2x2)の行列の列方向へ負を含むときのMax",
			fields: fields{Vec: [][]float32{{-2, -4}, {-1, -5}}},
			args:   args{axis: 0},
			want:   NewMx([][]float32{{-1, -4}}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Mx{
				Vec: tt.fields.Vec,
				t:   tt.fields.t,
			}
			if got := m.Max(tt.args.axis); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Mx.Max() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSub(t *testing.T) {
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
			name: "スカラーと行列のSub",
			args: args{
				a: NewMx([][]float32{{1, 2, 3}}),
				w: NewMx([][]float32{{10}}),
			},
			want: NewMx([][]float32{{-9, -8, -7}}),
		},
		{
			name: "行列と行列のSub",
			args: args{
				a: NewMx([][]float32{{1, 2, 3}, {4, 5, 6}}),
				w: NewMx([][]float32{{10, 10, 10}}),
			},
			want: NewMx([][]float32{{-9, -8, -7}, {-6, -5, -4}}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sub(tt.args.a, tt.args.w); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Sub() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMx_Sum(t *testing.T) {
	type fields struct {
		Vec [][]float32
		t   [][]float32
	}
	type args struct {
		axis int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Mx
	}{
		// TODO: Add test cases.
		{
			name:   "(2x3)の行列の列ごとにSumを計算",
			fields: fields{Vec: [][]float32{{1, 5, 3}, {2, 4, 6}}},
			args:   args{axis: 0},
			want:   NewMx([][]float32{{3, 9, 9}}),
		},
		{
			name:   "(2x3)の行列の行ごとにSumを計算",
			fields: fields{Vec: [][]float32{{1, 5, 3}, {2, 4, 6}}},
			args:   args{axis: 1},
			want:   NewMx([][]float32{{9}, {12}}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Mx{
				Vec: tt.fields.Vec,
				t:   tt.fields.t,
			}
			if got := m.Sum(tt.args.axis); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Mx.Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}
