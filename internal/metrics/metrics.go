package metrics

import "github.com/prometheus/client_golang/prometheus"

func Register() {
	prometheus.MustRegister(TracksCreated)
	prometheus.MustRegister(TracksStreams)
}

var(
	TracksCreated = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "tacks_created_total",
		Help: "Total tracks created",
	}) 
	TracksStreams = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "tracks_streamed_total",
		Help: "The total number of streams of a track",
	}, 
	[]string {"track_ID"},
	)
)
