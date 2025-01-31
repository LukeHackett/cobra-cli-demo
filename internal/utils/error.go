package utils

import (
	"fmt"
	"log/slog"
	"os"
)

func Die(err error, code int) {
	slog.Error(err.Error())
	fmt.Fprintln(os.Stderr, err)
	os.Exit(code)
}

func DieWithMessage(err error, message string, code int) {
	slog.Error(message, err.Error())
	fmt.Fprintln(os.Stderr, message, err)
	os.Exit(code)
}
