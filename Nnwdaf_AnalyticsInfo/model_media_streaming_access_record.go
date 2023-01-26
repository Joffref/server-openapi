/*
 * Nnwdaf_AnalyticsInfo
 *
 * Nnwdaf_AnalyticsInfo Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nnwdaf_AnalyticsInfo

type MediaStreamingAccessRecord struct {

	// string with format 'date-time' as defined in OpenAPI.
	Timestamp time.Time `json:"timestamp"`

	MediaStreamHandlerEndpointAddress EndpointAddress `json:"mediaStreamHandlerEndpointAddress"`

	ApplicationServerEndpointAddress EndpointAddress `json:"applicationServerEndpointAddress"`

	SessionIdentifier string `json:"sessionIdentifier,omitempty"`

	RequestMessage MediaStreamingAccessRecordAllOfRequestMessage `json:"requestMessage"`

	CacheStatus CacheStatus `json:"cacheStatus,omitempty"`

	ResponseMessage MediaStreamingAccessRecordAllOfResponseMessage `json:"responseMessage"`

	// string with format 'float' as defined in OpenAPI.
	ProcessingLatency float32 `json:"processingLatency"`

	ConnectionMetrics MediaStreamingAccessRecordAllOfConnectionMetrics `json:"connectionMetrics,omitempty"`
}
