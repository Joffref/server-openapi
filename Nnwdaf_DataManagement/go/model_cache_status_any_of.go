/*
 * Nnwdaf_DataManagement
 *
 * Nnwdaf_DataManagement API Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nnwdaf_DataManagement

type CacheStatusAnyOf string

// List of CacheStatusAnyOf
const (
	HIT CacheStatusAnyOf = "HIT"
	MISS CacheStatusAnyOf = "MISS"
	EXPIRED CacheStatusAnyOf = "EXPIRED"
)