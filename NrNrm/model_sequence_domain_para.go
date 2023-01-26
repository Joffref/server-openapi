/*
 * NR NRM
 *
 * OAS 3.0.1 specification of the NR NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 17.7.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_NrNrm

type SequenceDomainPara struct {

	NrofRIMRSSequenceCandidatesofRS1 int32 `json:"nrofRIMRSSequenceCandidatesofRS1,omitempty"`

	RimRSScrambleIdListofRS1 []int32 `json:"rimRSScrambleIdListofRS1,omitempty"`

	NrofRIMRSSequenceCandidatesofRS2 int32 `json:"nrofRIMRSSequenceCandidatesofRS2,omitempty"`

	RimRSScrambleIdListofRS2 []int32 `json:"rimRSScrambleIdListofRS2,omitempty"`

	EnableEnoughNotEnoughIndication string `json:"enableEnoughNotEnoughIndication,omitempty"`

	RIMRSScrambleTimerMultiplier int32 `json:"RIMRSScrambleTimerMultiplier,omitempty"`

	RIMRSScrambleTimerOffset int32 `json:"RIMRSScrambleTimerOffset,omitempty"`
}