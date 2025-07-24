package models

type TraitRarityRequest struct {
	CollectionId string            `json:"collectionId"`
	Properties   []TraitRarityProp `json:"properties"`
}

type TraitRarityProp struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type TraitRarity struct {
	Key    string `json:"key"`
	Value  string `json:"value"`
	Rarity string `json:"rarity"`
}

type TraitRarityResponse struct {
	Traits []TraitRarity `json:"traits"`
}
