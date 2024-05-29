package server

import (
	"context"
	"testing"

	pb "userdata/data"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestServer_GetData(t *testing.T) {
	s := &Server{
		Users: map[int32]*pb.Data{
			1: {Id: 1, Fname: "Steve", City: "LA", Phone: 1234567890, Height: 5.8, Married: true},
		},
	}

	t.Run("existing user", func(t *testing.T) {
		req := &pb.GetDataRequest{Id: 1}
		res, err := s.GetData(context.Background(), req)
		assert.NoError(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, int32(1), res.User.Id)
		assert.Equal(t, "Steve", res.User.Fname)
	})

	t.Run("non-existing user", func(t *testing.T) {
		req := &pb.GetDataRequest{Id: 2}
		res, err := s.GetData(context.Background(), req)
		assert.Error(t, err)
		assert.Nil(t, res)
		st, ok := status.FromError(err)
		assert.True(t, ok)
		assert.Equal(t, codes.NotFound, st.Code())
		assert.Equal(t, "user not found", st.Message())
	})
}

func TestServer_GetAllData(t *testing.T) {
	s := &Server{
		Users: map[int32]*pb.Data{
			1: {Id: 1, Fname: "Steve", City: "LA", Phone: 1234567890, Height: 5.8, Married: true},
			2: {Id: 2, Fname: "John", City: "NY", Phone: 9876543210, Height: 6.0, Married: false},
		},
	}

	t.Run("existing users", func(t *testing.T) {
		req := &pb.GetAllDataRequest{Ids: []int32{1, 2}}
		res, err := s.GetAllData(context.Background(), req)
		assert.NoError(t, err)
		assert.NotNil(t, res)
		assert.Len(t, res.Users, 2)
	})

	t.Run("some non-existing users", func(t *testing.T) {
		req := &pb.GetAllDataRequest{Ids: []int32{1, 3}}
		res, err := s.GetAllData(context.Background(), req)
		assert.NoError(t, err)
		assert.NotNil(t, res)
		assert.Len(t, res.Users, 1)
		assert.Equal(t, int32(1), res.Users[0].Id)
	})
}

func TestServer_SearchAllData(t *testing.T) {
	s := &Server{
		Users: map[int32]*pb.Data{
			1: {Id: 1, Fname: "Steve", City: "LA", Phone: 1234567890, Height: 5.8, Married: true},
			2: {Id: 2, Fname: "John", City: "NY", Phone: 9876543210, Height: 6.0, Married: false},
		},
	}

	t.Run("match search criteria", func(t *testing.T) {
		req := &pb.SearchAllDataRequest{City: "LA", Phone: 1234567890, Married: true}
		res, err := s.SearchAllData(context.Background(), req)
		assert.NoError(t, err)
		assert.NotNil(t, res)
		assert.Len(t, res.Users, 1)
		assert.Equal(t, int32(1), res.Users[0].Id)
	})

	t.Run("no match search criteria", func(t *testing.T) {
		req := &pb.SearchAllDataRequest{City: "LA", Phone: 1111111111, Married: false}
		res, err := s.SearchAllData(context.Background(), req)
		assert.NoError(t, err)
		assert.NotNil(t, res)
		assert.Len(t, res.Users, 0)
	})
}
