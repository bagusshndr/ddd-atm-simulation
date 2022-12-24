package usecase

import (
	"ddd-to-do-list/internal/valueobject"
)

type SummaryUsecase interface {
	GetSummary() (summaries valueobject.Summary, err error)
	GetSummaryByProductID(id uint64) (summaries valueobject.Summary, err error)
}
