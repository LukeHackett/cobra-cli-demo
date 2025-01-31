package utils

import (
	"context"

	"github.com/LukeHackett/cobra-cli-demo/internal/model"
)

type contextKey string

var (
	config  = contextKey("config")
	profile = contextKey("profile")
)

func SetConfig(ctx context.Context, value model.CliConfig) context.Context {
	return context.WithValue(ctx, config, value)
}

func GetConfig(ctx context.Context) model.CliConfig {
	return ctx.Value(config).(model.CliConfig)
}

func SetProfile(ctx context.Context, value string) context.Context {
	return context.WithValue(ctx, profile, value)
}

func GetProfile(ctx context.Context) string {
	return ctx.Value(profile).(string)
}
