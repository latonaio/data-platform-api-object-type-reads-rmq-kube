package dpfm_api_output_formatter

import (
	"data-platform-api-freight-agreement-reads-rmq-kube/DPFM_API_Caller/requests"
	"database/sql"
	"fmt"
)

func ConvertToHeader(rows *sql.Rows) (*[]Header, error) {
	defer rows.Close()
	header := make([]Header, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.Header{}

		err := rows.Scan(
			&pm.FreightAgreement,
			&pm.FreightAgreementDate,
			&pm.FreightAgreementType,
			&pm.SupplyChainRelationshipID,
			&pm.SupplyChainRelationshipFreightID,
			&pm.SupplyChainRelationshipFreightBillingID,
			&pm.SupplyChainRelationshipFreightPaymentID,
			&pm.Buyer,
			&pm.Seller,
			&pm.FreightPartner,
			&pm.FreightBillToParty,
			&pm.FreightBillFromParty,
			&pm.FreightBillToCountry,
			&pm.FreightBillFromCountry,
			&pm.FreightPayer,
			&pm.FreightPayee,
			&pm.CreationDate,
			&pm.LastChangeDate,
			&pm.ContractType,
			&pm.FreightAgreementValidityStartDate,
			&pm.FreightAgreementValidityEndDate,
			&pm.InvoicePeriodStartDate,
			&pm.InvoicePeriodEndDate,
			&pm.HeaderBillingStatus,
			&pm.HeaderDocReferenceStatus,
			&pm.TransactionCurrency,
			&pm.PricingDate,
			&pm.PriceDetnExchangeRate,
			&pm.Incoterms,
			&pm.PaymentTerms,
			&pm.PaymentMethod,
			&pm.ReferenceDocument,
			&pm.ReferenceDocumentItem,
			&pm.AccountAssignmentGroup,
			&pm.AccountingExchangeRate,
			&pm.InvoiceDocumentDate,
			&pm.IsExportImport,
			&pm.HeaderText,
			&pm.HeaderBlockStatus,
			&pm.HeaderBillingBlockStatus,
			&pm.IsCancelled,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &header, err
		}

		data := pm
		header = append(header, Header{
			FreightAgreement:                        data.FreightAgreement,
			FreightAgreementDate:                    data.FreightAgreementDate,
			FreightAgreementType:                    data.FreightAgreementType,
			SupplyChainRelationshipID:               data.SupplyChainRelationshipID,
			SupplyChainRelationshipFreightID:        data.SupplyChainRelationshipFreightID,
			SupplyChainRelationshipFreightBillingID: data.SupplyChainRelationshipFreightBillingID,
			SupplyChainRelationshipFreightPaymentID: data.SupplyChainRelationshipFreightPaymentID,
			Buyer:                                   data.Buyer,
			Seller:                                  data.Seller,
			FreightPartner:                          data.FreightPartner,
			FreightBillToParty:                      data.FreightBillToParty,
			FreightBillFromParty:                    data.FreightBillFromParty,
			FreightBillToCountry:                    data.FreightBillToCountry,
			FreightBillFromCountry:                  data.FreightBillFromCountry,
			FreightPayer:                            data.FreightPayer,
			FreightPayee:                            data.FreightPayee,
			CreationDate:                            data.CreationDate,
			LastChangeDate:                          data.LastChangeDate,
			ContractType:                            data.ContractType,
			FreightAgreementValidityStartDate:       data.FreightAgreementValidityStartDate,
			FreightAgreementValidityEndDate:         data.FreightAgreementValidityEndDate,
			InvoicePeriodStartDate:                  data.InvoicePeriodStartDate,
			InvoicePeriodEndDate:                    data.InvoicePeriodEndDate,
			HeaderBillingStatus:                     data.HeaderBillingStatus,
			HeaderDocReferenceStatus:                data.HeaderDocReferenceStatus,
			TransactionCurrency:                     data.TransactionCurrency,
			PricingDate:                             data.PricingDate,
			PriceDetnExchangeRate:                   data.PriceDetnExchangeRate,
			Incoterms:                               data.Incoterms,
			PaymentTerms:                            data.PaymentTerms,
			PaymentMethod:                           data.PaymentMethod,
			ReferenceDocument:                       data.ReferenceDocument,
			ReferenceDocumentItem:                   data.ReferenceDocumentItem,
			AccountAssignmentGroup:                  data.AccountAssignmentGroup,
			AccountingExchangeRate:                  data.AccountingExchangeRate,
			InvoiceDocumentDate:                     data.InvoiceDocumentDate,
			IsExportImport:                          data.IsExportImport,
			HeaderText:                              data.HeaderText,
			HeaderBlockStatus:                       data.HeaderBlockStatus,
			HeaderBillingBlockStatus:                data.HeaderBillingBlockStatus,
			IsCancelled:                             data.IsCancelled,
			IsMarkedForDeletion:                     data.IsMarkedForDeletion,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &header, nil
	}

	return &header, nil
}

func ConvertToItem(rows *sql.Rows) (*[]Item, error) {
	defer rows.Close()
	item := make([]Item, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.Item{}

		err := rows.Scan(
			&pm.FreightAgreement,
			&pm.FreightAgreementItem,
			&pm.FreightAgreementItemCategory,
			&pm.SupplyChainRelationshipID,
			&pm.SupplyChainRelationshipDeliveryID,
			&pm.SupplyChainRelationshipDeliveryPlantID,
			&pm.FreightAgreementItemText,
			&pm.Product,
			&pm.ProductStandardID,
			&pm.ProductGroup,
			&pm.BaseUnit,
			&pm.DeliverToParty,
			&pm.DeliverFromParty,
			&pm.CreationDate,
			&pm.LastChangeDate,
			&pm.DeliverToPlant,
			&pm.DeliverToPlantTimeZone,
			&pm.DeliverFromPlant,
			&pm.DeliverFromPlantTimeZone,
			&pm.Incoterms,
			&pm.TransactionTaxClassification,
			&pm.ProductTaxClassificationBillToCountry,
			&pm.ProductTaxClassificationBillFromCountry,
			&pm.DefinedTaxClassification,
			&pm.AccountAssignmentGroup,
			&pm.ProductAccountAssignmentGroup,
			&pm.PaymentTerms,
			&pm.DueCalculationBaseDate,
			&pm.PaymentDueDate,
			&pm.NetPaymentDays,
			&pm.PaymentMethod,
			&pm.Project,
			&pm.WBSElement,
			&pm.ItemBillingStatus,
			&pm.TaxCode,
			&pm.TaxRate,
			&pm.ItemBlockStatus,
			&pm.ItemBillingBlockStatus,
			&pm.IsCancelled,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &item, nil
		}

		data := pm
		item = append(item, Item{
			FreightAgreement:                        data.FreightAgreement,
			FreightAgreementItem:                    data.FreightAgreementItem,
			FreightAgreementItemCategory:            data.FreightAgreementItemCategory,
			SupplyChainRelationshipID:               data.SupplyChainRelationshipID,
			SupplyChainRelationshipDeliveryID:       data.SupplyChainRelationshipDeliveryID,
			SupplyChainRelationshipDeliveryPlantID:  data.SupplyChainRelationshipDeliveryPlantID,
			FreightAgreementItemText:                data.FreightAgreementItemText,
			Product:                                 data.Product,
			ProductStandardID:                       data.ProductStandardID,
			ProductGroup:                            data.ProductGroup,
			BaseUnit:                                data.BaseUnit,
			DeliverToParty:                          data.DeliverToParty,
			DeliverFromParty:                        data.DeliverFromParty,
			CreationDate:                            data.CreationDate,
			LastChangeDate:                          data.LastChangeDate,
			DeliverToPlant:                          data.DeliverToPlant,
			DeliverToPlantTimeZone:                  data.DeliverToPlantTimeZone,
			DeliverFromPlant:                        data.DeliverFromPlant,
			DeliverFromPlantTimeZone:                data.DeliverFromPlantTimeZone,
			Incoterms:                               data.Incoterms,
			TransactionTaxClassification:            data.TransactionTaxClassification,
			ProductTaxClassificationBillToCountry:   data.ProductTaxClassificationBillToCountry,
			ProductTaxClassificationBillFromCountry: data.ProductTaxClassificationBillFromCountry,
			DefinedTaxClassification:                data.DefinedTaxClassification,
			AccountAssignmentGroup:                  data.AccountAssignmentGroup,
			ProductAccountAssignmentGroup:           data.ProductAccountAssignmentGroup,
			PaymentTerms:                            data.PaymentTerms,
			DueCalculationBaseDate:                  data.DueCalculationBaseDate,
			PaymentDueDate:                          data.PaymentDueDate,
			NetPaymentDays:                          data.NetPaymentDays,
			PaymentMethod:                           data.PaymentMethod,
			Project:                                 data.Project,
			WBSElement:                              data.WBSElement,
			ItemBillingStatus:                       data.ItemBillingStatus,
			TaxCode:                                 data.TaxCode,
			TaxRate:                                 data.TaxRate,
			ItemBlockStatus:                         data.ItemBlockStatus,
			ItemBillingBlockStatus:                  data.ItemBillingBlockStatus,
			IsCancelled:                             data.IsCancelled,
			IsMarkedForDeletion:                     data.IsMarkedForDeletion,
		})
	}

	return &item, nil
}
