package vending_machine

import (
	"spark_networks_assessment/pkg/repositories/products"
	"testing"
)

type mockProductRepo struct {
	check int
	purchase bool
}

func (m *mockProductRepo) Add(product *products.Product) error {
	panic("implement me")
}

func (m *mockProductRepo) List() map[string][]*products.Product {
	panic("implement me")
}

func (m *mockProductRepo) Check(product *products.Product) int {
	return m.check
}

func (m *mockProductRepo) Purchase(product *products.Product) bool {
	return m.purchase
}

func Test_service_Charge(t *testing.T) {
	type fields struct {
		productRepo products.Repository
	}
	type args struct {
		coins []int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "charge 10 coins",
			fields: fields{
				productRepo: new(mockProductRepo),
			},
			args: args{
				coins: []int{10},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := New(tt.fields.productRepo)
			if got := s.Charge(tt.args.coins); got != tt.want {
				t.Errorf("Charge() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_Balance(t *testing.T) {
	type fields struct {
		UserCoins []int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "balance 20 coins",
			fields: fields{
				UserCoins: []int{10, 10},
			},
			want: 20,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				userCoins: tt.fields.UserCoins,
			}
			if got := s.Balance(); got != tt.want {
				t.Errorf("Balance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_Select(t *testing.T) {
	type fields struct {
		coins       []int
		productRepo products.Repository
	}
	type args struct {
		product *products.Product
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "select SparkPasta 1 available charging 35 coins",
			fields: fields{
				coins: []int{10, 25},
				productRepo: &mockProductRepo{
					check: 1,
					purchase: true,
				},
			},
			args: args{
				product: products.SparkPasta,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := New(tt.fields.productRepo)
			if len(tt.fields.coins) > 0 {
				s.Charge(tt.fields.coins)
			}
			if got := s.Select(tt.args.product); got != tt.want {
				t.Errorf("Select() = %v, want %v", got, tt.want)
			}
		})
	}
}
