package avito

type Transaction struct {
	Id        int     `json:"id" db:"id"`
	UserId    int     `json:"user_id" db:"user_id" binding:"required"`
	Type      int     `json:"type" db:"type" binding:"required"`
	Amount    float64 `json:"amount" db:"amount" binding:"required"`
	ServiceId int     `json:"service_id" db:"service_id" binding:"required"`
	OrderId   int     `json:"order_id" db:"order_id" binding:"required"`
	Date      string  `json:"date" db:"date"binding:"required"`
}
