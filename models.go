package avito

type Accrual struct {
	UserId int     `json:"user_id" db:"user_id"`
	Amount float64 `json:"amount" db:"amount"`
}

type Reservation struct {
	UserId    int     `json:"user_id" db:"user_id"`
	ServiceId int     `json:"service_id" db:"service_id"`
	OrderId   int     `json:"order_id" db:"order_id"`
	Amount    float64 `json:"amount" db:"amount"`
}

type Accounting struct {
	UserId    int     `json:"user_id" db:"user_id"`
	ServiceId int     `json:"service_id" db:"service_id"`
	OrderId   int     `json:"order_id" db:"order_id"`
	Amount    float64 `json:"amount" db:"amount"`
}

type Balance struct {
	UserId int `json:"user_id" db:"user_id"`
}
