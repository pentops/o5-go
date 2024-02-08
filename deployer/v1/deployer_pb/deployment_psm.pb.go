// Code generated by protoc-gen-go-psm. DO NOT EDIT.

package deployer_pb

import (
	context "context"
	fmt "fmt"
	psm "github.com/pentops/protostate/psm"
	proto "google.golang.org/protobuf/proto"
)

// StateObjectOptions: DeploymentPSM
type DeploymentPSMEventer = psm.Eventer[
	*DeploymentState,
	DeploymentStatus,
	*DeploymentEvent,
	DeploymentPSMEvent,
]

type DeploymentPSM = psm.StateMachine[
	*DeploymentState,
	DeploymentStatus,
	*DeploymentEvent,
	DeploymentPSMEvent,
]

func NewDeploymentPSM(db psm.Transactor, options ...psm.StateMachineOption[
	*DeploymentState,
	DeploymentStatus,
	*DeploymentEvent,
	DeploymentPSMEvent,
]) (*DeploymentPSM, error) {
	return psm.NewStateMachine[
		*DeploymentState,
		DeploymentStatus,
		*DeploymentEvent,
		DeploymentPSMEvent,
	](db, DeploymentPSMConverter{}, DefaultDeploymentPSMTableSpec, options...)
}

type DeploymentPSMTableSpec = psm.TableSpec[
	*DeploymentState,
	DeploymentStatus,
	*DeploymentEvent,
	DeploymentPSMEvent,
]

var DefaultDeploymentPSMTableSpec = DeploymentPSMTableSpec{
	StateTable: "deployment",
	EventTable: "deployment_event",
	PrimaryKey: func(event *DeploymentEvent) (map[string]interface{}, error) {
		return map[string]interface{}{
			"id": event.DeploymentId,
		}, nil
	},
	StateColumns: func(state *DeploymentState) (map[string]interface{}, error) {
		return map[string]interface{}{}, nil
	},
	EventColumns: func(event *DeploymentEvent) (map[string]interface{}, error) {
		metadata := event.Metadata
		return map[string]interface{}{
			"id":            metadata.EventId,
			"timestamp":     metadata.Timestamp,
			"actor":         metadata.Actor,
			"deployment_id": event.DeploymentId,
		}, nil
	},
	EventPrimaryKeyFieldPaths: []string{
		"metadata.event_id",
	},
	StatePrimaryKeyFieldPaths: []string{
		"deployment_id",
	},
}

type DeploymentPSMTransitionBaton = psm.TransitionBaton[*DeploymentEvent, DeploymentPSMEvent]

func DeploymentPSMFunc[SE DeploymentPSMEvent](cb func(context.Context, DeploymentPSMTransitionBaton, *DeploymentState, SE) error) psm.TransitionFunc[
	*DeploymentState,
	DeploymentStatus,
	*DeploymentEvent,
	DeploymentPSMEvent,
	SE,
] {
	return psm.TransitionFunc[
		*DeploymentState,
		DeploymentStatus,
		*DeploymentEvent,
		DeploymentPSMEvent,
		SE,
	](cb)
}

type DeploymentPSMEventKey string

const (
	DeploymentPSMEventNil             DeploymentPSMEventKey = "<nil>"
	DeploymentPSMEventCreated         DeploymentPSMEventKey = "created"
	DeploymentPSMEventTriggered       DeploymentPSMEventKey = "triggered"
	DeploymentPSMEventStackCreate     DeploymentPSMEventKey = "stack_create"
	DeploymentPSMEventStackWait       DeploymentPSMEventKey = "stack_wait"
	DeploymentPSMEventStackScale      DeploymentPSMEventKey = "stack_scale"
	DeploymentPSMEventStackTrigger    DeploymentPSMEventKey = "stack_trigger"
	DeploymentPSMEventStackUpsert     DeploymentPSMEventKey = "stack_upsert"
	DeploymentPSMEventStackStatus     DeploymentPSMEventKey = "stack_status"
	DeploymentPSMEventMigrateData     DeploymentPSMEventKey = "migrate_data"
	DeploymentPSMEventDbMigrateStatus DeploymentPSMEventKey = "db_migrate_status"
	DeploymentPSMEventDataMigrated    DeploymentPSMEventKey = "data_migrated"
	DeploymentPSMEventError           DeploymentPSMEventKey = "error"
	DeploymentPSMEventDone            DeploymentPSMEventKey = "done"
	DeploymentPSMEventTerminated      DeploymentPSMEventKey = "terminated"
)

