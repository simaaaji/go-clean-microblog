package createpost

import "context"

type Interaction interface {
	Execute(ctx context.Context, input Input) error
}

type Presenter interface {
	Present(ctx context.Context, output *Output) error
}

type Input struct {
	Content string
}

type Output struct {
	ID      int64
	Content string
}
