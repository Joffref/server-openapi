/*
 * NRF OAuth2
 *
 * NRF OAuth2 Authorization.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nnrf_AccessToken

// AccessTokenRsp - Contains information related to the access token response
type AccessTokenRsp struct {

	// JWS Compact Serialized representation of JWS signed JSON object (AccessTokenClaims) 
	AccessToken string `json:"access_token"`

	TokenType string `json:"token_type"`

	ExpiresIn int32 `json:"expires_in,omitempty"`

	Scope string `json:"scope,omitempty"`
}
