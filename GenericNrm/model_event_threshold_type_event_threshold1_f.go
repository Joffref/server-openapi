/*
 * Generic NRM
 *
 * OAS 3.0.1 definition of the Generic NRM © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 17.2.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_GenericNrm

type EventThresholdTypeEventThreshold1F struct {

	CPICH_RSCP int32 `json:"CPICH_RSCP,omitempty"`

	CPICHEcNo int32 `json:"CPICH_EcNo,omitempty"`

	PathLoss int32 `json:"PathLoss,omitempty"`
}
