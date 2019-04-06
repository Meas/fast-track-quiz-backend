package models

// Score - structure for single user score
type Score struct {
	ID     string
	Points int
}

// Result - structure for final quiz result
type Result struct {
	Points            int     `json:"points"`
	BetterThan        float32 `json:"betterThan"`
	BetterOrEqualThan float32 `json:"betterOrEqualThan"`
}
