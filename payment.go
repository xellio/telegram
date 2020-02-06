package telegram

//
// LabeledPrice - This object represents a portion of the price for goods or services.
// https://core.telegram.org/bots/api#labeledprice
// TODO
//
type LabeledPrice struct {
}

//
// Invoice - This object contains basic information about an invoice.
// https://core.telegram.org/bots/api#invoice
//
type Invoice struct {
	Title          string `json:"title"`
	Description    string `json:"description"`
	StartParameter string `json:"start_parameter"`
	Currency       string `json:"currency"`
	TotalAmount    int    `json:"total_amount"`
}

//
// ShippingAddress - This object represents a shipping address.
// https://core.telegram.org/bots/api#shippingaddress
//
type ShippingAddress struct {
	CountryCode string `json:"country_code"`
	State       string `json:"state"`
	City        string `json:"city"`
	StreetLine1 string `json:"street_line1"`
	StreetLine2 string `json:"street_line2"`
	PostCode    string `json:"post_code"`
}

//
// OrderInfo - This object represents information about an order.
// https://core.telegram.org/bots/api#orderinfo
//
type OrderInfo struct {
	Name            string           `json:"name"`
	PhoneNumber     string           `json:"phone_number"`
	Email           string           `json:"email"`
	ShippingAddress *ShippingAddress `json:"shipping_address"`
}

//
// ShippingOption - This object represents one shipping option.
// https://core.telegram.org/bots/api#shippingoption
// TODO
//
type ShippingOption struct {
}

//
// SuccessfulPayment - This object contains basic information about a successful payment.
// https://core.telegram.org/bots/api#successfulpayment
//
type SuccessfulPayment struct {
	Currency                string     `json:"currency"`
	TotalAmount             int        `json:"total_amount"`
	InvoicePayload          string     `json:"invoice_payload"`
	ShippingOptionID        string     `json:"shipping_option_id"`
	OrderInfo               *OrderInfo `json:"order_info"`
	TelegramPaymentChargeID string     `json:"telegram_payment_charge_id"`
	ProviderPaymentChargeID string     `json:"provider_payment_charge_id"`
}

//
// ShippingQuery - This object contains information about an incoming shipping query.
// https://core.telegram.org/bots/api#shippingquery
//
type ShippingQuery struct {
	ID              string           `json:"id"`
	From            *User            `json:"user"`
	InvoicePayload  string           `json:"invoice_payload"`
	ShippingAddress *ShippingAddress `json:"shipping_adress"`
}

//
// PreCheckoutQuery - This object contains information about an incoming pre-checkout query.
// https://core.telegram.org/bots/api#precheckoutquery
//
type PreCheckoutQuery struct {
	ID               string     `json:"id"`
	From             *User      `json:"from"`
	Currency         string     `json:"currency"`
	TotalAmount      int        `json:"total_amount"`
	InvoicePayload   string     `json:"invoice_payload"`
	ShippingOptionID string     `json:"shipping_option_id"`
	OrderInfo        *OrderInfo `json:"order_info"`
}
