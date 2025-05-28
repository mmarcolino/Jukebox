package api

import (
	"context"
)

func (h *Handler) Ping(ctx context.Context) (string, error) {

	return "Pong", nil
}
