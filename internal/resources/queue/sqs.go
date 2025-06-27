package queue

import (
	"context"
	"encoding/json"

	sqs "github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/marcolino/jukebox/internal/domain/entity"
	"github.com/marcolino/jukebox/internal/domain/gateway"
	"github.com/marcolino/jukebox/internal/utils"
)

type SQSClient interface {
	SendMessage(ctx context.Context, params *sqs.SendMessageInput, optFns ...func(*sqs.Options)) (*sqs.SendMessageOutput, error)
	ReceiveMessage(ctx context.Context, params *sqs.ReceiveMessageInput, optFns ...func(*sqs.Options)) (*sqs.ReceiveMessageOutput, error)
}

type SQS struct {
	SqsClient   SQSClient
	QueueUrl    string
	Region      string
	WaitTime    int32
	MaxMessages int32
}

var _ gateway.Queue = (*SQS)(nil)

func NewSQS(sqsClient SQSClient, queueUrl string, region string, WaitTime, MaxMessagens int32) *SQS {
	return &SQS{
		SqsClient: sqsClient,
		QueueUrl:  queueUrl,
		Region:    region,
	}
}

func (s *SQS) AddTrackToQueue(ctx context.Context, track entity.Track) error {
	trackAsBytes, err := json.Marshal(track)
	if err != nil {
		return err
	}

	trackAsString := string(trackAsBytes)

	_, err = s.SqsClient.SendMessage(ctx, &sqs.SendMessageInput{
		QueueUrl:    &s.QueueUrl,
		MessageBody: &trackAsString,
	})

	return err
}

func (s *SQS) ReceiveTracks(ctx context.Context) ([]entity.Track, error) {
	var tracks []entity.Track

	for {
		resp, err := s.SqsClient.ReceiveMessage(ctx, &sqs.ReceiveMessageInput{
			QueueUrl:            &s.QueueUrl,
			MaxNumberOfMessages: s.MaxMessages,
			WaitTimeSeconds:     s.WaitTime,
		})

		if err != nil {
			return nil, err
		}

		if len(resp.Messages) == 0 {
			break
		}

		for _, message := range resp.Messages {
			track := entity.Track{}

			err := json.Unmarshal([]byte(utils.PointerToString(message.Body)), &track)
			if err != nil {
				return nil, err
			}

			tracks = append(tracks, track)
		}
	}

	return tracks, nil

}