type DeploymentPSMEvent interface {
	proto.Message
	PSMEventKey() DeploymentPSMEventKey
}
type DeploymentPSMConverter struct{}

func (c DeploymentPSMConverter) Unwrap(e *DeploymentEvent) DeploymentPSMEvent {
	return e.UnwrapPSMEvent()
}

func (c DeploymentPSMConverter) StateLabel(s *DeploymentState) string {
	return s.Status.String()
}

func (c DeploymentPSMConverter) EventLabel(e DeploymentPSMEvent) string {
	return string(e.PSMEventKey())
}

func (c DeploymentPSMConverter) EmptyState(e *DeploymentEvent) *DeploymentState {
	return &DeploymentState{
		DeploymentId: e.DeploymentId,
	}
}
func (c DeploymentPSMConverter) CheckStateKeys(s *DeploymentState, e *DeploymentEvent) error {
	if s.DeploymentId != e.DeploymentId {
		return fmt.Errorf("state field 'DeploymentId' %q does not match event field %q", s.DeploymentId, e.DeploymentId)
	}
	return nil
}

func (ee *DeploymentEventType) UnwrapPSMEvent() DeploymentPSMEvent {
	if ee == nil {
		return nil
	}
	switch v := ee.Type.(type) {
	case *DeploymentEventType_Created_:
		return v.Created
	case *DeploymentEventType_Triggered_:
		return v.Triggered
	case *DeploymentEventType_StackCreate_:
		return v.StackCreate
	case *DeploymentEventType_StackWait_:
		return v.StackWait
	case *DeploymentEventType_StackScale_:
		return v.StackScale
	case *DeploymentEventType_StackTrigger_:
		return v.StackTrigger
	case *DeploymentEventType_StackUpsert_:
		return v.StackUpsert
	case *DeploymentEventType_StackStatus_:
		return v.StackStatus
	case *DeploymentEventType_MigrateData_:
		return v.MigrateData
	case *DeploymentEventType_DbMigrateStatus:
		return v.DbMigrateStatus
	case *DeploymentEventType_DataMigrated_:
		return v.DataMigrated
	case *DeploymentEventType_Error_:
		return v.Error
	case *DeploymentEventType_Done_:
		return v.Done
	case *DeploymentEventType_Terminated_:
		return v.Terminated
	default:
		return nil
	}
}
func (ee *DeploymentEventType) PSMEventKey() DeploymentPSMEventKey {
	tt := ee.UnwrapPSMEvent()
	if tt == nil {
		return DeploymentPSMEventNil
	}
	return tt.PSMEventKey()
}
func (ee *DeploymentEvent) PSMEventKey() DeploymentPSMEventKey {
	return ee.Event.PSMEventKey()
}
func (ee *DeploymentEvent) UnwrapPSMEvent() DeploymentPSMEvent {
	return ee.Event.UnwrapPSMEvent()
}
func (ee *DeploymentEvent) SetPSMEvent(inner DeploymentPSMEvent) {
	if ee.Event == nil {
		ee.Event = &DeploymentEventType{}
	}
	switch v := inner.(type) {
	case *DeploymentEventType_Created:
		ee.Event.Type = &DeploymentEventType_Created_{Created: v}
	case *DeploymentEventType_Triggered:
		ee.Event.Type = &DeploymentEventType_Triggered_{Triggered: v}
	case *DeploymentEventType_StackCreate:
		ee.Event.Type = &DeploymentEventType_StackCreate_{StackCreate: v}
	case *DeploymentEventType_StackWait:
		ee.Event.Type = &DeploymentEventType_StackWait_{StackWait: v}
	case *DeploymentEventType_StackScale:
		ee.Event.Type = &DeploymentEventType_StackScale_{StackScale: v}
	case *DeploymentEventType_StackTrigger:
		ee.Event.Type = &DeploymentEventType_StackTrigger_{StackTrigger: v}
	case *DeploymentEventType_StackUpsert:
		ee.Event.Type = &DeploymentEventType_StackUpsert_{StackUpsert: v}
	case *DeploymentEventType_StackStatus:
		ee.Event.Type = &DeploymentEventType_StackStatus_{StackStatus: v}
	case *DeploymentEventType_MigrateData:
		ee.Event.Type = &DeploymentEventType_MigrateData_{MigrateData: v}
	case *DeploymentEventType_DBMigrateStatus:
		ee.Event.Type = &DeploymentEventType_DbMigrateStatus{DbMigrateStatus: v}
	case *DeploymentEventType_DataMigrated:
		ee.Event.Type = &DeploymentEventType_DataMigrated_{DataMigrated: v}
	case *DeploymentEventType_Error:
		ee.Event.Type = &DeploymentEventType_Error_{Error: v}
	case *DeploymentEventType_Done:
		ee.Event.Type = &DeploymentEventType_Done_{Done: v}
	case *DeploymentEventType_Terminated:
		ee.Event.Type = &DeploymentEventType_Terminated_{Terminated: v}
	default:
		panic("invalid type")
	}
}
func (*DeploymentEventType_Created) PSMEventKey() DeploymentPSMEventKey {
	return DeploymentPSMEventCreated
}
func (*DeploymentEventType_Triggered) PSMEventKey() DeploymentPSMEventKey {
	return DeploymentPSMEventTriggered
}
func (*DeploymentEventType_StackCreate) PSMEventKey() DeploymentPSMEventKey {
	return DeploymentPSMEventStackCreate
}
func (*DeploymentEventType_StackWait) PSMEventKey() DeploymentPSMEventKey {
	return DeploymentPSMEventStackWait
}
func (*DeploymentEventType_StackScale) PSMEventKey() DeploymentPSMEventKey {
	return DeploymentPSMEventStackScale
}
func (*DeploymentEventType_StackTrigger) PSMEventKey() DeploymentPSMEventKey {
	return DeploymentPSMEventStackTrigger
}
func (*DeploymentEventType_StackUpsert) PSMEventKey() DeploymentPSMEventKey {
	return DeploymentPSMEventStackUpsert
}
func (*DeploymentEventType_StackStatus) PSMEventKey() DeploymentPSMEventKey {
	return DeploymentPSMEventStackStatus
}
func (*DeploymentEventType_MigrateData) PSMEventKey() DeploymentPSMEventKey {
	return DeploymentPSMEventMigrateData
}
func (*DeploymentEventType_DBMigrateStatus) PSMEventKey() DeploymentPSMEventKey {
	return DeploymentPSMEventDbMigrateStatus
}
func (*DeploymentEventType_DataMigrated) PSMEventKey() DeploymentPSMEventKey {
	return DeploymentPSMEventDataMigrated
}
func (*DeploymentEventType_Error) PSMEventKey() DeploymentPSMEventKey {
	return DeploymentPSMEventError
}
func (*DeploymentEventType_Done) PSMEventKey() DeploymentPSMEventKey {
	return DeploymentPSMEventDone
}
func (*DeploymentEventType_Terminated) PSMEventKey() DeploymentPSMEventKey {
	return DeploymentPSMEventTerminated
}
