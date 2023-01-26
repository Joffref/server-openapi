/*
 * NR NRM
 *
 * OAS 3.0.1 specification of the NR NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 17.7.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_NrNrm

type ManagementDataCollectionSingleAllOfAttributes struct {

	ManagementData ManagementData `json:"managementData,omitempty"`

	TargetNodeFilter NodeFilter `json:"targetNodeFilter,omitempty"`

	CollectionTimeWindow TimeWindow `json:"collectionTimeWindow,omitempty"`

	ReportingCtrl string `json:"reportingCtrl,omitempty"`

	DataScope string `json:"dataScope,omitempty"`
}
