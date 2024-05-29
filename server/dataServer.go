package server

import (
	"context"

	pb "userdata/data" // Replace with the actual path to your generated protobuf package

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

type Server struct {
	pb.UnimplementedUserDataServer
	Users map[int32]*pb.Data
}

func (s *Server) GetData(ctx context.Context, req *pb.GetDataRequest) (*pb.GetDataResponse, error) {
	user, exists := s.Users[req.Id]
	if !exists {
		return nil, grpc.Errorf(codes.NotFound, "user not found")
	}
	return &pb.GetDataResponse{User: user}, nil
}

func (s *Server) GetAllData(ctx context.Context, req *pb.GetAllDataRequest) (*pb.GetAllDataResponse, error) {
	var users []*pb.Data
	for _, id := range req.Ids {
		if user, exists := s.Users[id]; exists {
			users = append(users, user)
		}
	}
	return &pb.GetAllDataResponse{Users: users}, nil
}

func (s *Server) SearchAllData(ctx context.Context, req *pb.SearchAllDataRequest) (*pb.SearchAllDataResponse, error) {
	var users []*pb.Data
	for _, user := range s.Users {
		if user.City == req.City && user.Phone == req.Phone && user.Married == req.Married {
			users = append(users, user)
		}
	}
	return &pb.SearchAllDataResponse{Users: users}, nil
}
