package db

type Product struct {
	Id             string `db:"id" json:"id"`
	Id2            string `db:"id2" json:"id2"`                           // № изделия
	Name           string `db:"name" json:"name"`                         // Наименование
	SerialNumber   string `db:"serial_number" json:"serial_number"`       // Серийный номер
	Article        string `db:"article" json:"article"`                   // Артикул
	Date           string `db:"date" json:"date"`                         // Дата
	Quantity       string `db:"quantity" json:"quantity"`                 // Кол-во
	RetailPrice    string `db:"retail_price" json:"retail_price"`         // Розничная цена
	PurchasePrice  string `db:"purchase_price" json:"purchase_price"`     // Закупочная цена
	ExchangeRatePC string `db:"exchange_rate_pc" json:"exchange_rate_pc"` // Курс ПК
	ExchangeRatePR string `db:"exchange_rate_pr" json:"exchange_rate_pr"` // Курс ПР
	Warehouse      string `db:"warehouse" json:"warehouse"`               // Склад
	Location       string `db:"location" json:"location"`                 // Локация
	CustomerOrder  string `db:"customer_order" json:"customer_order"`     // Заказ клиента
	SupplierOrder  string `db:"supplier_order" json:"supplier_order"`     // Заказ поставщику
	Supplier       string `db:"supplier" json:"supplier"`                 // Поставщик
}

type Client struct {
	Id               string `db:"id" json:"id"`
	Id2              string `db:"id2" json:"id2"`
	Mark             string `db:"mark" json:"mark"`
	Contractor       string `db:"contractor" json:"contractor"`
	FullName         string `db:"full_name" json:"full_name"`
	Type             string `db:"type" json:"type"`
	Phones           string `db:"phones" json:"phones"`
	Email            string `db:"email" json:"email"`
	LegalAddress     string `db:"legal_address" json:"legal_address"`
	PhysicalAddress  string `db:"physical_address" json:"physical_address"`
	RegistrationDate string `db:"registration_date" json:"registration_date"`
	AdChannel        string `db:"ad_channel" json:"ad_channel"`
	RegData1         string `db:"reg_data_1" json:"reg_data_1"`
	RegData2         string `db:"reg_data_2" json:"reg_data_2"`
	Note             string `db:"note" json:"note"`
	RequestCount     string `db:"request_count" json:"request_count"`
	Birthday         string `db:"birthday" json:"birthday"`
	Income           string `db:"income" json:"income"`
}

type Call struct {
	Id               string `db:"id" json:"id"`
	Id2              string `db:"id2" json:"id2"`
}
