package repository

import (
    pb "github.com/Bekzodbekk/paris2024_livestream_protos/genproto/athletepb"
)

type AthleteRepository interface {
    CreateAthlete(req *pb.CreateAthleteRequest) (*pb.Athlete, error)
    GetAthlete(req *pb.GetAthleteRequest) (*pb.GetAthleteResponse, error)
    ListAthletes(req *pb.ListOfAthleteRequest) (*pb.ListOfAthleteResponse, error)
    UpdateAthlete(req *pb.UpdateAthleteRequest) (*pb.Athlete, error)
    DeleteAthlete(req *pb.DeleteAthleteRequest) (*pb.DeleteAthleteResponse, error)
}
