package api_processing_data_formatter

import (
	api_input_reader "data-platform-api-invoice-document-headers-creates-subfunc-rmq/API_Input_Reader"
	"data-platform-api-invoice-document-headers-creates-subfunc-rmq/DPFM_API_Caller/requests"
	"database/sql"
	"fmt"
)

// initializer
func (psdc *SDC) ConvertToMetaData(sdc *api_input_reader.SDC) (*MetaData, error) {
	pm := &requests.MetaData{
		BusinessPartnerID: sdc.BusinessPartnerID,
		ServiceLabel:      sdc.ServiceLabel,
	}

	data := pm
	metaData := MetaData{
		BusinessPartnerID: data.BusinessPartnerID,
		ServiceLabel:      data.ServiceLabel,
	}

	return &metaData, nil
}

func (psdc *SDC) ConvertToOrderIDByNumberSpecificationKey(sdc *api_input_reader.SDC, length int) (*OrderIDKey, error) {
	pm := &requests.OrderIDKey{
		HeaderCompleteDeliveryIsDefined: getBoolPtr(true),
		HeaderDeliveryStatus:            "CL",
		HeaderBillingStatus:             "CL",
		HeaderBillingBlockStatus:        getBoolPtr(false),
	}

	for i := 0; i < length; i++ {
		pm.BillFromParty = append(pm.BillFromParty, nil)
		pm.BillToParty = append(pm.BillToParty, nil)
	}

	data := pm
	orderIDKey := OrderIDKey{
		OrderID:                         data.OrderID,
		BillFromPartyFrom:               data.BillFromPartyFrom,
		BillFromPartyTo:                 data.BillFromPartyTo,
		BillToPartyFrom:                 data.BillToPartyFrom,
		BillToPartyTo:                   data.BillToPartyTo,
		BillFromParty:                   data.BillFromParty,
		BillToParty:                     data.BillToParty,
		HeaderCompleteDeliveryIsDefined: data.HeaderCompleteDeliveryIsDefined,
		HeaderDeliveryStatus:            data.HeaderDeliveryStatus,
		HeaderBillingStatus:             data.HeaderBillingStatus,
		HeaderBillingBlockStatus:        data.HeaderBillingBlockStatus,
	}

	return &orderIDKey, nil
}

func (psdc *SDC) ConvertToOrderIDByNumberSpecification(
	sdc *api_input_reader.SDC,
	rows *sql.Rows,
) (*[]OrderID, error) {
	var orderID []OrderID

	for i := 0; true; i++ {
		pm := &requests.OrderID{}

		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("DBに対象のレコードが存在しません。")
			} else {
				break
			}
		}
		err := rows.Scan(
			&pm.OrderID,
			&pm.BillFromParty,
			&pm.BillToParty,
			&pm.HeaderCompleteDeliveryIsDefined,
			&pm.HeaderDeliveryStatus,
			&pm.HeaderBillingStatus,
			&pm.HeaderBillingBlockStatus,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return nil, err
		}

		data := pm
		orderID = append(orderID, OrderID{
			InvoiceDocument:                 data.InvoiceDocument,
			OrderID:                         data.OrderID,
			BillFromPartyFrom:               data.BillFromPartyFrom,
			BillFromPartyTo:                 data.BillFromPartyTo,
			BillToPartyFrom:                 data.BillToPartyFrom,
			BillToPartyTo:                   data.BillToPartyTo,
			BillFromParty:                   data.BillFromParty,
			BillToParty:                     data.BillToParty,
			HeaderCompleteDeliveryIsDefined: data.HeaderCompleteDeliveryIsDefined,
			HeaderDeliveryStatus:            data.HeaderDeliveryStatus,
			HeaderBillingStatus:             data.HeaderBillingStatus,
			HeaderBillingBlockStatus:        data.HeaderBillingBlockStatus,
		})
	}

	return &orderID, nil
}

