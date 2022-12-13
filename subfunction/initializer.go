package subfunction

import (
	"context"
	api_input_reader "data-platform-api-invoice-document-headers-creates-subfunc-rmq/API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-invoice-document-headers-creates-subfunc-rmq/API_Output_Formatter"
	api_processing_data_formatter "data-platform-api-invoice-document-headers-creates-subfunc-rmq/API_Processing_Data_Formatter"
	"log"
	"strings"

	"sync"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	database "github.com/latonaio/golang-mysql-network-connector"
	"golang.org/x/xerrors"
)

type SubFunction struct {
	ctx context.Context
	db  *database.Mysql
	l   *logger.Logger
}

func NewSubFunction(ctx context.Context, db *database.Mysql, l *logger.Logger) *SubFunction {
	return &SubFunction{
		ctx: ctx,
		db:  db,
		l:   l,
	}
}

func (f *SubFunction) MetaData(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) (*api_processing_data_formatter.MetaData, error) {
	var err error
	var metaData *api_processing_data_formatter.MetaData

	metaData, err = psdc.ConvertToMetaData(sdc)
	if err != nil {
		return nil, err
	}

	return metaData, nil
}

func (f *SubFunction) OrderIDByNumberSpecification(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) (*[]api_processing_data_formatter.OrderID, error) {
	var args []interface{}

	billFromParty := sdc.InvoiceDocumentInputParameters.BillFromParty
	billToParty := sdc.InvoiceDocumentInputParameters.BillToParty

	if len(*billFromParty) != len(*billToParty) {
		return nil, nil
	}

	dataKey, err := psdc.ConvertToOrderIDByNumberSpecificationKey(sdc, len(*billFromParty))
	if err != nil {
		return nil, err
	}

	for i := range *billFromParty {
		dataKey.BillFromParty[i] = (*billFromParty)[i]
		dataKey.BillToParty[i] = (*billToParty)[i]
	}

	repeat := strings.Repeat("(?,?),", len(dataKey.BillFromParty)-1) + "(?,?)"
	for i := range dataKey.BillFromParty {
		args = append(args, dataKey.BillFromParty[i], dataKey.BillToParty[i])
	}

	args = append(
		args,
		dataKey.HeaderCompleteDeliveryIsDefined,
		dataKey.HeaderDeliveryStatus,
		dataKey.HeaderBillingBlockStatus,
		dataKey.HeaderBillingStatus,
	)

	var count *int
	err = f.db.QueryRow(
		`SELECT COUNT(*)
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_orders_header_data
		WHERE (BillFromParty, BillToParty) IN ( `+repeat+` )
		AND (HeaderCompleteDeliveryIsDefined, HeaderDeliveryStatus, HeaderBillingBlockStatus) = (?, ?, ?)
		AND HeaderBillingStatus <> ?;`, args...,
	).Scan(&count)
	if err != nil {
		log.Println(err)
	}
	if *count == 0 || *count > 1000 {
		return nil, xerrors.Errorf("OrderIDの検索結果がゼロ件または1,000件超です。")
	}

	rows, err := f.db.Query(
		`SELECT OrderID, BillFromParty, BillToParty, HeaderCompleteDeliveryIsDefined, HeaderDeliveryStatus, HeaderBillingStatus, HeaderBillingBlockStatus
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_orders_header_data
		WHERE (BillFromParty, BillToParty) IN ( `+repeat+` )
		AND (HeaderCompleteDeliveryIsDefined, HeaderDeliveryStatus, HeaderBillingBlockStatus) = (?, ?, ?)
		AND HeaderBillingStatus <> ?;`, args...,
	)
	if err != nil {
		return nil, err
	}

	data, err := psdc.ConvertToOrderIDByNumberSpecification(sdc, rows)
	if err != nil {
		return nil, err
	}

	return data, err
}

