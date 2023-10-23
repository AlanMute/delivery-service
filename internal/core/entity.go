package core

type Order struct {
	OrderUID          string   `gorm:"primary_key" json:"order_uid" faker:"uuid_digit"`
	TrackNumber       string   `json:"track_number"`
	Entry             string   `json:"entry"`
	DeliveryID        uint     `json:"-" faker:"-"`
	PaymentID         uint     `json:"-" faker:"-"`
	Delivery          Delivery `gorm:"foreignKey:DeliveryID"`
	Payment           Payment  `gorm:"foreignKey:PaymentID"`
	Items             []Item
	Locale            string `json:"locale" faker:"len=2"`
	InternalSignature string `json:"internal_signature" faker:"len=15"`
	CustomerID        string `json:"customer_id" faker:"len=15"`
	DeliveryService   string `json:"delivery_service" faker:"len=15"`
	ShardKey          string `json:"shardkey" faker:"len=10"`
	SMID              uint   `json:"sm_id"`
	DateCreated       string `json:"date_created" faker:"date"`
	OOFShard          string `json:"oof_shard" faker:"len=3"`
}

type Delivery struct {
	ID      uint   `gorm:"primary_key" json:"-" faker:"-"`
	Name    string `faker:"name"`
	Phone   string `faker:"phone_number"`
	Zip     string `faker:"len=7"`
	City    string `faker:"len=5"`
	Address string `faker:"len=10"`
	Region  string `faker:"len=10"`
	Email   string `faker:"email"`
}

type Payment struct {
	ID           uint   `gorm:"primary_key" json:"-" faker:"-"`
	Transaction  string `faker:"-"`
	RequestID    string `json:"request_id" faker:"len=5"`
	Currency     string `faker:"currency"`
	Provider     string `faker:"len=5"`
	Amount       uint
	PaymentDate  uint   `json:"payment_dt"`
	Bank         string `faker:"len=5"`
	DeliveryCost uint   `json:"delivery_cost"`
	GoodsTotal   uint   `json:"goods_total"`
	CustomFee    uint   `json:"custom_fee"`
}

type Item struct {
	ChrtID      uint   `gorm:"primary_key" json:"chrt_id"`
	TrackNumber string `json:"track_number" faker:"-"`
	Price       uint
	RID         string `json:"rid"`
	Name        string `faker:"oneof: Mobile, Notebook, Cherry, Racer"`
	Sale        uint
	Size        string `faker:"oneof: S, M, L, X, XL"`
	TotalPrice  uint   `json:"total_price"`
	NmID        uint   `json:"nm_id"`
	Brand       string `faker:"oneof: Apple, Samsung, Cherry, Grom"`
	Status      uint
}
