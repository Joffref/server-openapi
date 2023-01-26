/*
 * Nhss_imsUEAU
 *
 * Nhss UE Authentication Service for IMS.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nhss_imsUEAU

type SipAuthenticationSchemeAnyOf string

// List of SipAuthenticationSchemeAnyOf
const (
	DIGEST_AKAV1_MD5 SipAuthenticationSchemeAnyOf = "DIGEST-AKAV1-MD5"
	DIGEST_HTTP SipAuthenticationSchemeAnyOf = "DIGEST-HTTP"
	NBA SipAuthenticationSchemeAnyOf = "NBA"
	GIBA SipAuthenticationSchemeAnyOf = "GIBA"
	UNKNOWN SipAuthenticationSchemeAnyOf = "UNKNOWN"
)
