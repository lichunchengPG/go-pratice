package model


type Product struct {
	productName string
	productPrice float32
}

func (p *Product) SetProductName(_productName string)  {
	p.productName = _productName
}

func (p *Product) GetProductName() string {
	return p.productName
}