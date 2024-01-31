// Code generated by protoc-gen-go-psm. DO NOT EDIT.

package deployer_pb

import (
	context "context"
	fmt "fmt"
	psm "github.com/pentops/protostate/psm"
	proto "google.golang.org/protobuf/proto"
)

// StateObjectOptions: StackPSM
type StackPSMEventer = psm.Eventer[
	*StackState,
	StackStatus,
	*StackEvent,
	StackPSMEvent,
]

type StackPSM = psm.StateMachine[
	*StackState,
	StackStatus,
	*StackEvent,
	StackPSMEvent,
]

func NewStackPSM(db psm.Transactor, options ...psm.StateMachineOption[
	*StackState,
	StackStatus,
	*StackEvent,
	StackPSMEvent,
]) (*StackPSM, error) {
	return psm.NewStateMachine[
		*StackState,
		StackStatus,
		*StackEvent,
		StackPSMEvent,
	](db, StackPSMConverter{}, DefaultStackPSMTableSpec, options...)
}

type StackPSMTableSpec = psm.TableSpec[
	*StackState,
	StackStatus,
	*StackEvent,
	StackPSMEvent,
]

var DefaultStackPSMTableSpec = StackPSMTableSpec{
	StateTable: "stack",
	EventTable: "stack_event",
	PrimaryKey: func(event *StackEvent) (map[string]interface{}, error) {
		return map[string]interface{}{
			"id": event.StackId,
		}, nil
	},
	StateColumns: func(state *StackState) (map[string]interface{}, error) {
		return map[string]interface{}{}, nil
	},
	EventColumns: func(event *StackEvent) (map[string]interface{}, error) {
		metadata := event.Metadata
		return map[string]interface{}{
			"id":        metadata.EventId,
			"timestamp": metadata.Timestamp,
			"actor":     metadata.Actor,
			"stack_id":  event.StackId,
		}, nil
	},
	EventPrimaryKeyFieldPaths: []string{
		"metadata.event_id",
	},
	StatePrimaryKeyFieldPaths: []string{
		"stack_id",
	},
}

type StackPSMTransitionBaton = psm.TransitionBaton[*StackEvent, StackPSMEvent]

func StackPSMFunc[SE StackPSMEvent](cb func(context.Context, StackPSMTransitionBaton, *StackState, SE) error) psm.TransitionFunc[
	*StackState,
	StackStatus,
	*StackEvent,
	StackPSMEvent,
	SE,
] {
	return psm.TransitionFunc[
		*StackState,
		StackStatus,
		*StackEvent,
		StackPSMEvent,
		SE,
	](cb)
}

type StackPSMEventKey string

const (
	StackPSMEventNil                 StackPSMEventKey = "<nil>"
	StackPSMEventTriggered           StackPSMEventKey = "triggered"
	StackPSMEventDeploymentCompleted StackPSMEventKey = "deployment_completed"
	StackPSMEventDeploymentFailed    StackPSMEventKey = "deployment_failed"
	StackPSMEventAvailable           StackPSMEventKey = "available"
)

type StackPSMEvent interface {
	proto.Message
	PSMEventKey() StackPSMEventKey
}
type StackPSMConverter struct{}

func (c StackPSMConverter) Unwrap(e *StackEvent) StackPSMEvent {
	return e.UnwrapPSMEvent()
}

func (c StackPSMConverter) StateLabel(s *StackState) string {
	return s.Status.String()
}

func (c StackPSMConverter) EventLabel(e StackPSMEvent) string {
	return string(e.PSMEventKey())
}

func (c StackPSMConverter) EmptyState(e *StackEvent) *StackState {
	return &StackState{
		StackId: e.StackId,
	}
}
func (c StackPSMConverter) CheckStateKeys(s *StackState, e *StackEvent) error {
	if s.StackId != e.StackId {
		return fmt.Errorf("state field 'StackId' %q does not match event field %q", s.StackId, e.StackId)
	}
	return nil
}

func (ee *StackEventType) UnwrapPSMEvent() StackPSMEvent {
	if ee == nil {
		return nil
	}
	switch v := ee.Type.(type) {
	case *StackEventType_Triggered_:
		return v.Triggered
	case *StackEventType_DeploymentCompleted_:
		return v.DeploymentCompleted
	case *StackEventType_DeploymentFailed_:
		return v.DeploymentFailed
	case *StackEventType_Available_:
		return v.Available
	default:
		return nil
	}
}
func (ee *StackEventType) PSMEventKey() StackPSMEventKey {
	tt := ee.UnwrapPSMEvent()
	if tt == nil {
		return StackPSMEventNil
	}
	return tt.PSMEventKey()
}
func (ee *StackEvent) PSMEventKey() StackPSMEventKey {
	return ee.Event.PSMEventKey()
}
func (ee *StackEvent) UnwrapPSMEvent() StackPSMEvent {
	return ee.Event.UnwrapPSMEvent()
}
func (ee *StackEvent) SetPSMEvent(inner StackPSMEvent) {
	if ee.Event == nil {
		ee.Event = &StackEventType{}
	}
	switch v := inner.(type) {
	case *StackEventType_Triggered:
		ee.Event.Type = &StackEventType_Triggered_{Triggered: v}
	case *StackEventType_DeploymentCompleted:
		ee.Event.Type = &StackEventType_DeploymentCompleted_{DeploymentCompleted: v}
	case *StackEventType_DeploymentFailed:
		ee.Event.Type = &StackEventType_DeploymentFailed_{DeploymentFailed: v}
	case *StackEventType_Available:
		ee.Event.Type = &StackEventType_Available_{Available: v}
	default:
		panic("invalid type")
	}
}
func (*StackEventType_Triggered) PSMEventKey() StackPSMEventKey {
	return StackPSMEventTriggered
}
func (*StackEventType_DeploymentCompleted) PSMEventKey() StackPSMEventKey {
	return StackPSMEventDeploymentCompleted
}
func (*StackEventType_DeploymentFailed) PSMEventKey() StackPSMEventKey {
	return StackPSMEventDeploymentFailed
}
func (*StackEventType_Available) PSMEventKey() StackPSMEventKey {
	return StackPSMEventAvailable
}
