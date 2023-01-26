/*
 * nmbsf-mbs-ud-ingest
 *
 * API for MBS User Data Ingest Session Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nmbsf_MBSUserDataIngestSession

type EventOneOf string

// List of EventOneOf
const (
	USER_DATA_ING_SESS_STARTING EventOneOf = "USER_DATA_ING_SESS_STARTING"
	USER_DATA_ING_SESS_STARTED EventOneOf = "USER_DATA_ING_SESS_STARTED"
	USER_DATA_ING_SESS_TERMINATED EventOneOf = "USER_DATA_ING_SESS_TERMINATED"
	DIST_SESS_STARTING EventOneOf = "DIST_SESS_STARTING"
	DIST_SESS_STARTED EventOneOf = "DIST_SESS_STARTED"
	DIST_SESS_TERMINATED EventOneOf = "DIST_SESS_TERMINATED"
	DIST_SESS_SERV_MNGT_FAILURE EventOneOf = "DIST_SESS_SERV_MNGT_FAILURE"
	DIST_SESS_POL_CRTL_FAILURE EventOneOf = "DIST_SESS_POL_CRTL_FAILURE"
	DATA_INGEST_FAILURE EventOneOf = "DATA_INGEST_FAILURE"
	DELIVERY_STARTED EventOneOf = "DELIVERY_STARTED"
	SESSION_TERMINATED EventOneOf = "SESSION_TERMINATED"
)
