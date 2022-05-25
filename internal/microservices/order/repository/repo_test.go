package repository

import (
	"reflect"
	"testing"

	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/microservices/order/models"
	"github.com/jmoiron/sqlx"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestOrderRepo_Order(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")

	repo := OrderRepo{
		DB: sqlxDB,
	}

	testItem := &models.CreateOrderRepoReq{
		Address: "Москва, Измайловский проспект, 73А",
		Comment: "шаурму без лука",
		Cart:    []models.OrderPositionRepo{{Id: 1, Count: 2}, {Id: 2, Count: 1}},
	}

	rows := sqlmock.
		NewRows([]string{"order_id"})
	expect := []*models.CreateOrderRepoResp{
		{OrderId: 1},
	}
	for _, item := range expect {
		rows = rows.AddRow(item.OrderId)
	}

	// good query
	mock.
		ExpectQuery(`INSERT ...`).
		WithArgs(testItem.UserId, testItem.Address, testItem.Comment, testItem.Promocode, testItem.Cart[0].Id, testItem.Cart[0].Count, 0, testItem.Cart[1].Id, testItem.Cart[1].Count, 0).
		WillReturnRows(rows)
	item, err := repo.CreateOrder(testItem)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if item == nil {
		t.Errorf("bad id: want %v, have %v", item, 0)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
	if !reflect.DeepEqual(item, expect[0]) {
		t.Errorf("results not match, want %v, have %v", item, expect[0])
		return
	}
}

func TestOrderRepo_GetOrders(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")

	repo := NewOrderRepo(sqlxDB)

	testItem := &models.GetUserOrdersRepoReq{
		UserId: 1,
	}

	rows := sqlmock.
		NewRows([]string{"id", "date", "restaurant_name", "total_price_discount", "status", "delivery_price"})
	expect := []*models.ShortOrderRepo{
		{OrderId: 1, RestaurantName: "Name1", Date: "01.04.2001", TotalPrice: 1, Status: "Получен", DeliveryPrice: 1},
		{OrderId: 2, RestaurantName: "Name2", Date: "02.04.2001", TotalPrice: 2, Status: "Получен", DeliveryPrice: 2},
		{OrderId: 3, RestaurantName: "Name3", Date: "03.04.2001", TotalPrice: 3, Status: "Получен", DeliveryPrice: 3},
	}
	for _, item := range expect {
		rows = rows.AddRow(item.OrderId, item.Date, item.RestaurantName, item.TotalPrice, item.Status, item.DeliveryPrice)
	}

	// good query
	mock.
		ExpectQuery(`SELECT id, date, restaurant_name, total_price_discount, status, delivery_price FROM orders WHERE user_id = `).
		WithArgs(testItem.UserId).
		WillReturnRows(rows)
	ordersResp, err := repo.GetUserOrders(testItem)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if ordersResp == nil {
		t.Errorf("bad id: want %v, have %v", ordersResp, 0)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
	expectResp := &models.GetUserOrdersRepoResp{Orders: []models.ShortOrderRepo{*expect[0], *expect[1], *expect[2]}}
	if !reflect.DeepEqual(ordersResp, expectResp) {
		t.Errorf("results not match, want %v, have %v", ordersResp, expectResp)
		return
	}
}

func TestOrderRepo_GetOrderStatuses(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")

	repo := NewOrderRepo(sqlxDB)

	testItem := &models.GetUserOrderStatusesRepoReq{
		UserId: 1,
	}

	rows := sqlmock.
		NewRows([]string{"id", "status"})
	expect := []*models.OrderStatusRepo{
		{OrderId: 1, Status: "Получен"},
		{OrderId: 2, Status: "Получен"},
		{OrderId: 3, Status: "Получен"},
	}
	for _, item := range expect {
		rows = rows.AddRow(item.OrderId, item.Status)
	}

	// good query
	mock.
		ExpectQuery(`SELECT id, status FROM orders ... `).
		WithArgs(testItem.UserId).
		WillReturnRows(rows)
	ordersResp, err := repo.GetUserOrderStatuses(testItem)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if ordersResp == nil {
		t.Errorf("bad id: want %v, have %v", ordersResp, 0)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
	expectResp := &models.GetUserOrderStatusesRepoResp{OrderStatuses: []models.OrderStatusRepo{*expect[0], *expect[1], *expect[2]}}
	if !reflect.DeepEqual(ordersResp, expectResp) {
		t.Errorf("results not match, want %v, have %v", ordersResp, expectResp)
		return
	}
}

func TestOrderRepo_GetOrder(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")

	repo := NewOrderRepo(sqlxDB)

	testItem := &models.GetUserOrderRepoReq{
		OrderId: 1,
	}

	rows := sqlmock.
		NewRows([]string{"id", "address", "user_id", "date", "restaurant_name", "restaurant_slug", "total_price", "total_price_discount", "delivery_price", "status"})
	expect := []*models.GetUserOrderRepoResp{
		{OrderId: 1, Status: "Получен", Address: "Москваб Тверская улица, 2", UserId: 1, Date: "04.04.2001", RestaurantName: "Name", RestaurantSlug: " slug", TotalPrice: 10, TotalPriceDiscount: 1, DeliveryPrice: 2},
	}
	for _, item := range expect {
		rows = rows.AddRow(item.OrderId, item.Address, item.UserId, item.Date, item.RestaurantName, item.RestaurantSlug, item.TotalPrice, item.TotalPriceDiscount, item.DeliveryPrice, item.Status)
	}

	// good query
	mock.
		ExpectQuery(`SELECT id, address, user_id, date, restaurant_name, restaurant_slug, total_price, total_price_discount, delivery_price, status FROM orders WHERE ... `).
		WithArgs(testItem.OrderId).
		WillReturnRows(rows)

	rowsCart := sqlmock.
		NewRows([]string{"description", "name", "weight", "calories", "image_path", "price", "count"})
	expectCart := []*models.OrderPositionRepoResp{
		{Name: "name",
			ImagePath:   "img.png",
			Price:       10,
			Count:       10,
			Calories:    100,
			Weight:      50,
			Description: "description"}}

	for _, item := range expectCart {
		rowsCart = rowsCart.AddRow(item.Description, item.Name, item.Weight, item.Calories, item.ImagePath, item.Price, item.Count)
	}
	mock.
		ExpectQuery(`SELECT d.description, ... `).
		WithArgs(testItem.OrderId).
		WillReturnRows(rowsCart)
	ordersResp, err := repo.GetUserOrder(testItem)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if ordersResp == nil {
		t.Errorf("bad id: want %v, have %v", ordersResp, 0)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
	expectResp := expect[0]
	expectResp.Cart = append(expectResp.Cart, *expectCart[0])
	if !reflect.DeepEqual(ordersResp, expectResp) {
		t.Errorf("results not match, want %v, have %v", ordersResp, expectResp)
		return
	}
}

func TestOrderRepo_GetAddress(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")

	repo := NewOrderRepo(sqlxDB)

	testItem := &models.GetAddressRepoReq{
		City:   "Москва",
		Street: "Петровка",
		House:  "38",
	}

	cityRows := sqlmock.
		NewRows([]string{"id", "name"})
	expectCity := []struct {
		Name string
		Id   int
	}{
		{Id: 1, Name: "Москва"},
	}
	for _, item := range expectCity {
		cityRows = cityRows.AddRow(item.Id, item.Name)
	}

	streetRows := sqlmock.
		NewRows([]string{"id", "name"})
	expectStreet := []struct {
		Name string
		Id   int
	}{
		{Id: 1, Name: "Петровка"},
	}
	for _, item := range expectStreet {
		streetRows = streetRows.AddRow(item.Id, item.Name)
	}
	houseRows := sqlmock.
		NewRows([]string{"house"})
	expectHouse := []struct {
		House string
	}{
		{House: "38"},
	}
	for _, item := range expectHouse {
		houseRows = houseRows.AddRow(item.House)
	}

	// good query
	mock.
		ExpectQuery(`SELECT id, name FROM cities `).
		WithArgs(testItem.City).
		WillReturnRows(cityRows)
	mock.
		ExpectQuery(`SELECT id, name FROM streets `).
		WithArgs(1, testItem.Street, "").
		WillReturnRows(streetRows)
	mock.
		ExpectQuery(`SELECT house FROM houses `).
		WithArgs(1, testItem.House).
		WillReturnRows(houseRows)

	ordersResp, err := repo.GetAddress(testItem)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if ordersResp == nil {
		t.Errorf("bad id: want %v, have %v", ordersResp, 0)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
	expect := models.GetAddressRepoResp{
		City:   "Москва",
		Street: "Петровка",
		House:  "38",
	}
	if !reflect.DeepEqual(ordersResp, &expect) {
		t.Errorf("results not match, want %v, have %v", ordersResp, expect)
		return
	}
}
