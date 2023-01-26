/*
 * Unified Data Repository Service API file for subscription data
 *
 * Unified Data Repository Service (subscription data).   The API version is defined in 3GPP TS 29.504.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: -
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Subscription_Data

// BatteryIndication - Parameters \"replaceableInd\" and \"rechargeableInd\" are only included if the value of Parameter \"batteryInd\" is true. 
type BatteryIndication struct {

	// This IE shall indicate whether the UE is battery powered or not. true: the UE is battery powered; false or absent: the UE is not battery powered 
	BatteryInd bool `json:"batteryInd,omitempty"`

	// This IE shall indicate whether the battery of the UE is replaceable or not. true: the battery of the UE is replaceable; false or absent: the battery of the UE is not replaceable. 
	ReplaceableInd bool `json:"replaceableInd,omitempty"`

	// This IE shall indicate whether the battery of the UE is rechargeable or not. true: the battery of UE is rechargeable; false or absent: the battery of the UE is not rechargeable. 
	RechargeableInd bool `json:"rechargeableInd,omitempty"`
}