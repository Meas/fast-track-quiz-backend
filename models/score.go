package models

type Score struct {
	ID     string
	Points int
}

type Result struct {
	Points            int     `json:"points"`
	BetterThan        float32 `json:"betterThan"`
	BetterOrEqualThan float32 `json:"betterOrEqualThan"`
}
