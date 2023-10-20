package core

type Order struct {
	OrderUID          string   `gorm:"primary_key" json:"order_uid"`
	TrackNumber       string   `json:"track_number"`
	Entry             string   `json:"entry"`
	DeliveryID        uint     `json:"-"`
	PaymentID         uint     `json:"-"`
	Delivery          Delivery `gorm:"foreignKey:DeliveryID"`
	Payment           Payment  `gorm:"foreignKey:PaymentID"`
	Items             []Item
	Locale            string `json:"locale"`
	InternalSignature string `json:"internal_signature"`
	CustomerID        string `json:"customer_id"`
	DeliveryService   string `json:"delivery_service"`
	ShardKey          string `json:"shardkey"`
	SMID              int    `json:"sm_id"`
	DateCreated       string `json:"date_created"`
	OOFShard          string `json:"oof_shard"`
}

type Delivery struct {
	ID      uint `gorm:"primary_key" json:"-"`
	Name    string
	Phone   string
	Zip     string
	City    string
	Address string
	Region  string
	Email   string
}

type Payment struct {
	ID           uint `gorm:"primary_key" json:"-"`
	Transaction  string
	RequestID    string `json:"request_id"`
	Currency     string
	Provider     string
	Amount       int
	PaymentDate  int `json:"payment_dt"`
	Bank         string
	DeliveryCost int `json:"delivery_cost"`
	GoodsTotal   int `json:"goods_total"`
	CustomFee    int `json:"custom_fee"`
}

type Item struct {
	ChrtID      uint   `gorm:"primary_key" json:"chrt_id"`
	TrackNumber string `json:"track_number"`
	Price       int
	RID         string `json:"rid"`
	Name        string
	Sale        int
	Size        string
	TotalPrice  int `json:"total_price"`
	NmID        int `json:"nm_id"`
	Brand       string
	Status      int
}
