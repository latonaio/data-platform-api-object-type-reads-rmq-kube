package dpfm_api_input_reader

type EC_MC struct {
	ConnectionKey string `json:"connection_key"`
	Result        bool   `json:"result"`
	RedisKey      string `json:"redis_key"`
	Filepath      string `json:"filepath"`
	Document      struct {
		DocumentNo     string `json:"document_no"`
		DeliverTo      string `json:"deliver_to"`
		Quantity       string `json:"quantity"`
		PickedQuantity string `json:"picked_quantity"`
		Price          string `json:"price"`
		Batch          string `json:"batch"`
	} `json:"document"`
	BusinessPartner struct {
		DocumentNo           string `json:"document_no"`
		Status               string `json:"status"`
		DeliverTo            string `json:"deliver_to"`
		Quantity             string `json:"quantity"`
		CompletedQuantity    string `json:"completed_quantity"`
		PlannedStartDate     string `json:"planned_start_date"`
		PlannedValidatedDate string `json:"planned_validated_date"`
		ActualStartDate      string `json:"actual_start_date"`
		ActualValidatedDate  string `json:"actual_validated_date"`
		Batch                string `json:"batch"`
		Work                 struct {
			WorkNo                   string `json:"work_no"`
			Quantity                 string `json:"quantity"`
			CompletedQuantity        string `json:"completed_quantity"`
			ErroredQuantity          string `json:"errored_quantity"`
			Component                string `json:"component"`
			PlannedComponentQuantity string `json:"planned_component_quantity"`
			PlannedStartDate         string `json:"planned_start_date"`
			PlannedStartTime         string `json:"planned_start_time"`
			PlannedValidatedDate     string `json:"planned_validated_date"`
			PlannedValidatedTime     string `json:"planned_validated_time"`
			ActualStartDate          string `json:"actual_start_date"`
			ActualStartTime          string `json:"actual_start_time"`
			ActualValidatedDate      string `json:"actual_validated_date"`
			ActualValidatedTime      string `json:"actual_validated_time"`
		} `json:"work"`
	} `json:"business_partner"`
	APISchema     string   `json:"api_schema"`
	Accepter      []string `json:"accepter"`
	MaterialCode  string   `json:"material_code"`
	Plant         string   `json:"plant/supplier"`
	Stock         string   `json:"stock"`
	DocumentType  string   `json:"document_type"`
	DocumentNo    string   `json:"document_no"`
	PlannedDate   string   `json:"planned_date"`
	ValidatedDate string   `json:"validated_date"`
	Deleted       bool     `json:"deleted"`
}

type SDC struct {
	ConnectionKey                   string                          `json:"connection_key"`
	Result                          bool                            `json:"result"`
	RedisKey                        string                          `json:"redis_key"`
	Filepath                        string                          `json:"filepath"`
	APIStatusCode                   int                             `json:"api_status_code"`
	RuntimeSessionID                string                          `json:"runtime_session_id"`
	BusinessPartnerID               *int                            `json:"business_partner"`
	ServiceLabel                    string                          `json:"service_label"`
	APIType                         string                          `json:"APIType"`
	FreightAgreementInputParameters FreightAgreementInputParameters `json:"FreightAgreementInputParameters"`
	Header                          Header                          `json:"FreightAgreement"`
	APISchema                       string                          `json:"api_schema"`
	Accepter                        []string                        `json:"accepter"`
	Deleted                         bool                            `json:"deleted"`
}

type FreightAgreementInputParameters struct {
	ReferenceDocument     *int `json:"ReferenceDocument"`
	ReferenceDocumentItem *int `json:"ReferenceDocumentItem"`
}

