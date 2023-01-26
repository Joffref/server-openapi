/*
 * Generic NRM
 *
 * OAS 3.0.1 definition of the Generic NRM © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 17.2.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_GenericNrm

type SubNetworkNcO struct {

	ManagementNode []ManagementNodeSingle `json:"ManagementNode,omitempty"`

	MnsAgent []MnsAgentSingle `json:"MnsAgent,omitempty"`

	MeContext []MeContextSingle `json:"MeContext,omitempty"`

	PerfMetricJob []PerfMetricJobSingle `json:"PerfMetricJob,omitempty"`

	ThresholdMonitor []ThresholdMonitorSingle `json:"ThresholdMonitor,omitempty"`

	TraceJob []TraceJobSingle `json:"TraceJob,omitempty"`

	ManagementDataCollection []ManagementDataCollectionSingle `json:"ManagementDataCollection,omitempty"`

	NtfSubscriptionControl []NtfSubscriptionControlSingle `json:"NtfSubscriptionControl,omitempty"`

	AlarmList AlarmListSingle `json:"AlarmList,omitempty"`

	Files []FilesSingle `json:"Files,omitempty"`

	FileDownloadJob []FileDownloadJobSingle `json:"FileDownloadJob,omitempty"`

	MnsRegistry MnsRegistrySingle `json:"MnsRegistry,omitempty"`
}
