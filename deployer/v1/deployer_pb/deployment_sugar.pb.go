// Code generated by protoc-gen-go-sugar. DO NOT EDIT.

package deployer_pb

import (
	driver "database/sql/driver"
	fmt "fmt"
)

// DeploymentEventType is a oneof wrapper
type DeploymentEventTypeKey string

const (
	DeploymentEvent_Created         DeploymentEventTypeKey = "created"
	DeploymentEvent_Triggered       DeploymentEventTypeKey = "triggered"
	DeploymentEvent_StackCreate     DeploymentEventTypeKey = "stackCreate"
	DeploymentEvent_StackWait       DeploymentEventTypeKey = "stackWait"
	DeploymentEvent_StackScale      DeploymentEventTypeKey = "stackScale"
	DeploymentEvent_StackTrigger    DeploymentEventTypeKey = "stackTrigger"
	DeploymentEvent_StackUpsert     DeploymentEventTypeKey = "stackUpsert"
	DeploymentEvent_StackStatus     DeploymentEventTypeKey = "stackStatus"
	DeploymentEvent_MigrateData     DeploymentEventTypeKey = "migrateData"
	DeploymentEvent_DbMigrateStatus DeploymentEventTypeKey = "dbMigrateStatus"
	DeploymentEvent_DataMigrated    DeploymentEventTypeKey = "dataMigrated"
	DeploymentEvent_Error           DeploymentEventTypeKey = "error"
	DeploymentEvent_Done            DeploymentEventTypeKey = "done"
)

func (x *DeploymentEventType) TypeKey() (DeploymentEventTypeKey, bool) {
	switch x.Type.(type) {
	case *DeploymentEventType_Created_:
		return DeploymentEvent_Created, true
	case *DeploymentEventType_Triggered_:
		return DeploymentEvent_Triggered, true
	case *DeploymentEventType_StackCreate_:
		return DeploymentEvent_StackCreate, true
	case *DeploymentEventType_StackWait_:
		return DeploymentEvent_StackWait, true
	case *DeploymentEventType_StackScale_:
		return DeploymentEvent_StackScale, true
	case *DeploymentEventType_StackTrigger_:
		return DeploymentEvent_StackTrigger, true
	case *DeploymentEventType_StackUpsert_:
		return DeploymentEvent_StackUpsert, true
	case *DeploymentEventType_StackStatus_:
		return DeploymentEvent_StackStatus, true
	case *DeploymentEventType_MigrateData_:
		return DeploymentEvent_MigrateData, true
	case *DeploymentEventType_DbMigrateStatus:
		return DeploymentEvent_DbMigrateStatus, true
	case *DeploymentEventType_DataMigrated_:
		return DeploymentEvent_DataMigrated, true
	case *DeploymentEventType_Error_:
		return DeploymentEvent_Error, true
	case *DeploymentEventType_Done_:
		return DeploymentEvent_Done, true
	default:
		return "", false
	}
}

type IsDeploymentEventTypeWrappedType interface {
	TypeKey() DeploymentEventTypeKey
}

