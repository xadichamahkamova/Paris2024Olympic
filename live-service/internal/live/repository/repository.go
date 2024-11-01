package repository

import (
	pb "github.com/Bekzodbekk/paris2024_livestream_protos/genproto/livepb"
)

type LiveRepository interface {
	CreateLiveStream(req *pb.LiveStream) (*pb.ResponseMessage, error)
	GetLiveStream(req *pb.GetStreamRequest) (*pb.LiveStream, error)
}
