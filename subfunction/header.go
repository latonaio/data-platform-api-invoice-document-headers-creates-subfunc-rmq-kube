package subfunction

import (
	api_input_reader "data-platform-api-invoice-document-headers-creates-subfunc-rmq/API_Input_Reader"
	api_processing_data_formatter "data-platform-api-invoice-document-headers-creates-subfunc-rmq/API_Processing_Data_Formatter"
	"strings"

	"golang.org/x/xerrors"
)

func (f *SubFunction) CalculateInvoiceDocument(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) (*api_processing_data_formatter.CalculateInvoiceDocument, error) {
	metaData := psdc.MetaData
	dataKey, err := psdc.ConvertToCalculateInvoiceDocumentKey()
	if err != nil {
		return nil, err
	}

	dataKey.ServiceLabel = metaData.ServiceLabel

	rows, err := f.db.Query(
		`SELECT ServiceLabel, FieldNameWithNumberRange, LatestNumber
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_number_range_latest_number_data
		WHERE (ServiceLabel, FieldNameWithNumberRange) = (?, ?);`, dataKey.ServiceLabel, dataKey.FieldNameWithNumberRange,
	)
	if err != nil {
		return nil, err
	}

	dataQueryGets, err := psdc.ConvertToCalculateInvoiceDocumentQueryGets(sdc, rows)
	if err != nil {
		return nil, err
	}

	calculateInvoiceDocument := CalculateInvoiceDocument(*dataQueryGets.InvoiceDocumentLatestNumber)

	data, err := psdc.ConvertToCalculateInvoiceDocument(calculateInvoiceDocument)
	if err != nil {
		return nil, err
	}

	return data, err
}

func CalculateInvoiceDocument(latestNumber int) *int {
	res := latestNumber + 1
	return &res
}

func (f *SubFunction) HeaderOrdersHeader(
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

func (f *SubFunction) TotalNetAmount(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) (*api_processing_data_formatter.TotalNetAmount, error) {
	var err error
	var data *api_processing_data_formatter.TotalNetAmount

	// オーダー参照
	// TODO: nullの場合どうする？
	if sdc.InvoiceDocument.TotalNetAmount != nil {
		for i, v := range *psdc.HeaderOrdersHeader {
			if v.TotalNetAmount == *sdc.InvoiceDocument.TotalNetAmount {
				data, err = psdc.ConvertToTotalNetAmount(&v.TotalNetAmount)
				break
			}
			if i == len(*psdc.HeaderOrdersHeader)-1 {
				return nil, xerrors.Errorf("TotalNetAmountが一致しません。")
			}
		}
	}

	// // 入出荷伝票参照
	// rows, err := f.db.Query(
	// 	`SELECT InvoiceDocument, TotalNetAmount
	// 	FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_invoice_document_header_data
	// 	WHERE InvoiceDocument = ?;`, sdc.InvoiceDocument.InvoiceDocument,
	// )
	// if err != nil {
	// 	return nil, err
	// }

	// dataQueryGets, err := psdc.ConvertToTotalNetAmountQueryGets(sdc, rows)
	// if err != nil {
	// 	return nil, err
	// }

	// if sdc.InvoiceDocument.TotalNetAmount != nil {
	// 	if *dataQueryGets.TotalNetAmount == *sdc.InvoiceDocument.TotalNetAmount {
	// 		data, err = psdc.ConvertToTotalNetAmount(dataQueryGets.TotalNetAmount)
	// 	} else {
	// 		return nil, xerrors.Errorf("TotalNetAmountが一致しません。")
	// 	}
	// }

	return data, err
}
