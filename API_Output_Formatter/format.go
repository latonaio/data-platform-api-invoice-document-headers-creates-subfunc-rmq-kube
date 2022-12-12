package dpfm_api_output_formatter

import (
	api_input_reader "data-platform-api-invoice-document-headers-creates-subfunc-rmq/API_Input_Reader"
	api_processing_data_formatter "data-platform-api-invoice-document-headers-creates-subfunc-rmq/API_Processing_Data_Formatter"
	"encoding/json"
)

func ConvertToHeader(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) (*[]Header, error) {
	orderID := psdc.OrderID
	calculateInvoiceDocument := psdc.CalculateInvoiceDocument
	headerOrdersHeader := psdc.HeaderOrdersHeader
	headers := make([]Header, 0, len(*headerOrdersHeader))

	orderIDMap := make(map[int]api_processing_data_formatter.OrderID, len(*orderID))

	for _, v := range *orderID {
		orderIDMap[*v.OrderID] = v
	}

	for _, v := range *headerOrdersHeader {
		header := Header{}
		inputHeader := sdc.InvoiceDocument
		inputData, err := json.Marshal(inputHeader)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(inputData, &header)
		if err != nil {
			return nil, err
		}

		data, err := json.Marshal(v)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(data, &header)
		if err != nil {
			return nil, err
		}

		header.InvoiceDocument = calculateInvoiceDocument.InvoiceDocumentLatestNumber
		header.BillFromParty = orderIDMap[*v.OrderID].BillFromParty
		header.BillToParty = orderIDMap[*v.OrderID].BillToParty
		headers = append(headers, header)
	}

	return &headers, nil
}

func ConvertToHeaderPartner(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) (*[]HeaderPartner, error) {
	calculateInvoiceDocument := psdc.CalculateInvoiceDocument
	headerOrdersHeaderPartner := psdc.HeaderOrdersHeaderPartner
	headerPartners := make([]HeaderPartner, 0, len(*headerOrdersHeaderPartner))

	for _, v := range *headerOrdersHeaderPartner {
		headerPartner := HeaderPartner{}

		data, err := json.Marshal(v)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(data, &headerPartner)
		if err != nil {
			return nil, err
		}

		headerPartner.InvoiceDocument = calculateInvoiceDocument.InvoiceDocumentLatestNumber
		headerPartners = append(headerPartners, headerPartner)
	}

	return &headerPartners, nil
}
