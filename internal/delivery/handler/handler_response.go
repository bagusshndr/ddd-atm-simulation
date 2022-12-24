package handler

// import (
// 	"ddd-to-do-list/internal/aggregate"
// 	"ddd-to-do-list/internal/valueobject"
// )

// type ProductResponses []ProductResponse

// type ReviewResponses []ReviewResponse

// type SummaryResponse struct {
// 	TotalReview   float64 `json:"total_review"`
// 	AverageRating float64 `json:"average_rating"`
// 	FiveStar      int     `json:"5_star"`
// 	FourStar      int     `json:"4_star"`
// 	ThreeStar     int     `json:"3_star"`
// 	TwoStar       int     `json:"2_star"`
// 	OneStar       int     `json:"1_star"`
// }

// type ProductResponse struct {
// 	ID    uint64  `json:"id"`
// 	Name  string  `json:"name"`
// 	Price float64 `json:"price"`
// }

// type FailedActivityResponse struct {
// }

// type ReviewResponse struct {
// 	ID        uint64 `json:"id"`
// 	ProductID uint64 `json:"product_id"`
// 	Rating    int    `json:"rating"`
// }

// func (response SummaryResponse) Response(summary valueobject.Summary) SummaryResponse {

// 	response.TotalReview = summary.TotalReview
// 	response.AverageRating = summary.AverageRating
// 	response.FiveStar = summary.FiveStar
// 	response.FourStar = summary.FourStar
// 	response.ThreeStar = summary.ThreeStar
// 	response.TwoStar = summary.TwoStar
// 	response.OneStar = summary.OneStar

// 	return SummaryResponse{
// 		TotalReview:   response.TotalReview,
// 		AverageRating: response.AverageRating,
// 		FiveStar:      response.FiveStar,
// 		FourStar:      response.FourStar,
// 		ThreeStar:     response.ThreeStar,
// 		TwoStar:       response.TwoStar,
// 		OneStar:       response.OneStar,
// 	}
// }

// func (response ProductResponse) Responses(product aggregate.Products) (result ProductResponses) {
// 	for _, src := range product {
// 		response.ID = src.ID
// 		response.Name = src.Name
// 		response.Price = src.Price

// 		result = append(result, ProductResponse{
// 			response.ID,
// 			response.Name,
// 			response.Price,
// 		},
// 		)
// 	}
// 	return result
// }

// func (response ProductResponse) Response(product aggregate.Products) ProductResponse {
// 	for _, src := range product {
// 		response.ID = src.ID
// 		response.Name = src.Name
// 		response.Price = src.Price
// 	}
// 	return ProductResponse{
// 		ID:    response.ID,
// 		Name:  response.Name,
// 		Price: response.Price,
// 	}
// }

// func (response ReviewResponse) Responses(review aggregate.Reviews) (result ReviewResponses) {
// 	for _, src := range review {
// 		response.ID = src.ID
// 		response.ProductID = src.ProductID
// 		response.Rating = src.Rating

// 		result = append(result, ReviewResponse{
// 			response.ID,
// 			response.ProductID,
// 			response.Rating,
// 		},
// 		)
// 	}
// 	return result
// }

// func (response ReviewResponse) Response(review aggregate.Reviews) ReviewResponse {
// 	for _, src := range review {
// 		response.ID = src.ID
// 		response.ProductID = src.ProductID
// 		response.Rating = src.Rating
// 	}
// 	return ReviewResponse{
// 		response.ID,
// 		response.ProductID,
// 		response.Rating,
// 	}
// }
