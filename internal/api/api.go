package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/set0xc3/htmx/internal/db"
)

func FetchData(url string, result interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to fetch data: status code %d", resp.StatusCode)
	}

	return json.NewDecoder(resp.Body).Decode(result)
}

func FetchClients(count string) (data []db.Client) {
	err := FetchData("http://localhost:8090/api/v1/clients/"+count, &data)
	if err != nil {
		return nil
	}
	return data
}

func FetchProducts() (data []db.Product) {
	err := FetchData("http://localhost:8090/api/v1/products", &data)
	if err != nil {
		return nil
	}
	return data
}
