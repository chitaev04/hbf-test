package connection

import (
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func GRPC(t provider.T) *grpc.ClientConn {
	conn, err := grpc.NewClient(
		"localhost:9006",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	require.NoError(t, err)
	conn.Connect()
	return conn
}
