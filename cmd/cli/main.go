package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/set0xc3/htmx/internal/db"
	"github.com/xuri/excelize/v2"
)

var FILE_PATH string

func ReadProductFromFile(file_path string) {
	file, err := excelize.OpenFile(FILE_PATH)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Get all the rows in the Sheet1.
	rows, err := file.GetRows("Складские остатки")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Пропускаем заголовок и обрабатываем каждую строку
	for i, row := range rows {
		if i == 0 {
			continue
		}
		product := db.Product{}
		product.Id2 = row[0]
		if len(row) > 1 {
			product.Name = row[1]
		}
		if len(row) > 2 {
			product.SerialNumber = row[2]
		}
		if len(row) > 3 {
			product.Article = row[3]
		}
		if len(row) > 4 {
			product.Date = row[4]
		}
		if len(row) > 5 {
			product.Quantity = row[5]
		}
		if len(row) > 6 {
			product.RetailPrice = row[6]
		}
		if len(row) > 7 {
			product.PurchasePrice = row[7]
		}
		if len(row) > 8 {
			product.ExchangeRatePC = row[8]
		}
		if len(row) > 9 {
			product.ExchangeRatePR = row[9]
		}
		if len(row) > 10 {
			product.Warehouse = row[10]
		}
		if len(row) > 11 {
			product.Location = row[11]
		}
		if len(row) > 12 {
			product.CustomerOrder = row[12]
		}
		if len(row) > 13 {
			product.SupplierOrder = row[13]
		}
		if len(row) > 14 {
			product.Supplier = row[14]
		}

		SendToDB("http://localhost:8090/api/v1/product", product)
	}

}

func ReadClientFromFile(file_path string) {
	file, err := excelize.OpenFile(FILE_PATH)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Get all the rows in the Sheet1.
	rows, err := file.GetRows("Клиенты")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Пропускаем заголовок и обрабатываем каждую строку
	for i, row := range rows {
		if i == 0 {
			continue
		}
		client := db.Client{}
		client.Id2 = row[0]
		if len(row) > 1 {
			client.Mark = row[1]
		}
		if len(row) > 2 {
			client.Contractor = row[2]
		}
		if len(row) > 3 {
			client.FullName = row[3]
		}
		if len(row) > 4 {
			client.Type = row[4]
		}
		if len(row) > 5 {
			client.Phones = row[5]
		}
		if len(row) > 6 {
			client.Email = row[6]
		}
		if len(row) > 7 {
			client.LegalAddress = row[7]
		}
		if len(row) > 8 {
			client.PhysicalAddress = row[8]
		}
		if len(row) > 9 {
			client.RegistrationDate = row[9]
		}
		if len(row) > 10 {
			client.AdChannel = row[10]
		}
		if len(row) > 11 {
			client.RegData1 = row[11]
		}
		if len(row) > 12 {
			client.RegData2 = row[12]
		}
		if len(row) > 13 {
			client.Note = row[13]
		}
		if len(row) > 14 {
			client.RequestCount = row[14]
		}
		if len(row) > 15 {
			client.Birthday = row[15]
		}
		if len(row) > 16 {
			client.Income = row[16]
		}

		SendToDB("http://localhost:8090/api/v1/client", client)
	}

}

func SendToDB(url string, data interface{}) error {
	// Сериализуем данные в JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	// Отправляем POST-запрос с JSON-данными
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Проверяем статус ответа
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to add: %s", resp.Status)
	}

	return nil
}

func main() {
	if len(os.Args) > 1 {
		FILE_PATH = os.Args[1]
	}

	// ReadClientFromFile(FILE_PATH)
	ReadProductFromFile(FILE_PATH)
}
