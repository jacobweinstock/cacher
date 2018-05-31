package main

import (
	"database/sql"
	"net"
	"os"
	"strconv"
	"strings"

	"github.com/packethost/cacher/protos/cacher"
	"github.com/packethost/packngo"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var (
	api   = "https://api.packet.net/"
	sugar *zap.SugaredLogger
)

func getMaxErrs() int {
	sMaxErrs := os.Getenv("CACHER_MAX_ERRS")
	if sMaxErrs == "" {
		sMaxErrs = "5"
	}

	max, err := strconv.Atoi(sMaxErrs)
	if err != nil {
		panic("unable to convert CACHER_MAX_ERRS to int")
	}
	return max
}

func main() {
	log, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}

	sugar = log.Sugar()
	defer log.Sync()

	if url := os.Getenv("PACKET_API_URL"); url != "" && url != api {
		api = url
		if !strings.HasSuffix(api, "/") {
			api += "/"
		}
	}

	connStr := strings.Join([]string{
		"dbname=" + os.Getenv("POSTGRES_DB"),
		"host=" + os.Getenv("POSTGRES_HOST"),
		"password=" + os.Getenv("POSTGRES_PASSWORD"),
		"sslmode=disable",
		"user=" + os.Getenv("POSTGRES_USER"),
	}, " ")
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		sugar.Fatal(err)
	}
	if err := truncate(db); err != nil {
		panic(err)
	}

	client := packngo.NewClientWithAuth(os.Getenv("PACKET_CONSUMER_TOKEN"), os.Getenv("PACKET_API_AUTH_TOKEN"), nil)

	facility := os.Getenv("FACILITY")
	sugar.Infow("starting fetch")
	data, err := fetchFacility(client, api, facility)
	sugar.Info("done fetching")
	if err != nil {
		sugar.Info(err)
	}

	sugar.Info("copying")
	if err = copyin(db, data); err != nil {
		sugar.Info(err)
	}
	sugar.Info("done copying")

	lis, err := net.Listen("tcp", clientPort)
	if err != nil {
		sugar.Fatalf("failed to listen: %v", err)
	}

	tc, err := credentials.NewServerTLSFromFile("/certs/"+facility+"/server.pem", "/certs/"+facility+"/server-key.pem")
	if err != nil {
		sugar.Fatalf("failed to read TLS files: %v", err)
	}
	s := grpc.NewServer(grpc.Creds(tc))
	cacher.RegisterCacherServer(s, &server{
		db:     db,
		packet: client,
	})

	sugar.Info("serving grpc")
	if err := s.Serve(lis); err != nil {
		sugar.Fatalf("failed to serve: %v", err)
	}
}
