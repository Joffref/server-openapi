/*
 * coslaNrm
 *
 * OAS 3.0.1 specification of the Cosla NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 17.3.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_CoslaNrm

type ManagementNodeSingleAllOfAttributes struct {

	UserLabel string `json:"userLabel,omitempty"`

	ManagedElements []string `json:"managedElements,omitempty"`

	VendorName string `json:"vendorName,omitempty"`

	UserDefinedState string `json:"userDefinedState,omitempty"`

	LocationName string `json:"locationName,omitempty"`

	SwVersion string `json:"swVersion,omitempty"`
}
