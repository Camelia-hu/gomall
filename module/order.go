package module

type Order struct {
	Uid          uint32  `json:"uid"`
	UserCurrency string  `json:"userCurrency"`
	Address      string  `json:"address"`
	Email        string  `json:"email"`
	UUid         string  `json:"UUid"`
	Cost         float32 `json:"cost"`
	ProductId    uint32  `json:"productId"`
	Quantity     int32   `json:"quantity"`
}
