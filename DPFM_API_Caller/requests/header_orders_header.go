package requests

type HeaderOrdersHeader struct {
	InvoiceDocument          *int    `json:"InvoiceDocument"`
	OrderID                  *int    `json:"OrderID"`
	OrderType                string  `json:"OrderType"`
	Buyer                    int     `json:"Buyer"`
	Seller                   int     `json:"Seller"`
	ContractType             *string `json:"ContractType"`
	VaridityStartDate        *string `json:"VaridityStartDate"`
	VaridityEndDate          *string `json:"ValidityEndDate"`
	InvoiceScheduleStartDate *string `json:"InvoiceScheduleStartDate"`
	InvoiceScheduleEndDate   *string `json:"InvoiceScheduleEndDate"`
	TotalNetAmount           float32 `json:"TotalNetAmount"`
	TotalTaxAmount           float32 `json:"TotalTaxAmount"`
	TotalGrossAmount         float32 `json:"TotalGrossAmount"`
	TransactionCurrency      string  `json:"TransactionCurrency"`
	PricingDate              string  `json:"PricingDate"`
	Incoterms                *string `json:"Incoterms"`
	BillFromCountry          *string `json:"BillFromCountry"`
	BillToCountry            *string `json:"BillToCountry"`
	Payer                    *int    `json:"Payer"`
	Payee                    *int    `json:"Payee"`
	PaymentTerms             *string `json:"PaymentTerms"`
	PaymentMethod            *string `json:"PaymentMethod"`
	IsExportImportDelivery   *bool   `json:"IsExportImportDelivery"`
}
