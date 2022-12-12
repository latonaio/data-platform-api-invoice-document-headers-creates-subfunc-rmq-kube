package requests

type OrderIDKey struct {
	OrderID                         *int   `json:"OrderID"`
	BillFromPartyFrom               *int   `json:"BillFromPartyFrom"`
	BillFromPartyTo                 *int   `json:"BillFromPartyTo"`
	BillToPartyFrom                 *int   `json:"BillToPartyFrom"`
	BillToPartyTo                   *int   `json:"BillToPartyTo"`
	BillFromParty                   []*int `json:"BillFromParty"`
	BillToParty                     []*int `json:"BillToParty"`
	HeaderCompleteDeliveryIsDefined *bool  `json:"HeaderCompleteDeliveryIsDefined"`
	HeaderDeliveryStatus            string `json:"HeaderDeliveryStatus"`
	HeaderBillingStatus             string `json:"HeaderBillingStatus"`
	HeaderBillingBlockStatus        *bool  `json:"HeaderBillingBlockStatus"`
}
