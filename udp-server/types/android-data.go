package types

type AndroidData struct {
	Timestamp int64    `json:"timestamp"`
	Values    []float64 `json:"values"`
	Type      string    `json:"type"`
}
