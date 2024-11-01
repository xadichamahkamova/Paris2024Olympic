package service

import (
	"context"

	pbAthlete "github.com/Bekzodbekk/paris2024_livestream_protos/genproto/athletepb"
	pbCountry "github.com/Bekzodbekk/paris2024_livestream_protos/genproto/countrypb"
	pbEvent "github.com/Bekzodbekk/paris2024_livestream_protos/genproto/eventpb"
	"github.com/Bekzodbekk/paris2024_livestream_protos/genproto/livepb"
	pbLive "github.com/Bekzodbekk/paris2024_livestream_protos/genproto/livepb"
	pbMedal "github.com/Bekzodbekk/paris2024_livestream_protos/genproto/medalspb"
	pbUser "github.com/Bekzodbekk/paris2024_livestream_protos/genproto/userpb"
)

type ServiceRepositoryClient struct {
	userClient    pbUser.UserServiceClient
	medalClient   pbMedal.MedalServiceClient
	countryClient pbCountry.CountryServiceClient
	eventClient   pbEvent.EventServiceClient
	athleteClient pbAthlete.AthleteServiceClient
	liveClient    pbLive.LiveStreamServiceClient
}

func NewServiceRepositoryClient(
	conn1 *pbUser.UserServiceClient,
	conn2 *pbMedal.MedalServiceClient,
	conn3 *pbCountry.CountryServiceClient,
	conn4 *pbEvent.EventServiceClient,
	conn5 *pbAthlete.AthleteServiceClient,
	conn6 *pbLive.LiveStreamServiceClient,
) *ServiceRepositoryClient {
	return &ServiceRepositoryClient{
		userClient:    *conn1,
		medalClient:   *conn2,
		countryClient: *conn3,
		eventClient:   *conn4,
		athleteClient: *conn5,
		liveClient:    *conn6,
	}
}

//User methods

func (s *ServiceRepositoryClient) Register(ctx context.Context, req *pbUser.CreateUserRequest) (*pbUser.CreateUserResponse, error) {
	return s.userClient.Register(ctx, req)
}

func (s *ServiceRepositoryClient) Login(ctx context.Context, req *pbUser.LoginRequest) (*pbUser.LoginResponse, error) {
	return s.userClient.Login(ctx, req)
}

func (s *ServiceRepositoryClient) RefreshToken(ctx context.Context, req *pbUser.RefreshTokenRequest) (*pbUser.RefreshTokenResponse, error) {
	return s.userClient.RefreshToken(ctx, req)
}

func (s *ServiceRepositoryClient) UpdateUser(ctx context.Context, req *pbUser.UpdateUserRequest) (*pbUser.UpdateUserResponse, error) {
	return s.userClient.UpdateUser(ctx, req)
}

func (s *ServiceRepositoryClient) DeleteUser(ctx context.Context, req *pbUser.DeleteUserRequest) (*pbUser.DeleteUserResponse, error) {
	return s.userClient.DeleteUser(ctx, req)
}

func (s *ServiceRepositoryClient) GetUserById(ctx context.Context, req *pbUser.GetUserRequest) (*pbUser.GetUserResponse, error) {
	return s.userClient.GetUserById(ctx, req)
}

func (s *ServiceRepositoryClient) GetUsers(ctx context.Context, req *pbUser.Void) (*pbUser.GetUsersResponse, error) {
	return s.userClient.GetUsers(ctx, req)
}

func (s *ServiceRepositoryClient) GetUserByFilter(ctx context.Context, req *pbUser.UserFilter) (*pbUser.GetUsersResponse, error) {
	return s.userClient.GetUserByFilter(ctx, req)
}

// Medal methods
func (s *ServiceRepositoryClient) CreateMedal(ctx context.Context, req *pbMedal.CreateMedalRequest) (*pbMedal.CreateMedalResponse, error) {
	return s.medalClient.CreateMedal(ctx, req)
}

func (s *ServiceRepositoryClient) UpdateMedal(ctx context.Context, req *pbMedal.UpdateMedalRequest) (*pbMedal.UpdateMedalResponse, error) {
	return s.medalClient.UpdateMedal(ctx, req)
}

func (s *ServiceRepositoryClient) DeleteMedal(ctx context.Context, req *pbMedal.DeleteMedalRequest) (*pbMedal.DeleteMedalResponse, error) {
	return s.medalClient.DeleteMedal(ctx, req)
}