func (psdc *SDC) ConvertToOrderIDByRangeSpecificationKey(sdc *api_input_reader.SDC) (*OrderIDKey, error) {
	pm := &requests.OrderIDKey{
		HeaderCompleteDeliveryIsDefined: getBoolPtr(true),
		HeaderDeliveryStatus:            "CL",
		HeaderBillingStatus:             "CL",
		HeaderBillingBlockStatus:        getBoolPtr(false),
	}

	data := pm
	orderIDKey := OrderIDKey{
		OrderID:                         data.OrderID,
		BillFromPartyFrom:               data.BillFromPartyFrom,
		BillFromPartyTo:                 data.BillFromPartyTo,
		BillToPartyFrom:                 data.BillToPartyFrom,
		BillToPartyTo:                   data.BillToPartyTo,
		BillFromParty:                   data.BillFromParty,
		BillToParty:                     data.BillToParty,
		HeaderCompleteDeliveryIsDefined: data.HeaderCompleteDeliveryIsDefined,
		HeaderDeliveryStatus:            data.HeaderDeliveryStatus,
		HeaderBillingStatus:             data.HeaderBillingStatus,
		HeaderBillingBlockStatus:        data.HeaderBillingBlockStatus,
	}

	return &orderIDKey, nil
}

func (psdc *SDC) ConvertToOrderIDByRangeSpecification(
	sdc *api_input_reader.SDC,
	rows *sql.Rows,
) (*[]OrderID, error) {
	var orderID []OrderID

	for i := 0; true; i++ {
		pm := &requests.OrderID{}

		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("DBに対象のレコードが存在しません。")
			} else {
				break
			}
		}
		err := rows.Scan(
			&pm.OrderID,
			&pm.BillFromParty,
			&pm.BillToParty,
			&pm.HeaderCompleteDeliveryIsDefined,
			&pm.HeaderDeliveryStatus,
			&pm.HeaderBillingStatus,
			&pm.HeaderBillingBlockStatus,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return nil, err
		}

		data := pm
		orderID = append(orderID, OrderID{
			InvoiceDocument:                 data.InvoiceDocument,
			OrderID:                         data.OrderID,
			BillFromPartyFrom:               data.BillFromPartyFrom,
			BillFromPartyTo:                 data.BillFromPartyTo,
			BillToPartyFrom:                 data.BillToPartyFrom,
			BillToPartyTo:                   data.BillToPartyTo,
			BillFromParty:                   data.BillFromParty,
			BillToParty:                     data.BillToParty,
			HeaderCompleteDeliveryIsDefined: data.HeaderCompleteDeliveryIsDefined,
			HeaderDeliveryStatus:            data.HeaderDeliveryStatus,
			HeaderBillingStatus:             data.HeaderBillingStatus,
			HeaderBillingBlockStatus:        data.HeaderBillingBlockStatus,
		})
	}

	return &orderID, nil
}

func (psdc *SDC) ConvertToOrderIDByReferenceDocumentKey(sdc *api_input_reader.SDC) (*OrderIDKey, error) {
	pm := &requests.OrderIDKey{
		HeaderCompleteDeliveryIsDefined: getBoolPtr(false),
		HeaderDeliveryStatus:            "CL",
		HeaderBillingBlockStatus:        getBoolPtr(false),
	}

	data := pm
	orderIDKey := OrderIDKey{
		OrderID:                         data.OrderID,
		BillFromParty:                   data.BillFromParty,
		BillToParty:                     data.BillToParty,
		HeaderCompleteDeliveryIsDefined: data.HeaderCompleteDeliveryIsDefined,
		HeaderDeliveryStatus:            data.HeaderDeliveryStatus,
		HeaderBillingStatus:             data.HeaderBillingStatus,
		HeaderBillingBlockStatus:        data.HeaderBillingBlockStatus,
	}

	return &orderIDKey, nil
}

