package core

type Order struct {
	OrderUID          string `json:"order_uid" gorm:"primaryKey"`
	TrackNumber       string `json:"track_number"`
	Entry             string `json:"entry"`
	Delivery          Delivery
	Payment           Payment
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
	ID      uint `gorm:"primaryKey" json:"-"`
	Name    string
	Phone   string
	Zip     string
	City    string
	Address string
	Region  string
	Email   string
}

type Payment struct {
	Transaction  string `gorm:"primaryKey"`
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
	ChrtID      int    `json:"chrt_id" gorm:"primaryKey"`
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
