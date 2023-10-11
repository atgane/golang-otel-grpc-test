package internal

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// 3.1. grpc 클라이언트를 생성하기 위한 코드를 작성합니다.
// 테스트용으로 구성하여 secure옵션은 따로 설정하지 않았습니다.
func CreateClient(ctx context.Context, addr string) (conn *grpc.ClientConn, err error) {
	return grpc.DialContext(ctx, addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
}
