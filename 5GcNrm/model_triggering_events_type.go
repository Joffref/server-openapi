/*
 * 3GPP 5GC NRM
 *
 * OAS 3.0.1 specification of the 5GC NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 17.7.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_5GcNrm

// TriggeringEventsType - Specifies when to start a Trace Recording Session and which message shall be recorded first, when to stop a Trace Recording Session and which message shall be recorded last respectively. See 3GPP TS 32.422 clause 5.1 for additional detials.
type TriggeringEventsType struct {

	MSC_SERVER []string `json:"MSC_SERVER,omitempty"`

	SGSN []string `json:"SGSN,omitempty"`

	MGW []string `json:"MGW,omitempty"`

	GGSN []string `json:"GGSN,omitempty"`

	IMS []string `json:"IMS,omitempty"`

	BM_SC []string `json:"BM_SC,omitempty"`

	MME []string `json:"MME,omitempty"`

	SGW []string `json:"SGW,omitempty"`

	PGW []string `json:"PGW,omitempty"`

	AMF []string `json:"AMF,omitempty"`

	SMF []string `json:"SMF,omitempty"`

	PCF []string `json:"PCF,omitempty"`

	UPF []string `json:"UPF,omitempty"`

	AUSF []string `json:"AUSF,omitempty"`

	NEF []string `json:"NEF,omitempty"`

	NRF []string `json:"NRF,omitempty"`

	NSSF []string `json:"NSSF,omitempty"`

	SMSF []string `json:"SMSF,omitempty"`

	UDM []string `json:"UDM,omitempty"`
}
