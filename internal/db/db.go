package db

type Product struct {
	Id       string `db:"id" json:"id"`
	Name     string `db:"name" json:"name"`
	Price    string `db:"price" json:"price"`
	Quantity int    `db:"quantity" json:"quantity"`
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