func (f *SubFunction) OrderIDByRangeSpecification(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) (*[]api_processing_data_formatter.OrderID, error) {
	dataKey, err := psdc.ConvertToOrderIDByRangeSpecificationKey(sdc)
	if err != nil {
		return nil, err
	}

	dataKey.BillFromPartyFrom = sdc.InvoiceDocumentInputParameters.BillFromPartyFrom
	dataKey.BillFromPartyTo = sdc.InvoiceDocumentInputParameters.BillFromPartyTo
	dataKey.BillToPartyFrom = sdc.InvoiceDocumentInputParameters.BillToPartyFrom
	dataKey.BillToPartyTo = sdc.InvoiceDocumentInputParameters.BillToPartyTo

	var count *int
	err = f.db.QueryRow(
		`SELECT COUNT(*)
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_orders_header_data
		WHERE BillFromParty BETWEEN ? AND ?
		AND BillToParty BETWEEN ? AND ?
		AND (HeaderCompleteDeliveryIsDefined, HeaderDeliveryStatus, HeaderBillingBlockStatus) = (?, ?, ?)
		AND HeaderBillingStatus <> ?;`, dataKey.BillFromPartyFrom, dataKey.BillFromPartyTo, dataKey.BillToPartyFrom, dataKey.BillToPartyTo, dataKey.HeaderCompleteDeliveryIsDefined, dataKey.HeaderDeliveryStatus, dataKey.HeaderBillingBlockStatus, dataKey.HeaderBillingStatus,
	).Scan(&count)
	if err != nil {
		log.Println(err)
	}
	if *count == 0 || *count > 1000 {
		return nil, xerrors.Errorf("OrderIDの検索結果がゼロ件または1,000件超です。")
	}

	rows, err := f.db.Query(
		`SELECT OrderID, BillFromParty, BillToParty, HeaderCompleteDeliveryIsDefined, HeaderDeliveryStatus, HeaderBillingStatus, HeaderBillingBlockStatus
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_orders_header_data
		WHERE BillFromParty BETWEEN ? AND ?
		AND BillToParty BETWEEN ? AND ?
		AND (HeaderCompleteDeliveryIsDefined, HeaderDeliveryStatus, HeaderBillingBlockStatus) = (?, ?, ?)
		AND HeaderBillingStatus <> ?;`, dataKey.BillFromPartyFrom, dataKey.BillFromPartyTo, dataKey.BillToPartyFrom, dataKey.BillToPartyTo, dataKey.HeaderCompleteDeliveryIsDefined, dataKey.HeaderDeliveryStatus, dataKey.HeaderBillingBlockStatus, dataKey.HeaderBillingStatus,
	)
	if err != nil {
		return nil, err
	}

	data, err := psdc.ConvertToOrderIDByRangeSpecification(sdc, rows)
	if err != nil {
		return nil, err
	}

	return data, err
}

func (f *SubFunction) OrderIDByReferenceDocument(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) (*[]api_processing_data_formatter.OrderID, error) {
	dataKey, err := psdc.ConvertToOrderIDByReferenceDocumentKey(sdc)
	if err != nil {
		return nil, err
	}

	dataKey.OrderID = sdc.InvoiceDocumentInputParameters.ReferenceDocument
	dataKey.BillFromParty = append(dataKey.BillFromParty, (*sdc.InvoiceDocumentInputParameters.BillFromParty)[0])
	dataKey.BillToParty = append(dataKey.BillToParty, (*sdc.InvoiceDocumentInputParameters.BillToParty)[0])

	rows, err := f.db.Query(
		`SELECT OrderID, BillFromParty, BillToParty, HeaderCompleteDeliveryIsDefined, HeaderDeliveryStatus, HeaderBillingBlockStatus
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_orders_header_data
		WHERE (OrderID, BillFromParty, BillToParty, HeaderCompleteDeliveryIsDefined, HeaderBillingBlockStatus) = (?, ?, ?, ?, ?)
		AND HeaderDeliveryStatus <> ?;`, dataKey.OrderID, dataKey.BillFromParty[0], dataKey.BillToParty[0], dataKey.HeaderCompleteDeliveryIsDefined, dataKey.HeaderBillingBlockStatus, dataKey.HeaderDeliveryStatus,
	)
	if err != nil {
		return nil, err
	}

	data, err := psdc.ConvertToOrderIDByReferenceDocument(sdc, rows)
	if err != nil {
		return nil, err
	}

	return data, err
}

func (f *SubFunction) OrdersHeaderPartner(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) (*[]api_processing_data_formatter.OrdersHeaderPartner, error) {
	var args []interface{}

	orderID := psdc.OrderID

	repeat := strings.Repeat("?,", len(*orderID)-1) + "?"
	for _, tag := range *orderID {
		args = append(args, tag.OrderID)
	}

	rows, err := f.db.Query(
		`SELECT OrderID, PartnerFunction, BusinessPartner
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_orders_header_partner_data
		WHERE OrderID IN ( `+repeat+` );`, args...,
	)
	if err != nil {
		return nil, err
	}

	data, err := psdc.ConvertToOrdersHeaderPartner(sdc, rows)
	if err != nil {
		return nil, err
	}

	return data, err
}

