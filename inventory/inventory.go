package inventory

import (
	product "ApiStore/Product"
)

type ProductJson struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
	Price    int    `json:"price"`
}

type NewProduct struct {
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
	Price    int    `json:"price"`
}

type CheckJson struct {
	Bought []ProductJson
	Sum    int
}

func (p ProductJson) MarshAl(pr []product.Product) []ProductJson {
	pj := []ProductJson{}
	for _, v := range pr {
		pj = append(pj, ProductJson{Id: v.Id, Name: v.Name, Quantity: v.Quantity, Price: v.Price})
	}
	return pj
}

func (p NewProduct) MarshAl() product.Product {
	return product.Product{Name: p.Name, Quantity: p.Quantity, Price: p.Price}
}
