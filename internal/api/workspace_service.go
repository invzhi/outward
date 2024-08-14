package api

import (
	"context"

	"connectrpc.com/connect"
	"github.com/samber/lo"

	"github.com/invzhi/outward/config"
	"github.com/invzhi/outward/internal/sqlc"
	pbv1 "github.com/invzhi/outward/proto/outward/v1"
)

type WorkspaceService struct {
	*config.AppContext
}

func NewWorkspaceService(appctx *config.AppContext) *WorkspaceService {
	return &WorkspaceService{AppContext: appctx}
}

func (s *WorkspaceService) CreateWorkspace(ctx context.Context, c *connect.Request[pbv1.CreateWorkspaceRequest]) (*connect.Response[pbv1.CreateWorkspaceResponse], error) {
	workspace, err := s.Queries.CreateWorkspace(ctx, sqlc.CreateWorkspaceParams{
		ID:     sqlc.NewID(),
		Name:   c.Msg.Name,
		Region: int32(c.Msg.Region),
	})
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(&pbv1.CreateWorkspaceResponse{
		Workspace: &pbv1.Workspace{
			Name:   workspace.Name,
			Region: pbv1.WorkspaceRegion(workspace.Region),
		},
	}), nil
}

func (s *WorkspaceService) GetWorkspaceList(ctx context.Context, c *connect.Request[pbv1.GetWorkspaceListRequest]) (*connect.Response[pbv1.GetWorkspaceListResponse], error) {
	workspaces, err := s.Queries.GetWorkspaces(ctx, 526844145155710908) // TODO
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(&pbv1.GetWorkspaceListResponse{
		Workspaces: lo.Map(workspaces, func(workspace sqlc.Workspace, _ int) *pbv1.Workspace {
			return &pbv1.Workspace{
				Name:   workspace.Name,
				Region: pbv1.WorkspaceRegion(workspace.Region),
			}
		}),
	}), nil
}
