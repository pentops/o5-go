// Code generated by protoc-gen-go-psm. DO NOT EDIT.

package dante_pb

import (
	context "context"
	fmt "fmt"
	psm "github.com/pentops/protostate/psm"
	proto "google.golang.org/protobuf/proto"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

// StateObjectOptions: DeadmessagePSM
type DeadmessagePSM = psm.StateMachine[
	*DeadMessageState,
	MessageStatus,
	*DeadMessageEvent,
	DeadmessagePSMEvent,
]

type DeadmessagePSMDB = psm.DBStateMachine[
	*DeadMessageState,
	MessageStatus,
	*DeadMessageEvent,
	DeadmessagePSMEvent,
]

type DeadmessagePSMEventer = psm.Eventer[
	*DeadMessageState,
	MessageStatus,
	*DeadMessageEvent,
	DeadmessagePSMEvent,
]

func DefaultDeadmessagePSMConfig() *psm.StateMachineConfig[
	*DeadMessageState,
	MessageStatus,
	*DeadMessageEvent,
	DeadmessagePSMEvent,
] {
	return psm.NewStateMachineConfig[
		*DeadMessageState,
		MessageStatus,
		*DeadMessageEvent,
		DeadmessagePSMEvent,
	](DeadmessagePSMConverter{}, DefaultDeadmessagePSMTableSpec)
}

func NewDeadmessagePSM(config *psm.StateMachineConfig[
	*DeadMessageState,
	MessageStatus,
	*DeadMessageEvent,
	DeadmessagePSMEvent,
]) (*DeadmessagePSM, error) {
	return psm.NewStateMachine[
		*DeadMessageState,
		MessageStatus,
		*DeadMessageEvent,
		DeadmessagePSMEvent,
	](config)
}

type DeadmessagePSMTableSpec = psm.PSMTableSpec[
	*DeadMessageState,
	MessageStatus,
	*DeadMessageEvent,
	DeadmessagePSMEvent,
]

var DefaultDeadmessagePSMTableSpec = DeadmessagePSMTableSpec{
	State: psm.TableSpec[*DeadMessageState]{
		TableName:  "deadmessage",
		DataColumn: "state",
		StoreExtraColumns: func(state *DeadMessageState) (map[string]interface{}, error) {
			return map[string]interface{}{}, nil
		},
		PKFieldPaths: []string{
			"message_id",
		},
	},
	Event: psm.TableSpec[*DeadMessageEvent]{
		TableName:  "deadmessage_event",
		DataColumn: "data",
		StoreExtraColumns: func(event *DeadMessageEvent) (map[string]interface{}, error) {
			metadata := event.Metadata
			return map[string]interface{}{
				"id":         metadata.EventId,
				"timestamp":  metadata.Timestamp,
				"actor":      metadata.Actor,
				"message_id": event.MessageId,
			}, nil
		},
		PKFieldPaths: []string{
			"metadata.event_id",
		},
	},
	PrimaryKey: func(event *DeadMessageEvent) (map[string]interface{}, error) {
		return map[string]interface{}{
			"message_id": event.MessageId,
		}, nil
	},
}

type DeadmessagePSMTransitionBaton = psm.TransitionBaton[*DeadMessageEvent, DeadmessagePSMEvent]

func DeadmessagePSMFunc[SE DeadmessagePSMEvent](cb func(context.Context, DeadmessagePSMTransitionBaton, *DeadMessageState, SE) error) psm.TransitionFunc[
	*DeadMessageState,
	MessageStatus,
	*DeadMessageEvent,
	DeadmessagePSMEvent,
	SE,
] {
	return psm.TransitionFunc[
		*DeadMessageState,
		MessageStatus,
		*DeadMessageEvent,
		DeadmessagePSMEvent,
		SE,
	](cb)
}

type DeadmessagePSMEventKey = string

const (
	DeadmessagePSMEventNil      DeadmessagePSMEventKey = "<nil>"
	DeadmessagePSMEventCreated  DeadmessagePSMEventKey = "created"
	DeadmessagePSMEventUpdated  DeadmessagePSMEventKey = "updated"
	DeadmessagePSMEventReplayed DeadmessagePSMEventKey = "replayed"
	DeadmessagePSMEventRejected DeadmessagePSMEventKey = "rejected"
)

type DeadmessagePSMEvent interface {
	proto.Message
	PSMEventKey() DeadmessagePSMEventKey
}

type DeadmessagePSMConverter struct{}

func (c DeadmessagePSMConverter) EmptyState(e *DeadMessageEvent) *DeadMessageState {
	return &DeadMessageState{
		MessageId: e.MessageId,
	}
}

func (c DeadmessagePSMConverter) DeriveChainEvent(e *DeadMessageEvent, systemActor psm.SystemActor, eventKey string) *DeadMessageEvent {
	metadata := &Metadata{
		EventId:   systemActor.NewEventID(e.Metadata.EventId, eventKey),
		Timestamp: timestamppb.Now(),
	}
	actorProto := systemActor.ActorProto()
	refl := metadata.ProtoReflect()
	refl.Set(refl.Descriptor().Fields().ByName("actor"), actorProto)
	return &DeadMessageEvent{
		Metadata:  metadata,
		MessageId: e.MessageId,
	}
}

func (c DeadmessagePSMConverter) CheckStateKeys(s *DeadMessageState, e *DeadMessageEvent) error {
	if s.MessageId != e.MessageId {
		return fmt.Errorf("state field 'MessageId' %q does not match event field %q", s.MessageId, e.MessageId)
	}
	return nil
}

func (etw *DeadMessageEventType) UnwrapPSMEvent() DeadmessagePSMEvent {
	if etw == nil {
		return nil
	}
	switch v := etw.Type.(type) {
	case *DeadMessageEventType_Created_:
		return v.Created
	case *DeadMessageEventType_Updated_:
		return v.Updated
	case *DeadMessageEventType_Replayed_:
		return v.Replayed
	case *DeadMessageEventType_Rejected_:
		return v.Rejected
	default:
		return nil
	}
}
func (etw *DeadMessageEventType) PSMEventKey() DeadmessagePSMEventKey {
	tt := etw.UnwrapPSMEvent()
	if tt == nil {
		return DeadmessagePSMEventNil
	}
	return tt.PSMEventKey()
}
func (etw *DeadMessageEventType) SetPSMEvent(inner DeadmessagePSMEvent) {
	switch v := inner.(type) {
	case *DeadMessageEventType_Created:
		etw.Type = &DeadMessageEventType_Created_{Created: v}
	case *DeadMessageEventType_Updated:
		etw.Type = &DeadMessageEventType_Updated_{Updated: v}
	case *DeadMessageEventType_Replayed:
		etw.Type = &DeadMessageEventType_Replayed_{Replayed: v}
	case *DeadMessageEventType_Rejected:
		etw.Type = &DeadMessageEventType_Rejected_{Rejected: v}
	default:
		panic("invalid type")
	}
}

func (ee *DeadMessageEvent) PSMEventKey() DeadmessagePSMEventKey {
	return ee.Event.PSMEventKey()
}

func (ee *DeadMessageEvent) UnwrapPSMEvent() DeadmessagePSMEvent {
	return ee.Event.UnwrapPSMEvent()
}

func (ee *DeadMessageEvent) SetPSMEvent(inner DeadmessagePSMEvent) {
	if ee.Event == nil {
		ee.Event = &DeadMessageEventType{}
	}
	ee.Event.SetPSMEvent(inner)
}

func (*DeadMessageEventType_Created) PSMEventKey() DeadmessagePSMEventKey {
	return DeadmessagePSMEventCreated
}
func (*DeadMessageEventType_Updated) PSMEventKey() DeadmessagePSMEventKey {
	return DeadmessagePSMEventUpdated
}
func (*DeadMessageEventType_Replayed) PSMEventKey() DeadmessagePSMEventKey {
	return DeadmessagePSMEventReplayed
}
func (*DeadMessageEventType_Rejected) PSMEventKey() DeadmessagePSMEventKey {
	return DeadmessagePSMEventRejected
}
