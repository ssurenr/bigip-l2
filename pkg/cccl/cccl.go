package cccl

import "fmt"
import "encoding/json"

// Defines an ARP entry.
type ArpType struct {
	// IpAddress corresponds to the JSON schema field "ipAddress".
	IpAddress string `json:"ipAddress"`

	// MacAddress corresponds to the JSON schema field "macAddress".
	MacAddress string `json:"macAddress"`

	// Name corresponds to the JSON schema field "name".
	Name string `json:"name"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ArpType) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["ipAddress"]; !ok || v == nil {
		return fmt.Errorf("field ipAddress: required")
	}
	if v, ok := raw["macAddress"]; !ok || v == nil {
		return fmt.Errorf("field macAddress: required")
	}
	if v, ok := raw["name"]; !ok || v == nil {
		return fmt.Errorf("field name: required")
	}
	type Plain ArpType
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = ArpType(plain)
	return nil
}

// Defines an FDB tunnel.
type FdbTunnelType struct {
	// Name corresponds to the JSON schema field "name".
	Name string `json:"name"`

	// Records corresponds to the JSON schema field "records".
	Records interface{} `json:"records"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *FdbTunnelType) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["name"]; !ok || v == nil {
		return fmt.Errorf("field name: required")
	}
	if v, ok := raw["records"]; !ok || v == nil {
		return fmt.Errorf("field records: required")
	}
	type Plain FdbTunnelType
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = FdbTunnelType(plain)
	return nil
}

type RecordType struct {
	// Endpoint corresponds to the JSON schema field "endpoint".
	Endpoint string `json:"endpoint"`

	// Name of the record (MAC address).
	Name string `json:"name"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *RecordType) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["endpoint"]; !ok || v == nil {
		return fmt.Errorf("field endpoint: required")
	}
	if v, ok := raw["name"]; !ok || v == nil {
		return fmt.Errorf("field name: required")
	}
	type Plain RecordType
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = RecordType(plain)
	return nil
}

// The CCCL "Cecil" library allows clients to define services that describe
// NET resources on a managed partition of the BIG-IP.  The managed resources
// are defined in this schema definitions section.
// The structure of the service definition is a collection of lists of
// supported resources.  Initially this is ARPs and FDB tunnel records.
// Where appropriate some basic constraints are defined by the schema; however,
// not all actual constraints can be enforced by the schema.  It is the
// responsibility of the client application to ensure that all dependencies
// among the specified resources are met; otherwise, the service will be deployed
// in a degraded state.
//
type CcclNet struct {
	// List of all ARP resources that should exist
	Arps []ArpType `json:"arps,omitempty"`

	// List of all FDB tunnel resources that should exist
	FdbTunnels []FdbTunnelType `json:"fdbTunnels,omitempty"`

	// List of user-created FDB tunnel resources to be updated. These are expected to
	// be administratively created beforehand. CCCL will perform updates only on these
	// tunnels, no deletion or creation.
	//
	UserFdbTunnels []FdbTunnelType `json:"userFdbTunnels,omitempty"`
}
