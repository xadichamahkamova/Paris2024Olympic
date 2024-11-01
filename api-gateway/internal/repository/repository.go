package repository

import (
	"context"

	pbUserAthlete "github.com/Bekzodbekk/paris2024_livestream_protos/genproto/athletepb"
	pbUserCountry "github.com/Bekzodbekk/paris2024_livestream_protos/genproto/countrypb"
	pbUserEvent "github.com/Bekzodbekk/paris2024_livestream_protos/genproto/eventpb"
	"github.com/Bekzodbekk/paris2024_livestream_protos/genproto/livepb"
	pbMedal "github.com/Bekzodbekk/paris2024_livestream_protos/genproto/medalspb"
	pbUser "github.com/Bekzodbekk/paris2024_livestream_protos/genproto/userpb"
)

type ServiceRepository interface {
	//User methods
	Register(ctx context.Context, req *pbUser.CreateUserRequest) (*pbUser.CreateUserResponse, error)
	Login(ctx context.Context, req *pbUser.LoginRequest) (*pbUser.LoginResponse, error)
	RefreshToken(ctx context.Context, req *pbUser.RefreshTokenRequest) (*pbUser.RefreshTokenResponse, error)
	UpdateUser(ctx context.Context, req *pbUser.UpdateUserRequest) (*pbUser.UpdateUserResponse, error)
	DeleteUser(ctx context.Context, req *pbUser.DeleteUserRequest) (*pbUser.DeleteUserResponse, error)
	GetUserById(ctx context.Context, req *pbUser.GetUserRequest) (*pbUser.GetUserResponse, error)
	GetUsers(ctx context.Context, req *pbUser.Void) (*pbUser.GetUsersResponse, error)
	GetUserByFilter(ctx context.Context, req *pbUser.UserFilter) (*pbUser.GetUsersResponse, error)

	//Model methods
	CreateMedal(req *pbMedal.CreateMedalRequest) (*pbMedal.CreateMedalResponse, error)
	UpdateMedal(req *pbMedal.UpdateMedalRequest) (*pbMedal.UpdateMedalResponse, error)
	DeleteMedal(req *pbMedal.DeleteMedalRequest) (*pbMedal.DeleteMedalResponse, error)
	GetMedalById(req *pbMedal.GetMedalByIdRequest) (*pbMedal.GetMedalByIdResponse, error)
	GetMedals(req *pbMedal.VoidMedal) (*pbMedal.GetMedalsResponse, error)
	GetMedalByFilter(req *pbMedal.GetMedalByFilterRequest) (pbMedal.GetMedalByFilterResponse, error)

	// Country methods
	CreateCountry(req *pbUserCountry.CreateCountryRequest) (*pbUserCountry.Country, error)
	GetCountry(req *pbUserCountry.GetCountryRequest) (*pbUserCountry.Country, error)
	ListOfCountry(req *pbUserCountry.ListOfCountryRequest) (*pbUserCountry.ListOfCountryResponse, error)
	UpdateCountry(req *pbUserCountry.UpdateCountryRequest) (*pbUserCountry.Country, error)
	DeleteCountry(req *pbUserCountry.DeleteCountryRequest) (*pbUserCountry.DeleteCountryResponse, error)

	// Event methods
	CreateEvent(req *pbUserEvent.CreateEventRequest) (*pbUserEvent.Event, error)
	GetEvent(req *pbUserEvent.GetEventRequest) (*pbUserEvent.Event, error)
	ListOfEvent(req *pbUserEvent.ListOfEventRequest) (*pbUserEvent.ListOfEventResponse, error)
	UpdateEvent(req *pbUserEvent.UpdateEventRequest) (*pbUserEvent.Event, error)
	DeleteEvent(req *pbUserEvent.DeleteEventRequest) (*pbUserEvent.DeleteEventResponse, error)

	// Athlete methods
	CreateAthlete(req *pbUserAthlete.CreateAthleteRequest) (*pbUserAthlete.Athlete, error)
	GetAthlete(req *pbUserAthlete.GetAthleteRequest) (*pbUserAthlete.Athlete, error)
	ListAthletes(req *pbUserAthlete.ListOfAthleteRequest) (*pbUserAthlete.ListOfAthleteResponse, error)
	UpdateAthlete(req *pbUserAthlete.UpdateAthleteRequest) (*pbUserAthlete.Athlete, error)
	DeleteAthlete(req *pbUserAthlete.DeleteAthleteRequest) (*pbUserAthlete.DeleteAthleteResponse, error)

	// Live methods
	CreateLiveStream(req *livepb.LiveStream) (*livepb.ResponseMessage, error)
	GetLiveStream(req *livepb.GetStreamRequest) (*livepb.LiveStream, error)
}
