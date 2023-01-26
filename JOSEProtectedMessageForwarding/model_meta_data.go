/*
 * JOSE Protected Message Forwarding API
 *
 * N32-f Message Forwarding Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_JOSEProtectedMessageForwarding

// MetaData - Contains the meta data information needed for replay protection
type MetaData struct {

	N32fContextId string `json:"n32fContextId"`

	MessageId string `json:"messageId"`

	AuthorizedIpxId string `json:"authorizedIpxId"`
}