package db

type Product struct {
	Id       string `db:"id" json:"id"`
	Name     string `db:"name" json:"name"`
	Price    string `db:"price" json:"price"`
	Quantity int    `db:"quantity" json:"quantity"`
}

type Client struct {
	Id    string `db:"id" json:"id"`
	Id2   string `db:"id2" json:"id2"`
	Name  string `db:"name" json:"name"`
	Phone string `db:"phone" json:"phone"`
}
