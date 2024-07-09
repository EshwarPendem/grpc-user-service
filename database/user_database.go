package database

import (
	"sync"
	"github.com/eshwarpendem/grpc-user-service/proto"
)

var db *userDatabase = nil

type userDatabase struct {
	users map[uint32]*proto.User
	mu          sync.RWMutex
}

func MockDatabaseConnection() *userDatabase {
	if db == nil {
		db = &userDatabase{}
		db.PopulateMockDatabase()
	}
	return db
}

func (udb *userDatabase) PopulateMockDatabase() {
	udb.users = make(map[uint32]*proto.User)

	udb.users[1] = &proto.User{
		Id: 1,
		Fname: "Alice",
		City: "New York",
		Phone: 1234567890,
		Height: 5.6,
		Married: true,
	}
	udb.users[2] = &proto.User{
		Id: 2,
		Fname: "Bob",
		City: "New York",
		Phone: 1234567890,
		Height: 5.9,
		Married: false,
	}
	udb.users[3] = &proto.User{
		Id: 3,
		Fname: "Carol",
		City: "New York",
		Phone: 1234567890,
		Height: 5.4,
		Married: true,
	}
	udb.users[4] = &proto.User{
		Id: 4,
		Fname: "David",
		City: "Los Angeles",
		Phone: 9876543210,
		Height: 6.0,
		Married: false,
	}
	udb.users[5] = &proto.User{
		Id: 5,
		Fname: "Emily",
		City: "Los Angeles",
		Phone: 9876543210,
		Height: 6.1,
		Married: true,
	}
	udb.users[6] = &proto.User{
		Id: 6,
		Fname: "Frank",
		City: "Chicago",
		Phone: 5551112222,
		Height: 5.8,
		Married: false,
	}
	udb.users[7] = &proto.User{
		Id: 7,
		Fname: "Grace",
		City: "Houston",
		Phone: 9998887777,
		Height: 5.5,
		Married: true,
	}
	udb.users[8] = &proto.User{
		Id: 8,
		Fname: "Henry",
		City: "Miami",
		Phone: 3334445555,
		Height: 5.9,
		Married: false,
	}
	udb.users[9] = &proto.User{
		Id: 9,
		Fname: "Isabella",
		City: "Seattle",
		Phone: 1112223333,
		Height: 5.6,
		Married: true,
	}
	udb.users[10] = &proto.User{
		Id: 10,
		Fname: "Jack",
		City: "San Francisco",
		Phone: 7778889999,
		Height: 6.2,
		Married: false,
	}
	udb.users[11] = &proto.User{
		Id: 11,
		Fname: "Kate",
		City: "Boston",
		Phone: 4445556666,
		Height: 5.7,
		Married: true,
	}
	udb.users[12] = &proto.User{
		Id: 12,
		Fname: "Liam",
		City: "Denver",
		Phone: 6667778888,
		Height: 5.8,
		Married: false,
	}
	udb.users[13] = &proto.User{
		Id: 13,
		Fname: "Mira",
		City: "Portland",
		Phone: 2223334444,
		Height: 5.4,
		Married: true,
	}
	udb.users[14] = &proto.User{
		Id: 14,
		Fname: "Noah",
		City: "Austin",
		Phone: 8889990000,
		Height: 6.0,
		Married: false,
	}
	udb.users[15] = &proto.User{
		Id: 15,
		Fname: "Olivia",
		City: "Phoenix",
		Phone: 5556667777,
		Height: 5.5,
		Married: true,
	}
}

func (udb *userDatabase) GetUserById(userId uint32) (*proto.User, bool) {
	udb.mu.Lock()
	defer udb.mu.Unlock()
	user, found := udb.users[userId]
	return user, found
}

func (udb *userDatabase) Search(req *proto.SearchRequest) ([]*proto.User, bool) {
	var userList []*proto.User
	udb.mu.Lock()
	defer udb.mu.Unlock()
	for _, user := range udb.users {
		//the case where all the criteria is nil is handled in service package
		if req.UserId == nil || req.UserId.Value == user.Id {
			if req.Fname == nil || req.Fname.Value == user.Fname {
				if req.City == nil || req.City.Value == user.City {
					if req.Phone == nil || req.Phone.Value == user.Phone {
						if req.Height == nil || req.Height.Value == user.Height {
							if req.Married == nil || req.Married.Value == user.Married {
								userList = append(userList, user)
							}
						}
					}
				}
			}
		}
	}
	if len(userList) > 0 {
		return userList, true
	}
	return nil, false
}