func (x *DeploymentEventType) Set(val IsDeploymentEventTypeWrappedType) {
	switch v := val.(type) {
	case *DeploymentEventType_Created:
		x.Type = &DeploymentEventType_Created_{Created: v}
	case *DeploymentEventType_Triggered:
		x.Type = &DeploymentEventType_Triggered_{Triggered: v}
	case *DeploymentEventType_StackCreate:
		x.Type = &DeploymentEventType_StackCreate_{StackCreate: v}
	case *DeploymentEventType_StackWait:
		x.Type = &DeploymentEventType_StackWait_{StackWait: v}
	case *DeploymentEventType_StackScale:
		x.Type = &DeploymentEventType_StackScale_{StackScale: v}
	case *DeploymentEventType_StackTrigger:
		x.Type = &DeploymentEventType_StackTrigger_{StackTrigger: v}
	case *DeploymentEventType_StackUpsert:
		x.Type = &DeploymentEventType_StackUpsert_{StackUpsert: v}
	case *DeploymentEventType_StackStatus:
		x.Type = &DeploymentEventType_StackStatus_{StackStatus: v}
	case *DeploymentEventType_MigrateData:
		x.Type = &DeploymentEventType_MigrateData_{MigrateData: v}
	case *DeploymentEventType_DBMigrateStatus:
		x.Type = &DeploymentEventType_DbMigrateStatus{DbMigrateStatus: v}
	case *DeploymentEventType_DataMigrated:
		x.Type = &DeploymentEventType_DataMigrated_{DataMigrated: v}
	case *DeploymentEventType_Error:
		x.Type = &DeploymentEventType_Error_{Error: v}
	case *DeploymentEventType_Done:
		x.Type = &DeploymentEventType_Done_{Done: v}
	}
}
func (x *DeploymentEventType) Get() IsDeploymentEventTypeWrappedType {
	switch v := x.Type.(type) {
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
	default:
		return nil
	}
}
func (x *DeploymentEventType_Created) TypeKey() DeploymentEventTypeKey {
	return DeploymentEvent_Created
}
func (x *DeploymentEventType_Triggered) TypeKey() DeploymentEventTypeKey {
	return DeploymentEvent_Triggered
}
func (x *DeploymentEventType_StackCreate) TypeKey() DeploymentEventTypeKey {
	return DeploymentEvent_StackCreate
}
func (x *DeploymentEventType_StackWait) TypeKey() DeploymentEventTypeKey {
	return DeploymentEvent_StackWait
}
func (x *DeploymentEventType_StackScale) TypeKey() DeploymentEventTypeKey {
	return DeploymentEvent_StackScale
}
func (x *DeploymentEventType_StackTrigger) TypeKey() DeploymentEventTypeKey {
	return DeploymentEvent_StackTrigger
}
func (x *DeploymentEventType_StackUpsert) TypeKey() DeploymentEventTypeKey {
	return DeploymentEvent_StackUpsert
}
func (x *DeploymentEventType_StackStatus) TypeKey() DeploymentEventTypeKey {
	return DeploymentEvent_StackStatus
}
func (x *DeploymentEventType_MigrateData) TypeKey() DeploymentEventTypeKey {
	return DeploymentEvent_MigrateData
}
func (x *DeploymentEventType_DBMigrateStatus) TypeKey() DeploymentEventTypeKey {
	return DeploymentEvent_DbMigrateStatus
}
func (x *DeploymentEventType_DataMigrated) TypeKey() DeploymentEventTypeKey {
	return DeploymentEvent_DataMigrated
}
func (x *DeploymentEventType_Error) TypeKey() DeploymentEventTypeKey {
	return DeploymentEvent_Error
}
func (x *DeploymentEventType_Done) TypeKey() DeploymentEventTypeKey {
	return DeploymentEvent_Done
}

type IsDeploymentEventType_Type = isDeploymentEventType_Type

// DeploymentStatus
const (
	DeploymentStatus_UNSPECIFIED    DeploymentStatus = 0
	DeploymentStatus_QUEUED         DeploymentStatus = 1
	DeploymentStatus_TRIGGERED      DeploymentStatus = 2
	DeploymentStatus_WAITING        DeploymentStatus = 3
	DeploymentStatus_AVAILABLE      DeploymentStatus = 4
	DeploymentStatus_SCALING_DOWN   DeploymentStatus = 5
	DeploymentStatus_SCALED_DOWN    DeploymentStatus = 6
	DeploymentStatus_INFRA_MIGRATE  DeploymentStatus = 7
	DeploymentStatus_INFRA_MIGRATED DeploymentStatus = 8
	DeploymentStatus_DB_MIGRATING   DeploymentStatus = 9
	DeploymentStatus_DB_MIGRATED    DeploymentStatus = 10
	DeploymentStatus_SCALING_UP     DeploymentStatus = 11
	DeploymentStatus_SCALED_UP      DeploymentStatus = 12
	DeploymentStatus_CREATING       DeploymentStatus = 13
	DeploymentStatus_UPSERTING      DeploymentStatus = 15
	DeploymentStatus_UPSERTED       DeploymentStatus = 16
	DeploymentStatus_DONE           DeploymentStatus = 100
	DeploymentStatus_FAILED         DeploymentStatus = 101
)

