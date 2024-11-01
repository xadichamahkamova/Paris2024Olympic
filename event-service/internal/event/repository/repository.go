package repository 

import (
	pb "github.com/Bekzodbekk/paris2024_livestream_protos/genproto/eventpb"
)
type EventRepository interface {
	CreateEvent(req *pb.CreateEventRequest) (*pb.Event, error)
	GetEvent(req *pb.GetEventRequest) (*pb.Event, error)
	ListOfEvent(req *pb.ListOfEventRequest) (*pb.ListOfEventResponse, error)
	UpdateEvent(*pb.UpdateEventRequest) (*pb.Event, error)
	DeleteEvent(req *pb.DeleteEventRequest) (*pb.DeleteEventResponse, error)
}