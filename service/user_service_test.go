package service

import (
	"context"
	"testing"

	"github.com/eshwarpendem/grpc-user-service/proto"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

// For an existing user id
func TestGetUserByIdExist(t *testing.T) {
	server := &userServiceServer{}

	userID := uint32(1)
	request := &proto.UserRequest{UserId: &wrapperspb.UInt32Value{Value: userID}}
	response, err := server.GetUserById(context.Background(), request)

	if err != nil {
		t.Errorf("Error should be nil, got %v", err)
	}

	if response.GetId() != userID {
		t.Errorf("Expected user ID %d, got %d", userID, response.GetId())
	}
}

// For non existing user id
func TestGetUserByIdNotExist(t *testing.T) {
	server := &userServiceServer{}

	nonExistentUserID := &wrapperspb.UInt32Value{Value: uint32(20)}
	nonExistentRequest := &proto.UserRequest{UserId: nonExistentUserID}
	nonExistentResponse, err := server.GetUserById(context.Background(), nonExistentRequest)

	if err == nil {
		t.Errorf("Expected error, got nil")
	}

	if nonExistentResponse != nil {
		t.Errorf("nonExistentResponse should be nil, got %s", nonExistentResponse.String())
	}
}

// For all users exist
func TestGetUsersByIdsExist(t *testing.T) {
	server := &userServiceServer{}

	userIDs := []uint32{10, 2, 4}
	request := &proto.UserListRequest{UserIds: userIDs}
	response, err := server.GetUsersByIds(context.Background(), request)

	if err != nil {
		t.Errorf("Error should be nil, got %v", err)
	}

	if len(response.GetUsers()) != len(userIDs) {
		t.Errorf("Expected %d users, got %d", len(userIDs), len(response.GetUsers()))
	}
}

// For atleast one user not exists
func TestGetUsersByIdsNotExist(t *testing.T) {
	server := &userServiceServer{}

	invalidUserIDs := []uint32{0, 10, 24}
	invalidRequest := &proto.UserListRequest{UserIds: invalidUserIDs}
	invalidResponse, err := server.GetUsersByIds(context.Background(), invalidRequest)

	if err == nil {
		t.Errorf("Expected error, got nil")
	}

	if invalidResponse != nil {
		t.Errorf("invalidResponse should be nil, got %s", invalidResponse.String())
	}
}
