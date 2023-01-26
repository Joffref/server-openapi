/*
 * Npcf_PolicyAuthorization Service API
 *
 * PCF Policy Authorization Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Npcf_PolicyAuthorization

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// PcscfRestoration - Indicates P-CSCF restoration and does not create an Individual Application Session Context
func PcscfRestoration(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}