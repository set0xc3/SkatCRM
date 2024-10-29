package db

import (
	"encoding/json"
	"net/http"
)

type Product struct {
	Id       string `db:"id" json:"id"`
	Name     string `db:"name" json:"name"`
	Price    string `db:"price" json:"price"`
	Quantity int    `db:"quantity" json:"quantity"`
}

func FetchProducts() []Product {
	resp, err := http.Get("http://localhost:8090/api/v1/products")
	if err != nil {
		return nil
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil
	}

	var products []Product
	if err := json.NewDecoder(resp.Body).Decode(&products); err != nil {
		return nil
	}

	return products
}

func FetchClients() []Product {
	resp, err := http.Get("http://localhost:8090/api/v1/clients")
	if err != nil {
		return nil
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil
	}

	var products []Product
	if err := json.NewDecoder(resp.Body).Decode(&products); err != nil {
		return nil
	}

	return products
}
