package models

type Transaction struct {
	ID int64
	Date string
	Time string
	OperationAmount int64
	AccountNumber int64
	ReceiverAccountNumber int64
	AvailableLimit int64
}