package basket

import (
	"ApiStore/Product"
	"time"
)

type Basket struct {
	Id         string
	Products   map[int]product.Product
	CreateTime time.Time
}
