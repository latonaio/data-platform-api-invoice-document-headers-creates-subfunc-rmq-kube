package api_processing_data_formatter

type SDC struct {
	MetaData                            *MetaData                              `json:"MetaData"`
	OrderID                             *[]OrderID                             `json:"OrderID"`
	OrdersHeaderPartner                 *[]OrdersHeaderPartner                 `json:"OrdersHeaderPartner"`
	DeliveryDocument                    *[]DeliveryDocument                    `json:"DeliveryDocument"`
	DeliveryDocumentHeaderPartner       *[]DeliveryDocumentHeaderPartner       `json:"DeliveryDocumentHeaderPartner"`
	CalculateInvoiceDocument            *CalculateInvoiceDocument              `json:"CalculateInvoiceDocument"`
	HeaderOrdersHeader                  *[]HeaderOrdersHeader                  `json:"HeaderOrdersHeader"`
	HeaderDeliveryDocumentHeader        *[]HeaderDeliveryDocumentHeader        `json:"HeaderDeliveryDocumentHeader"`
	HeaderOrdersHeaderPartner           *[]HeaderOrdersHeaderPartner           `json:"HeaderOrdersHeaderPartner"`
	HeaderDeliveryDocumentHeaderPartner *[]HeaderDeliveryDocumentHeaderPartner `json:"HeaderDeliveryDocumentHeaderPartner"`
}

type MetaData struct {
	BusinessPartnerID *int   `json:"business_partner"`
	ServiceLabel      string `json:"service_label"`
}

type OrderIDKey struct {
	OrderID                         *int   `json:"OrderID"`
	BillFromPartyFrom               *int   `json:"BillFromPartyFrom"`
	BillFromPartyTo                 *int   `json:"BillFromPartyTo"`
	BillToPartyFrom                 *int   `json:"BillToPartyFrom"`
	BillToPartyTo                   *int   `json:"BillToPartyTo"`
	BillFromParty                   []*int `json:"BillFromParty"`
	BillToParty                     []*int `json:"BillToParty"`
	HeaderCompleteDeliveryIsDefined *bool  `json:"HeaderCompleteDeliveryIsDefined"`
	HeaderDeliveryStatus            string `json:"HeaderDeliveryStatus"`
	HeaderBillingStatus             string `json:"HeaderBillingStatus"`
	HeaderBillingBlockStatus        *bool  `json:"HeaderBillingBlockStatus"`
}

type OrderID struct {
	InvoiceDocument                 *int   `json:"InvoiceDocument"`
	OrderID                         *int   `json:"OrderID"`
	BillFromPartyFrom               *int   `json:"BillFromPartyFrom"`
	BillFromPartyTo                 *int   `json:"BillFromPartyTo"`
	BillToPartyFrom                 *int   `json:"BillToPartyFrom"`
	BillToPartyTo                   *int   `json:"BillToPartyTo"`
	BillFromParty                   *int   `json:"BillFromParty"`
	BillToParty                     *int   `json:"BillToParty"`
	HeaderCompleteDeliveryIsDefined *bool  `json:"HeaderCompleteDeliveryIsDefined"`
	HeaderDeliveryStatus            string `json:"HeaderDeliveryStatus"`
	HeaderBillingStatus             string `json:"HeaderBillingStatus"`
	HeaderBillingBlockStatus        *bool  `json:"HeaderBillingBlockStatus"`
}

type OrdersHeaderPartner struct {
	InvoiceDocument *int   `json:"InvoiceDocument"`
	OrderID         int    `json:"OrderID"`
	PartnerFunction string `json:"PartnerFunction"`
	BusinessPartner int    `json:"BusinessPartner"`
}

