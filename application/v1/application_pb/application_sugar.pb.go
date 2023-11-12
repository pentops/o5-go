// Code generated by protoc-gen-go-sugar. DO NOT EDIT.

package application_pb

import (
	driver "database/sql/driver"
	fmt "fmt"
)

type IsDatabase_Engine = isDatabase_Engine
type IsContainer_Source = isContainer_Source
type IsEnvironmentVariable_Spec = isEnvironmentVariable_Spec
type IsBlobstoreEnvVar_Format = isBlobstoreEnvVar_Format

// RouteProtocol
const (
	RouteProtocol_UNSPECIFIED RouteProtocol = 0
	RouteProtocol_HTTP        RouteProtocol = 1
	RouteProtocol_GRPC        RouteProtocol = 2
)

var (
	RouteProtocol_name_short = map[int32]string{
		0: "UNSPECIFIED",
		1: "HTTP",
		2: "GRPC",
	}
	RouteProtocol_value_short = map[string]int32{
		"UNSPECIFIED": 0,
		"HTTP":        1,
		"GRPC":        2,
	}
	RouteProtocol_value_either = map[string]int32{
		"UNSPECIFIED":                0,
		"ROUTE_PROTOCOL_UNSPECIFIED": 0,
		"HTTP":                       1,
		"ROUTE_PROTOCOL_HTTP":        1,
		"GRPC":                       2,
		"ROUTE_PROTOCOL_GRPC":        2,
	}
)

// ShortString returns the un-prefixed string representation of the enum value
func (x RouteProtocol) ShortString() string {
	return RouteProtocol_name_short[int32(x)]
}
func (x RouteProtocol) Value() (driver.Value, error) {
	return []uint8(x.ShortString()), nil
}
func (x *RouteProtocol) Scan(value interface{}) error {
	var strVal string
	switch vt := value.(type) {
	case []uint8:
		strVal = string(vt)
	case string:
		strVal = vt
	default:
		return fmt.Errorf("invalid type %T", value)
	}
	val := RouteProtocol_value_either[strVal]
	*x = RouteProtocol(val)
	return nil
}

// RouteGroup
const (
	RouteGroup_UNSPECIFIED RouteGroup = 0
	RouteGroup_FIRST       RouteGroup = 1
	RouteGroup_NORMAL      RouteGroup = 2
	RouteGroup_FALLBACK    RouteGroup = 3
)

var (
	RouteGroup_name_short = map[int32]string{
		0: "UNSPECIFIED",
		1: "FIRST",
		2: "NORMAL",
		3: "FALLBACK",
	}
	RouteGroup_value_short = map[string]int32{
		"UNSPECIFIED": 0,
		"FIRST":       1,
		"NORMAL":      2,
		"FALLBACK":    3,
	}
	RouteGroup_value_either = map[string]int32{
		"UNSPECIFIED":             0,
		"ROUTE_GROUP_UNSPECIFIED": 0,
		"FIRST":                   1,
		"ROUTE_GROUP_FIRST":       1,
		"NORMAL":                  2,
		"ROUTE_GROUP_NORMAL":      2,
		"FALLBACK":                3,
		"ROUTE_GROUP_FALLBACK":    3,
	}
)

// ShortString returns the un-prefixed string representation of the enum value
func (x RouteGroup) ShortString() string {
	return RouteGroup_name_short[int32(x)]
}
func (x RouteGroup) Value() (driver.Value, error) {
	return []uint8(x.ShortString()), nil
}
func (x *RouteGroup) Scan(value interface{}) error {
	var strVal string
	switch vt := value.(type) {
	case []uint8:
		strVal = string(vt)
	case string:
		strVal = vt
	default:
		return fmt.Errorf("invalid type %T", value)
	}
	val := RouteGroup_value_either[strVal]
	*x = RouteGroup(val)
	return nil
}

// Demand
const (
	Demand_UNSPECIFIED Demand = 0
	Demand_LIGHT       Demand = 1
	Demand_MEDIUM      Demand = 2
	Demand_HEAVY       Demand = 3
)

var (
	Demand_name_short = map[int32]string{
		0: "UNSPECIFIED",
		1: "LIGHT",
		2: "MEDIUM",
		3: "HEAVY",
	}
	Demand_value_short = map[string]int32{
		"UNSPECIFIED": 0,
		"LIGHT":       1,
		"MEDIUM":      2,
		"HEAVY":       3,
	}
	Demand_value_either = map[string]int32{
		"UNSPECIFIED":        0,
		"DEMAND_UNSPECIFIED": 0,
		"LIGHT":              1,
		"DEMAND_LIGHT":       1,
		"MEDIUM":             2,
		"DEMAND_MEDIUM":      2,
		"HEAVY":              3,
		"DEMAND_HEAVY":       3,
	}
)

// ShortString returns the un-prefixed string representation of the enum value
func (x Demand) ShortString() string {
	return Demand_name_short[int32(x)]
}
func (x Demand) Value() (driver.Value, error) {
	return []uint8(x.ShortString()), nil
}
func (x *Demand) Scan(value interface{}) error {
	var strVal string
	switch vt := value.(type) {
	case []uint8:
		strVal = string(vt)
	case string:
		strVal = vt
	default:
		return fmt.Errorf("invalid type %T", value)
	}
	val := Demand_value_either[strVal]
	*x = Demand(val)
	return nil
}