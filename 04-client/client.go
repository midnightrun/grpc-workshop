package main

import (
	"context"

	surveypb "github.com/midnightrun/grpc-workshop/01-protobuffer"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

const (
	target = "localhost:50051"
)

func main() {
	log.Info("gRPC Client starting ...")

	options := grpc.WithInsecure()

	conn, err := grpc.Dial(target, options)
	if err != nil {
		log.WithFields(log.Fields{
			"target":  target,
			"options": options,
			"error":   err,
		}).Fatalf("Failed to create client connection")
	}

	client := surveypb.NewFeedbackServiceClient(conn)

	request := &surveypb.FeedbackRequest{
		Feedback: []*surveypb.Feedback{
			&surveypb.Feedback{
				Expectation: "Learn about gRPC",
				Rating:      6,
				Message:     "Didn't learn anything",
			},
		},
	}

	res, err := client.Feedback(context.Background(), request)
	if err != nil {
		log.WithFields(log.Fields{
			"request": request,
			"error":   err,
		}).Fatalf("Failed to receive response from server")
	}

	log.Infof("Receive response from server: %s", res.GetResult())

	log.Info("gRPC Client shutdown ...")
}
