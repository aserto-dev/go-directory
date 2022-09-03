package main

import (
	"fmt"
	"log"
	"net"

	"github.com/aserto-dev/go-directory/aserto/directory/exporter/v2"
	"github.com/aserto-dev/go-directory/aserto/directory/importer/v2"
	"github.com/aserto-dev/go-directory/aserto/directory/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const port = 12345

type server struct {
	directory.UnimplementedDirectoryServer
	importer.UnimplementedImporterServer
	exporter.UnimplementedExporterServer
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	directory.RegisterDirectoryServer(s, &server{})
	exporter.RegisterExporterServer(s, &server{})
	importer.RegisterImporterServer(s, &server{})

	reflection.Register(s)

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
