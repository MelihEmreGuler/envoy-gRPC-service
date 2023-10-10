package main

import (
	"bufio"
	"flag"
	"fmt"
	pb "github.com/MelihEmreGuler/envoy-gRPC-service/instancepb"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
	"io/ioutil"
	"log"
	"net"
	"os"
)

var (
	port = flag.Int("port", 50051, "The server port")
	// Create a channel to signal when instances are ready
	//instancesReadyCh = make(chan struct{})
)

type instanceServer struct {
	pb.UnimplementedInstanceServer
	instances []*pb.Instance
}

func readJSONFile() (error, []byte) {
	// Open the JSON file
	filePath := "server/instance.json"
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Error opening JSON file %v", err)
		return err, nil
	}
	defer file.Close() // Ensure the file is closed when done.

	// Create a buffered reader to read the JSON data
	reader := bufio.NewReader(file)

	// Read the JSON data from the file.
	jsonData, err := ioutil.ReadAll(reader)
	if err != nil {
		log.Fatalf("error reading JSON: %v", err)
		return err, nil
	}

	// Return the JSON data
	return nil, jsonData
}

func unmarshalJSON(data []byte) (error, []*pb.Instance) {

	// Create a response object
	response := &pb.GetInstancesByRegionResponse{}

	//Unmarshal the JSON data into the interface.
	err := protojson.Unmarshal(data, response)
	if err != nil {
		return err, nil
	}

	//return the instances
	return nil, response.Instances
}

func (s *instanceServer) loadInstances() {
	// Read the JSON data from the file.
	err, jsonByte := readJSONFile()
	if err != nil {
		log.Fatalf("error reading JSON: %v", err)
	}

	// Unmarshal the JSON data into the proto message.
	err, message := unmarshalJSON(jsonByte)
	if err != nil {
		log.Fatalf("error unmarshalling JSON: %v", err)
	}

	s.instances = message
}

func newServer() *instanceServer {
	s := &instanceServer{instances: make([]*pb.Instance, 0)}
	s.loadInstances()
	fmt.Println("instances loaded", s.instances)
	return s
}

// GetInstancesByRegion returns the instances in the given region.
func (s *instanceServer) GetInstancesByRegion(req *pb.GetInstancesByRegionRequest, stream pb.Instance_GetInstancesByRegionServer) error {
	log.Printf("Received GetInstancesByRegion request from client: \n %+v\n", req)

	// Simulate some work to retrieve instances (replace this with your actual logic)
	/*go func() {
		time.Sleep(1 * time.Second) // Simulate work
		// When instances are ready, signal it
		close(instancesReadyCh)
	}()*/

	// Send the instances to the client with a stream
	err := stream.Send(&pb.GetInstancesByRegionResponse{Instances: s.instances})
	if err != nil {
		return err
	}
	// Wait for the instances to send to the client
	//<-instancesReadyCh
	return nil
}

// SendStatusUpdates sends status updates to the client.
func (s *instanceServer) SendStatusUpdates(req *pb.GetInstancesByRegionRequest, stream pb.Instance_SendStatusUpdatesServer) error {
	log.Printf("Received SendStatusUpdates request from client: \n %+v\n", req)

	// Notify the client that scanning is starting
	statusUpdate := &pb.StatusUpdate{Message: "Scanning started..."}
	if err := stream.Send(statusUpdate); err != nil {
		return err
	}

	//communicate with the GetInstancesByRegion function and wait here until sending the instances to client
	// Wait for instances to be ready
	//<-instancesReadyCh

	// Notify the client that scanning is complete
	statusUpdate = &pb.StatusUpdate{Message: "Scanning completed."}
	if err := stream.Send(statusUpdate); err != nil {
		return err
	}

	return nil
}

func main() {
	// parse the port flag
	flag.Parse()

	// Create a listener on port 50051
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption

	// Register your gRPC service implementation
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterInstanceServer(grpcServer, newServer())

	// Serve gRPC listener
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
