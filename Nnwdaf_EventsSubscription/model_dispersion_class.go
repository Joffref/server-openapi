/*
 * Nnwdaf_EventsSubscription
 *
 * Nnwdaf_EventsSubscription Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nnwdaf_EventsSubscription

// DispersionClass - Possible values are: - FIXED: Dispersion class as fixed UE its data or transaction usage at a location or a slice, is higher than its class threshold set for its all data or transaction usage. - CAMPER: Dispersion class as camper UE, its data or transaction usage at a location or a slice, is higher than its class threshold and lower than the fixed class threshold set for its all data or transaction usage.. - TRAVELLER: Dispersion class as traveller UE, its data or transaction usage at a location or a slice, is lower than the camper class threshold set for its all data or transaction usage. - TOP_HEAVY: Dispersion class as Top_Heavy UE, who's dispersion percentile rating at a location or a slice, is higher than its class threshold. 
type DispersionClass struct {
}
