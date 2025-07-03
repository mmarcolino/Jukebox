package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/go-co-op/gocron/v2"
	"github.com/marcolino/jukebox/internal/domain/service"
	"github.com/marcolino/jukebox/internal/metrics"
	"github.com/marcolino/jukebox/internal/resources/queue"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	awsRegion := os.Getenv("AWS_REGION")
	awsEndpoint := os.Getenv("AWS_ENDPOINT")
	queueUrl := os.Getenv("QUEUE_URL")

	sqsWaitTime, err := strconv.Atoi(os.Getenv("SQS_TIMEOUT"))
	if err != nil {
		log.Fatalf("could not convert sqs timeout to int: %v", err)
	}

	maxMessages, err := strconv.Atoi(os.Getenv("SQS_MAXMESSAGES"))
	if err != nil {
		log.Fatalf("could not convert sqs max messages to int: %v", err)
	}

	ctx := context.Background()

	cfg, err := config.LoadDefaultConfig(ctx,
		config.WithRegion(awsRegion),
		config.WithCredentialsProvider(aws.NewCredentialsCache(credentials.NewStaticCredentialsProvider("test", "test", ""))),
		config.WithEndpointResolverWithOptions(
			aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
				return aws.Endpoint{
					URL:           awsEndpoint,
					SigningRegion: awsRegion,
				}, nil
			}),
		),
	)
	if err != nil {
		log.Fatalf("error while creating aws config: %v", err)
	}

	client := sqs.NewFromConfig(cfg)
	queueClient := queue.NewSQS(client, queueUrl, awsRegion, int32(sqsWaitTime), int32(maxMessages))

	metrics.Register()
	go func() {
		http.Handle("/metrics", promhttp.Handler())
		http.ListenAndServe(":2113", nil)
	}()
	worker := service.NewWorker(queueClient)
	
	scheduler, err := gocron.NewScheduler()
	
	if err != nil {
		log.Fatal(err)
	}

	_, err = scheduler.NewJob(gocron.DurationJob(3*time.Second), gocron.NewTask(func ()  {
		err := worker.Run(ctx)
		if err != nil{
			log.Printf("error while executing worker: %v", err)
		}
	}))
	
	if err != nil {
		log.Fatal(err)
	}

	scheduler.Start()
	select{}
}
