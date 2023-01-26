/*
 * 3gpp-service-parameter
 *
 * API for AF service paramter   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.1.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_ServiceParameter

// Failure - Possible values are: - UNSPECIFIED: Indicates the PCF received the UE sent UE policy delivery service cause #111 (Protocol error, unspecified). - UE_NOT_REACHABLE: Indicates the PCF received the notification from the AMF that the UE is not reachable. - UNKNOWN: Indicates unknown reasons upon no response from the UE, e.g. UPDS message type is not defined or not implemented by the UE, or not compatible with the UPDS state, in which the UE shall ignore the UPDS message. - UE_TEMP_UNREACHABLE: Indicates the PCF received the notification from the AMF that the UE is not reachable but the PCF will retry again. 
type Failure struct {
}
