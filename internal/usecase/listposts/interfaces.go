package listposts

import "go-clean-microblog/internal/usecase"

type Interaction interface {
	Execute(ctx usecase.Context) error
}

type Presenter interface {
	Present(ctx usecase.PresenterContext, output *Output) error
}

type Output struct {
	Posts []PostOutput
}

type PostOutput struct {
	ID      int64
	Content string
}
