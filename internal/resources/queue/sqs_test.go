package queue_test

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
	"github.com/marcolino/jukebox/internal/domain/entity"
	"github.com/marcolino/jukebox/internal/resources/queue"
	"github.com/marcolino/jukebox/internal/utils"
	"github.com/marcolino/jukebox/test/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestAddTrackToQueue(t *testing.T) {
	ctx := context.Background()
	queueUrl := "url"

	trackSent := entity.Track{
		ID:       "01JX3872K622GTRCCVXHXVP8ZY",
		Title:    "Next Semester",
		Artist:   "Twenty One Pilots",
		Album:    "Clancy",
		Genre:    "Rock",
		Duration: 249,
	}
	for name, scenario := range map[string]struct {
		track         entity.Track
		trackAsString string
		expectedError error
	}{
		"success": {
			track:         trackSent,
			trackAsString: "{\"ID\":\"01JX3872K622GTRCCVXHXVP8ZY\",\"Title\":\"Next Semester\",\"Artist\":\"Twenty One Pilots\",\"Album\":\"Clancy\",\"Genre\":\"Rock\",\"Duration\":249}",
			expectedError: nil,
		},
		"failure": {
			track:         entity.Track{},
			trackAsString: "{\"ID\":\"\",\"Title\":\"\",\"Artist\":\"\",\"Album\":\"\",\"Genre\":\"\",\"Duration\":0}",
			expectedError: entity.GenericErr,
		},
	} {
		t.Run(name, func(t *testing.T) {
			sqsClientMock := mocks.NewMockSQSClient(t)

			sqsClientMock.On("SendMessage", ctx, mock.AnythingOfType("*sqs.SendMessageInput")).Return(nil, scenario.expectedError).
				Run(func(args mock.Arguments) {
					messageInput, ok := args.Get(1).(*sqs.SendMessageInput)
					require.True(t, ok)

					assert.Equal(t, scenario.trackAsString, utils.PointerToString(messageInput.MessageBody))

				})
			sqsQueue := queue.NewSQS(sqsClientMock, queueUrl, "region", 1, 1)
			err := sqsQueue.AddTrackToQueue(ctx, scenario.track)
			assert.ErrorIs(t, err, scenario.expectedError)
		})
	}
}

func TestReceiveTracks(t *testing.T) {
	ctx := context.Background()

	queueUrl := "queueUrl"
	maxNumberOfMessages := 1
	waitTimeSeconds := 1

	tracks := []entity.Track{
		{
			ID:       "01JX3872K622GTRCCVXHXVP8ZY",
			Title:    "Next Semester",
			Artist:   "Twenty One Pilots",
			Album:    "Clancy",
			Genre:    "Rock",
			Duration: 249,
		},
	}

	for name, scenario := range map[string]struct{
		messageBody string
		expectedError error
		expectedTracks []entity.Track
	}{
		"success":{
			messageBody: "{\"ID\":\"01JX3872K622GTRCCVXHXVP8ZY\",\"Title\":\"Next Semester\",\"Artist\":\"Twenty One Pilots\",\"Album\":\"Clancy\",\"Genre\":\"Rock\",\"Duration\":249}",
			expectedError: nil,
			expectedTracks: tracks,
		},
		"failure":{
			messageBody: "",
			expectedError: entity.GenericErr,
			expectedTracks: nil,
		},
	}{
		t.Run(name, func(t *testing.T) {
			sqsClientMock := mocks.NewMockSQSClient(t)
			sqsClientMock.On("ReceiveMessage", ctx, mock.AnythingOfType("*sqs.ReceiveMessageInput")).Return(&sqs.ReceiveMessageOutput{Messages: []types.Message{{Body: &scenario.messageBody}}}, scenario.expectedError).Times(1)
			
			if scenario.expectedError == nil{
				sqsClientMock.On("ReceiveMessage", ctx, mock.AnythingOfType("*sqs.ReceiveMessageInput")).Return(&sqs.ReceiveMessageOutput{Messages: []types.Message{}}, scenario.expectedError).Times(1)
			}
			
			sqs := queue.NewSQS(sqsClientMock, queueUrl, "region", int32(waitTimeSeconds), int32(maxNumberOfMessages))
			tracks, err := sqs.ReceiveTracks(ctx)
			assert.ErrorIs(t, err, scenario.expectedError)
			assert.Equal(t, scenario.expectedTracks, tracks)
		})
	}
}
