/*
 * 3gpp-mbs-session
 *
 * API for MBS Session Management.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package MBSSession

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// DeleteIndMBSSession - Request the Deletion of an existing Individual MBS Session resource.
func DeleteIndMBSSession(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// ModifyIndMBSSession - Request the modification of an existing Individual MBS Session resource.
func ModifyIndMBSSession(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}