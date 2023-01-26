/*
 * Nsmf_PDUSession
 *
 * SMF PDU Session Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.3.0-alpha.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nsmf_PDUSession

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Route is the information for every URI.
type Route struct {
	// Name is the name of this Route.
	Name        string
	// Method is the string for the HTTP method. ex) GET, POST etc..
	Method      string
	// Pattern is the pattern of the URI.
	Pattern     string
	// HandlerFunc is the handler function of this route.
	HandlerFunc gin.HandlerFunc
}

// Routes is the list of the generated Route.
type Routes []Route

// NewRouter returns a new router.
func NewRouter() *gin.Engine {
	router := gin.Default()
	for _, route := range routes {
		switch route.Method {
		case http.MethodGet:
			router.GET(route.Pattern, route.HandlerFunc)
		case http.MethodPost:
			router.POST(route.Pattern, route.HandlerFunc)
		case http.MethodPut:
			router.PUT(route.Pattern, route.HandlerFunc)
		case http.MethodPatch:
			router.PATCH(route.Pattern, route.HandlerFunc)
		case http.MethodDelete:
			router.DELETE(route.Pattern, route.HandlerFunc)
		}
	}

	return router
}

// Index is the index handler.
func Index(c *gin.Context) {
	c.String(http.StatusOK, "Hello World!")
}

var routes = Routes{
	{
		"Index",
		http.MethodGet,
		"/nsmf-pdusession/v1/",
		Index,
	},

	{
		"ReleasePduSession",
		http.MethodPost,
		"/nsmf-pdusession/v1/pdu-sessions/:pduSessionRef/release",
		ReleasePduSession,
	},

	{
		"RetrievePduSession",
		http.MethodPost,
		"/nsmf-pdusession/v1/pdu-sessions/:pduSessionRef/retrieve",
		RetrievePduSession,
	},

	{
		"TransferMoData",
		http.MethodPost,
		"/nsmf-pdusession/v1/pdu-sessions/:pduSessionRef/transfer-mo-data",
		TransferMoData,
	},

	{
		"UpdatePduSession",
		http.MethodPost,
		"/nsmf-pdusession/v1/pdu-sessions/:pduSessionRef/modify",
		UpdatePduSession,
	},

	{
		"ReleaseSmContext",
		http.MethodPost,
		"/nsmf-pdusession/v1/sm-contexts/:smContextRef/release",
		ReleaseSmContext,
	},

	{
		"RetrieveSmContext",
		http.MethodPost,
		"/nsmf-pdusession/v1/sm-contexts/:smContextRef/retrieve",
		RetrieveSmContext,
	},

	{
		"SendMoData",
		http.MethodPost,
		"/nsmf-pdusession/v1/sm-contexts/:smContextRef/send-mo-data",
		SendMoData,
	},

	{
		"UpdateSmContext",
		http.MethodPost,
		"/nsmf-pdusession/v1/sm-contexts/:smContextRef/modify",
		UpdateSmContext,
	},

	{
		"PostPduSessions",
		http.MethodPost,
		"/nsmf-pdusession/v1/pdu-sessions",
		PostPduSessions,
	},

	{
		"PostSmContexts",
		http.MethodPost,
		"/nsmf-pdusession/v1/sm-contexts",
		PostSmContexts,
	},
}