type Header struct {
	FreightAgreement                        int      `json:"FreightAgreement"`
	FreightAgreementDate                    *string  `json:"FreightAgreementDate"`
	FreightAgreementType                    *string  `json:"FreightAgreementType"`
	SupplyChainRelationshipID               *int     `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipFreightID        *int     `json:"SupplyChainRelationshipFreightID"`
	SupplyChainRelationshipFreightBillingID *int     `json:"SupplyChainRelationshipFreightBillingID"`
	SupplyChainRelationshipFreightPaymentID *int     `json:"SupplyChainRelationshipFreightPaymentID"`
	Buyer                                   *int     `json:"Buyer"`
	Seller                                  *int     `json:"Seller"`
	FreightPartner                          *int     `json:"FreightPartner"`
	FreightBillToParty                      *int     `json:"FreightBillToParty"`
	FreightBillFromParty                    *int     `json:"FreightBillFromParty"`
	FreightBillToCountry                    *string  `json:"FreightBillToCountry"`
	FreightBillFromCountry                  *string  `json:"FreightBillFromCountry"`
	FreightPayer                            *int     `json:"FreightPayer"`
	FreightPayee                            *int     `json:"FreightPayee"`
	CreationDate                            *string  `json:"CreationDate"`
	LastChangeDate                          *string  `json:"LastChangeDate"`
	ContractType                            *string  `json:"ContractType"`
	FreightAgreementValidityStartDate       *string  `json:"FreightAgreementValidityStartDate"`
	FreightAgreementValidityEndDate         *string  `json:"FreightAgreementValidityEndDate"`
	InvoicePeriodStartDate                  *string  `json:"InvoicePeriodStartDate"`
	InvoicePeriodEndDate                    *string  `json:"InvoicePeriodEndDate"`
	HeaderBillingStatus                     *string  `json:"HeaderBillingStatus"`
	HeaderDocReferenceStatus                *string  `json:"HeaderDocReferenceStatus"`
	TransactionCurrency                     *string  `json:"TransactionCurrency"`
	PricingDate                             *string  `json:"PricingDate"`
	PriceDetnExchangeRate                   *float32 `json:"PriceDetnExchangeRate"`
	Incoterms                               *string  `json:"Incoterms"`
	PaymentTerms                            *string  `json:"PaymentTerms"`
	PaymentMethod                           *string  `json:"PaymentMethod"`
	ReferenceDocument                       *int     `json:"ReferenceDocument"`
	ReferenceDocumentItem                   *int     `json:"ReferenceDocumentItem"`
	AccountAssignmentGroup                  *string  `json:"AccountAssignmentGroup"`
	AccountingExchangeRate                  *float32 `json:"AccountingExchangeRate"`
	InvoiceDocumentDate                     *string  `json:"InvoiceDocumentDate"`
	IsExportImport                          *bool    `json:"IsExportImport"`
	HeaderText                              *string  `json:"HeaderText"`
	HeaderBlockStatus                       *bool    `json:"HeaderBlockStatus"`
	HeaderBillingBlockStatus                *bool    `json:"HeaderBillingBlockStatus"`
	IsCancelled                             *bool    `json:"IsCancelled"`
	IsMarkedForDeletion                     *bool    `json:"IsMarkedForDeletion"`
	Item                                    []Item   `json:"Item"`
}

type Item struct {
	FreightAgreement                        int      `json:"FreightAgreement"`
	FreightAgreementItem                    int      `json:"FreightAgreementItem"`
	FreightAgreementItemCategory            *string  `json:"FreightAgreementItemCategory"`
	SupplyChainRelationshipID               *int     `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipDeliveryID       *int     `json:"SupplyChainRelationshipDeliveryID"`
	SupplyChainRelationshipDeliveryPlantID  *int     `json:"SupplyChainRelationshipDeliveryPlantID"`
	FreightAgreementItemText                *string  `json:"FreightAgreementItemText"`
	Product                                 *string  `json:"Product"`
	ProductStandardID                       *string  `json:"ProductStandardID"`
	ProductGroup                            *string  `json:"ProductGroup"`
	BaseUnit                                *string  `json:"BaseUnit"`
	DeliverToParty                          *int     `json:"DeliverToParty"`
	DeliverFromParty                        *int     `json:"DeliverFromParty"`
	CreationDate                            *string  `json:"CreationDate"`
	LastChangeDate                          *string  `json:"LastChangeDate"`
	DeliverToPlant                          *string  `json:"DeliverToPlant"`
	DeliverToPlantTimeZone                  *string  `json:"DeliverToPlantTimeZone"`
	DeliverFromPlant                        *string  `json:"DeliverFromPlant"`
	DeliverFromPlantTimeZone                *string  `json:"DeliverFromPlantTimeZone"`
	Incoterms                               *string  `json:"Incoterms"`
	TransactionTaxClassification            *string  `json:"TransactionTaxClassification"`
	ProductTaxClassificationBillToCountry   *string  `json:"ProductTaxClassificationBillToCountry"`
	ProductTaxClassificationBillFromCountry *string  `json:"ProductTaxClassificationBillFromCountry"`
	DefinedTaxClassification                *string  `json:"DefinedTaxClassification"`
	AccountAssignmentGroup                  *string  `json:"AccountAssignmentGroup"`
	ProductAccountAssignmentGroup           *string  `json:"ProductAccountAssignmentGroup"`
	PaymentTerms                            *string  `json:"PaymentTerms"`
	DueCalculationBaseDate                  *string  `json:"DueCalculationBaseDate"`
	PaymentDueDate                          *string  `json:"PaymentDueDate"`
	NetPaymentDays                          *int     `json:"NetPaymentDays"`
	PaymentMethod                           *string  `json:"PaymentMethod"`
	Project                                 *int     `json:"Project"`
	WBSElement                              *int     `json:"WBSElement"`
	ItemBillingStatus                       *string  `json:"ItemBillingStatus"`
	TaxCode                                 *string  `json:"TaxCode"`
	TaxRate                                 *float32 `json:"TaxRate"`
	ItemBlockStatus                         *bool    `json:"ItemBlockStatus"`
	ItemBillingBlockStatus                  *bool    `json:"ItemBillingBlockStatus"`
	IsCancelled                             *bool    `json:"IsCancelled"`
	IsMarkedForDeletion                     *bool    `json:"IsMarkedForDeletion"`
}
