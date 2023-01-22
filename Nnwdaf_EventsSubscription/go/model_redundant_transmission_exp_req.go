/*
 * Nnwdaf_EventsSubscription
 *
 * Nnwdaf_EventsSubscription Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nnwdaf_EventsSubscription

// RedundantTransmissionExpReq - Represents other redundant transmission experience analytics requirements.
type RedundantTransmissionExpReq struct {

	RedTOrderCriter RedTransExpOrderingCriterion `json:"redTOrderCriter,omitempty"`

	Order MatchingDirection `json:"order,omitempty"`
}