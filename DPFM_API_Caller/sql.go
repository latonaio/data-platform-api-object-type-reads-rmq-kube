package dpfm_api_caller

import (
	"context"
	dpfm_api_input_reader "data-platform-api-freight-agreement-reads-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-freight-agreement-reads-rmq-kube/DPFM_API_Output_Formatter"
	"fmt"
	"sync"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

func (c *DPFMAPICaller) readSqlProcess(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	accepter []string,
	errs *[]error,
	log *logger.Logger,
) interface{} {
	var header *[]dpfm_api_output_formatter.Header
	var item *[]dpfm_api_output_formatter.Item

	for _, fn := range accepter {
		switch fn {
		case "Header":
			func() {
				header = c.Header(mtx, input, output, errs, log)
			}()
		case "HeadersByBuyer":
			func() {
				header = c.HeadersByBuyer(mtx, input, output, errs, log)
			}()
		case "HeadersBySeller":
			func() {
				header = c.HeadersBySeller(mtx, input, output, errs, log)
			}()
		case "Item":
			func() {
				item = c.Item(mtx, input, output, errs, log)
			}()
		case "Items":
			func() {
				item = c.Items(mtx, input, output, errs, log)
			}()

		default:
		}
		if len(*errs) != 0 {
			break
		}
	}

	data := &dpfm_api_output_formatter.Message{
		Header: header,
		Item:   item,
	}

	return data
}

func (c *DPFMAPICaller) Header(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Header {
	where := fmt.Sprintf("WHERE header.FreightAgreement = %d ", input.Header.FreightAgreement)

	if input.Header.IsCancelled != nil {
		where = fmt.Sprintf("%s\nAND header.IsCancelled = %v ", where, *input.Header.IsCancelled)
	}
	if input.Header.IsMarkedForDeletion != nil {
		where = fmt.Sprintf("%s\nAND header.IsMarkedForDeletion = %v", where, *input.Header.IsMarkedForDeletion)
	}

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_freight_agreement_header_data AS header
		` + where + ` ORDER BY header.IsMarkedForDeletion ASC, header.IsCancelled ASC, header.FreightAgreement DESC;`,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToHeader(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) HeadersByBuyer(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Header {
	where := "WHERE 1 = 1"
	if input.Header.Buyer != nil {
		where = fmt.Sprintf("%s\nAND Buyer = %v", where, *input.Header.Buyer)
	}
	if input.Header.HeaderBlockStatus != nil {
		where = fmt.Sprintf("%s\nAND HeaderBlockStatus = %t", where, *input.Header.HeaderBlockStatus)
	}
	if input.Header.HeaderBillingStatus != nil {
		where = fmt.Sprintf("%s\nAND HeaderBillingStatus != '%s'", where, *input.Header.HeaderBillingStatus)
	}
	if input.Header.IsCancelled != nil {
		where = fmt.Sprintf("%s\nAND IsCancelled = %t", where, *input.Header.IsCancelled)
	}
	if input.Header.IsMarkedForDeletion != nil {
		where = fmt.Sprintf("%s\nAND IsMarkedForDeletion = %t", where, *input.Header.IsMarkedForDeletion)
	}

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_freight_agreement_header_data
		` + where + `;`,
	)

	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	data, err := dpfm_api_output_formatter.ConvertToHeader(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) HeadersBySeller(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Header {
	where := "WHERE 1 = 1"
	if input.Header.Seller != nil {
		where = fmt.Sprintf("%s\nAND Seller = %v", where, *input.Header.Seller)
	}
	if input.Header.HeaderBlockStatus != nil {
		where = fmt.Sprintf("%s\nAND HeaderBlockStatus = %t", where, *input.Header.HeaderBlockStatus)
	}
	if input.Header.HeaderBillingStatus != nil {
		where = fmt.Sprintf("%s\nAND HeaderBillingStatus != '%s'", where, *input.Header.HeaderBillingStatus)
	}
	if input.Header.IsCancelled != nil {
		where = fmt.Sprintf("%s\nAND IsCancelled = %t", where, *input.Header.IsCancelled)
	}
	if input.Header.IsMarkedForDeletion != nil {
		where = fmt.Sprintf("%s\nAND IsMarkedForDeletion = %t", where, *input.Header.IsMarkedForDeletion)
	}

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_freight_agreement_header_data
		` + where + `;`,
	)

	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	data, err := dpfm_api_output_formatter.ConvertToHeader(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) Item(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Item {
	var args []interface{}
	where := fmt.Sprintf("WHERE FreightAgreement = %d ", input.Header.FreightAgreement)

	itemIDs := ""
	for _, v := range input.Header.Item {
		itemIDs = fmt.Sprintf("%s, %d", itemIDs, v.FreightAgreementItem)
	}

	if len(itemIDs) != 0 {
		where = fmt.Sprintf("%s\nAND FreightAgreementItem IN ( %s ) ", where, itemIDs[1:])
	}
	rows, err := c.db.Query(
		`SELECT
    		*
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_freight_agreement_item_data
		`+where+` ORDER BY IsMarkedForDeletion ASC, IsCancelled ASC, FreightAgreement DESC, FreightAgreementItem ASC ;`, args...,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToItem(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) Items(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Item {
	item := &dpfm_api_input_reader.Item{}
	if len(input.Header.Item) > 0 {
		item = &input.Header.Item[0]
	}
	where := "WHERE 1 = 1"

	if item != nil {
		if item.ItemBlockStatus != nil {
			where = fmt.Sprintf("%s\nAND item.ItemBlockStatus = '%v'", where, *item.ItemBlockStatus)
		}
		if item.ItemBillingBlockStatus != nil {
			where = fmt.Sprintf("%s\nAND item.ItemBillingBlockStatus = %v", where, *item.ItemBillingBlockStatus)
		}
		if item.IsCancelled != nil {
			where = fmt.Sprintf("%s\nAND item.IsCancelled = %v", where, *item.IsCancelled)
		}
		if item.IsMarkedForDeletion != nil {
			where = fmt.Sprintf("%s\nAND item.IsMarkedForDeletion = %v", where, *item.IsMarkedForDeletion)
		}
	}

	rows, err := c.db.Query(
		`SELECT 
			*
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_freight_agreement_item_data as item
		` + where + ` ORDER BY item.IsMarkedForDeletion ASC, item.IsCancelled ASC, item.FreightAgreement DESC, item.FreightAgreementItem ASC ;`)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToItem(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}
