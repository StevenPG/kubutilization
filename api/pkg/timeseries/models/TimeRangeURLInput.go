package models

// TimeRangeURLInput ... struct that contains input parameters for Time Range query parameters
type TimeRangeURLInput struct {
	From int64 `form:"from" json:"from"`
	To   int64 `form:"to" json:"to"`
}
