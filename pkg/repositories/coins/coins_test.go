package coins

import (
	"errors"
	"reflect"
	"testing"
)

func TestRepository_Add(t *testing.T) {
	tests := []struct {
		name string
		args int
		want error
	}{
		{
			name: "add 1 coins",
			args: 1,
			want: nil,
		},
		{
			name: "add 5 coins",
			args: 5,
			want: nil,
		},
		{
			name: "add 10 coins",
			args: 10,
			want: nil,
		},
		{
			name: "add 25 coins",
			args: 25,
			want: nil,
		},
		{
			name: "add 7 coins",
			args: 7,
			want: errors.New(invalidAmountError),
		},
		{
			name: "add -1 coins",
			args: -1,
			want: errors.New(negativeAmountError),
		},
		{
			name: "add 0 coins",
			args: 0,
			want: errors.New(zeroAmountError),
		},
	}
	c := New()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := c.Add(tt.args)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("error got %v want %v", got, tt.want)
			}
		})
	}
}

func TestRepository_List(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want []int
	}{
		{
			name: "intial amount",
			args: make([]int, 0),
			want: make([]int, 0),
		},
		{
			name: "add 10 coins",
			args: []int{10},
			want: []int{10},
		},
		{
			name: "add 5 coins",
			args: []int{5},
			want: []int{5},
		},
		{
			name: "add 1, 5, 10 coins",
			args: []int{1, 5, 10},
			want: []int{1, 5, 10},
		},
		{
			name: "add 5, 5, 5 coins",
			args: []int{5, 5, 5},
			want: []int{5, 5, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := New()
			for _, arg := range tt.args {
				err := c.Add(arg)
				if err != nil {
					t.Fatal(err)
				}
			}
			got := c.List()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("list return got %v want %v", got, tt.want)
			}
		})
	}
}
