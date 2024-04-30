package myapp

import (
	"context"

	"github.com/morikuni/failure/v2"
)

type App struct{}

func NewApp() *App {
	return &App{}
}

func (a *App) Run(ctx context.Context) error {
	return failure.New(ErrNotImplemented, failure.Message("not implemented yet"))
}
