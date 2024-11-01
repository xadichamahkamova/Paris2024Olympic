package service

import (
	"context"
	pb "github.com/Bekzodbekk/paris2024_livestream_protos/genproto/eventpb"
	"event-service/internal/event/repository"
)

type EventService struct {
	pb.UnimplementedEventServiceServer
	Repo repository.EventRepository
} 

func NewEventService(repo repository.EventRepository) *EventService {
	return &EventService{
		Repo: repo,
	}
}

func(s *EventService) CreateEvent(ctx context.Context,req *pb.CreateEventRequest) (*pb.Event, error) {
	return s.Repo.CreateEvent(req)
} 

func(s *EventService) GetEvent(ctx context.Context,req *pb.GetEventRequest) (*pb.Event, error) {
	return s.Repo.GetEvent(req)
}

func(s *EventService) ListOfEvent(ctx context.Context,req *pb.ListOfEventRequest) (*pb.ListOfEventResponse, error) {
	return s.Repo.ListOfEvent(req)
}

func(s *EventService) UpdateEvent(ctx context.Context, req *pb.UpdateEventRequest) (*pb.Event, error) {
	return s.Repo.UpdateEvent(req)
}

func(s *EventService) DeleteEvent(ctx context.Context,req *pb.DeleteEventRequest) (*pb.DeleteEventResponse, error) {
	return s.Repo.DeleteEvent(req)
}