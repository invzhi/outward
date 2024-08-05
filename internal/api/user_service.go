package api

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/samber/lo"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/invzhi/outward/config"
	"github.com/invzhi/outward/internal/sqlc"
	pbv1 "github.com/invzhi/outward/proto/outward/v1"
)

type UserServer struct {
	pbv1.UnimplementedUserServiceServer
	*config.AppContext
}

func NewUserServer(appctx *config.AppContext) *UserServer {
	return &UserServer{AppContext: appctx}
}

func (s *UserServer) CreateUser(ctx context.Context, req *pbv1.CreateUserRequest) (*pbv1.CreateUserResponse, error) {
	var passwordHash pgtype.Text
	if len(req.Password) > 0 {
		hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "cannot generate password hash: %v", err)
		}
		passwordHash = pgtype.Text{String: string(hash), Valid: true}
	}

	user, err := s.Queries.CreateUser(ctx, sqlc.CreateUserParams{
		ID:           sqlc.NewID(),
		Email:        req.Email,
		FirstName:    req.FirstName,
		LastName:     req.LastName,
		PasswordHash: passwordHash,
	})
	if err != nil {
		return nil, err
	}

	return &pbv1.CreateUserResponse{
		User: &pbv1.User{
			Id:        user.ID,
			Email:     user.Email,
			FirstName: user.FirstName,
			LastName:  user.LastName,
		},
	}, nil
}

func (s *UserServer) GetUserList(ctx context.Context, req *pbv1.GetUserListRequest) (*pbv1.GetUserListResponse, error) {
	params := sqlc.GetWorkspaceMembersParams{
		WorkspaceID: req.WorkspaceId,
		Limit:       req.PageSize,
	}
	if len(req.PageToken) > 0 {
		pageToken, err := ParsePageToken(req.PageToken)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "invalid page token: %v", err)
		}

		params.Limit = pageToken.PageSize
		params.Cursor = pageToken.Cursor
	}

	users, err := s.Queries.GetWorkspaceMembers(ctx, params)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot get workspace members: %v", err)
	}

	var nextPageToken string
	if len(users) == int(req.PageSize) {
		nextPageToken, err = NewPageToken(&pbv1.PageToken{PageSize: req.PageSize, Cursor: users[len(users)-1].ID})
		if err != nil {
			return nil, status.Errorf(codes.Internal, "cannot get next page token: %v", err)
		}
	}

	return &pbv1.GetUserListResponse{
		NextPageToken: nextPageToken,
		Users: lo.Map(users, func(user sqlc.User, _ int) *pbv1.User {
			return &pbv1.User{
				Id:        user.ID,
				Email:     user.Email,
				FirstName: user.FirstName,
				LastName:  user.LastName,
			}
		}),
	}, nil
}

func (s *UserServer) CreateWorkspaceMember(ctx context.Context, req *pbv1.CreateWorkspaceMemberRequest) (*pbv1.CreateWorkspaceMemberResponse, error) {
	err := s.Queries.CreateWorkspaceMember(ctx, sqlc.CreateWorkspaceMemberParams{
		WorkspaceID: req.WorkspaceId,
		UserID:      req.UserId,
		Role:        int32(req.Role),
	})
	return &pbv1.CreateWorkspaceMemberResponse{}, err
}

func (s *UserServer) Login(ctx context.Context, req *pbv1.LoginRequest) (*pbv1.LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}

func (s *UserServer) Logout(ctx context.Context, req *pbv1.LogoutRequest) (*pbv1.LogoutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Logout not implemented")
}
