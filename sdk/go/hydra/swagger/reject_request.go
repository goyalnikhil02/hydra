/*
 * ORY Hydra
 *
 * Welcome to the ORY Hydra HTTP API documentation. You will find documentation for all HTTP APIs here.
 *
 * OpenAPI spec version: latest
 *
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package swagger

type RejectRequest struct {
	Error_ string `json:"error,omitempty"`

	ErrorDebug string `json:"error_debug,omitempty"`

	ErrorDescription string `json:"error_description,omitempty"`

	ErrorHint string `json:"error_hint,omitempty"`

	StatusCode int64 `json:"status_code,omitempty"`
}