type DeliveryDocumentKey struct {
	DeliveryDocument                *int   `json:"DeliveryDocument"`
	BillFromPartyFrom               *int   `json:"BillFromPartyFrom"`
	BillFromPartyTo                 *int   `json:"BillFromPartyTo"`
	BillToPartyFrom                 *int   `json:"BillToPartyFrom"`
	BillToPartyTo                   *int   `json:"BillToPartyTo"`
	BillFromParty                   []*int `json:"BillFromParty"`
	BillToParty                     []*int `json:"BillToParty"`
	HeaderCompleteDeliveryIsDefined *bool  `json:"HeaderCompleteDeliveryIsDefined"`
	HeaderDeliveryStatus            string `json:"HeaderDeliveryStatus"`
	HeaderBillingStatus             string `json:"HeaderBillingStatus"`
	HeaderBillingBlockStatus        *bool  `json:"HeaderBillingBlockStatus"`
}

type DeliveryDocument struct {
	InvoiceDocument                 *int   `json:"InvoiceDocument"`
	DeliveryDocument                *int   `json:"DeliveryDocument"`
	BillFromPartyFrom               *int   `json:"BillFromPartyFrom"`
	BillFromPartyTo                 *int   `json:"BillFromPartyTo"`
	BillToPartyFrom                 *int   `json:"BillToPartyFrom"`
	BillToPartyTo                   *int   `json:"BillToPartyTo"`
	BillFromParty                   *int   `json:"BillFromParty"`
	BillToParty                     *int   `json:"BillToParty"`
	HeaderCompleteDeliveryIsDefined *bool  `json:"HeaderCompleteDeliveryIsDefined"`
	HeaderDeliveryStatus            string `json:"HeaderDeliveryStatus"`
	HeaderBillingStatus             string `json:"HeaderBillingStatus"`
	HeaderBillingBlockStatus        *bool  `json:"HeaderBillingBlockStatus"`
}

type DeliveryDocumentHeaderPartner struct {
	InvoiceDocument  *int   `json:"InvoiceDocument"`
	DeliveryDocument int    `json:"DeliveryDocument"`
	PartnerFunction  string `json:"PartnerFunction"`
	BusinessPartner  int    `json:"BusinessPartner"`
}

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

type CalculateInvoiceDocumentKey struct {
	ServiceLabel             string `json:"service_label"`
	FieldNameWithNumberRange string `json:"FieldNameWithNumberRange"`
}

type CalculateInvoiceDocumentQueryGets struct {
	ServiceLabel                string `json:"service_label"`
	FieldNameWithNumberRange    string `json:"FieldNameWithNumberRange"`
	InvoiceDocumentLatestNumber *int   `json:"InvoiceDocumentLatestNumber"`
}

type CalculateInvoiceDocument struct {
	InvoiceDocumentLatestNumber *int `json:"InvoiceDocumentLatestNumber"`
	InvoiceDocument             *int `json:"InvoiceDocument"`
}

type HeaderOrdersHeaderPartner struct {
	InvoiceDocument         *int    `json:"InvoiceDocument"`
	OrderID                 *int    `json:"OrderID"`
	PartnerFunction         *string `json:"PartnerFunction"`
	BusinessPartner         *int    `json:"BusinessPartner"`
	BusinessPartnerFullName *string `json:"BusinessPartnerFullName"`
	BusinessPartnerName     *string `json:"BusinessPartnerName"`
	Organization            *string `json:"Organization"`
	Country                 *string `json:"Country"`
	Language                *string `json:"Language"`
	Currency                *string `json:"Currency"`
	ExternalDocumentID      *string `json:"ExternalDocumentID"`
	AddressID               *int    `json:"AddressID"`
}

type HeaderDeliveryDocumentHeaderPartner struct {
	InvoiceDocument         *int    `json:"InvoiceDocument"`
	DeliveryDocument        *int    `json:"DeliveryDocument"`
	PartnerFunction         *string `json:"PartnerFunction"`
	BusinessPartner         *int    `json:"BusinessPartner"`
	BusinessPartnerFullName *string `json:"BusinessPartnerFullName"`
	BusinessPartnerName     *string `json:"BusinessPartnerName"`
	Organization            *string `json:"Organization"`
	Country                 *string `json:"Country"`
	Language                *string `json:"Language"`
	Currency                *string `json:"Currency"`
	ExternalDocumentID      *string `json:"ExternalDocumentID"`
	AddressID               *int    `json:"AddressID"`
}
