/*
 * coslaNrm
 *
 * OAS 3.0.1 specification of the Cosla NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 17.3.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_CoslaNrm

type ThresholdMonitorSingleAllOfAttributes struct {

	AdministrativeState AdministrativeState `json:"administrativeState,omitempty"`

	OperationalState OperationalState `json:"operationalState,omitempty"`

	PerformanceMetrics []string `json:"performanceMetrics,omitempty"`

	ThresholdInfoList []ThresholdInfo `json:"thresholdInfoList,omitempty"`

	MonitorGranularityPeriod int32 `json:"monitorGranularityPeriod,omitempty"`

	ObjectInstances []string `json:"objectInstances,omitempty"`

	RootObjectInstances []string `json:"rootObjectInstances,omitempty"`
}