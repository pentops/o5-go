// Code generated by Bprotoc-gen-go-o5-messaging . DO NOT EDIT.
// versions:
// - protoc-gen-go-o5-messaging 0.0.0
// source: o5/messaging/v1/topic/dead.proto

package messaging_tpb

import (
	context "context"
	o5msg "github.com/pentops/o5-messaging.go/o5msg"
)

// Service: DeadMessageTopic
type DeadMessageTopicSender[C any] struct {
	Sender o5msg.Sender[C]
}

func NewDeadMessageTopicSender[C any](sender o5msg.Sender[C]) *DeadMessageTopicSender[C] {
	sender.Register(o5msg.TopicDescriptor{
		Service: "o5.messaging.v1.topic.DeadMessageTopic",
		Methods: []o5msg.MethodDescriptor{
			{
				Name:    "Dead",
				Message: (*DeadMessage).ProtoReflect(nil).Descriptor(),
			},
		},
	})
	return &DeadMessageTopicSender[C]{Sender: sender}
}

type DeadMessageTopicCollector[C any] struct {
	Collector o5msg.Collector[C]
}

func NewDeadMessageTopicCollector[C any](collector o5msg.Collector[C]) *DeadMessageTopicCollector[C] {
	collector.Register(o5msg.TopicDescriptor{
		Service: "o5.messaging.v1.topic.DeadMessageTopic",
		Methods: []o5msg.MethodDescriptor{
			{
				Name:    "Dead",
				Message: (*DeadMessage).ProtoReflect(nil).Descriptor(),
			},
		},
	})
	return &DeadMessageTopicCollector[C]{Collector: collector}
}

// Method: Dead

func (msg *DeadMessage) O5MessageHeader() o5msg.Header {
	header := o5msg.Header{
		GrpcService:      "o5.messaging.v1.topic.DeadMessageTopic",
		GrpcMethod:       "Dead",
		Headers:          map[string]string{},
		DestinationTopic: "dead-letter",
	}
	return header
}

func (send DeadMessageTopicSender[C]) Dead(ctx context.Context, sendContext C, msg *DeadMessage) error {
	return send.Sender.Send(ctx, sendContext, msg)
}

func (collect DeadMessageTopicCollector[C]) Dead(sendContext C, msg *DeadMessage) {
	collect.Collector.Collect(sendContext, msg)
}