func (f *SubFunction) DeliveryDocumentByNumberSpecification(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) (*[]api_processing_data_formatter.DeliveryDocument, error) {
	var args []interface{}

	billFromParty := sdc.InvoiceDocumentInputParameters.BillFromParty
	billToParty := sdc.InvoiceDocumentInputParameters.BillToParty

	if len(*billFromParty) != len(*billToParty) {
		return nil, nil
	}

	dataKey, err := psdc.ConvertToDeliveryDocumentByNumberSpecificationKey(sdc, len(*billFromParty))
	if err != nil {
		return nil, err
	}

	for i := range *billFromParty {
		dataKey.BillFromParty[i] = (*billFromParty)[i]
		dataKey.BillToParty[i] = (*billToParty)[i]
	}

	repeat := strings.Repeat("(?,?),", len(dataKey.BillFromParty)-1) + "(?,?)"
	for i := range dataKey.BillFromParty {
		args = append(args, dataKey.BillFromParty[i], dataKey.BillToParty[i])
	}

	args = append(
		args,
		dataKey.HeaderCompleteDeliveryIsDefined,
		dataKey.HeaderDeliveryStatus,
		dataKey.HeaderBillingBlockStatus,
		dataKey.HeaderBillingStatus,
	)

	var count *int
	err = f.db.QueryRow(
		`SELECT COUNT(*)
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_delivery_document_header_data
		WHERE (BillFromParty, BillToParty) IN ( `+repeat+` )
		AND (HeaderCompleteDeliveryIsDefined, HeaderDeliveryStatus, HeaderBillingBlockStatus) = (?, ?, ?)
		AND HeaderBillingStatus <> ?;`, args...,
	).Scan(&count)
	if err != nil {
		log.Println(err)
	}
	if *count == 0 || *count > 1000 {
		return nil, xerrors.Errorf("OrderIDの検索結果がゼロ件または1,000件超です。")
	}

	rows, err := f.db.Query(
		`SELECT DeliveryDocument, BillFromParty, BillToParty, HeaderCompleteDeliveryIsDefined, HeaderDeliveryStatus, HeaderBillingStatus, HeaderBillingBlockStatus
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_delivery_document_header_data
		WHERE (BillFromParty, BillToParty) IN ( `+repeat+` )
		AND (HeaderCompleteDeliveryIsDefined, HeaderDeliveryStatus, HeaderBillingBlockStatus) = (?, ?, ?)
		AND HeaderBillingStatus <> ?;`, args...,
	)
	if err != nil {
		return nil, err
	}

	data, err := psdc.ConvertToDeliveryDocumentByNumberSpecification(sdc, rows)
	if err != nil {
		return nil, err
	}

	return data, err
}

