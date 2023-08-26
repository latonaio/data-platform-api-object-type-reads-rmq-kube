package requests

type Header struct {
	FreightAgreement                        int      `json:"FreightAgreement"`
	FreightAgreementDate                    string   `json:"FreightAgreementDate"`
	FreightAgreementType                    string   `json:"FreightAgreementType"`
	SupplyChainRelationshipID               int      `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipFreightID        int      `json:"SupplyChainRelationshipFreightID"`
	SupplyChainRelationshipFreightBillingID *int     `json:"SupplyChainRelationshipFreightBillingID"`
	SupplyChainRelationshipFreightPaymentID *int     `json:"SupplyChainRelationshipFreightPaymentID"`
	Buyer                                   int      `json:"Buyer"`
	Seller                                  int      `json:"Seller"`
	FreightPartner                          int      `json:"FreightPartner"`
	FreightBillToParty                      *int     `json:"FreightBillToParty"`
	FreightBillFromParty                    *int     `json:"FreightBillFromParty"`
	FreightBillToCountry                    *string  `json:"FreightBillToCountry"`
	FreightBillFromCountry                  *string  `json:"FreightBillFromCountry"`
	FreightPayer                            *int     `json:"FreightPayer"`
	FreightPayee                            *int     `json:"FreightPayee"`
	CreationDate                            string   `json:"CreationDate"`
	LastChangeDate                          string   `json:"LastChangeDate"`
	ContractType                            *string  `json:"ContractType"`
	FreightAgreementValidityStartDate       *string  `json:"FreightAgreementValidityStartDate"`
	FreightAgreementValidityEndDate         *string  `json:"FreightAgreementValidityEndDate"`
	InvoicePeriodStartDate                  *string  `json:"InvoicePeriodStartDate"`
	InvoicePeriodEndDate                    *string  `json:"InvoicePeriodEndDate"`
	HeaderBillingStatus                     string   `json:"HeaderBillingStatus"`
	HeaderDocReferenceStatus                string   `json:"HeaderDocReferenceStatus"`
	TransactionCurrency                     string   `json:"TransactionCurrency"`
	PricingDate                             string   `json:"PricingDate"`
	PriceDetnExchangeRate                   *float32 `json:"PriceDetnExchangeRate"`
	Incoterms                               *string  `json:"Incoterms"`
	PaymentTerms                            string   `json:"PaymentTerms"`
	PaymentMethod                           string   `json:"PaymentMethod"`
	ReferenceDocument                       *int     `json:"ReferenceDocument"`
	ReferenceDocumentItem                   *int     `json:"ReferenceDocumentItem"`
	AccountAssignmentGroup                  string   `json:"AccountAssignmentGroup"`
	AccountingExchangeRate                  *float32 `json:"AccountingExchangeRate"`
	InvoiceDocumentDate                     string   `json:"InvoiceDocumentDate"`
	IsExportImport                          *bool    `json:"IsExportImport"`
	HeaderText                              *string  `json:"HeaderText"`
	HeaderBlockStatus                       *bool    `json:"HeaderBlockStatus"`
	HeaderBillingBlockStatus                *bool    `json:"HeaderBillingBlockStatus"`
	IsCancelled                             *bool    `json:"IsCancelled"`
	IsMarkedForDeletion                     *bool    `json:"IsMarkedForDeletion"`
}
