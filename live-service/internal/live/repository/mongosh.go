package repository

import (
	"context"
	"fmt"
	"live-service/internal/live/pkg/mongosh"
	"live-service/logger"

	pb "github.com/Bekzodbekk/paris2024_livestream_protos/genproto/livepb"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

type MongoshLiveRepository struct {
	Client mongosh.Mongo
}

func NewMongoshLiveRepository(client mongosh.Mongo) LiveRepository {
	return &MongoshLiveRepository{
		Client: client,
	}
}

func (db *MongoshLiveRepository) CreateLiveStream(req *pb.LiveStream) (*pb.ResponseMessage, error) {
	_, err := db.Client.Collection.InsertOne(context.Background(), req)
	if err != nil {
		logger.Error("Failed to create live stream: ", logrus.Fields{
			"error":err,
		})
		return nil, fmt.Errorf("failed to create live stream: %v", err)
	}

	logger.Info("Live stream created successfully: ", logrus.Fields{
		"left_side":req.LeftSide,
		"right_side":req.RightSide,
	})
	return &pb.ResponseMessage{
		Status:  "success",
		Message: "Live stream created successfully",
	}, nil
}

func (db *MongoshLiveRepository) GetLiveStream(req *pb.GetStreamRequest) (*pb.LiveStream, error) {
	var result pb.LiveStream
	err := db.Client.Collection.FindOne(context.Background(), bson.M{"event_id": req.Id}).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			logger.Warn("Live stream not found: ", logrus.Fields{
				"event_id":req.Id,
			})
			return nil, fmt.Errorf("live stream not found")
		}
		logger.Error("Failed to get live stream: ", logrus.Fields{
				"event_id":req.Id,
		})
		return nil, fmt.Errorf("failed to get live stream: %v", err)
	}
	logger.Info("Get live stream is successfully complete: ", logrus.Fields{
		"left_side": result.LeftSide,
		"right_side": result.RightSide,
	})
	return &result, nil
}