func (s *ServiceRepositoryClient) GetMedalById(ctx context.Context, req *pbMedal.GetMedalByIdRequest) (*pbMedal.GetMedalByIdResponse, error) {
	return s.medalClient.GetMedalById(ctx, req)
}

func (s *ServiceRepositoryClient) GetMedals(ctx context.Context, req *pbMedal.VoidMedal) (*pbMedal.GetMedalsResponse, error) {
	return s.medalClient.GetMedals(ctx, req)
}

func (s *ServiceRepositoryClient) GetMedalByFilter(ctx context.Context, req *pbMedal.GetMedalByFilterRequest) (*pbMedal.GetMedalByFilterResponse, error) {
	return s.medalClient.GetMedalByFilter(ctx, req)
}

// Country methods
func (s *ServiceRepositoryClient) CreateCountry(req *pbCountry.CreateCountryRequest) (*pbCountry.Country, error) {
	return s.countryClient.CreateCountry(context.Background(), req)
}

func (s *ServiceRepositoryClient) GetCountry(req *pbCountry.GetCountryRequest) (*pbCountry.Country, error) {
	return s.countryClient.GetCountry(context.Background(), req)
}

func (s *ServiceRepositoryClient) ListOfCountry(req *pbCountry.ListOfCountryRequest) (*pbCountry.ListOfCountryResponse, error) {
	return s.countryClient.ListOfCountry(context.Background(), req)
}

func (s *ServiceRepositoryClient) UpdateCountry(req *pbCountry.UpdateCountryRequest) (*pbCountry.Country, error) {
	return s.countryClient.UpdateCountry(context.Background(), req)
}

func (s *ServiceRepositoryClient) DeleteCountry(req *pbCountry.DeleteCountryRequest) (*pbCountry.DeleteCountryResponse, error) {
	return s.countryClient.DeleteCountry(context.Background(), req)
}

// Event methods
func (s *ServiceRepositoryClient) CreateEvent(req *pbEvent.CreateEventRequest) (*pbEvent.Event, error) {
	return s.eventClient.CreateEvent(context.Background(), req)
}

func (s *ServiceRepositoryClient) GetEvent(req *pbEvent.GetEventRequest) (*pbEvent.Event, error) {
	return s.eventClient.GetEvent(context.Background(), req)
}

func (s *ServiceRepositoryClient) ListOfEvent(req *pbEvent.ListOfEventRequest) (*pbEvent.ListOfEventResponse, error) {
	return s.eventClient.ListOfEvent(context.Background(), req)
}

func (s *ServiceRepositoryClient) UpdateEvent(req *pbEvent.UpdateEventRequest) (*pbEvent.Event, error) {
	return s.eventClient.UpdateEvent(context.Background(), req)
}

func (s *ServiceRepositoryClient) DeleteEvent(req *pbEvent.DeleteEventRequest) (*pbEvent.DeleteEventResponse, error) {
	return s.eventClient.DeleteEvent(context.Background(), req)
}

// Athlete methods
func (s *ServiceRepositoryClient) CreateAthlete(req *pbAthlete.CreateAthleteRequest) (*pbAthlete.Athlete, error) {
	return s.athleteClient.CreateAthlete(context.Background(), req)
}

func (s *ServiceRepositoryClient) GetAthlete(req *pbAthlete.GetAthleteRequest) (*pbAthlete.GetAthleteResponse, error) {
	return s.athleteClient.GetAthlete(context.Background(), req)
}

func (s *ServiceRepositoryClient) ListOfAthlete(req *pbAthlete.ListOfAthleteRequest) (*pbAthlete.ListOfAthleteResponse, error) {
	return s.athleteClient.ListOfAthlete(context.Background(), req)
}

func (s *ServiceRepositoryClient) UpdateAthlete(req *pbAthlete.UpdateAthleteRequest) (*pbAthlete.Athlete, error) {
	return s.athleteClient.UpdateAthlete(context.Background(), req)
}

func (s *ServiceRepositoryClient) DeleteAthlete(req *pbAthlete.DeleteAthleteRequest) (*pbAthlete.DeleteAthleteResponse, error) {
	return s.athleteClient.DeleteAthlete(context.Background(), req)
}

// Live methods

func(s *ServiceRepositoryClient) CreateLive(req *livepb.LiveStream) (*livepb.ResponseMessage, error){
	return s.liveClient.CreateLiveStream(context.Background(), req)
}

func(s *ServiceRepositoryClient) GetLive(req *livepb.GetStreamRequest) (*livepb.LiveStream, error){
	return s.liveClient.GetLiveStream(context.Background(), req)
}