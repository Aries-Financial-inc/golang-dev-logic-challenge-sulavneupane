package model_response

// GraphPoint represents a pair of X and Y values
type GraphPoint struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

// AnalysisResponse represents the data structure of the analysis result
type AnalysisResponse struct {
	XYValues        []GraphPoint `json:"xy_values"`
	MaxProfit       float64      `json:"max_profit"`
	MaxLoss         float64      `json:"max_loss"`
	BreakEvenPoints []float64    `json:"break_even_points"`
}
