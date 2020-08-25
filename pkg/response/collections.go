package response

type Collection struct {
	ID     int    `json:"id"`
	Label  string `json:"label"`
	Views  int    `json:"views"`
	Public int    `json:"public"`
	Count  int    `json:"count"`
}

type CollectionsResponse struct {
	Collections []Collection `json:"data"` 
}
