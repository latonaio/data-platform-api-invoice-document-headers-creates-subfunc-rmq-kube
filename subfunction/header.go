package subfunction

import (
	api_input_reader "data-platform-api-invoice-document-headers-creates-subfunc-rmq/API_Input_Reader"
	api_processing_data_formatter "data-platform-api-invoice-document-headers-creates-subfunc-rmq/API_Processing_Data_Formatter"
	"strings"
)

func (f *SubFunction) OrdersHeader(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) (*[]api_processing_data_formatter.HeaderOrdersHeader, error) {
	var args []interface{}

	orderID := psdc.OrderID
	repeat := strings.Repeat("?,", len(*orderID)-1) + "?"
	for _, tag := range *orderID {
		args = append(args, tag.OrderID)
	}

	rows, err := f.db.Query(
		`SELECT OrderID, OrderType, Buyer, Seller, ContractType, VaridityStartDate, VaridityEndDate, InvoiceScheduleStartDate,
		InvoiceScheduleEndDate, TotalNetAmount, TotalTaxAmount, TotalGrossAmount, TransactionCurrency, PricingDate, Incoterms,
		BillFromCountry, BillToCountry, Payer, Payee, PaymentTerms, PaymentMethod, IsExportImportDelivery
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_orders_header_data
		WHERE OrderID IN ( `+repeat+` );`, args...,
	)
	if err != nil {
		return nil, err
	}

	data, err := psdc.ConvertToHeaderOrdersHeader(sdc, rows)
	if err != nil {
		return nil, err
	}

	return data, err
}

func (f *SubFunction) HeaderDeliveryDocumentHeader(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) (*[]api_processing_data_formatter.HeaderDeliveryDocumentHeader, error) {
	var args []interface{}

	deliveryDocument := psdc.DeliveryDocument
	repeat := strings.Repeat("?,", len(*deliveryDocument)-1) + "?"
	for _, tag := range *deliveryDocument {
		args = append(args, tag.DeliveryDocument)
	}

	rows, err := f.db.Query(
		`SELECT DeliveryDocument, Buyer, Seller, OrderID, OrderItem, ContractType, OrderValidityStartDate,
		OrderValidityEndDate, InvoiceScheduleStartDate, InvoiceScheduleEndDate, GoodsIssueOrReceiptSlipNumber,
		Incoterms, BillToCountry, BillFromCountry, Payer, Payee, IsExportImportDelivery, TransactionCurrency
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_delivery_document_header_data
		WHERE DeliveryDocument IN ( `+repeat+` );`, args...,
	)
	if err != nil {
		return nil, err
	}

	data, err := psdc.ConvertToHeaderDeliveryDocumentHeader(sdc, rows)
	if err != nil {
		return nil, err
	}

	return data, err
}
