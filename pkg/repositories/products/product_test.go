package products

import (
	"reflect"
	"testing"
)

func Test_repository_List(t *testing.T) {
	tests := []struct {
		name string
		want map[string][]*Product
	}{
		{
			name: "initial state",
			want: make(map[string][]*Product),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := New()
			if got := r.List(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("List() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_repository_Check(t *testing.T) {
	type args struct {
		availableProducts []*Product
		checkProduct      *Product
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "check SparkPasta with 1 available",
			args: args{
				availableProducts: []*Product{
					SparkPasta,
				},
				checkProduct: SparkPasta,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := New()
			for _, product := range tt.args.availableProducts {
				err := r.Add(product)
				if err != nil {
					t.Fatal(err)
				}
			}
			if got := r.Check(tt.args.checkProduct); got != tt.want {
				t.Errorf("Check() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_repository_Purchase(t *testing.T) {
	type fields struct {
		AvailableProducts map[string][]*Product
	}
	type args struct {
		product *Product
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "buy SparkPasta",
			fields: fields{
				AvailableProducts: map[string][]*Product{
					SparkPastaName: {
						SparkPasta,
					},
				},
			},
			args: args{
				product: SparkPasta,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := New()
			for _, fields := range tt.fields.AvailableProducts {
				for _, product := range fields {
					err := r.Add(product)
					if err != nil {
						t.Fatal(err)
					}
				}
			}
			got := r.Purchase(tt.args.product)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("purchate got %v want %v", got, tt.want)
			}
		})
	}
}
