package grpc

import (
	context "context"

	"github.com/vscode-debug-specs/go/usecase"
	grpc "google.golang.org/grpc"
)

type sorterService struct {
	UnimplementedSorterServer
}

func New() *grpc.Server {
	sv := &sorterService{}
	grpcServer := grpc.NewServer()
	RegisterSorterServer(grpcServer, sv)
	return grpcServer
}

func (sv *sorterService) SortInts(ctx context.Context, items *IntItems) (*IntItems, error) {
	input := make([]int, len(items.Items))
	itemMap := make(map[int]*IntItem)
	for i := range items.Items {
		input[i] = int(items.Items[i].Num)
		itemMap[int(items.Items[i].Num)] = items.Items[i]
	}
	output := usecase.SortInts(input)
	result := make([]*IntItem, len(items.Items))
	for i := range output {
		result[i] = itemMap[output[i]]
	}
	return &IntItems{
		Items: result,
	}, nil
}

func (sv *sorterService) SortText(ctx context.Context, items *TextItems) (*TextItems, error) {
	input := make([]string, len(items.Items))
	itemMap := make(map[string]*TextItem)
	for i := range items.Items {
		input[i] = items.Items[i].Text
		itemMap[items.Items[i].Text] = items.Items[i]
	}
	output := usecase.SortText(input)
	result := make([]*TextItem, len(items.Items))
	for i := range output {
		result[i] = itemMap[output[i]]
	}
	return &TextItems{
		Items: result,
	}, nil
}
