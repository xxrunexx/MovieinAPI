package response

import "movie-api/features/paymentmethod"

type RespPaymentmethod struct {
	ID   uint   `json:"id"`
	Name string `jsong:"name"`
}

func ToPaymentmethodResponse(pm paymentmethod.PaymentmethodCore) RespPaymentmethod {
	return RespPaymentmethod{
		ID:   pm.ID,
		Name: pm.Name,
	}
}
