package products

const (
	SparkPastaName  = "SparkPastaName"
	SparkPastaValue = 35
)

var (
	SparkPasta = &Product{
		Name:  SparkPastaName,
		Value: SparkPastaValue,
	}
)

type Product struct {
	Name  string
	Value int
}

type Repository interface {
	Add(*Product) error
	List() map[string][]*Product
	Check(*Product) int
	Purchase(*Product) bool
}

type repository struct {
	AvailableProducts map[string][]*Product
}

func New() Repository {
	return &repository{
		AvailableProducts: make(map[string][]*Product),
	}
}

func (r *repository) Add(product *Product) error {
	//TODO product validation

	r.AvailableProducts[product.Name] = append(r.AvailableProducts[product.Name], product)

	return nil
}

func (r *repository) List() map[string][]*Product {
	return r.AvailableProducts
}

func (r *repository) Check(product *Product) int {
	if _, ok := r.AvailableProducts[product.Name]; ok {
		return len(r.AvailableProducts[product.Name])
	}

	return 0
}

func (r repository) Purchase(product *Product) bool {
	if r.Check(product) > 0 {

		list := r.AvailableProducts[product.Name]
		r.AvailableProducts[product.Name] = list[:len(list)-1]

		return true
	}

	return false
}
