/*
 * Nnwdaf_EventsSubscription
 *
 * Nnwdaf_EventsSubscription Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nnwdaf_EventsSubscription

// AnalyticsContextIdentifier - Contains information about available analytics contexts.
type AnalyticsContextIdentifier struct {

	// The identifier of a subscription.
	SubscriptionId string `json:"subscriptionId,omitempty"`

	// List of analytics types for which NF related analytics contexts can be retrieved. 
	NfAnaCtxts []NwdafEvent `json:"nfAnaCtxts,omitempty"`

	// List of objects that indicate for which SUPI and analytics types combinations analytics  context can be retrieved. 
	UeAnaCtxts []UeAnalyticsContextDescriptor `json:"ueAnaCtxts,omitempty"`
}
