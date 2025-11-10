package usecase

import (
	"context"
)

type PresenterContext any

type Context struct {
	Ctx          context.Context
	PresenterCtx PresenterContext
}
