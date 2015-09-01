package context

import (
	"github.com/Sirupsen/logrus"
	"golang.org/x/net/context"
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
