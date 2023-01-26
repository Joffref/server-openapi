/*
 * coslaNrm
 *
 * OAS 3.0.1 specification of the Cosla NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 17.3.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_CoslaNrm

type MnsInfoSingle struct {

	MnsLabel string `json:"mnsLabel,omitempty"`

	MnsType string `json:"mnsType,omitempty"`

	MnsVersion string `json:"mnsVersion,omitempty"`

	MnsAddress string `json:"mnsAddress,omitempty"`

	// List of the managed object instances that can be accessed using the MnS. If a complete SubNetwork can be accessed using the MnS, this attribute may contain the DN of the SubNetwork instead of the DNs of the individual managed entities within the SubNetwork.
	MnsScope []string `json:"mnsScope,omitempty"`
}
