// Code generated by protoc-gen-go-psm. DO NOT EDIT.

package deployer_pb

import (
	context "context"
	fmt "fmt"
	psm "github.com/pentops/protostate/psm"
	proto "google.golang.org/protobuf/proto"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

// StateObjectOptions: EnvironmentPSM
type EnvironmentPSMEventer = psm.Eventer[
	*EnvironmentState,
	EnvironmentStatus,
	*EnvironmentEvent,
	EnvironmentPSMEvent,
]

type EnvironmentPSM = psm.StateMachine[
	*EnvironmentState,
	EnvironmentStatus,
	*EnvironmentEvent,
	EnvironmentPSMEvent,
]

type EnvironmentPSMDB = psm.DBStateMachine[
	*EnvironmentState,
	EnvironmentStatus,
	*EnvironmentEvent,
	EnvironmentPSMEvent,
]

func DefaultEnvironmentPSMConfig() *psm.StateMachineConfig[
	*EnvironmentState,
	EnvironmentStatus,
	*EnvironmentEvent,
	EnvironmentPSMEvent,
] {
	return psm.NewStateMachineConfig[
		*EnvironmentState,
		EnvironmentStatus,
		*EnvironmentEvent,
		EnvironmentPSMEvent,
	](EnvironmentPSMConverter{}, DefaultEnvironmentPSMTableSpec)
}

func NewEnvironmentPSM(config *psm.StateMachineConfig[
	*EnvironmentState,
	EnvironmentStatus,
	*EnvironmentEvent,
	EnvironmentPSMEvent,
]) (*EnvironmentPSM, error) {
	return psm.NewStateMachine[
		*EnvironmentState,
		EnvironmentStatus,
		*EnvironmentEvent,
		EnvironmentPSMEvent,
	](config)
}

type EnvironmentPSMTableSpec = psm.TableSpec[
	*EnvironmentState,
	EnvironmentStatus,
	*EnvironmentEvent,
	EnvironmentPSMEvent,
]

var DefaultEnvironmentPSMTableSpec = EnvironmentPSMTableSpec{
	StateTable: "environment",
	EventTable: "environment_event",
	PrimaryKey: func(event *EnvironmentEvent) (map[string]interface{}, error) {
		return map[string]interface{}{
			"id": event.EnvironmentId,
		}, nil
	},
	StateColumns: func(state *EnvironmentState) (map[string]interface{}, error) {
		return map[string]interface{}{}, nil
	},
	EventColumns: func(event *EnvironmentEvent) (map[string]interface{}, error) {
		metadata := event.Metadata
		return map[string]interface{}{
			"id":             metadata.EventId,
			"timestamp":      metadata.Timestamp,
			"actor":          metadata.Actor,
			"environment_id": event.EnvironmentId,
		}, nil
	},
	EventPrimaryKeyFieldPaths: []string{
		"metadata.event_id",
	},
	StatePrimaryKeyFieldPaths: []string{
		"environment_id",
	},
}

type EnvironmentPSMTransitionBaton = psm.TransitionBaton[*EnvironmentEvent, EnvironmentPSMEvent]

func EnvironmentPSMFunc[SE EnvironmentPSMEvent](cb func(context.Context, EnvironmentPSMTransitionBaton, *EnvironmentState, SE) error) psm.TransitionFunc[
	*EnvironmentState,
	EnvironmentStatus,
	*EnvironmentEvent,
	EnvironmentPSMEvent,
	SE,
] {
	return psm.TransitionFunc[
		*EnvironmentState,
		EnvironmentStatus,
		*EnvironmentEvent,
		EnvironmentPSMEvent,
		SE,
	](cb)
}

type EnvironmentPSMEventKey string

const (
	EnvironmentPSMEventNil        EnvironmentPSMEventKey = "<nil>"
	EnvironmentPSMEventConfigured EnvironmentPSMEventKey = "configured"
)

type EnvironmentPSMEvent interface {
	proto.Message
	PSMEventKey() EnvironmentPSMEventKey
}
type EnvironmentPSMConverter struct{}

func (c EnvironmentPSMConverter) Unwrap(e *EnvironmentEvent) EnvironmentPSMEvent {
	return e.UnwrapPSMEvent()
}

func (c EnvironmentPSMConverter) StateLabel(s *EnvironmentState) string {
	return s.Status.String()
}

func (c EnvironmentPSMConverter) EventLabel(e EnvironmentPSMEvent) string {
	return string(e.PSMEventKey())
}

func (c EnvironmentPSMConverter) EmptyState(e *EnvironmentEvent) *EnvironmentState {
	return &EnvironmentState{
		EnvironmentId: e.EnvironmentId,
	}
}

func (c EnvironmentPSMConverter) DeriveChainEvent(e *EnvironmentEvent, systemActor psm.SystemActor, eventKey string) *EnvironmentEvent {
	metadata := &EventMetadata{
		EventId:   systemActor.NewEventID(e.Metadata.EventId, eventKey),
		Timestamp: timestamppb.Now(),
	}
	actorProto := systemActor.ActorProto()
	refl := metadata.ProtoReflect()
	refl.Set(refl.Descriptor().Fields().ByName("actor"), actorProto)
	return &EnvironmentEvent{
		Metadata:      metadata,
		EnvironmentId: e.EnvironmentId,
	}
}

func (c EnvironmentPSMConverter) CheckStateKeys(s *EnvironmentState, e *EnvironmentEvent) error {
	if s.EnvironmentId != e.EnvironmentId {
		return fmt.Errorf("state field 'EnvironmentId' %q does not match event field %q", s.EnvironmentId, e.EnvironmentId)
	}
	return nil
}

func (ee *EnvironmentEventType) UnwrapPSMEvent() EnvironmentPSMEvent {
	if ee == nil {
		return nil
	}
	switch v := ee.Type.(type) {
	case *EnvironmentEventType_Configured_:
		return v.Configured
	default:
		return nil
	}
}
func (ee *EnvironmentEventType) PSMEventKey() EnvironmentPSMEventKey {
	tt := ee.UnwrapPSMEvent()
	if tt == nil {
		return EnvironmentPSMEventNil
	}
	return tt.PSMEventKey()
}
func (ee *EnvironmentEvent) PSMEventKey() EnvironmentPSMEventKey {
	return ee.Event.PSMEventKey()
}
func (ee *EnvironmentEvent) UnwrapPSMEvent() EnvironmentPSMEvent {
	return ee.Event.UnwrapPSMEvent()
}
func (ee *EnvironmentEvent) SetPSMEvent(inner EnvironmentPSMEvent) {
	if ee.Event == nil {
		ee.Event = &EnvironmentEventType{}
	}
	switch v := inner.(type) {
	case *EnvironmentEventType_Configured:
		ee.Event.Type = &EnvironmentEventType_Configured_{Configured: v}
	default:
		panic("invalid type")
	}
}
func (*EnvironmentEventType_Configured) PSMEventKey() EnvironmentPSMEventKey {
	return EnvironmentPSMEventConfigured
}
