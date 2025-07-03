package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/marcolino/jukebox/gen/openapi"
	"github.com/marcolino/jukebox/internal/api"
	"github.com/marcolino/jukebox/internal/metrics"
	"github.com/marcolino/jukebox/internal/resources/database/postgres"
	"github.com/marcolino/jukebox/internal/resources/queue"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	// db_url := os.Getenv("DATABASE_URL")
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}

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


	postgresHandler := postgres.New(db)
	handler := api.NewHandler(postgresHandler, postgresHandler, queueClient)

	server, err := openapi.NewServer(handler)
	if err != nil {
		log.Fatal(err)
	}
	
	metrics.Register()
	go func() {
		http.Handle("/metrics", promhttp.Handler())
		http.ListenAndServe(":2112", nil)
	}()

	if err = http.ListenAndServe(":9090", server); err != nil {
		log.Fatal(err)
	}

}
