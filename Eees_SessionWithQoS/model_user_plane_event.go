/*
 * EES Session with QoS API
 *
 * API for EES Session with Qos service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Eees_SessionWithQoS

// UserPlaneEvent - Possible values are - SESSION_TERMINATION: Indicates that Rx session is terminated. - LOSS_OF_BEARER : Indicates a loss of a bearer. - RECOVERY_OF_BEARER: Indicates a recovery of a bearer. - RELEASE_OF_BEARER: Indicates a release of a bearer. - USAGE_REPORT: Indicates the usage report event. - FAILED_RESOURCES_ALLOCATION: Indicates the resource allocation is failed. - QOS_GUARANTEED: The QoS targets of one or more SDFs are guaranteed again. - QOS_NOT_GUARANTEED: The QoS targets of one or more SDFs are not being guaranteed. - QOS_MONITORING: Indicates a QoS monitoring event. - SUCCESSFUL_RESOURCES_ALLOCATION: Indicates the resource allocation is successful. - ACCESS_TYPE_CHANGE: Indicates an Access type change. - PLMN_CHG: Indicates a PLMN change. 
type UserPlaneEvent struct {
}
