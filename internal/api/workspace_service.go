package api

import (
	"context"

	"github.com/samber/lo"

	"github.com/invzhi/outward/config"
	"github.com/invzhi/outward/internal/sqlc"
	pbv1 "github.com/invzhi/outward/proto/outward/v1"
)

type WorkspaceServer struct {
	pbv1.UnimplementedWorkspaceServiceServer
	*config.AppContext
}

func NewWorkspaceServer(appctx *config.AppContext) *WorkspaceServer {
	return &WorkspaceServer{AppContext: appctx}
}

func (s *WorkspaceServer) CreateWorkspace(ctx context.Context, req *pbv1.CreateWorkspaceRequest) (*pbv1.CreateWorkspaceResponse, error) {
	workspace, err := s.Queries.CreateWorkspace(ctx, sqlc.CreateWorkspaceParams{
		ID:     sqlc.NewID(),
		Name:   req.Name,
		Region: int32(req.Region),
	})
	if err != nil {
		return nil, err
	}

	return &pbv1.CreateWorkspaceResponse{
		Workspace: &pbv1.Workspace{
			Name:   workspace.Name,
			Region: req.Region,
		},
	}, nil
}

func (s *WorkspaceServer) GetWorkspaceList(ctx context.Context, req *pbv1.GetWorkspaceListRequest) (*pbv1.GetWorkspaceListResponse, error) {
	workspaces, err := s.Queries.GetWorkspaces(ctx, 526844145155710908) // TODO
	if err != nil {
		return nil, err
	}

	return &pbv1.GetWorkspaceListResponse{
		Workspaces: lo.Map(workspaces, func(workspace sqlc.Workspace, _ int) *pbv1.Workspace {
			return &pbv1.Workspace{
				Name:   workspace.Name,
				Region: pbv1.WorkspaceRegion(workspace.Region),
			}
		}),
	}, nil
}
