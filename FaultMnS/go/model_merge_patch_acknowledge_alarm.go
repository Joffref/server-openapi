/*
 * Fault Supervision MnS
 *
 * OAS 3.0.1 definition of the Fault Supervision MnS © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 17.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package FaultMnS

// MergePatchAcknowledgeAlarm - Patch document acknowledging or unacknowledging a single alarm. For acknowleding an alarm the value of ackState is ACKNOWLEDGED, for unacknowleding an alarm the value of ackState is UNACKNOWLEDGED.
type MergePatchAcknowledgeAlarm struct {

	AckUserId string `json:"ackUserId"`

	AckSystemId string `json:"ackSystemId,omitempty"`

	AckState AckState `json:"ackState"`
}