 //+build debug 

/*
 * Nsmf_PDUSession
 *
 * SMF PDU Session Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nsmf_PDUSession

import (
	"crypto/tls"
	"golang.org/x/net/http2"
	"net/http"
	"free5gc/lib/http2_util"
)

// APIClient manages communication with the Nsmf_PDUSession API v1.0.0
// In most cases there should be only one, shared, APIClient.
type APIClient struct {
	cfg    *Configuration
	common service // Reuse a single struct instead of allocating one for each service on the heap.

	// API Services
	IndividualPDUSessionHSMFApi *IndividualPDUSessionHSMFApiService
	IndividualSMContextApi      *IndividualSMContextApiService
	PDUSessionsCollectionApi    *PDUSessionsCollectionApiService
	SMContextsCollectionApi     *SMContextsCollectionApiService
}

type service struct {
	client *APIClient
}

// NewAPIClient creates a new API client. Requires a userAgent string describing your application.
// optionally a custom http.Client to allow for advanced features such as caching.
func NewAPIClient(cfg *Configuration) *APIClient {
	if cfg.httpClient == nil {
		cfg.httpClient = http.DefaultClient
		cfg.httpClient.Transport = &http2.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
				Rand:         http2_util.ZeroSource{},
			},
		}
	}

	c := &APIClient{}
	c.cfg = cfg
	c.common.client = c

	// API Services
	c.IndividualPDUSessionHSMFApi = (*IndividualPDUSessionHSMFApiService)(&c.common)
	c.IndividualSMContextApi = (*IndividualSMContextApiService)(&c.common)
	c.PDUSessionsCollectionApi = (*PDUSessionsCollectionApiService)(&c.common)
	c.SMContextsCollectionApi = (*SMContextsCollectionApiService)(&c.common)

	return c
}