func (psdc *SDC) ConvertToOrderIDByReferenceDocument(
	sdc *api_input_reader.SDC,
	rows *sql.Rows,
) (*[]OrderID, error) {
	var orderID []OrderID

	for i := 0; true; i++ {
		pm := &requests.OrderID{}

		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("DBに対象のレコードが存在しません。")
			} else {
				break
			}
		}
		err := rows.Scan(
			&pm.OrderID,
			&pm.BillFromParty,
			&pm.BillToParty,
			&pm.HeaderCompleteDeliveryIsDefined,
			&pm.HeaderDeliveryStatus,
			&pm.HeaderBillingBlockStatus,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return nil, err
		}

		data := pm
		orderID = append(orderID, OrderID{
			InvoiceDocument:                 data.InvoiceDocument,
			OrderID:                         data.OrderID,
			BillFromParty:                   data.BillFromParty,
			BillToParty:                     data.BillToParty,
			HeaderCompleteDeliveryIsDefined: data.HeaderCompleteDeliveryIsDefined,
			HeaderDeliveryStatus:            data.HeaderDeliveryStatus,
			HeaderBillingStatus:             data.HeaderBillingStatus,
			HeaderBillingBlockStatus:        data.HeaderBillingBlockStatus,
		})
	}

	return &orderID, nil
}

func (psdc *SDC) ConvertToOrdersHeaderPartner(
	sdc *api_input_reader.SDC,
	rows *sql.Rows,
) (*[]OrdersHeaderPartner, error) {
	var ordersHeaderPartner []OrdersHeaderPartner

	for i := 0; true; i++ {
		pm := &requests.OrdersHeaderPartner{}

		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("DBに対象のレコードが存在しません。")
			} else {
				break
			}
		}
		err := rows.Scan(
			&pm.OrderID,
			&pm.PartnerFunction,
			&pm.BusinessPartner,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return nil, err
		}

		data := pm
		ordersHeaderPartner = append(ordersHeaderPartner, OrdersHeaderPartner{
			InvoiceDocument: data.InvoiceDocument,
			OrderID:         data.OrderID,
			PartnerFunction: data.PartnerFunction,
			BusinessPartner: data.BusinessPartner,
		})
	}

	return &ordersHeaderPartner, nil
}

func (psdc *SDC) ConvertToDeliveryDocumentByNumberSpecificationKey(sdc *api_input_reader.SDC, length int) (*DeliveryDocumentKey, error) {
	pm := &requests.DeliveryDocumentKey{
		HeaderCompleteDeliveryIsDefined: getBoolPtr(true),
		HeaderDeliveryStatus:            "CL",
		HeaderBillingStatus:             "CL",
		HeaderBillingBlockStatus:        getBoolPtr(false),
	}

	for i := 0; i < length; i++ {
		pm.BillFromParty = append(pm.BillFromParty, nil)
		pm.BillToParty = append(pm.BillToParty, nil)
	}

	data := pm
	deliveryDocumentKey := DeliveryDocumentKey{
		DeliveryDocument:                data.DeliveryDocument,
		BillFromPartyFrom:               data.BillFromPartyFrom,
		BillFromPartyTo:                 data.BillFromPartyTo,
		BillToPartyFrom:                 data.BillToPartyFrom,
		BillToPartyTo:                   data.BillToPartyTo,
		BillFromParty:                   data.BillFromParty,
		BillToParty:                     data.BillToParty,
		HeaderCompleteDeliveryIsDefined: data.HeaderCompleteDeliveryIsDefined,
		HeaderDeliveryStatus:            data.HeaderDeliveryStatus,
		HeaderBillingStatus:             data.HeaderBillingStatus,
		HeaderBillingBlockStatus:        data.HeaderBillingBlockStatus,
	}

	return &deliveryDocumentKey, nil
}

