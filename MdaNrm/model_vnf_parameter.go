/*
 * MDA NRM
 *
 * OAS 3.0.1 specification of the MDA NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 17.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_MdaNrm

type VnfParameter struct {

	VnfInstanceId string `json:"vnfInstanceId,omitempty"`

	VnfdId string `json:"vnfdId,omitempty"`

	FlavourId string `json:"flavourId,omitempty"`

	AutoScalable bool `json:"autoScalable,omitempty"`
}