func (f *SubFunction) DeliveryDocumentByRangeSpecification(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) (*[]api_processing_data_formatter.DeliveryDocument, error) {
	dataKey, err := psdc.ConvertToDeliveryDocumentByRangeSpecificationKey(sdc)
	if err != nil {
		return nil, err
	}

	dataKey.BillFromPartyFrom = sdc.InvoiceDocumentInputParameters.BillFromPartyFrom
	dataKey.BillFromPartyTo = sdc.InvoiceDocumentInputParameters.BillFromPartyTo
	dataKey.BillToPartyFrom = sdc.InvoiceDocumentInputParameters.BillToPartyFrom
	dataKey.BillToPartyTo = sdc.InvoiceDocumentInputParameters.BillToPartyTo

	var count *int
	err = f.db.QueryRow(
		`SELECT COUNT(*)
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_delivery_document_header_data
		WHERE BillFromParty BETWEEN ? AND ?
		AND BillToParty BETWEEN ? AND ?
		AND (HeaderCompleteDeliveryIsDefined, HeaderDeliveryStatus, HeaderBillingBlockStatus) = (?, ?, ?)
		AND HeaderBillingStatus <> ?;`, dataKey.BillFromPartyFrom, dataKey.BillFromPartyTo, dataKey.BillToPartyFrom, dataKey.BillToPartyTo, dataKey.HeaderCompleteDeliveryIsDefined, dataKey.HeaderDeliveryStatus, dataKey.HeaderBillingBlockStatus, dataKey.HeaderBillingStatus,
	).Scan(&count)
	if err != nil {
		log.Println(err)
	}
	if *count == 0 || *count > 1000 {
		return nil, xerrors.Errorf("OrderIDの検索結果がゼロ件または1,000件超です。")
	}

	rows, err := f.db.Query(
		`SELECT DeliveryDocument, BillFromParty, BillToParty, HeaderCompleteDeliveryIsDefined, HeaderDeliveryStatus, HeaderBillingStatus, HeaderBillingBlockStatus
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_delivery_document_header_data
		WHERE BillFromParty BETWEEN ? AND ?
		AND BillToParty BETWEEN ? AND ?
		AND (HeaderCompleteDeliveryIsDefined, HeaderDeliveryStatus, HeaderBillingBlockStatus) = (?, ?, ?)
		AND HeaderBillingStatus <> ?;`, dataKey.BillFromPartyFrom, dataKey.BillFromPartyTo, dataKey.BillToPartyFrom, dataKey.BillToPartyTo, dataKey.HeaderCompleteDeliveryIsDefined, dataKey.HeaderDeliveryStatus, dataKey.HeaderBillingBlockStatus, dataKey.HeaderBillingStatus,
	)
	if err != nil {
		return nil, err
	}

	data, err := psdc.ConvertToDeliveryDocumentByRangeSpecification(sdc, rows)
	if err != nil {
		return nil, err
	}

	return data, err
}

func (f *SubFunction) DeliveryDocumentByReferenceDocument(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) (*[]api_processing_data_formatter.DeliveryDocument, error) {
	dataKey, err := psdc.ConvertToDeliveryDocumentByReferenceDocumentKey(sdc)
	if err != nil {
		return nil, err
	}

	dataKey.DeliveryDocument = sdc.InvoiceDocumentInputParameters.ReferenceDocument
	dataKey.BillFromParty = append(dataKey.BillFromParty, (*sdc.InvoiceDocumentInputParameters.BillFromParty)[0])
	dataKey.BillToParty = append(dataKey.BillToParty, (*sdc.InvoiceDocumentInputParameters.BillToParty)[0])

	rows, err := f.db.Query(
		`SELECT DeliveryDocument, BillFromParty, BillToParty, HeaderCompleteDeliveryIsDefined, HeaderDeliveryStatus, HeaderBillingStatus, HeaderBillingBlockStatus
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_delivery_document_header_data
		WHERE (DeliveryDocument, BillFromParty, BillToParty, HeaderCompleteDeliveryIsDefined, HeaderDeliveryStatus, HeaderBillingBlockStatus) = (?, ?, ?, ?, ?, ?)
		AND HeaderBillingStatus <> ?;`, dataKey.DeliveryDocument, dataKey.BillFromParty[0], dataKey.BillToParty[0], dataKey.HeaderCompleteDeliveryIsDefined, dataKey.HeaderDeliveryStatus, dataKey.HeaderBillingBlockStatus, dataKey.HeaderBillingStatus,
	)
	if err != nil {
		return nil, err
	}

	data, err := psdc.ConvertToDeliveryDocumentByReferenceDocument(sdc, rows)
	if err != nil {
		return nil, err
	}

	return data, err
}

func (f *SubFunction) DeliveryDocumentHeaderPartner(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) (*[]api_processing_data_formatter.DeliveryDocumentHeaderPartner, error) {
	var args []interface{}

	deliveryDocument := psdc.DeliveryDocument

	repeat := strings.Repeat("?,", len(*deliveryDocument)-1) + "?"
	for _, tag := range *deliveryDocument {
		args = append(args, tag.DeliveryDocument)
	}

	rows, err := f.db.Query(
		`SELECT DeliveryDocument, PartnerFunction, BusinessPartner
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_delivery_document_header_partner_data
		WHERE DeliveryDocument IN ( `+repeat+` );`, args...,
	)
	if err != nil {
		return nil, err
	}

	data, err := psdc.ConvertToDeliveryDocumentHeaderPartner(sdc, rows)
	if err != nil {
		return nil, err
	}

	return data, err
}

