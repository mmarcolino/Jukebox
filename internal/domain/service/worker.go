package service

import (
	"context"
	"fmt"

	"github.com/marcolino/jukebox/internal/domain/gateway"
	"github.com/marcolino/jukebox/internal/metrics"
)

type Worker struct {
	queueClient gateway.Queue
}

func NewWorker(queueClient gateway.Queue) *Worker {
	return &Worker{
		queueClient: queueClient,
	}
}

func (w *Worker) Run(ctx context.Context) error {
	tracks, err := w.queueClient.ReceiveTracks(ctx)
	if err != nil {
		return err
	}

	for _, track := range tracks {
		fmt.Printf("Now Playing %s - %s", track.Title, track.Artist)
		metrics.TracksStreams.WithLabelValues(track.ID).Inc()
	}

	return nil
}
