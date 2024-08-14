package api

import (
	"context"
	"fmt"

	"connectrpc.com/connect"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/samber/lo"
	"golang.org/x/crypto/bcrypt"

	"github.com/invzhi/outward/config"
	"github.com/invzhi/outward/internal/sqlc"
	pbv1 "github.com/invzhi/outward/proto/outward/v1"
)

type UserService struct {
	*config.AppContext
}

func NewUserService(appctx *config.AppContext) *UserService {
	return &UserService{AppContext: appctx}
}

func (s *UserService) CreateUser(ctx context.Context, req *connect.Request[pbv1.CreateUserRequest]) (*connect.Response[pbv1.CreateUserResponse], error) {
	var passwordHash pgtype.Text
	if len(req.Msg.Password) > 0 {
		hash, err := bcrypt.GenerateFromPassword([]byte(req.Msg.Password), bcrypt.DefaultCost)
		if err != nil {
			return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("cannot generate password hash: %v", err))
		}
		passwordHash = pgtype.Text{String: string(hash), Valid: true}
	}

	user, err := s.Queries.CreateUser(ctx, sqlc.CreateUserParams{
		ID:           sqlc.NewID(),
		Email:        req.Msg.Email,
		FirstName:    req.Msg.FirstName,
		LastName:     req.Msg.LastName,
		PasswordHash: passwordHash,
	})
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(&pbv1.CreateUserResponse{
		User: &pbv1.User{
			Id:        user.ID,
			Email:     user.Email,
			FirstName: user.FirstName,
			LastName:  user.LastName,
		},
	}), nil
}

func (s *UserService) GetUserList(ctx context.Context, req *connect.Request[pbv1.GetUserListRequest]) (*connect.Response[pbv1.GetUserListResponse], error) {
	params := sqlc.GetWorkspaceMembersParams{
		WorkspaceID: req.Msg.WorkspaceId,
		Limit:       req.Msg.PageSize,
	}
	if len(req.Msg.PageToken) > 0 {
		pageToken, err := ParsePageToken(req.Msg.PageToken)
		if err != nil {
			return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("invalid page token: %v", err))
		}

		params.Limit = pageToken.PageSize
		params.Cursor = pageToken.Cursor
	}

	users, err := s.Queries.GetWorkspaceMembers(ctx, params)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("cannot get workspace members: %v", err))
	}

	var nextPageToken string
	if len(users) == int(req.Msg.PageSize) {
		nextPageToken, err = NewPageToken(&pbv1.PageToken{PageSize: req.Msg.PageSize, Cursor: users[len(users)-1].ID})
		if err != nil {
			return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("cannot get next page token: %v", err))
		}
	}

	return connect.NewResponse(&pbv1.GetUserListResponse{
		NextPageToken: nextPageToken,
		Users: lo.Map(users, func(user sqlc.User, _ int) *pbv1.User {
			return &pbv1.User{
				Id:        user.ID,
				Email:     user.Email,
				FirstName: user.FirstName,
				LastName:  user.LastName,
			}
		}),
	}), nil
}

func (s *UserService) CreateWorkspaceMember(ctx context.Context, c *connect.Request[pbv1.CreateWorkspaceMemberRequest]) (*connect.Response[pbv1.CreateWorkspaceMemberResponse], error) {
	err := s.Queries.CreateWorkspaceMember(ctx, sqlc.CreateWorkspaceMemberParams{
		WorkspaceID: c.Msg.WorkspaceId,
		UserID:      c.Msg.UserId,
		Role:        int32(c.Msg.Role),
	})
	return connect.NewResponse(&pbv1.CreateWorkspaceMemberResponse{}), err
}

func (s *UserService) Login(ctx context.Context, c *connect.Request[pbv1.LoginRequest]) (*connect.Response[pbv1.LoginResponse], error) {
	// TODO implement me
	panic("implement me")
}

func (s *UserService) Logout(ctx context.Context, c *connect.Request[pbv1.LogoutRequest]) (*connect.Response[pbv1.LogoutResponse], error) {
	// TODO implement me
	panic("implement me")
}