func (f *SubFunction) CreateSdc(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
	osdc *dpfm_api_output_formatter.SDC,
) error {
	var err error
	var e error

	wg := sync.WaitGroup{}
	wg.Add(3)

	psdc.MetaData, err = f.MetaData(sdc, psdc)
	if err != nil {
		return err
	}

	go func(wg *sync.WaitGroup) {
		defer wg.Done()

		// // I-1-1. OrderIDの絞り込み、および、入力パラメータによる請求元と請求先の絞り込み
		// psdc.OrderID, e = f.OrderIDByNumberSpecification(sdc, psdc)
		// if e != nil {
		// 	err = e
		// 	return
		// }

		// I-1-1. OrderIDの絞り込み、および、入力パラメータによる請求元と請求先の絞り込み
		psdc.OrderID, e = f.OrderIDByRangeSpecification(sdc, psdc)
		if e != nil {
			err = e
			return
		}

		// // II-1-1. OrderIDが未請求対象であることの確認
		// psdc.OrderID, e = f.OrderIDByReferenceDocument(sdc, psdc)
		// if e != nil {
		// 	err = e
		// 	return
		// }

		// I-1-2. ヘッダパートナのデータ取得
		psdc.OrdersHeaderPartner, e = f.OrdersHeaderPartner(sdc, psdc)
		if e != nil {
			err = e
			return
		}

		// 1-I-1.オーダー参照レコード・値の取得（オーダーヘッダ）
		psdc.HeaderOrdersHeader, e = f.HeaderOrdersHeader(sdc, psdc)
		if e != nil {
			err = e
			return
		}

		// 1-I-2. オーダー参照レコード・値の取得（オーダーヘッダパートナ）
		psdc.HeaderOrdersHeaderPartner, e = f.HeaderOrdersHeaderPartner(sdc, psdc)
		if e != nil {
			err = e
			return
		}

		// 2-5. TotalNetAmount
		psdc.TotalNetAmount, e = f.TotalNetAmount(sdc, psdc)
		if e != nil {
			err = e
			return
		}
	}(&wg)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		// // I-2-1. Delivery Document Headerの絞り込み、および、入力パラメータによる請求元と請求先の絞り込み
		// psdc.DeliveryDocument, e = f.DeliveryDocumentByNumberSpecification(sdc, psdc)
		// if e != nil {
		// 	err = e
		// 	return
		// }

		// I-2-1. Delivery Document Headerの絞り込み、および、入力パラメータによる請求元と請求先の絞り込み
		psdc.DeliveryDocument, e = f.DeliveryDocumentByRangeSpecification(sdc, psdc)
		if e != nil {
			err = e
			return
		}

		// // II-2-1. Delivery Document Headerの絞り込み、および、入力パラメータによる請求元と請求先の絞り込み
		// psdc.DeliveryDocument, e = f.DeliveryDocumentByReferenceDocument(sdc, psdc)
		// if e != nil {
		// 	err = e
		// 	return
		// }

		// II-1-2. ヘッダパートナのデータ取得
		psdc.DeliveryDocumentHeaderPartner, e = f.DeliveryDocumentHeaderPartner(sdc, psdc)
		if e != nil {
			err = e
			return
		}

		// 1-II-1. 入出荷伝票参照レコード・値の取得（入出荷伝票ヘッダ）
		psdc.HeaderDeliveryDocumentHeader, e = f.HeaderDeliveryDocumentHeader(sdc, psdc)
		if e != nil {
			err = e
			return
		}

		// 1-II-2. 入出荷伝票参照レコード・値の取得（入出荷伝票ヘッダパートナ）
		psdc.HeaderDeliveryDocumentHeaderPartner, e = f.HeaderDeliveryDocumentHeaderPartner(sdc, psdc)
		if e != nil {
			err = e
			return
		}

		// // 2-5. TotalNetAmount
		// psdc.TotalNetAmount, e = f.TotalNetAmount(sdc, psdc)
		// if e != nil {
		// 	err = e
		// 	return
		// }
	}(&wg)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		// 1-1 InvoiceDocument
		psdc.CalculateInvoiceDocument, e = f.CalculateInvoiceDocument(sdc, psdc)
		if e != nil {
			err = e
			return
		}
	}(&wg)

	wg.Wait()
	if err != nil {
		return err
	}

	f.l.Info(psdc)
	osdc, err = f.SetValue(sdc, psdc, osdc)
	if err != nil {
		return err
	}

	return nil
}
