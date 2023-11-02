package main

import (
    "context"
    "log"
    "net"
    "google.golang.org/grpc"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"
    "user"
)

var users = map[int64]*user.User{
    1: {
        Id:      1,
        Fname:   "Steve",
        City:    "LA",
        Phone:   1234567890,
        Height:  5.8,
        Married: true,
    }, 
	2: {
        Id:      2,
        Fname:   "James",
        City:    "NY",
        Phone:   9876543210,
        Height:  5.5,
        Married: false,
    },
    3: {
        Id:      3,
        Fname:   "Mark",
        City:    "LA",
        Phone:   1234567899,
        Height:  6.0,
        Married: true,
    },
}

type userServiceServer struct {
    user.UnimplementedUserServiceServer
}

func (s *userServiceServer) GetUserById(ctx context.Context, req *user.UserRequest) (*user.User, error) {
    id := req.Id

    if id <= 0 {
        return nil, status.Errorf(codes.InvalidArgument, "Invalid user ID")
    }

    user, ok := users[id]
    if !ok {
        return nil, status.Errorf(codes.NotFound, "User not found")
    }
    return user, nil
}

func (s *userServiceServer) GetUsersByIds(ctx context.Context, req *user.UserIdsRequest) (*user.User, error) {
    var userSlice []*user.User
    for _, id := range req.Ids {
        if id <= 0 {
            return nil, status.Errorf(codes.InvalidArgument, "Invalid user ID")
        }

        u, ok := users[id]
        if ok {
            userSlice = append(userSlice, u)
        }
    }
    return &user.User{Users: userSlice}, nil
}

func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("Failed to listen: %v", err)
    }

    s := grpc.NewServer()
    user.RegisterUserServiceServer(s, &userServiceServer{})

    if err := s.Serve(lis); err != nil {
        log.Fatalf("Failed to serve: %v", err)
    }
}
