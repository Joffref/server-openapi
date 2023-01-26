/*
 * NR NRM
 *
 * OAS 3.0.1 specification of the NR NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 17.7.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_NrNrm

import (
	"time"
)

type AlarmListSingleAllOfAttributes struct {

	AdministrativeState AdministrativeState `json:"administrativeState,omitempty"`

	OperationalState OperationalState `json:"operationalState,omitempty"`

	NumOfAlarmRecords int32 `json:"numOfAlarmRecords,omitempty"`

	LastModification time.Time `json:"lastModification,omitempty"`

	// This resource represents a map of alarm records. The alarmIds are used as keys in the map.
	AlarmRecords map[string]AlarmRecord `json:"alarmRecords,omitempty"`
}
