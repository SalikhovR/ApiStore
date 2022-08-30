package main

import (
	api "ApiStore/Api"
	"ApiStore/repository"
	"ApiStore/repository/posgres"
	"net/http"
)

func main() {
	db := posgres.Connect()
	psql := repository.New(db)
	r := api.NewRouter(*psql)
	http.ListenAndServe(":8080", r)
}
