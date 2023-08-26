package requests

type Item struct {
	FreightAgreement                        int      `json:"FreightAgreement"`
	FreightAgreementItem                    int      `json:"FreightAgreementItem"`
	FreightAgreementItemCategory            string   `json:"FreightAgreementItemCategory"`
	SupplyChainRelationshipID               int      `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipDeliveryID       *int     `json:"SupplyChainRelationshipDeliveryID"`
	SupplyChainRelationshipDeliveryPlantID  *int     `json:"SupplyChainRelationshipDeliveryPlantID"`
	FreightAgreementItemText                string   `json:"FreightAgreementItemText"`
	Product                                 string   `json:"Product"`
	ProductStandardID                       string   `json:"ProductStandardID"`
	ProductGroup                            *string  `json:"ProductGroup"`
	BaseUnit                                string   `json:"BaseUnit"`
	DeliverToParty                          *int     `json:"DeliverToParty"`
	DeliverFromParty                        *int     `json:"DeliverFromParty"`
	CreationDate                            string   `json:"CreationDate"`
	LastChangeDate                          string   `json:"LastChangeDate"`
	DeliverToPlant                          *string  `json:"DeliverToPlant"`
	DeliverToPlantTimeZone                  *string  `json:"DeliverToPlantTimeZone"`
	DeliverFromPlant                        *string  `json:"DeliverFromPlant"`
	DeliverFromPlantTimeZone                *string  `json:"DeliverFromPlantTimeZone"`
	Incoterms                               *string  `json:"Incoterms"`
	TransactionTaxClassification            string   `json:"TransactionTaxClassification"`
	ProductTaxClassificationBillToCountry   string   `json:"ProductTaxClassificationBillToCountry"`
	ProductTaxClassificationBillFromCountry string   `json:"ProductTaxClassificationBillFromCountry"`
	DefinedTaxClassification                string   `json:"DefinedTaxClassification"`
	AccountAssignmentGroup                  string   `json:"AccountAssignmentGroup"`
	ProductAccountAssignmentGroup           string   `json:"ProductAccountAssignmentGroup"`
	PaymentTerms                            string   `json:"PaymentTerms"`
	DueCalculationBaseDate                  *string  `json:"DueCalculationBaseDate"`
	PaymentDueDate                          *string  `json:"PaymentDueDate"`
	NetPaymentDays                          *int     `json:"NetPaymentDays"`
	PaymentMethod                           string   `json:"PaymentMethod"`
	Project                                 *int     `json:"Project"`
	WBSElement                              *int     `json:"WBSElement"`
	ItemBillingStatus                       *string  `json:"ItemBillingStatus"`
	TaxCode                                 *string  `json:"TaxCode"`
	TaxRate                                 *float32 `json:"TaxRate"`
	ItemBlockStatus                         bool     `json:"ItemBlockStatus"`
	ItemBillingBlockStatus                  bool     `json:"ItemBillingBlockStatus"`
	IsCancelled                             bool     `json:"IsCancelled"`
	IsMarkedForDeletion                     bool     `json:"IsMarkedForDeletion"`
}
