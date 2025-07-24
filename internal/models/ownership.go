package models

type Ownership struct {
	ID         string `json:"id"`
	Contract   string `json:"contract"`
	TokenID    string `json:"tokenId"`
	Owner      string `json:"owner"`
	Value      string `json:"value"`
	Blockchain string `json:"blockchain"`
}
