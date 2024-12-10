package repository

import "tugas/model"

func GetOrders() []model.Order {
	return []model.Order{
		{OrderID: "001", Customer: "John Doe", Amount: 100.50, OrderDate: "2024-12-01"},
		{OrderID: "002", Customer: "Jane Smith", Amount: 200.00, OrderDate: "2024-12-02"},
		{OrderID: "003", Customer: "Bob Johnson", Amount: 150.75, OrderDate: "2024-12-03"},
	}
}
