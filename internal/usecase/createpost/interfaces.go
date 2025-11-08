package createpost

import "go-clean-microblog/internal/usecase"

type Interaction interface {
	Execute(ctx usecase.Context, input Input) error
}

type Presenter interface {
	Present(ctx usecase.PresenterContext, output *Output) error
}

type Input struct {
	Content string
}

type Output struct {
	ID      int64
	Content string
}
