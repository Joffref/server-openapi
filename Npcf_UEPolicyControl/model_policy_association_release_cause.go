/*
 * Npcf_UEPolicyControl
 *
 * UE Policy Control Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Npcf_UEPolicyControl

// PolicyAssociationReleaseCause - Possible values are: - UNSPECIFIED: This value is used for unspecified reasons. - UE_SUBSCRIPTION: This value is used to indicate that the policy association needs to be   terminated because the subscription of UE has changed (e.g. was removed). - INSUFFICIENT_RES: This value is used to indicate that the server is overloaded and needs   to abort the policy association. 
type PolicyAssociationReleaseCause struct {
}