func (psdc *SDC) ConvertToDeliveryDocumentByNumberSpecification(
	sdc *api_input_reader.SDC,
	rows *sql.Rows,
) (*[]DeliveryDocument, error) {
	var deliveryDocument []DeliveryDocument

	for i := 0; true; i++ {
		pm := &requests.DeliveryDocument{}

		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("DBに対象のレコードが存在しません。")
			} else {
				break
			}
		}
		err := rows.Scan(
			&pm.DeliveryDocument,
			&pm.BillFromParty,
			&pm.BillToParty,
			&pm.HeaderCompleteDeliveryIsDefined,
			&pm.HeaderDeliveryStatus,
			&pm.HeaderBillingStatus,
			&pm.HeaderBillingBlockStatus,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return nil, err
		}

		data := pm
		deliveryDocument = append(deliveryDocument, DeliveryDocument{
			InvoiceDocument:                 data.InvoiceDocument,
			DeliveryDocument:                data.DeliveryDocument,
			BillFromPartyFrom:               data.BillFromPartyFrom,
			BillFromPartyTo:                 data.BillFromPartyTo,
			BillToPartyFrom:                 data.BillToPartyFrom,
			BillToPartyTo:                   data.BillToPartyTo,
			BillFromParty:                   data.BillFromParty,
			BillToParty:                     data.BillToParty,
			HeaderCompleteDeliveryIsDefined: data.HeaderCompleteDeliveryIsDefined,
			HeaderDeliveryStatus:            data.HeaderDeliveryStatus,
			HeaderBillingStatus:             data.HeaderBillingStatus,
			HeaderBillingBlockStatus:        data.HeaderBillingBlockStatus,
		})
	}

	return &deliveryDocument, nil
}

func (psdc *SDC) ConvertToDeliveryDocumentByRangeSpecificationKey(sdc *api_input_reader.SDC) (*DeliveryDocumentKey, error) {
	pm := &requests.DeliveryDocumentKey{
		HeaderCompleteDeliveryIsDefined: getBoolPtr(true),
		HeaderDeliveryStatus:            "CL",
		HeaderBillingStatus:             "CL",
		HeaderBillingBlockStatus:        getBoolPtr(false),
	}

	data := pm
	deliveryDocumentKey := DeliveryDocumentKey{
		DeliveryDocument:                data.DeliveryDocument,
		BillFromPartyFrom:               data.BillFromPartyFrom,
		BillFromPartyTo:                 data.BillFromPartyTo,
		BillToPartyFrom:                 data.BillToPartyFrom,
		BillToPartyTo:                   data.BillToPartyTo,
		BillFromParty:                   data.BillFromParty,
		BillToParty:                     data.BillToParty,
		HeaderCompleteDeliveryIsDefined: data.HeaderCompleteDeliveryIsDefined,
		HeaderDeliveryStatus:            data.HeaderDeliveryStatus,
		HeaderBillingStatus:             data.HeaderBillingStatus,
		HeaderBillingBlockStatus:        data.HeaderBillingBlockStatus,
	}

	return &deliveryDocumentKey, nil
}

func (psdc *SDC) ConvertToDeliveryDocumentByRangeSpecification(
	sdc *api_input_reader.SDC,
	rows *sql.Rows,
) (*[]DeliveryDocument, error) {
	var deliveryDocument []DeliveryDocument

	for i := 0; true; i++ {
		pm := &requests.DeliveryDocument{}

		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("DBに対象のレコードが存在しません。")
			} else {
				break
			}
		}
		err := rows.Scan(
			&pm.DeliveryDocument,
			&pm.BillFromParty,
			&pm.BillToParty,
			&pm.HeaderCompleteDeliveryIsDefined,
			&pm.HeaderDeliveryStatus,
			&pm.HeaderBillingStatus,
			&pm.HeaderBillingBlockStatus,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return nil, err
		}

		data := pm
		deliveryDocument = append(deliveryDocument, DeliveryDocument{
			InvoiceDocument:                 data.InvoiceDocument,
			DeliveryDocument:                data.DeliveryDocument,
			BillFromPartyFrom:               data.BillFromPartyFrom,
			BillFromPartyTo:                 data.BillFromPartyTo,
			BillToPartyFrom:                 data.BillToPartyFrom,
			BillToPartyTo:                   data.BillToPartyTo,
			BillFromParty:                   data.BillFromParty,
			BillToParty:                     data.BillToParty,
			HeaderCompleteDeliveryIsDefined: data.HeaderCompleteDeliveryIsDefined,
			HeaderDeliveryStatus:            data.HeaderDeliveryStatus,
			HeaderBillingStatus:             data.HeaderBillingStatus,
			HeaderBillingBlockStatus:        data.HeaderBillingBlockStatus,
		})
	}

	return &deliveryDocument, nil
}

