package listposts

import "context"

type Interaction interface {
	Execute(ctx context.Context) error
}

type Presenter interface {
	Present(ctx context.Context, output *Output) error
}

type Output struct {
	Posts []PostOutput
}

type PostOutput struct {
	ID      int64
	Content string
}
