/*
 * Intent NRM
 *
 * OAS 3.0.1 definition of the Intent NRM © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 17.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_IntentNrm

type VsDataContainerSingleAttributes struct {

	VsDataType string `json:"vsDataType,omitempty"`

	VsDataFormatVersion string `json:"vsDataFormatVersion,omitempty"`

	VsData *interface{} `json:"vsData,omitempty"`
}
