package main

import (
	"context"
	"net"

	surveypb "github.com/midnightrun/grpc-workshop/01-protobuffer"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

const (
	network = "tcp"
	address = "localhost:50051"
)

type server struct{}

func (s server) Feedback(ctx context.Context, r *surveypb.FeedbackRequest) (*surveypb.FeedbackResponse, error) {
	feedbacks := r.GetFeedback()

	for _, feedback := range feedbacks {
		expectation := feedback.GetExpectation()
		message := feedback.GetMessage()
		rating := feedback.GetRating()

		log.WithFields(log.Fields{
			"expectation": expectation,
			"message":     message,
			"rating":      rating,
		}).Info("Rating")
	}

	response := &surveypb.FeedbackResponse{
		Result: "Message received and processed",
	}

	return response, nil
}

func main() {
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.WithFields(log.Fields{
			"network": network,
			"address": address,
		}).Fatal("Failed to create Listener")
	}

	s := grpc.NewServer()
	surveypb.RegisterFeedbackServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to accept incoming requests: %+v", err)
	}
}
