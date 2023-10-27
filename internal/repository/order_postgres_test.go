package repository

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/KrizzMU/delivery-service/internal/core"
	"github.com/go-playground/assert/v2"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func TestGetAllOrders(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	gormDB, err := gorm.Open("postgres", db)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a GORM database connection", err)
	}

	//gormDB.LogMode(true)

	orderPostgres := &OrderPostgres{db: gormDB}

	mock.ExpectQuery(`SELECT \* FROM "orders"`).WillReturnRows(sqlmock.NewRows([]string{
		"order_uid", "track_number", "entry", "delivery_id", "payment_id", "locale", "internal_signature",
		"customer_id", "delivery_service", "shard_key", "sm_id", "date_created", "oof_shard",
	}).AddRow("order1", "123", "entry1", 1, 1, "en", "sig1", "customer1", "service1", "shard1", 1, "2023-10-26", "oof1").
		AddRow("order2", "456", "entry2", 2, 2, "fr", "sig2", "customer2", "service2", "shard2", 2, "2023-10-27", "oof2"))

	mock.ExpectQuery(`SELECT \* FROM "deliveries" WHERE \("id" IN \(\$1,\$2\)\)`).WillReturnRows(sqlmock.NewRows([]string{"id", "name", "phone", "zip", "city", "address", "region", "email"}).
		AddRow(1, "Delivery1", "1234567890", "1234567", "City1", "Address1", "Region1", "email1").
		AddRow(2, "Delivery2", "9876543210", "7654321", "City2", "Address2", "Region2", "email2"))

	mock.ExpectQuery(`SELECT \* FROM "payments" WHERE \("id" IN \(\$1,\$2\)\)`).WillReturnRows(sqlmock.NewRows([]string{"id", "transaction", "request_id", "currency", "provider", "amount", "payment_date", "bank", "delivery_cost", "goods_total", "custom_fee"}).
		AddRow(1, "Transaction1", "Req1", "USD", "Provider1", 100, 1635237600, "Bank1", 10, 90, 5).
		AddRow(2, "Transaction2", "Req2", "EUR", "Provider2", 200, 1635237600, "Bank2", 20, 180, 10))

	mock.ExpectQuery(`SELECT \* FROM "items" WHERE \(track_number = \$1\)`).
		WithArgs("123").WillReturnRows(sqlmock.NewRows([]string{"chrt_id", "track_number", "price", "r_id", "name", "sale", "size", "total_price", "nm_id", "brand", "status"}).
		AddRow(1, "123", 50, "rid1", "Mobile", 5, "M", 250, 1, "Apple", 1).
		AddRow(2, "123", 60, "rid2", "Notebook", 10, "L", 600, 2, "Samsung", 2))

	mock.ExpectQuery(`SELECT \* FROM "items" WHERE \(track_number = \$1\)`).
		WithArgs("456").WithArgs("456").WillReturnRows(sqlmock.NewRows([]string{"chrt_id", "track_number", "price", "r_id", "name", "sale", "size", "total_price", "nm_id", "brand", "status"}).
		AddRow(3, "456", 70, "rid3", "Cherry", 15, "S", 1050, 3, "Cherry", 3).
		AddRow(4, "456", 80, "rid4", "Racer", 20, "XL", 1600, 4, "Grom", 4))

	orders := orderPostgres.GetAll() //

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	expectedOrders := []core.Order{
		{
			OrderUID:          "order1",
			TrackNumber:       "123",
			Entry:             "entry1",
			DeliveryID:        1,
			PaymentID:         1,
			Locale:            "en",
			InternalSignature: "sig1",
			CustomerID:        "customer1",
			DeliveryService:   "service1",
			ShardKey:          "shard1",
			SMID:              1,
			DateCreated:       "2023-10-26",
			OOFShard:          "oof1",
			Delivery: core.Delivery{
				ID:      1,
				Name:    "Delivery1",
				Phone:   "1234567890",
				Zip:     "1234567",
				City:    "City1",
				Address: "Address1",
				Region:  "Region1",
				Email:   "email1",
			},
			Payment: core.Payment{
				ID:           1,
				Transaction:  "Transaction1",
				RequestID:    "Req1",
				Currency:     "USD",
				Provider:     "Provider1",
				Amount:       100,
				PaymentDate:  1635237600,
				Bank:         "Bank1",
				DeliveryCost: 10,
				GoodsTotal:   90,
				CustomFee:    5,
			},
			Items: []core.Item{
				{
					ChrtID:      1,
					TrackNumber: "123",
					Price:       50,
					RID:         "rid1",
					Name:        "Mobile",
					Sale:        5,
					Size:        "M",
					TotalPrice:  250,
					NmID:        1,
					Brand:       "Apple",
					Status:      1,
				},
				{
					ChrtID:      2,
					TrackNumber: "123",
					Price:       60,
					RID:         "rid2",
					Name:        "Notebook",
					Sale:        10,
					Size:        "L",
					TotalPrice:  600,
					NmID:        2,
					Brand:       "Samsung",
					Status:      2,
				},
			},
		},
		{
			OrderUID:          "order2",
			TrackNumber:       "456",
			Entry:             "entry2",
			DeliveryID:        2,
			PaymentID:         2,
			Locale:            "fr",
			InternalSignature: "sig2",
			CustomerID:        "customer2",
			DeliveryService:   "service2",
			ShardKey:          "shard2",
			SMID:              2,
			DateCreated:       "2023-10-27",
			OOFShard:          "oof2",
			Delivery: core.Delivery{
				ID:      2,
				Name:    "Delivery2",
				Phone:   "9876543210",
				Zip:     "7654321",
				City:    "City2",
				Address: "Address2",
				Region:  "Region2",
				Email:   "email2",
			},
			Payment: core.Payment{
				ID:           2,
				Transaction:  "Transaction2",
				RequestID:    "Req2",
				Currency:     "EUR",
				Provider:     "Provider2",
				Amount:       200,
				PaymentDate:  1635237600,
				Bank:         "Bank2",
				DeliveryCost: 20,
				GoodsTotal:   180,
				CustomFee:    10,
			},
			Items: []core.Item{
				{
					ChrtID:      3,
					TrackNumber: "456",
					Price:       70,
					RID:         "rid3",
					Name:        "Cherry",
					Sale:        15,
					Size:        "S",
					TotalPrice:  1050,
					NmID:        3,
					Brand:       "Cherry",
					Status:      3,
				},
				{
					ChrtID:      4,
					TrackNumber: "456",
					Price:       80,
					RID:         "rid4",
					Name:        "Racer",
					Sale:        20,
					Size:        "XL",
					TotalPrice:  1600,
					NmID:        4,
					Brand:       "Grom",
					Status:      4,
				},
			},
		},
	}

	assert.Equal(t, expectedOrders, orders)
}
