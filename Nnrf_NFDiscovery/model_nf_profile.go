/*
 * NRF NFDiscovery Service
 *
 * NRF NFDiscovery Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.3.0-alpha.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nnrf_NFDiscovery

import (
	"time"
)

// NfProfile - Information of an NF Instance discovered by the NRF
type NfProfile struct {

	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.  
	NfInstanceId string `json:"nfInstanceId"`

	NfInstanceName string `json:"nfInstanceName,omitempty"`

	NfType NfType `json:"nfType"`

	NfStatus NfStatus `json:"nfStatus"`

	CollocatedNfInstances []CollocatedNfInstance `json:"collocatedNfInstances,omitempty"`

	PlmnList []PlmnId `json:"plmnList,omitempty"`

	SNssais []ExtSnssai `json:"sNssais,omitempty"`

	PerPlmnSnssaiList []PlmnSnssai `json:"perPlmnSnssaiList,omitempty"`

	NsiList []string `json:"nsiList,omitempty"`

	// Fully Qualified Domain Name
	Fqdn string `json:"fqdn,omitempty"`

	// Fully Qualified Domain Name
	InterPlmnFqdn string `json:"interPlmnFqdn,omitempty"`

	Ipv4Addresses []string `json:"ipv4Addresses,omitempty"`

	Ipv6Addresses []Ipv6Addr `json:"ipv6Addresses,omitempty"`

	Capacity int32 `json:"capacity,omitempty"`

	Load int32 `json:"load,omitempty"`

	// string with format 'date-time' as defined in OpenAPI.
	LoadTimeStamp time.Time `json:"loadTimeStamp,omitempty"`

	Locality string `json:"locality,omitempty"`

	// A map (list of key-value pairs) where a (unique) valid JSON string serves  as key representing a type of locality 
	ExtLocality map[string]string `json:"extLocality,omitempty"`

	Priority int32 `json:"priority,omitempty"`

	UdrInfo UdrInfo `json:"udrInfo,omitempty"`

	// A map (list of key-value pairs) where a (unique) valid JSON string serves as key of UdrInfo 
	UdrInfoList map[string]UdrInfo `json:"udrInfoList,omitempty"`

	UdmInfo UdmInfo `json:"udmInfo,omitempty"`

	// A map (list of key-value pairs) where a (unique) valid JSON string serves as key of UdmInfo 
	UdmInfoList map[string]UdmInfo `json:"udmInfoList,omitempty"`

	AusfInfo AusfInfo `json:"ausfInfo,omitempty"`

	// A map (list of key-value pairs) where a (unique) valid JSON string serves as key of AusfInfo 
	AusfInfoList map[string]AusfInfo `json:"ausfInfoList,omitempty"`

	AmfInfo AmfInfo `json:"amfInfo,omitempty"`

	// A map (list of key-value pairs) where a (unique) valid JSON string serves as key of AmfInfo 
	AmfInfoList map[string]AmfInfo `json:"amfInfoList,omitempty"`

	SmfInfo SmfInfo `json:"smfInfo,omitempty"`

	// A map (list of key-value pairs) where a (unique) valid JSON string serves as key of SmfInfo 
	SmfInfoList map[string]SmfInfo `json:"smfInfoList,omitempty"`

	UpfInfo UpfInfo `json:"upfInfo,omitempty"`

	// A map (list of key-value pairs) where a (unique) valid JSON string serves as key of UpfInfo 
	UpfInfoList map[string]UpfInfo `json:"upfInfoList,omitempty"`

	PcfInfo PcfInfo `json:"pcfInfo,omitempty"`

	// A map (list of key-value pairs) where a (unique) valid JSON string serves as key of PcfInfo 
	PcfInfoList map[string]PcfInfo `json:"pcfInfoList,omitempty"`

	BsfInfo BsfInfo `json:"bsfInfo,omitempty"`

	// A map (list of key-value pairs) where a (unique) valid JSON string serves as key of BsfInfo 
	BsfInfoList map[string]BsfInfo `json:"bsfInfoList,omitempty"`

	ChfInfo ChfInfo `json:"chfInfo,omitempty"`

	// A map (list of key-value pairs) where a (unique) valid JSON string serves as key of ChfInfo 
	ChfInfoList map[string]ChfInfo `json:"chfInfoList,omitempty"`

	UdsfInfo UdsfInfo `json:"udsfInfo,omitempty"`

	// A map (list of key-value pairs) where a (unique) valid JSON string serves as key of UdsfInfo 
	UdsfInfoList map[string]UdsfInfo `json:"udsfInfoList,omitempty"`

	NwdafInfo NwdafInfo `json:"nwdafInfo,omitempty"`

	// A map (list of key-value pairs) where a (unique) valid JSON string serves as key of NwdafInfo 
	NwdafInfoList map[string]NwdafInfo `json:"nwdafInfoList,omitempty"`

	NefInfo NefInfo `json:"nefInfo,omitempty"`

	// A map (list of key-value pairs) where a (unique) valid JSON string serves as key of PcscfInfo 
	PcscfInfoList map[string]PcscfInfo `json:"pcscfInfoList,omitempty"`

	// A map (list of key-value pairs) where a (unique) valid JSON string serves as key of HssInfo 
	HssInfoList map[string]HssInfo `json:"hssInfoList,omitempty"`

	CustomInfo map[string]interface{} `json:"customInfo,omitempty"`

	// string with format 'date-time' as defined in OpenAPI.
	RecoveryTime time.Time `json:"recoveryTime,omitempty"`

	NfServicePersistence bool `json:"nfServicePersistence,omitempty"`

	// Deprecated
	NfServices []NfService `json:"nfServices,omitempty"`

	// A map (list of key-value pairs) where serviceInstanceId serves as key of NFService 
	NfServiceList map[string]NfService `json:"nfServiceList,omitempty"`

	DefaultNotificationSubscriptions []DefaultNotificationSubscription `json:"defaultNotificationSubscriptions,omitempty"`

	LmfInfo LmfInfo `json:"lmfInfo,omitempty"`

	GmlcInfo GmlcInfo `json:"gmlcInfo,omitempty"`

	SnpnList []PlmnIdNid `json:"snpnList,omitempty"`

	NfSetIdList []string `json:"nfSetIdList,omitempty"`

	ServingScope []string `json:"servingScope,omitempty"`

	LcHSupportInd bool `json:"lcHSupportInd,omitempty"`

	OlcHSupportInd bool `json:"olcHSupportInd,omitempty"`

	// A map (list of key-value pairs) where NfSetId serves as key of DateTime
	NfSetRecoveryTimeList map[string]time.Time `json:"nfSetRecoveryTimeList,omitempty"`

	// A map (list of key-value pairs) where NfServiceSetId serves as key of DateTime 
	ServiceSetRecoveryTimeList map[string]time.Time `json:"serviceSetRecoveryTimeList,omitempty"`

	ScpDomains []string `json:"scpDomains,omitempty"`

	ScpInfo ScpInfo `json:"scpInfo,omitempty"`

	SeppInfo SeppInfo `json:"seppInfo,omitempty"`

	// Vendor ID of the NF Service instance (Private Enterprise Number assigned by IANA)
	VendorId string `json:"vendorId,omitempty"`

	// The key of the map is the IANA-assigned SMI Network Management Private Enterprise Codes 
	SupportedVendorSpecificFeatures map[string][]VendorSpecificFeature `json:"supportedVendorSpecificFeatures,omitempty"`

	// A map (list of key-value pairs) where a (unique) valid JSON string serves as key of AanfInfo 
	AanfInfoList map[string]AanfInfo `json:"aanfInfoList,omitempty"`

	MfafInfo MfafInfo `json:"mfafInfo,omitempty"`

	// A map(list of key-value pairs) where a (unique) valid JSON string serves as key of EasdfInfo 
	EasdfInfoList map[string]EasdfInfo `json:"easdfInfoList,omitempty"`

	DccfInfo DccfInfo `json:"dccfInfo,omitempty"`

	// A map (list of key-value pairs) where a (unique) valid JSON string serves as key of NsacfInfo 
	NsacfInfoList map[string]NsacfInfo `json:"nsacfInfoList,omitempty"`

	// A map (list of key-value pairs) where a (unique) valid JSON string serves as key of MbSmfInfo 
	MbSmfInfoList map[string]MbSmfInfo `json:"mbSmfInfoList,omitempty"`

	// A map (list of key-value pairs) where a (unique) valid JSON string serves as key of TsctsfInfo 
	TsctsfInfoList map[string]TsctsfInfo `json:"tsctsfInfoList,omitempty"`

	// A map (list of key-value pairs) where a (unique) valid JSON string serves as key of MbUpfInfo 
	MbUpfInfoList map[string]MbUpfInfo `json:"mbUpfInfoList,omitempty"`

	TrustAfInfo TrustAfInfo `json:"trustAfInfo,omitempty"`

	NssaafInfo NssaafInfo `json:"nssaafInfo,omitempty"`

	HniList []string `json:"hniList,omitempty"`

	IwmscInfo IwmscInfo `json:"iwmscInfo,omitempty"`

	MnpfInfo MnpfInfo `json:"mnpfInfo,omitempty"`
}
