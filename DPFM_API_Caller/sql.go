package dpfm_api_caller

import (
	"context"
	dpfm_api_input_reader "data-platform-api-object-type-reads-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-object-type-reads-rmq-kube/DPFM_API_Output_Formatter"
	"fmt"
	"strings"
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
	var objectType *[]dpfm_api_output_formatter.ObjectType
	var text *[]dpfm_api_output_formatter.Text
	for _, fn := range accepter {
		switch fn {
		case "ObjectType":
			func() {
				objectType = c.ObjectType(mtx, input, output, errs, log)
			}()
		case "ObjectTypes":
			func() {
				objectType = c.ObjectTypes(mtx, input, output, errs, log)
			}()
		case "Text":
			func() {
				text = c.Text(mtx, input, output, errs, log)
			}()
		case "Texts":
			func() {
				text = c.Texts(mtx, input, output, errs, log)
			}()
		default:
		}
	}

	data := &dpfm_api_output_formatter.Message{
		ObjectType: objectType,
		Text:		text,
	}

	return data
}

func (c *DPFMAPICaller) ObjectType(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.ObjectType {
	where := fmt.Sprintf("WHERE ObjectType = '%s'", input.ObjectType.ObjectType)

	if input.ObjectType.IsMarkedForDeletion != nil {
		where = fmt.Sprintf("%s\nAND IsMarkedForDeletion = %v", where, *input.ObjectType.IsMarkedForDeletion)
	}

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_object_type_object_type_data
		` + where + ` ORDER BY IsMarkedForDeletion ASC, ObjectType DESC;`,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToObjectType(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) ObjectTypes(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.ObjectType {
	where := fmt.Sprintf("WHERE ObjectType = '%s'", input.ObjectType.ObjectType)

	if input.ObjectType.IsMarkedForDeletion != nil {
		where = fmt.Sprintf("%s\nAND IsMarkedForDeletion = %v", where, *input.ObjectType.IsMarkedForDeletion)
	}

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_object_type_object_type_data
		` + where + ` ORDER BY IsMarkedForDeletion ASC, ObjectType DESC;`,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToObjectType(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) Text(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Text {
	var args []interface{}
	objectType := input.ObjectType.ObjectType
	text := input.ObjectType.Text

	cnt := 0
	for _, v := range text {
		args = append(args, objectType, v.Language)
		cnt++
	}

	repeat := strings.Repeat("(?,?),", cnt-1) + "(?,?)"
	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_object_type_text_data
		WHERE (ObjectType, Language) IN ( `+repeat+` );`, args...,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToText(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) Texts(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Text {
	var args []interface{}
	text := input.ObjectType.Text

	cnt := 0
	for _, v := range text {
		args = append(args, v.Language)
		cnt++
	}

	repeat := strings.Repeat("(?),", cnt-1) + "(?)"
	rows, err := c.db.Query(
		`SELECT * 
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_object_type_text_data
		WHERE Language IN ( `+repeat+` );`, args...,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	//
	data, err := dpfm_api_output_formatter.ConvertToText(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}
