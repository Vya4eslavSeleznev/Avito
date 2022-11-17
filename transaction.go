package avito

type Transaction struct {
	Id        int     `json:"id" db:"id"`
	UserId    int     `json:"user_id" db:"user_id"`
	Type      int     `json:"type" db:"type"`
	Amount    float64 `json:"amount" db:"amount"`
	ServiceId int     `json:"service_id" db:"service_id"`
	OrderId   int     `json:"order_id" db:"order_id"`
	Date      string  `json:"date" db:"date"`
}
