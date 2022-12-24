package usecase

import (
	"ddd-atm-simulation/internal/valueobject"
)

type SummaryUsecase interface {
	GetSummary() (summaries valueobject.Summary, err error)
	GetSummaryByProductID(id uint64) (summaries valueobject.Summary, err error)
}
