package dpfm_api_output_formatter

import (
	"data-platform-api-object-type-reads-rmq-kube/DPFM_API_Caller/requests"
	"database/sql"
	"fmt"
)

func ConvertToObjectType(rows *sql.Rows) (*[]ObjectType, error) {
	defer rows.Close()
	objectType := make([]ObjectType, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.ObjectType{}

		err := rows.Scan(
			&pm.ObjectType,
			&pm.CreationDate,
			&pm.LastChangeDate,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &objectType, nil
		}

		data := pm
		objectType = append(objectType, ObjectType{
			ObjectType:				data.ObjectType,
			CreationDate:			data.CreationDate,
			LastChangeDate:			data.LastChangeDate,
			IsMarkedForDeletion:	data.IsMarkedForDeletion,
		})
	}

	return &objectType, nil
}

func ConvertToText(rows *sql.Rows) (*[]Text, error) {
	defer rows.Close()
	text := make([]Text, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.Text{}

		err := rows.Scan(
			&pm.ObjectType,
			&pm.Language,
			&pm.ObjectTypeName,
			&pm.CreationDate,
			&pm.LastChangeDate,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &text, err
		}

		data := pm
		text = append(text, Text{
			ObjectType:     		data.ObjectType,
			Language:          		data.Language,
			ObjectTypeName:			data.ObjectTypeName,
			CreationDate:			data.CreationDate,
			LastChangeDate:			data.LastChangeDate,
			IsMarkedForDeletion:	data.IsMarkedForDeletion,
		})
	}

	return &text, nil
}
