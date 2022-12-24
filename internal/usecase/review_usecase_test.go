package usecase

// import (
// 	"ddd-to-do-list/internal/aggregate"
// 	"ddd-to-do-list/internal/repository"
// 	"errors"
// 	"testing"

// 	"github.com/stretchr/testify/suite"
// )

// type reviewUsecaseTest struct {
// 	suite.Suite
// 	review  *aggregate.Review
// 	repo    *repository.ReviewMock
// 	usecase ReviewUsecase
// }

// func (t *reviewUsecaseTest) SetupSuite() {
// 	t.repo = new(repository.ReviewMock)
// 	t.review = aggregate.RebuildReviews(1, 1, 5)
// 	t.usecase = NewreviewUsecase(t.repo)
// }

// func (t *reviewUsecaseTest) TestGetReview() {
// 	reviews := aggregate.Reviews{aggregate.RebuildReviews(1, 1, 5)}
// 	t.Run("success", func() {
// 		t.repo.On("GetReview").Return(reviews, nil).Once()
// 		product, err := t.usecase.GetReview()
// 		t.NotNil(product)
// 		t.NoError(err)
// 	})

// 	t.Run("failed", func() {
// 		t.repo.On("GetReview").Return((aggregate.Reviews)(nil), errors.New("error")).Once()
// 		product, err := t.usecase.GetReview()
// 		t.Nil(product)
// 		t.Error(err)
// 	})
// }

// func (t *reviewUsecaseTest) TestGetReviewByID() {
// 	reviews := aggregate.Reviews{aggregate.RebuildReviews(1, 1, 5)}
// 	t.Run("success", func() {
// 		t.repo.On("GetReviewByID", reviews[0].ID).Return(reviews, nil).Once()
// 		product, err := t.usecase.GetReviewByID(reviews[0].ID)
// 		t.NotNil(product)
// 		t.NoError(err)
// 	})

// 	t.Run("failed", func() {
// 		t.repo.On("GetReviewByID", reviews[0].ID).Return((aggregate.Reviews)(nil), errors.New("error")).Once()
// 		product, err := t.usecase.GetReviewByID(reviews[0].ID)
// 		t.Nil(product)
// 		t.Error(err)
// 	})
// }

// func (t *reviewUsecaseTest) TestDelete() {
// 	reviews := aggregate.Reviews{aggregate.RebuildReviews(1, 1, 5)}
// 	t.Run("success", func() {
// 		t.repo.On("DeleteReview", reviews[0].ID).Return(nil).Once()
// 		err := t.usecase.DeleteReview(reviews[0].ID)
// 		t.NoError(err)
// 	})

// 	t.Run("failed", func() {
// 		t.repo.On("DeleteReview", reviews[0].ID).Return(errors.New("error")).Once()
// 		err := t.usecase.DeleteReview(reviews[0].ID)
// 		t.Error(err)
// 	})
// }

// func TestReviewUsecase(t *testing.T) {
// 	suite.Run(t, new(reviewUsecaseTest))
// }
