/*
 * 3GPP Edge NRM
 *
 * OAS 3.0.1 specification of the Edge NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 17.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_EdgeNrm

type NfServiceType string

// List of NfServiceType
const (
	NAMF_COMMUNICATION NfServiceType = "Namf_Communication"
	NAMF_EVENT_EXPOSURE NfServiceType = "Namf_EventExposure"
	NAMF_MT NfServiceType = "Namf_MT"
	NAMF_LOCATION NfServiceType = "Namf_Location"
	NSMF_PDU_SESSION NfServiceType = "Nsmf_PDUSession"
	NSMF_EVENT_EXPOSURE NfServiceType = "Nsmf_EventExposure"
	OTHERS NfServiceType = "Others"
)