func (psdc *SDC) ConvertToDeliveryDocumentByReferenceDocumentKey(sdc *api_input_reader.SDC) (*DeliveryDocumentKey, error) {
	pm := &requests.DeliveryDocumentKey{
		HeaderCompleteDeliveryIsDefined: getBoolPtr(true),
		HeaderDeliveryStatus:            "CL",
		HeaderBillingStatus:             "CL",
		HeaderBillingBlockStatus:        getBoolPtr(false),
	}

	data := pm
	deliveryDocumentKey := DeliveryDocumentKey{
		DeliveryDocument:                data.DeliveryDocument,
		BillFromParty:                   data.BillFromParty,
		BillToParty:                     data.BillToParty,
		HeaderCompleteDeliveryIsDefined: data.HeaderCompleteDeliveryIsDefined,
		HeaderDeliveryStatus:            data.HeaderDeliveryStatus,
		HeaderBillingStatus:             data.HeaderBillingStatus,
		HeaderBillingBlockStatus:        data.HeaderBillingBlockStatus,
	}

	return &deliveryDocumentKey, nil
}

func (psdc *SDC) ConvertToDeliveryDocumentByReferenceDocument(
	sdc *api_input_reader.SDC,
	rows *sql.Rows,
) (*[]DeliveryDocument, error) {
	var deliveryDocument []DeliveryDocument

	for i := 0; true; i++ {
		pm := &requests.DeliveryDocument{}

		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("DBに対象のレコードが存在しません。")
			} else {
				break
			}
		}
		err := rows.Scan(
			&pm.DeliveryDocument,
			&pm.BillFromParty,
			&pm.BillToParty,
			&pm.HeaderCompleteDeliveryIsDefined,
			&pm.HeaderDeliveryStatus,
			&pm.HeaderBillingStatus,
			&pm.HeaderBillingBlockStatus,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return nil, err
		}

		data := pm
		deliveryDocument = append(deliveryDocument, DeliveryDocument{
			InvoiceDocument:                 data.InvoiceDocument,
			DeliveryDocument:                data.DeliveryDocument,
			BillFromParty:                   data.BillFromParty,
			BillToParty:                     data.BillToParty,
			HeaderCompleteDeliveryIsDefined: data.HeaderCompleteDeliveryIsDefined,
			HeaderDeliveryStatus:            data.HeaderDeliveryStatus,
			HeaderBillingStatus:             data.HeaderBillingStatus,
			HeaderBillingBlockStatus:        data.HeaderBillingBlockStatus,
		})
	}

	return &deliveryDocument, nil
}

func (psdc *SDC) ConvertToDeliveryDocumentHeaderPartner(
	sdc *api_input_reader.SDC,
	rows *sql.Rows,
) (*[]DeliveryDocumentHeaderPartner, error) {
	var deliveryDocumentHeaderPartner []DeliveryDocumentHeaderPartner

	for i := 0; true; i++ {
		pm := &requests.DeliveryDocumentHeaderPartner{}

		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("DBに対象のレコードが存在しません。")
			} else {
				break
			}
		}
		err := rows.Scan(
			&pm.DeliveryDocument,
			&pm.PartnerFunction,
			&pm.BusinessPartner,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return nil, err
		}

		data := pm
		deliveryDocumentHeaderPartner = append(deliveryDocumentHeaderPartner, DeliveryDocumentHeaderPartner{
			InvoiceDocument:  data.InvoiceDocument,
			DeliveryDocument: data.DeliveryDocument,
			PartnerFunction:  data.PartnerFunction,
			BusinessPartner:  data.BusinessPartner,
		})
	}

	return &deliveryDocumentHeaderPartner, nil
}

