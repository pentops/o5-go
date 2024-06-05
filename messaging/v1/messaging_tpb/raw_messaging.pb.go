// Code generated by protoc-gen-go-messaging. DO NOT EDIT.

package messaging_tpb

import (
	messaging_pb "github.com/pentops/o5-go/messaging/v1/messaging_pb"
	protojson "google.golang.org/protobuf/encoding/protojson"
)

// Service: RawMessageTopic
// Method: Raw

func (msg *RawMessage) O5Message(id string) (*messaging_pb.Message, error) {
	body, err := protojson.Marshal(msg)
	if err != nil {
		return nil, err
	}
	return &messaging_pb.Message{
		DestinationTopic: "raw",
		MessageId:        id,
		GrpcService:      "o5.messaging.v1.topic.RawMessageTopic",
		GrpcMethod:       "Raw",
		Body: &messaging_pb.Any{
			TypeUrl:  "type.googleapis.com/o5.messaging.v1.topic.RawMessage",
			Value:    body,
			Encoding: messaging_pb.WireEncoding_WIRE_ENCODING_PROTOJSON,
		},
		Headers: map[string]string{},
	}, nil
}

func (msg *RawMessage) MessagingTopic() string {
	return "raw"
}
func (msg *RawMessage) MessagingHeaders() map[string]string {
	headers := map[string]string{
		"grpc-service": "/o5.messaging.v1.topic.RawMessageTopic/Raw",
		"grpc-message": "o5.messaging.v1.topic.RawMessage",
	}
	return headers
}