var (
	DeploymentStatus_name_short = map[int32]string{
		0:   "UNSPECIFIED",
		1:   "QUEUED",
		2:   "TRIGGERED",
		3:   "WAITING",
		4:   "AVAILABLE",
		5:   "SCALING_DOWN",
		6:   "SCALED_DOWN",
		7:   "INFRA_MIGRATE",
		8:   "INFRA_MIGRATED",
		9:   "DB_MIGRATING",
		10:  "DB_MIGRATED",
		11:  "SCALING_UP",
		12:  "SCALED_UP",
		13:  "CREATING",
		15:  "UPSERTING",
		16:  "UPSERTED",
		100: "DONE",
		101: "FAILED",
	}
	DeploymentStatus_value_short = map[string]int32{
		"UNSPECIFIED":    0,
		"QUEUED":         1,
		"TRIGGERED":      2,
		"WAITING":        3,
		"AVAILABLE":      4,
		"SCALING_DOWN":   5,
		"SCALED_DOWN":    6,
		"INFRA_MIGRATE":  7,
		"INFRA_MIGRATED": 8,
		"DB_MIGRATING":   9,
		"DB_MIGRATED":    10,
		"SCALING_UP":     11,
		"SCALED_UP":      12,
		"CREATING":       13,
		"UPSERTING":      15,
		"UPSERTED":       16,
		"DONE":           100,
		"FAILED":         101,
	}
	DeploymentStatus_value_either = map[string]int32{
		"UNSPECIFIED":                      0,
		"DEPLOYMENT_STATUS_UNSPECIFIED":    0,
		"QUEUED":                           1,
		"DEPLOYMENT_STATUS_QUEUED":         1,
		"TRIGGERED":                        2,
		"DEPLOYMENT_STATUS_TRIGGERED":      2,
		"WAITING":                          3,
		"DEPLOYMENT_STATUS_WAITING":        3,
		"AVAILABLE":                        4,
		"DEPLOYMENT_STATUS_AVAILABLE":      4,
		"SCALING_DOWN":                     5,
		"DEPLOYMENT_STATUS_SCALING_DOWN":   5,
		"SCALED_DOWN":                      6,
		"DEPLOYMENT_STATUS_SCALED_DOWN":    6,
		"INFRA_MIGRATE":                    7,
		"DEPLOYMENT_STATUS_INFRA_MIGRATE":  7,
		"INFRA_MIGRATED":                   8,
		"DEPLOYMENT_STATUS_INFRA_MIGRATED": 8,
		"DB_MIGRATING":                     9,
		"DEPLOYMENT_STATUS_DB_MIGRATING":   9,
		"DB_MIGRATED":                      10,
		"DEPLOYMENT_STATUS_DB_MIGRATED":    10,
		"SCALING_UP":                       11,
		"DEPLOYMENT_STATUS_SCALING_UP":     11,
		"SCALED_UP":                        12,
		"DEPLOYMENT_STATUS_SCALED_UP":      12,
		"CREATING":                         13,
		"DEPLOYMENT_STATUS_CREATING":       13,
		"UPSERTING":                        15,
		"DEPLOYMENT_STATUS_UPSERTING":      15,
		"UPSERTED":                         16,
		"DEPLOYMENT_STATUS_UPSERTED":       16,
		"DONE":                             100,
		"DEPLOYMENT_STATUS_DONE":           100,
		"FAILED":                           101,
		"DEPLOYMENT_STATUS_FAILED":         101,
	}
)

// ShortString returns the un-prefixed string representation of the enum value
func (x DeploymentStatus) ShortString() string {
	return DeploymentStatus_name_short[int32(x)]
}
func (x DeploymentStatus) Value() (driver.Value, error) {
	return []uint8(x.ShortString()), nil
}
func (x *DeploymentStatus) Scan(value interface{}) error {
	var strVal string
	switch vt := value.(type) {
	case []uint8:
		strVal = string(vt)
	case string:
		strVal = vt
	default:
		return fmt.Errorf("invalid type %T", value)
	}
	val := DeploymentStatus_value_either[strVal]
	*x = DeploymentStatus(val)
	return nil
}

// StackLifecycle
const (
	StackLifecycle_UNSPECIFIED   StackLifecycle = 0
	StackLifecycle_PROGRESS      StackLifecycle = 1
	StackLifecycle_COMPLETE      StackLifecycle = 2
	StackLifecycle_ROLLING_BACK  StackLifecycle = 3
	StackLifecycle_CREATE_FAILED StackLifecycle = 4
	StackLifecycle_TERMINAL      StackLifecycle = 5
	StackLifecycle_ROLLED_BACK   StackLifecycle = 7
	StackLifecycle_MISSING       StackLifecycle = 6
)