// Header
func (psdc *SDC) ConvertToHeaderOrdersHeader(
	sdc *api_input_reader.SDC,
	rows *sql.Rows,
) (*[]HeaderOrdersHeader, error) {
	var headerOrdersHeader []HeaderOrdersHeader
	pm := &requests.HeaderOrdersHeader{}

	for i := 0; true; i++ {
		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("DBに対象のレコードが存在しません。")
			} else {
				break
			}
		}
		err := rows.Scan(
			&pm.OrderID,
			&pm.OrderType,
			&pm.Buyer,
			&pm.Seller,
			&pm.ContractType,
			&pm.VaridityStartDate,
			&pm.VaridityEndDate,
			&pm.InvoiceScheduleStartDate,
			&pm.InvoiceScheduleEndDate,
			&pm.TotalNetAmount,
			&pm.TotalTaxAmount,
			&pm.TotalGrossAmount,
			&pm.TransactionCurrency,
			&pm.PricingDate,
			&pm.Incoterms,
			&pm.BillFromCountry,
			&pm.BillToCountry,
			&pm.Payer,
			&pm.Payee,
			&pm.PaymentTerms,
			&pm.PaymentMethod,
			&pm.IsExportImportDelivery,
		)
		if err != nil {
			return nil, err
		}

		data := pm
		headerOrdersHeader = append(headerOrdersHeader, HeaderOrdersHeader{
			InvoiceDocument:          data.InvoiceDocument,
			OrderID:                  data.OrderID,
			OrderType:                data.OrderType,
			Buyer:                    data.Buyer,
			Seller:                   data.Seller,
			ContractType:             data.ContractType,
			VaridityStartDate:        data.VaridityStartDate,
			VaridityEndDate:          data.VaridityEndDate,
			InvoiceScheduleStartDate: data.InvoiceScheduleStartDate,
			InvoiceScheduleEndDate:   data.InvoiceScheduleEndDate,
			TotalNetAmount:           data.TotalNetAmount,
			TotalTaxAmount:           data.TotalTaxAmount,
			TotalGrossAmount:         data.TotalGrossAmount,
			TransactionCurrency:      data.TransactionCurrency,
			PricingDate:              data.PricingDate,
			Incoterms:                data.Incoterms,
			BillFromCountry:          data.BillFromCountry,
			BillToCountry:            data.BillToCountry,
			Payer:                    data.Payer,
			Payee:                    data.Payee,
			PaymentTerms:             data.PaymentTerms,
			PaymentMethod:            data.PaymentMethod,
			IsExportImportDelivery:   data.IsExportImportDelivery,
		})
	}

	return &headerOrdersHeader, nil
}

func (psdc *SDC) ConvertToCalculateInvoiceDocumentKey() (*CalculateInvoiceDocumentKey, error) {
	pm := &requests.CalculateInvoiceDocumentKey{
		ServiceLabel:             "",
		FieldNameWithNumberRange: "InvoiceDocument",
	}

	data := pm
	calculateInvoiceDocumentKey := CalculateInvoiceDocumentKey{
		ServiceLabel:             data.ServiceLabel,
		FieldNameWithNumberRange: data.FieldNameWithNumberRange,
	}

	return &calculateInvoiceDocumentKey, nil
}

func (psdc *SDC) ConvertToCalculateInvoiceDocumentQueryGets(
	sdc *api_input_reader.SDC,
	rows *sql.Rows,
) (*CalculateInvoiceDocumentQueryGets, error) {
	pm := &requests.CalculateInvoiceDocumentQueryGets{}

	for i := 0; true; i++ {
		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("DBに対象のレコードが存在しません。")
			} else {
				break
			}
		}
		err := rows.Scan(
			&pm.ServiceLabel,
			&pm.FieldNameWithNumberRange,
			&pm.InvoiceDocumentLatestNumber,
		)
		if err != nil {
			return nil, err
		}
	}

	data := pm
	calculateInvoiceDocumentQueryGets := CalculateInvoiceDocumentQueryGets{
		ServiceLabel:                data.ServiceLabel,
		FieldNameWithNumberRange:    data.FieldNameWithNumberRange,
		InvoiceDocumentLatestNumber: data.InvoiceDocumentLatestNumber,
	}

	return &calculateInvoiceDocumentQueryGets, nil
}

