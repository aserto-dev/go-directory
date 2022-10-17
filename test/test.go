package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"testing"

	"github.com/aserto-dev/go-directory/aserto/directory/common/v2"
	"github.com/aserto-dev/go-directory/aserto/directory/exporter/v2"
	"github.com/aserto-dev/go-directory/aserto/directory/importer/v2"
	"github.com/aserto-dev/go-directory/aserto/directory/reader/v2"
	"github.com/aserto-dev/go-directory/aserto/directory/writer/v2"
	"github.com/aserto-dev/go-directory/pkg/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

type server struct {
	reader.UnimplementedReaderServer
	writer.UnimplementedWriterServer
	importer.UnimplementedImporterServer
	exporter.UnimplementedExporterServer
}

var port int

func main() {
	flag.IntVar(&port, "port", 0, "port number")
	flag.Parse()

	if port == 0 {
		testObjectType()
		testPermission()
		testRelationType()
		testObject()
		testRelation()
		os.Exit(0)
	}

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	reader.RegisterReaderServer(s, &server{})
	writer.RegisterWriterServer(s, &server{})
	exporter.RegisterExporterServer(s, &server{})
	importer.RegisterImporterServer(s, &server{})

	reflection.Register(s)

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func testObjectType() {
	req := &common.ObjectType{
		Id:          10,
		Name:        "book",
		DisplayName: "book object",
		IsSubject:   false,
		Ordinal:     1234,
		Status:      1,
		Schema:      pb.NewStruct(),
		Hash:        "01234567890",
	}
	fmt.Println(ProtoToStr(req))
}

func testPermission() {
	req := &common.Permission{
		Id:          "69595fa6-5437-4837-a091-0d282636429c",
		Name:        "book:sell",
		DisplayName: "sell book",
		Hash:        "01234567890",
	}
	fmt.Println(ProtoToStr(req))
}

func testRelationType() {
	req := &common.RelationType{
		Id:          11,
		Name:        "owner",
		ObjectType:  "book",
		DisplayName: "book::owner",
		Ordinal:     0,
		Status:      0,
		Unions:      []string{},
		Permissions: []string{"book:sell"},
		Hash:        "01234567890",
	}
	fmt.Println(ProtoToStr(req))
}

func testObject() {
	req := &common.Object{
		Id:          "b51983e9-085c-486a-acf9-da9508ff415b",
		Key:         "978-1260440218",
		Type:        "book",
		DisplayName: "Java: A Beginner's Guide, Eighth Edition 8th Edition",
		Properties:  pb.NewStruct(),
		Hash:        "01234567890",
	}
	fmt.Println(ProtoToStr(req))
}

func testRelation() {
	req := &common.Relation{
		Subject: &common.ObjectIdentifier{
			Id:   nil, //proto.String(""),
			Key:  proto.String("euang@acmecorp.com"),
			Type: proto.String("user"),
		},
		Relation: "owner",
		Object: &common.ObjectIdentifier{
			Id:   nil, //proto.String(""),
			Key:  proto.String("978-1260440218"),
			Type: proto.String("book"),
		},
		Hash: "01234567890",
	}
	fmt.Println(ProtoToStr(req))
}

func ProtoToStr(msg proto.Message) string {
	return protojson.MarshalOptions{
		Multiline:       false,
		Indent:          "  ",
		AllowPartial:    false,
		UseProtoNames:   true,
		UseEnumNumbers:  false,
		EmitUnpopulated: false,
	}.Format(msg)
}

func TestGetManyObject(t *testing.T) {
	_ = reader.GetObjectManyRequest{
		Param: []*common.ObjectIdentifier{
			{Id: proto.String("ID1")},
			{Id: proto.String("ID2")},
		},
	}
}
