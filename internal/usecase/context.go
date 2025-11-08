package usecase

import (
	"context"
)

type PresenterContext interface {
	Set(key string, value any)
}

type Context struct {
	Ctx          context.Context
	PresenterCtx PresenterContext
}
