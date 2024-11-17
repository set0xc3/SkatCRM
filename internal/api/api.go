package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

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

func FetchClients(count int, offset int) (data []db.Client) {
	err := FetchData("http://localhost:8090/api/v1/clients/"+strconv.Itoa(count)+"/"+strconv.Itoa(offset), &data)
	if err != nil {
		return nil
	}
	return data
}

func FetchClientCount() (data int) {
	err := FetchData("http://localhost:8090/api/v1/clients", &data)
	if err != nil {
		return 0
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
