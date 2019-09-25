package main

import (
	"context"
	"net"
	"os"
	"os/signal"

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
		}).Info("Request")
	}

	response := &surveypb.FeedbackResponse{
		Result: "Message received and processed",
	}

	return response, nil
}

func main() {
	log.Info("gRPC Server starting ...")

	shutdown := make(chan os.Signal, 1)

	signal.Notify(shutdown, os.Interrupt)
	signal.Notify(shutdown, os.Kill)

	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.WithFields(log.Fields{
			"network": network,
			"address": address,
		}).Fatal("Failed to create Listener")
	}

	s := grpc.NewServer()
	surveypb.RegisterFeedbackServiceServer(s, &server{})

	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Failed to accept incoming requests: %+v", err)
		}
	}()

	<-shutdown

	log.Info("Initiate graceful shutdown here")
	os.Exit(0)
}
