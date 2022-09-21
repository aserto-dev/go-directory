package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/aserto-dev/go-directory/aserto/directory/common/v2"
	"github.com/aserto-dev/go-directory/aserto/directory/exporter/v2"
	"github.com/aserto-dev/go-directory/aserto/directory/importer/v2"
	"github.com/aserto-dev/go-directory/aserto/directory/v2"
	"github.com/aserto-dev/go-directory/aserto/directory/writer/v2"
	"github.com/aserto-dev/go-utils/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

type server struct {
	directory.UnimplementedDirectoryServer
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

	directory.RegisterDirectoryServer(s, &server{})
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
		Id:          nil, //proto.Int32(10),
		Name:        proto.String("book"),
		DisplayName: proto.String("book object"),
		IsSubject:   proto.Bool(false),
		Ordinal:     proto.Int32(1234),
		Status:      proto.Uint32(1),
		Schema:      pb.NewStruct(),
		Hash:        nil, //proto.String("01234567890"),
	}
	fmt.Println(ProtoToStr(req))
}

func testPermission() {
	req := &common.Permission{
		Id:          nil, //proto.String("69595fa6-5437-4837-a091-0d282636429c"),
		Name:        proto.String("book:sell"),
		DisplayName: proto.String("sell book"),
		Hash:        nil, //proto.String("01234567890"),
	}
	fmt.Println(ProtoToStr(req))
}

func testRelationType() {
	req := &common.RelationType{
		Id:          nil, //proto.Int32(11),
		Name:        proto.String("owner"),
		ObjectType:  proto.String("book"),
		DisplayName: proto.String("book::owner"),
		Ordinal:     proto.Int32(0),
		Status:      proto.Uint32(0),
		Unions:      []string{},
		Permissions: []string{"book:sell"},
		Hash:        nil, //proto.String("01234567890"),
	}
	fmt.Println(ProtoToStr(req))
}

func testObject() {
	req := &common.Object{
		Id:          nil, //proto.String("b51983e9-085c-486a-acf9-da9508ff415b"),
		Key:         proto.String("978-1260440218"),
		Type:        proto.String("book"),
		DisplayName: proto.String("Java: A Beginner's Guide, Eighth Edition 8th Edition"),
		Properties:  pb.NewStruct(),
		Hash:        nil, //proto.String("01234567890"),
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
		Relation: proto.String("owner"),
		Object: &common.ObjectIdentifier{
			Id:   nil, //proto.String(""),
			Key:  proto.String("978-1260440218"),
			Type: proto.String("book"),
		},
		Hash: nil, //proto.String("01234567890"),
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
