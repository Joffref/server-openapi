/*
 * AI/ML NRM
 *
 * OAS 3.0.1 specification of the AI/ML NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 17.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package AiMlNrm

type PerfMetricJobSingleAllOfAttributes struct {

	AdministrativeState AdministrativeState `json:"administrativeState,omitempty"`

	OperationalState OperationalState `json:"operationalState,omitempty"`

	JobId string `json:"jobId,omitempty"`

	PerformanceMetrics []string `json:"performanceMetrics,omitempty"`

	GranularityPeriod int32 `json:"granularityPeriod,omitempty"`

	ObjectInstances []string `json:"objectInstances,omitempty"`

	RootObjectInstances []string `json:"rootObjectInstances,omitempty"`

	ReportingCtrl ReportingCtrl `json:"reportingCtrl,omitempty"`
}