func (psdc *SDC) ConvertToCalculateInvoiceDocument(
	invoiceDocumentLatestNumber *int,
) (*CalculateInvoiceDocument, error) {
	pm := &requests.CalculateInvoiceDocument{}

	pm.InvoiceDocumentLatestNumber = invoiceDocumentLatestNumber

	data := pm
	calculateInvoiceDocument := CalculateInvoiceDocument{
		InvoiceDocumentLatestNumber: data.InvoiceDocumentLatestNumber,
		InvoiceDocument:             data.InvoiceDocument,
	}

	return &calculateInvoiceDocument, nil
}

func (psdc *SDC) ConvertToTotalNetAmountQueryGets(
	sdc *api_input_reader.SDC,
	rows *sql.Rows,
) (*TotalNetAmountQueryGets, error) {
	pm := &requests.TotalNetAmountQueryGets{}

	for i := 0; true; i++ {
		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("DBに対象のレコードが存在しません。")
			} else {
				break
			}
		}
		err := rows.Scan(
			&pm.InvoiceDocument,
			&pm.TotalNetAmount,
		)
		if err != nil {
			return nil, err
		}
	}

	data := pm
	totalNetAmountQueryGets := TotalNetAmountQueryGets{
		InvoiceDocument: data.InvoiceDocument,
		TotalNetAmount:  data.TotalNetAmount,
	}

	return &totalNetAmountQueryGets, nil
}

func (psdc *SDC) ConvertToTotalNetAmount(
	inputTotalNetAmount *float32,
) (*TotalNetAmount, error) {
	pm := &requests.TotalNetAmount{}

	pm.TotalNetAmount = inputTotalNetAmount

	data := pm
	totalNetAmount := TotalNetAmount{
		InvoiceDocument: data.InvoiceDocument,
		TotalNetAmount:  data.TotalNetAmount,
	}

	return &totalNetAmount, nil
}

// HeaderPartner
func (psdc *SDC) ConvertToHeaderOrdersHeaderPartner(
	sdc *api_input_reader.SDC,
	rows *sql.Rows,
) (*[]HeaderOrdersHeaderPartner, error) {
	var headerOrdersHeaderPartner []HeaderOrdersHeaderPartner
	pm := &requests.HeaderOrdersHeaderPartner{}

	for i := 0; true; i++ {
		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("DBに対象のレコードが存在しません。")
			} else {
				break
			}
		}
		err := rows.Scan(
			&pm.OrderID,
			&pm.PartnerFunction,
			&pm.BusinessPartner,
			&pm.BusinessPartnerFullName,
			&pm.BusinessPartnerName,
			&pm.Organization,
			&pm.Country,
			&pm.Language,
			&pm.Currency,
			&pm.ExternalDocumentID,
			&pm.AddressID,
		)
		if err != nil {
			return nil, err
		}

		data := pm
		headerOrdersHeaderPartner = append(headerOrdersHeaderPartner, HeaderOrdersHeaderPartner{
			InvoiceDocument:         data.InvoiceDocument,
			OrderID:                 data.OrderID,
			PartnerFunction:         data.PartnerFunction,
			BusinessPartner:         data.BusinessPartner,
			BusinessPartnerFullName: data.BusinessPartnerFullName,
			BusinessPartnerName:     data.BusinessPartnerName,
			Organization:            data.Organization,
			Country:                 data.Country,
			Language:                data.Language,
			Currency:                data.Currency,
			ExternalDocumentID:      data.ExternalDocumentID,
			AddressID:               data.AddressID,
		})
	}

	return &headerOrdersHeaderPartner, nil
}

