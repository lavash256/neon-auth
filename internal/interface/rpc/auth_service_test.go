package rpc

import (
	"context"
	"log"
	"neon-auth/internal/interface/persistence"
	rpc "neon-auth/internal/interface/rpc/protocol"
	"neon-auth/internal/usecase"
	"neon-auth/tools"
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

func dialer() func(context.Context, string) (net.Conn, error) {
	listener := bufconn.Listen(1024 * 1024)
	server := grpc.NewServer()
	accountMemoryRepository := persistence.MemoryAccountRepository{}
	stubLogger := tools.LoggerStub{}
	accountUsecase := usecase.NewAccountUsecase(&accountMemoryRepository, &stubLogger)
	rpc.RegisterAuthServiceServer(server, NewAccountService(accountUsecase))
	go func() {
		if err := server.Serve(listener); err != nil {
			log.Fatal(err)
		}
	}()

	return func(context.Context, string) (net.Conn, error) {
		return listener.Dial()
	}

}

func TestAccountService(t *testing.T) {
	tests := []struct {
		name     string
		email    string
		password string
		err      string
	}{
		{
			"Validating incorrect input values",
			"test",
			"test",
			"rpc error: code = Unknown desc = email: must be a valid email address.",
		},
		{
			"Validating empty value handling",
			"",
			"",
			"rpc error: code = Unknown desc = email: cannot be blank; password: cannot be blank.",
		},
		{
			"Correct input values",
			"test@test.ru",
			"test",
			"",
		},
	}

	ctx := context.Background()

	conn, err := grpc.DialContext(ctx, "", grpc.WithInsecure(), grpc.WithContextDialer(dialer()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := rpc.NewAuthServiceClient(conn)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := &rpc.CreateAccountRequest{Email: tt.email, Password: tt.password}

			_, err := client.CreateAccount(ctx, request)
			if err != nil {
				assert.Equal(t, tt.err, err.Error())

			} else {
				assert.Equal(t, tt.err, "")
			}

		})
	}
}
