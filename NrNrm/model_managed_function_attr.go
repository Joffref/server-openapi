/*
 * NR NRM
 *
 * OAS 3.0.1 specification of the NR NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 17.7.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_NrNrm

type ManagedFunctionAttr struct {

	UserLabel string `json:"userLabel,omitempty"`

	VnfParametersList []VnfParameter `json:"vnfParametersList,omitempty"`

	PeeParametersList []PeeParameter `json:"peeParametersList,omitempty"`

	PriorityLabel int32 `json:"priorityLabel,omitempty"`

	SupportedPerfMetricGroups []SupportedPerfMetricGroup `json:"supportedPerfMetricGroups,omitempty"`

	SupportedTraceMetrics []string `json:"supportedTraceMetrics,omitempty"`
}
