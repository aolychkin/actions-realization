package board

import (
	"context"

	board_v1 "github.com/aolychkin/actions-contract/gen/go/board"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Обработка запросов, описанных в server.go
func (s *serverAPI) GetBoard(
	ctx context.Context,
	req *board_v1.GetBoardRequest,
) (*board_v1.GetBoardResponse, error) {
	if req.Id == "" {
		return nil, status.Error(codes.InvalidArgument, "board id is required")
	}
	board, err := s.board.GetBoard(ctx, req.GetId())

	return &board_v1.GetBoardResponse{
		Board: board,
	}, err
}
