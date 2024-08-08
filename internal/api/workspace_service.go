package api

import (
	"context"

	"github.com/samber/lo"

	"github.com/invzhi/outward/config"
	"github.com/invzhi/outward/internal/id"
	"github.com/invzhi/outward/internal/sqlc"
	proto "github.com/invzhi/outward/proto/outward/v1"
)

type WorkspaceServer struct {
	proto.UnimplementedWorkspaceServiceServer
	*config.AppContext
}

func NewWorkspaceServer(appctx *config.AppContext) *WorkspaceServer {
	return &WorkspaceServer{AppContext: appctx}
}

func (s *WorkspaceServer) CreateWorkspace(ctx context.Context, req *proto.CreateWorkspaceRequest) (*proto.Workspace, error) {
	workspace, err := s.Queries.CreateWorkspace(ctx, sqlc.CreateWorkspaceParams{
		ID:     id.New(),
		Name:   req.Name,
		Region: req.Region,
	})
	if err != nil {
		return nil, err
	}

	return &proto.Workspace{Name: workspace.Name, Region: req.Region}, nil
}

func (s *WorkspaceServer) GetWorkspaceList(ctx context.Context, req *proto.GetWorkspaceListRequest) (*proto.GetWorkspaceListResponse, error) {
	workspaces, err := s.Queries.GetWorkspaces(ctx, sqlc.GetWorkspacesParams{
		UserID: 1, // TODO
		Limit:  req.PageSize,
	})
	if err != nil {
		return nil, err
	}

	return &proto.GetWorkspaceListResponse{
		Workspaces: lo.Map(workspaces, func(workspace sqlc.Workspace, _ int) *proto.Workspace {
			return &proto.Workspace{
				Name:   workspace.Name,
				Region: workspace.Region,
			}
		}),
		NextPageToken: "fdlf",
	}, nil
}
