package repository

import (
	pb "github.com/Bekzodbekk/paris2024_livestream_protos/genproto/countrypb"
)

type CountryRepository interface {
	CreateCountry(req *pb.CreateCountryRequest) (*pb.Country, error)
	GetCountry(req *pb.GetCountryRequest) (*pb.Country, error)
	ListOfCountry(req *pb.ListOfCountryRequest) (*pb.ListOfCountryResponse, error)
	UpdateCountry(req *pb.UpdateCountryRequest) (*pb.Country, error)
	DeleteCountry(req *pb.DeleteCountryRequest) (*pb.DeleteCountryResponse, error)
}
