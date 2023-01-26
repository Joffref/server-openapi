/*
 * 3GPP Edge NRM
 *
 * OAS 3.0.1 specification of the Edge NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 17.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_EdgeNrm

import (
	"time"
)

// AlarmRecord - The alarmId is not a property of an alarm record. It is used as key in the map of alarm records instead.
type AlarmRecord struct {

	ObjectInstance string `json:"objectInstance,omitempty"`

	NotificationId int32 `json:"notificationId,omitempty"`

	AlarmRaisedTime time.Time `json:"alarmRaisedTime,omitempty"`

	AlarmChangedTime time.Time `json:"alarmChangedTime,omitempty"`

	AlarmClearedTime time.Time `json:"alarmClearedTime,omitempty"`

	AlarmType AlarmType `json:"alarmType,omitempty"`

	ProbableCause ProbableCause `json:"probableCause,omitempty"`

	SpecificProblem SpecificProblem `json:"specificProblem,omitempty"`

	PerceivedSeverity PerceivedSeverity `json:"perceivedSeverity,omitempty"`

	BackedUpStatus bool `json:"backedUpStatus,omitempty"`

	BackUpObject string `json:"backUpObject,omitempty"`

	TrendIndication TrendIndication `json:"trendIndication,omitempty"`

	Thresholdinfo ThresholdInfo1 `json:"thresholdinfo,omitempty"`

	CorrelatedNotifications []CorrelatedNotification `json:"correlatedNotifications,omitempty"`

	// The first array item contains the attribute name value pairs with the new values, and the second array item the attribute name value pairs with the optional old values.
	StateChangeDefinition []map[string]interface{} `json:"stateChangeDefinition,omitempty"`

	// The key of this map is the attribute name, and the value the attribute value.
	MonitoredAttributes map[string]interface{} `json:"monitoredAttributes,omitempty"`

	ProposedRepairActions string `json:"proposedRepairActions,omitempty"`

	AdditionalText string `json:"additionalText,omitempty"`

	// The key of this map is the attribute name, and the value the attribute value.
	AdditionalInformation map[string]interface{} `json:"additionalInformation,omitempty"`

	RootCauseIndicator bool `json:"rootCauseIndicator,omitempty"`

	AckTime time.Time `json:"ackTime,omitempty"`

	AckUserId string `json:"ackUserId,omitempty"`

	AckSystemId string `json:"ackSystemId,omitempty"`

	AckState AckState `json:"ackState,omitempty"`

	ClearUserId string `json:"clearUserId,omitempty"`

	ClearSystemId string `json:"clearSystemId,omitempty"`

	ServiceUser string `json:"serviceUser,omitempty"`

	ServiceProvider string `json:"serviceProvider,omitempty"`

	SecurityAlarmDetector string `json:"securityAlarmDetector,omitempty"`
}
