package context

import (
	"context"
	"github.com/Sirupsen/logrus"
)

var Canceled = context.Canceled

type Context interface {
	context.Context
	Logger
}

var _ Context = &ctx{}

type ctx struct {
	context.Context
	*logrus.Entry
}
