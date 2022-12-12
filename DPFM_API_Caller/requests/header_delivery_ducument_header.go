package requests

type HeaderDeliveryDocumentHeader struct {
	InvoiceDocument               *int    `json:"InvoiceDocument"`
	DeliveryDocument              *int    `json:"DeliveryDocument"`
	Buyer                         *int    `json:"Buyer"`
	Seller                        *int    `json:"Seller"`
	OrderID                       *string `json:"OrderID"`
	OrderItem                     *string `json:"OrderItem"`
	ContractType                  *string `json:"ContractType"`
	OrderValidityStartDate        *string `json:"OrderValidityStartDate"`
	OrderValidityEndDate          *string `json:"OrderValidityEndDate"`
	InvoiceScheduleStartDate      *string `json:"InvoiceScheduleStartDate"`
	InvoiceScheduleEndDate        *string `json:"InvoiceScheduleEndDate"`
	GoodsIssueOrReceiptSlipNumber *string `json:"GoodsIssueOrReceiptSlipNumber"`
	Incoterms                     *string `json:"Incoterms"`
	BillFromCountry               *string `json:"BillFromCountry"`
	BillToCountry                 *string `json:"BillToCountry"`
	Payer                         *int    `json:"Payer"`
	Payee                         *int    `json:"Payee"`
	IsExportImportDelivery        *bool   `json:"IsExportImportDelivery"`
	TransactionCurrency           *string `json:"TransactionCurrency"`
}
