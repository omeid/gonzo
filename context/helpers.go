package context

import (
	"time"

	"context"

	"github.com/sirupsen/logrus"
)

func Background() Context {
	return ctx{
		Context: context.Background(),
		Entry:   std,
	}
}

//func TODO() Context
func WithCancel(parent Context) (Context, context.CancelFunc) {
	c, cancel := context.WithCancel(parent)

	var entry *logrus.Entry
	if ctx, ok := parent.(*ctx); ok {
		entry = ctx.Entry
	} else {
		entry = std
	}

	ctx := ctx{
		Context: c,
		Entry:   entry,
	}

	return &ctx, cancel
}

func WithDeadline(parent Context, deadline time.Time) (Context, context.CancelFunc) {
	c, cancel := context.WithDeadline(parent, deadline)

	var entry *logrus.Entry
	if ctx, ok := parent.(*ctx); ok {
		entry = ctx.Entry
	} else {
		entry = std
	}

	ctx := &ctx{
		Context: c,
		Entry:   entry,
	}

	return ctx, cancel
}

func WithTimeout(parent Context, timeout time.Duration) (Context, context.CancelFunc) {
	c, cancel := context.WithTimeout(parent, timeout)

	var entry *logrus.Entry
	if ctx, ok := parent.(*ctx); ok {
		entry = ctx.Entry
	} else {
		entry = std
	}

	ctx := &ctx{
		Context: c,
		Entry:   entry,
	}

	return ctx, cancel
}

func WithValue(parent Context, key string, value interface{}) Context {
	c := context.WithValue(parent, key, value)

	var entry *logrus.Entry
	if ctx, ok := parent.(*ctx); ok {
		entry = ctx.Entry
	} else {
		entry = std
	}

	ctx := &ctx{
		Context: c,
		Entry:   entry.WithField(key, value),
	}

	return ctx
}
