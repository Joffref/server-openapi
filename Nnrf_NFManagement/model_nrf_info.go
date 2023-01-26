/*
 * NRF NFManagement Service
 *
 * NRF NFManagement Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.3.0-alpha.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nnrf_NFManagement

// NrfInfo - Information of an NRF NF Instance, used in hierarchical NRF deployments
type NrfInfo struct {

	// A map (list of key-value pairs) where nfInstanceId serves as key
	ServedUdrInfo map[string]NrfInfoServedUdrInfoValue `json:"servedUdrInfo,omitempty"`

	// A map (list of key-value pairs) where nfInstanceId serves as key
	ServedUdrInfoList map[string]map[string]NrfInfoServedUdrInfoValue `json:"servedUdrInfoList,omitempty"`

	// A map (list of key-value pairs) where nfInstanceId serves as key
	ServedUdmInfo map[string]NrfInfoServedUdmInfoValue `json:"servedUdmInfo,omitempty"`

	// A map (list of key-value pairs) where nfInstanceId serves as key
	ServedUdmInfoList map[string]map[string]NrfInfoServedUdmInfoValue `json:"servedUdmInfoList,omitempty"`

	// A map (list of key-value pairs) where nfInstanceId serves as key
	ServedAusfInfo map[string]NrfInfoServedAusfInfoValue `json:"servedAusfInfo,omitempty"`

	// A map (list of key-value pairs) where nfInstanceId serves as key
	ServedAusfInfoList map[string]map[string]NrfInfoServedAusfInfoValue `json:"servedAusfInfoList,omitempty"`

	// A map (list of key-value pairs) where nfInstanceId serves as key
	ServedAmfInfo map[string]NrfInfoServedAmfInfoValue `json:"servedAmfInfo,omitempty"`

	// A map (list of key-value pairs) where nfInstanceId serves as key
	ServedAmfInfoList map[string]map[string]NrfInfoServedAmfInfoValue `json:"servedAmfInfoList,omitempty"`

	// A map (list of key-value pairs) where nfInstanceId serves as key
	ServedSmfInfo map[string]NrfInfoServedSmfInfoValue `json:"servedSmfInfo,omitempty"`

	// A map (list of key-value pairs) where nfInstanceId serves as key
	ServedSmfInfoList map[string]map[string]NrfInfoServedSmfInfoValue `json:"servedSmfInfoList,omitempty"`

	// A map (list of key-value pairs) where nfInstanceId serves as key
	ServedUpfInfo map[string]NrfInfoServedUpfInfoValue `json:"servedUpfInfo,omitempty"`

	// A map (list of key-value pairs) where nfInstanceId serves as key
	ServedUpfInfoList map[string]map[string]NrfInfoServedUpfInfoValue `json:"servedUpfInfoList,omitempty"`

	// A map (list of key-value pairs) where nfInstanceId serves as key
	ServedPcfInfo map[string]NrfInfoServedPcfInfoValue `json:"servedPcfInfo,omitempty"`

	// A map (list of key-value pairs) where nfInstanceId serves as key
	ServedPcfInfoList map[string]map[string]NrfInfoServedPcfInfoValue `json:"servedPcfInfoList,omitempty"`

	// A map (list of key-value pairs) where nfInstanceId serves as key
	ServedBsfInfo map[string]NrfInfoServedBsfInfoValue `json:"servedBsfInfo,omitempty"`

	// A map (list of key-value pairs) where nfInstanceId serves as key
	ServedBsfInfoList map[string]map[string]NrfInfoServedBsfInfoValue `json:"servedBsfInfoList,omitempty"`

	// A map (list of key-value pairs) where nfInstanceId serves as key
	ServedChfInfo map[string]NrfInfoServedChfInfoValue `json:"servedChfInfo,omitempty"`

	// A map (list of key-value pairs) where nfInstanceId serves as key
	ServedChfInfoList map[string]map[string]NrfInfoServedChfInfoValue `json:"servedChfInfoList,omitempty"`

	// A map (list of key-value pairs) where nfInstanceId serves as key
	ServedNefInfo map[string]NrfInfoServedNefInfoValue `json:"servedNefInfo,omitempty"`

	// A map (list of key-value pairs) where nfInstanceId serves as key
	ServedNwdafInfo map[string]NrfInfoServedNwdafInfoValue `json:"servedNwdafInfo,omitempty"`

	// A map (list of key-value pairs) where NF Instance Id serves as key
	ServedNwdafInfoList map[string]map[string]NwdafInfo `json:"servedNwdafInfoList,omitempty"`

	// A map (list of key-value pairs) where nfInstanceId serves as key
	ServedPcscfInfoList map[string]map[string]NrfInfoServedPcscfInfoListValueValue `json:"servedPcscfInfoList,omitempty"`

	// A map (list of key-value pairs) where nfInstanceId serves as key
	ServedGmlcInfo map[string]NrfInfoServedGmlcInfoValue `json:"servedGmlcInfo,omitempty"`

	// A map (list of key-value pairs) where nfInstanceId serves as key
	ServedLmfInfo map[string]NrfInfoServedLmfInfoValue `json:"servedLmfInfo,omitempty"`

	// A map (list of key-value pairs) where nfInstanceId serves as key
	ServedNfInfo map[string]NfInfo `json:"servedNfInfo,omitempty"`

	// A map (list of key-value pairs) where nfInstanceId serves as key
	ServedHssInfoList map[string]map[string]NrfInfoServedHssInfoListValueValue `json:"servedHssInfoList,omitempty"`

	// A map (list of key-value pairs) where nfInstanceId serves as key
	ServedUdsfInfo map[string]NrfInfoServedUdsfInfoValue `json:"servedUdsfInfo,omitempty"`

	// A map (list of key-value pairs) where nfInstanceId serves as key
	ServedUdsfInfoList map[string]map[string]NrfInfoServedUdsfInfoValue `json:"servedUdsfInfoList,omitempty"`

	// A map (list of key-value pairs) where nfInstanceId serves as key
	ServedScpInfoList map[string]NrfInfoServedScpInfoListValue `json:"servedScpInfoList,omitempty"`

	// A map (list of key-value pairs) where nfInstanceId serves as key
	ServedSeppInfoList map[string]NrfInfoServedSeppInfoListValue `json:"servedSeppInfoList,omitempty"`

	// A map (list of key-value pairs) where NF Instance Id serves as key
	ServedAanfInfoList map[string]map[string]NrfInfoServedAanfInfoListValueValue `json:"servedAanfInfoList,omitempty"`

	Served5gDdnmfInfo map[string]Model5GDdnmfInfo `json:"served5gDdnmfInfo,omitempty"`

	// A map (list of key-value pairs) where NF Instance Id serves as key
	ServedMfafInfoList map[string]MfafInfo `json:"servedMfafInfoList,omitempty"`

	// A map (list of key-value pairs) where NF Instance Id serves as key
	ServedEasdfInfoList map[string]map[string]EasdfInfo `json:"servedEasdfInfoList,omitempty"`

	// A map (list of key-value pairs) where NF Instance Id serves as key
	ServedDccfInfoList map[string]DccfInfo `json:"servedDccfInfoList,omitempty"`

	// A map (list of key-value pairs) where nfInstanceId serves as key
	ServedMbSmfInfoList map[string]map[string]NrfInfoServedMbSmfInfoListValueValue `json:"servedMbSmfInfoList,omitempty"`

	// A map (list of key-value pairs) where NF Instance Id serves as key
	ServedTsctsfInfoList map[string]map[string]TsctsfInfo `json:"servedTsctsfInfoList,omitempty"`

	// A map (list of key-value pairs) where NF Instance Id serves as key
	ServedMbUpfInfoList map[string]map[string]MbUpfInfo `json:"servedMbUpfInfoList,omitempty"`

	// A map (list of key-value pairs) where NF Instance Id serves as key
	ServedTrustAfInfo map[string]TrustAfInfo `json:"servedTrustAfInfo,omitempty"`

	// A map (list of key-value pairs) where NF Instance Id serves as key
	ServedNssaafInfo map[string]NssaafInfo `json:"servedNssaafInfo,omitempty"`
}
