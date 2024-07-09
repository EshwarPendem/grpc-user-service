package service

import (
	"context"
	"errors"
	"fmt"
	"math"

	"github.com/eshwarpendem/grpc-user-service/database"
	"github.com/eshwarpendem/grpc-user-service/proto"
)

type userServiceServer struct {
	proto.UserServiceServer
}

func NewUserServiceServer() *userServiceServer {
	return &userServiceServer{}
}

// GetUserById returns the requested single user detail
func (s *userServiceServer) GetUserById(ctx context.Context, req *proto.UserRequest) (*proto.User, error) {
	userId := req.GetUserId()
	if userId == nil {
		return nil, errors.New("\"user_id\" field is not provided in requst")
	}

	db := database.MockDatabaseConnection()

	user, found := db.GetUserById(userId.Value)
	if !found {
		return nil, fmt.Errorf("user not found for ID: %d", userId.Value)
	}
	return user, nil
}

// GetUsersByIds returns requested multiple user details
func (s *userServiceServer) GetUsersByIds(ctx context.Context, req *proto.UserListRequest) (*proto.UserList, error) {
	userIds := req.GetUserIds()
	if len(userIds) == 0 {
		return nil, errors.New("user IDs list is empty or not provided in request")
	}
	var userListDetails []*proto.User
	var err error = nil

	db := database.MockDatabaseConnection()

	for _, userId := range userIds {
		user, found := db.GetUserById(userId)
		if !found {
			err = errors.Join(err, fmt.Errorf("user not found for ID: %d", userId))
			continue
		}
		userListDetails = append(userListDetails, user)
	}

	if err != nil {
		return nil, err
	}

	return &proto.UserList{
		Users: userListDetails,
	}, nil
}


func (s *userServiceServer) SearchUser(ctx context.Context, req *proto.SearchRequest) (*proto.UserList, error) {
	if req.UserId==nil && req.City==nil && req.Fname==nil && req.Height==nil && req.Phone==nil && req.Married==nil {
		return nil, fmt.Errorf("provide at lead one attribue of user as a filter")
	}
	if	req.Phone != nil && (int)(math.Log10((float64)(req.Phone.Value))) != 9 {
		return nil, fmt.Errorf("provide a valid non-negative 10 digit number as phone number")
	}
	if req.Height != nil && req.Height.Value <= 0 {
		return nil, fmt.Errorf("provide a positive value for height")
	}
	var userList []*proto.User

	db := database.MockDatabaseConnection()

	userList, found := db.Search(req)
	if !found {
		return nil, fmt.Errorf("no users found with the given criteria")
	}
	return &proto.UserList{
		Users: userList,
	},nil
}