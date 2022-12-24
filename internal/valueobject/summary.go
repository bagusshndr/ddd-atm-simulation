package valueobject

type Summaries []Summary

type Summary struct {
	TotalReview   float64
	AverageRating float64
	FiveStar      int
	FourStar      int
	ThreeStar     int
	TwoStar       int
	OneStar       int
}

func NewSummary(TotalReview, AverageRating float64) (*Summary, error) {
	return &Summary{
		TotalReview:   TotalReview,
		AverageRating: AverageRating,
		FiveStar:      0,
		FourStar:      0,
		ThreeStar:     0,
		TwoStar:       0,
		OneStar:       0,
	}, nil
}
