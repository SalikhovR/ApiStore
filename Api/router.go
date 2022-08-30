package api

import (
	"ApiStore/inventory"
	"ApiStore/repository"
	"encoding/json"
	"github.com/go-chi/chi"
	"io"
	"net/http"
	"strconv"
)

type Handle struct {
	psql repository.PsqlRepository
}

func (h Handle) Products(w http.ResponseWriter, r *http.Request) {
	pr := h.psql.Products()
	jp := inventory.ProductJson{}.MarshAl(pr)
	encoder := json.NewEncoder(w)
	encoder.SetIndent(" ", " ")
	if err := encoder.Encode(jp); err != nil {
		panic(err)
	}

}
func (h Handle) NewBasket(w http.ResponseWriter, r *http.Request) {
	id := h.psql.NewBasket()
	_, err := io.WriteString(w, id)
	if err != nil {
		panic(err)
	}
}
func (h Handle) CreateProduct(w http.ResponseWriter, r *http.Request) {
	np := inventory.NewProduct{}
	json.NewDecoder(r.Body).Decode(&np)
	h.psql.CreateProduct(np.MarshAl())
}
func (h Handle) Shop(w http.ResponseWriter, r *http.Request) {
	val := r.URL.Query()
	pid := val.Get("product_id")
	bid := val.Get("basket_id")
	q, err := strconv.Atoi(val.Get("quantity"))

	if err != nil {
		panic(err)
	}
	h.psql.Shop(pid, bid, q)

}
func (h Handle) Check(w http.ResponseWriter, r *http.Request) {
	val := r.URL.Query()
	ch := h.psql.Check(val.Get("id"))
	encoder := json.NewEncoder(w)
	encoder.SetIndent(" ", " ")
	if err := encoder.Encode(ch); err != nil {
		panic(err)
	}
}

func NewRouter(psql repository.PsqlRepository) http.Handler {
	r := chi.NewRouter()
	h := Handle{psql: psql}
	r.Get("/Products", h.Products)
	r.Get("/NewBasket", h.NewBasket)
	r.Post("/CreateProduct", h.CreateProduct)
	r.Post("/Shop", h.Shop)
	r.Get("/Chek", h.Check)
	return r
}