var (
	StackLifecycle_name_short = map[int32]string{
		0: "UNSPECIFIED",
		1: "PROGRESS",
		2: "COMPLETE",
		3: "ROLLING_BACK",
		4: "CREATE_FAILED",
		5: "TERMINAL",
		7: "ROLLED_BACK",
		6: "MISSING",
	}
	StackLifecycle_value_short = map[string]int32{
		"UNSPECIFIED":   0,
		"PROGRESS":      1,
		"COMPLETE":      2,
		"ROLLING_BACK":  3,
		"CREATE_FAILED": 4,
		"TERMINAL":      5,
		"ROLLED_BACK":   7,
		"MISSING":       6,
	}
	StackLifecycle_value_either = map[string]int32{
		"UNSPECIFIED":                   0,
		"STACK_LIFECYCLE_UNSPECIFIED":   0,
		"PROGRESS":                      1,
		"STACK_LIFECYCLE_PROGRESS":      1,
		"COMPLETE":                      2,
		"STACK_LIFECYCLE_COMPLETE":      2,
		"ROLLING_BACK":                  3,
		"STACK_LIFECYCLE_ROLLING_BACK":  3,
		"CREATE_FAILED":                 4,
		"STACK_LIFECYCLE_CREATE_FAILED": 4,
		"TERMINAL":                      5,
		"STACK_LIFECYCLE_TERMINAL":      5,
		"ROLLED_BACK":                   7,
		"STACK_LIFECYCLE_ROLLED_BACK":   7,
		"MISSING":                       6,
		"STACK_LIFECYCLE_MISSING":       6,
	}
)

// ShortString returns the un-prefixed string representation of the enum value
func (x StackLifecycle) ShortString() string {
	return StackLifecycle_name_short[int32(x)]
}
func (x StackLifecycle) Value() (driver.Value, error) {
	return []uint8(x.ShortString()), nil
}
func (x *StackLifecycle) Scan(value interface{}) error {
	var strVal string
	switch vt := value.(type) {
	case []uint8:
		strVal = string(vt)
	case string:
		strVal = vt
	default:
		return fmt.Errorf("invalid type %T", value)
	}
	val := StackLifecycle_value_either[strVal]
	*x = StackLifecycle(val)
	return nil
}

// DatabaseMigrationStatus
const (
	DatabaseMigrationStatus_UNSPECIFIED DatabaseMigrationStatus = 0
	DatabaseMigrationStatus_PENDING     DatabaseMigrationStatus = 1
	DatabaseMigrationStatus_RUNNING     DatabaseMigrationStatus = 2
	DatabaseMigrationStatus_CLEANUP     DatabaseMigrationStatus = 3
	DatabaseMigrationStatus_COMPLETED   DatabaseMigrationStatus = 4
	DatabaseMigrationStatus_FAILED      DatabaseMigrationStatus = 5
)

var (
	DatabaseMigrationStatus_name_short = map[int32]string{
		0: "UNSPECIFIED",
		1: "PENDING",
		2: "RUNNING",
		3: "CLEANUP",
		4: "COMPLETED",
		5: "FAILED",
	}
	DatabaseMigrationStatus_value_short = map[string]int32{
		"UNSPECIFIED": 0,
		"PENDING":     1,
		"RUNNING":     2,
		"CLEANUP":     3,
		"COMPLETED":   4,
		"FAILED":      5,
	}
	DatabaseMigrationStatus_value_either = map[string]int32{
		"UNSPECIFIED":                           0,
		"DATABASE_MIGRATION_STATUS_UNSPECIFIED": 0,
		"PENDING":                               1,
		"DATABASE_MIGRATION_STATUS_PENDING":     1,
		"RUNNING":                               2,
		"DATABASE_MIGRATION_STATUS_RUNNING":     2,
		"CLEANUP":                               3,
		"DATABASE_MIGRATION_STATUS_CLEANUP":     3,
		"COMPLETED":                             4,
		"DATABASE_MIGRATION_STATUS_COMPLETED":   4,
		"FAILED":                                5,
		"DATABASE_MIGRATION_STATUS_FAILED":      5,
	}
)

// ShortString returns the un-prefixed string representation of the enum value
func (x DatabaseMigrationStatus) ShortString() string {
	return DatabaseMigrationStatus_name_short[int32(x)]
}
func (x DatabaseMigrationStatus) Value() (driver.Value, error) {
	return []uint8(x.ShortString()), nil
}
func (x *DatabaseMigrationStatus) Scan(value interface{}) error {
	var strVal string
	switch vt := value.(type) {
	case []uint8:
		strVal = string(vt)
	case string:
		strVal = vt
	default:
		return fmt.Errorf("invalid type %T", value)
	}
	val := DatabaseMigrationStatus_value_either[strVal]
	*x = DatabaseMigrationStatus(val)
	return nil
}