func (psdc *SDC) ConvertToHeaderDeliveryDocumentHeader(
	sdc *api_input_reader.SDC,
	rows *sql.Rows,
) (*[]HeaderDeliveryDocumentHeader, error) {
	var headerdeliveryDocumentHeader []HeaderDeliveryDocumentHeader
	pm := &requests.HeaderDeliveryDocumentHeader{}

	for i := 0; true; i++ {
		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("DBに対象のレコードが存在しません。")
			} else {
				break
			}
		}
		err := rows.Scan(
			&pm.DeliveryDocument,
			&pm.Buyer,
			&pm.Seller,
			&pm.OrderID,
			&pm.OrderItem,
			&pm.ContractType,
			&pm.OrderValidityStartDate,
			&pm.OrderValidityEndDate,
			&pm.InvoiceScheduleStartDate,
			&pm.InvoiceScheduleEndDate,
			&pm.GoodsIssueOrReceiptSlipNumber,
			&pm.Incoterms,
			&pm.BillToCountry,
			&pm.BillFromCountry,
			&pm.Payer,
			&pm.Payee,
			&pm.IsExportImportDelivery,
			&pm.TransactionCurrency,
		)
		if err != nil {
			return nil, err
		}

		data := pm
		headerdeliveryDocumentHeader = append(headerdeliveryDocumentHeader, HeaderDeliveryDocumentHeader{
			InvoiceDocument:               data.InvoiceDocument,
			DeliveryDocument:              data.DeliveryDocument,
			Buyer:                         data.Buyer,
			Seller:                        data.Seller,
			OrderID:                       data.OrderID,
			OrderItem:                     data.OrderItem,
			ContractType:                  data.ContractType,
			OrderValidityStartDate:        data.OrderValidityStartDate,
			OrderValidityEndDate:          data.OrderValidityEndDate,
			InvoiceScheduleStartDate:      data.InvoiceScheduleStartDate,
			InvoiceScheduleEndDate:        data.InvoiceScheduleEndDate,
			GoodsIssueOrReceiptSlipNumber: data.GoodsIssueOrReceiptSlipNumber,
			Incoterms:                     data.Incoterms,
			BillToCountry:                 data.BillToCountry,
			BillFromCountry:               data.BillFromCountry,
			Payer:                         data.Payer,
			Payee:                         data.Payee,
			IsExportImportDelivery:        data.IsExportImportDelivery,
			TransactionCurrency:           data.TransactionCurrency,
		})
	}

	return &headerdeliveryDocumentHeader, nil
}

func (psdc *SDC) ConvertToHeaderDeliveryDocumentHeaderPartner(
	sdc *api_input_reader.SDC,
	rows *sql.Rows,
) (*[]HeaderDeliveryDocumentHeaderPartner, error) {
	var headerDeliveryDocumentHeaderPartner []HeaderDeliveryDocumentHeaderPartner
	pm := &requests.HeaderDeliveryDocumentHeaderPartner{}

	for i := 0; true; i++ {
		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("DBに対象のレコードが存在しません。")
			} else {
				break
			}
		}
		err := rows.Scan(
			&pm.DeliveryDocument,
			&pm.PartnerFunction,
			&pm.BusinessPartner,
			&pm.BusinessPartnerFullName,
			&pm.BusinessPartnerName,
			&pm.Organization,
			&pm.Country,
			&pm.Language,
			&pm.Currency,
			&pm.ExternalDocumentID,
			&pm.AddressID,
		)
		if err != nil {
			return nil, err
		}

		data := pm
		headerDeliveryDocumentHeaderPartner = append(headerDeliveryDocumentHeaderPartner, HeaderDeliveryDocumentHeaderPartner{
			InvoiceDocument:         data.InvoiceDocument,
			DeliveryDocument:        data.DeliveryDocument,
			PartnerFunction:         data.PartnerFunction,
			BusinessPartner:         data.BusinessPartner,
			BusinessPartnerFullName: data.BusinessPartnerFullName,
			BusinessPartnerName:     data.BusinessPartnerName,
			Organization:            data.Organization,
			Country:                 data.Country,
			Language:                data.Language,
			Currency:                data.Currency,
			ExternalDocumentID:      data.ExternalDocumentID,
			AddressID:               data.AddressID,
		})
	}

	return &headerDeliveryDocumentHeaderPartner, nil
}

func getBoolPtr(b bool) *bool {
	return &b
}
