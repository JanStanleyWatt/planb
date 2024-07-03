package server

import (
	"context"
	"os"

	"connectrpc.com/connect"

	v1c "github.com/JanStanleyWatt/planb/pkg/autogen/go/api/v1"
)

type PlanBServiceServer struct{}

func (s *PlanBServiceServer) FileInfo(ctx context.Context, req *connect.Request[v1c.FileInfoRequest]) (*connect.Response[v1c.FileInfoResponse], error) {

	if err := ctx.Err(); err != nil {
		return nil, err // automatically coded correctly
	}

	path := req.Msg.Path

	info, err := os.Stat(path)
	if err != nil {
		return nil, err
	}

	return &connect.Response[v1c.FileInfoResponse]{
		Msg: &v1c.FileInfoResponse{
			Path:    path,
			Name:    info.Name(),
			Size:    info.Size(),
			ModTime: info.ModTime().Unix(),
			IsDir:   info.IsDir(),
		},
	}, nil
}

func (s PlanBServiceServer) UpdateFiles(ctx context.Context, req *connect.Request[v1c.UpdateFilesRequest]) (*connect.Response[v1c.UpdateFilesResponse], error) {
	return nil, nil
}
