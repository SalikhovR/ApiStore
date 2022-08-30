package repository

import (
	product "ApiStore/Product"
	"ApiStore/inventory"
	"database/sql"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"time"
)

type PsqlRepository struct {
	db *sql.DB
}

func New(db *sql.DB) *PsqlRepository {
	return &PsqlRepository{db: db}
}

func (p *PsqlRepository) CreateProduct(Pr product.Product) {
	Pr.Id = uuid.New().String()
	r, err := p.db.Query(`SELECT * FROM products`)
	if err != nil {
		panic(err)
	}
	products := make([]product.Product, 0)
	for r.Next() {
		var pr product.Product
		err = r.Scan(&pr.Id, &pr.Name, &pr.Quantity, &pr.Price, &pr.OriginalPrice)
		products = append(products, pr)
	}
	for _, v := range products {
		if v.Name == Pr.Name {
			_, err = p.db.Exec(`UPDATE products SET quantity =$1  WHERE id = $2`, v.Quantity+Pr.Quantity, v.Id)
			if err != nil {
				panic(err)
			}
			return
		}
	}
	s := fmt.Sprintf("INSERT INTO products VALUES ('%s','%s',%d,%d,%d)", Pr.Id, Pr.Name, Pr.Quantity, Pr.Price+Pr.Price/3, Pr.Price)
	print("\n", s)
	_, err = p.db.Exec(s)
	if err != nil {
		panic(err)
	}
}
func (p PsqlRepository) Products() []product.Product {
	r, err := p.db.Query(`SELECT * FROM products`)
	if err != nil {
		panic(err)
	}
	products := make([]product.Product, 0)
	for r.Next() {
		var pr = product.Product{}
		err = r.Scan(&pr.Id, &pr.Name, &pr.Quantity, &pr.Price, &pr.OriginalPrice)
		products = append(products, pr)
	}
	return products
}

func (p *PsqlRepository) NewBasket() string {
	id := uuid.New().String()
	_, err := p.db.Exec(`INSERT INTO basket VALUES ($1,$2)`, id, time.Now())
	if err != nil {
		panic(err)
	}
	return id
}

func (p *PsqlRepository) Shop(Pid, Bid string, q int) error {
	pr := product.Product{}
	r1 := p.db.QueryRow(`SELECT * FROM products WHERE id = $1`, Pid)
	//r2 := p.db.QueryRow(`SELECT id FROM basket WHERE id = $1`, Bid)
	if err := r1.Scan(&pr.Id, &pr.Name, &pr.Quantity, &pr.Price, &pr.OriginalPrice); err != nil {
		return err
	}
	fmt.Println(pr)
	if pr.Quantity > q {
		_, err := p.db.Exec(`INSERT INTO check VALUES ($1,$2,$3)`, Pid, Bid, q)
		if err != nil {
			return err
		}
		_, err = p.db.Exec(`UPDATE products SET quantity = quantity - $1 WHERE id = $2`, q, pr.Id)
		if err != nil {
			return err
		}
	} else if pr.Quantity == q {
		_, err := p.db.Exec(`INSERT INTO check VALUES ($1,$2,$3)`, Pid, Bid, q)
		if err != nil {
			return err
		}
		_, err = p.db.Exec(`DELETE FROM products WHERE id = $1`, pr.Id)
		if err != nil {
			return err
		}
	} else {
		return errors.New("kam")
	}
	//r1 := p.db.QueryRow(`SELECT price FROM products WHERE id = $1`, Pid)
	_, err := p.db.Exec(`UPDATE store SET profit = profit + $1`, pr.Price*q)
	if err != nil {
		return err
	}

	return nil
}

func (p PsqlRepository) Check(id string) inventory.CheckJson {
	r, err := p.db.Query(`SELECT * FROM check WHERE basket_id = $1`, id)
	if err != nil {
		panic(err)
	}
	sum := 0
	products := make([]inventory.ProductJson, 0)
	for r.Next() {
		var pr inventory.ProductJson
		var pid, bid string
		var q int
		err = r.Scan(&pid, &bid, &q)
		if err != nil {
			panic(err)
		}
		or := 0
		err = p.db.QueryRow(`SELECT * FROM products WHERE id = $1`, pid).Scan(&pr.Id, &pr.Name, &pr.Quantity, &pr.Price, &or)
		if err != nil {
			panic(err)
		}
		pr.Quantity = q
		pr.Price *= q
		sum += pr.Price
		products = append(products, pr)
	}
	var check = inventory.CheckJson{
		Bought: products,
		Sum:    sum,
	}
	return check
}
