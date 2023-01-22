/*
 * AUSF API
 *
 * AUSF UE Authentication Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nausf_UEAuthentication
// AuthResult : Indicates the result of the authentication.
type AuthResult string

// List of AuthResult
const (
	SUCCESS AuthResult = "AUTHENTICATION_SUCCESS"
	FAILURE AuthResult = "AUTHENTICATION_FAILURE"
	ONGOING AuthResult = "AUTHENTICATION_ONGOING"
)