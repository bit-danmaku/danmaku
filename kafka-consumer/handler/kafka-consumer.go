package handler

import (
	"context"
	"io"
	"time"

	log "github.com/asim/go-micro/v3/logger"

	pb "github.com/bit-danmaku/danmaku/proto/kafkaconsumer"
)

type KafkaConsumer struct{}

func (e *KafkaConsumer) Call(ctx context.Context, req *pb.CallRequest, rsp *pb.CallResponse) error {
	log.Infof("Received KafkaConsumer.Call request: %v", req)
	rsp.Msg = "Hello " + req.Name
	return nil
}

func (e *KafkaConsumer) ClientStream(ctx context.Context, stream pb.KafkaConsumer_ClientStreamStream) error {
	var count int64
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			log.Infof("Got %v pings total", count)
			return stream.SendMsg(&pb.ClientStreamResponse{Count: count})
		}
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		count++
	}
}

func (e *KafkaConsumer) ServerStream(ctx context.Context, req *pb.ServerStreamRequest, stream pb.KafkaConsumer_ServerStreamStream) error {
	log.Infof("Received KafkaConsumer.ServerStream request: %v", req)
	for i := 0; i < int(req.Count); i++ {
		log.Infof("Sending %d", i)
		if err := stream.Send(&pb.ServerStreamResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
		time.Sleep(time.Millisecond * 250)
	}
	return nil
}

func (e *KafkaConsumer) BidiStream(ctx context.Context, stream pb.KafkaConsumer_BidiStreamStream) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&pb.BidiStreamResponse{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
