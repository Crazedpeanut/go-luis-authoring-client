// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetUserExternalKeysParams creates a new GetUserExternalKeysParams object
// with the default values initialized.
func NewGetUserExternalKeysParams() *GetUserExternalKeysParams {

	return &GetUserExternalKeysParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetUserExternalKeysParamsWithTimeout creates a new GetUserExternalKeysParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetUserExternalKeysParamsWithTimeout(timeout time.Duration) *GetUserExternalKeysParams {

	return &GetUserExternalKeysParams{

		timeout: timeout,
	}
}

// NewGetUserExternalKeysParamsWithContext creates a new GetUserExternalKeysParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetUserExternalKeysParamsWithContext(ctx context.Context) *GetUserExternalKeysParams {

	return &GetUserExternalKeysParams{

		Context: ctx,
	}
}

// NewGetUserExternalKeysParamsWithHTTPClient creates a new GetUserExternalKeysParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetUserExternalKeysParamsWithHTTPClient(client *http.Client) *GetUserExternalKeysParams {

	return &GetUserExternalKeysParams{
		HTTPClient: client,
	}
}

/*GetUserExternalKeysParams contains all the parameters to send to the API endpoint
for the get user external keys operation typically these are written to a http.Request
*/
type GetUserExternalKeysParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get user external keys params
func (o *GetUserExternalKeysParams) WithTimeout(timeout time.Duration) *GetUserExternalKeysParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get user external keys params
func (o *GetUserExternalKeysParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get user external keys params
func (o *GetUserExternalKeysParams) WithContext(ctx context.Context) *GetUserExternalKeysParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get user external keys params
func (o *GetUserExternalKeysParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get user external keys params
func (o *GetUserExternalKeysParams) WithHTTPClient(client *http.Client) *GetUserExternalKeysParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get user external keys params
func (o *GetUserExternalKeysParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetUserExternalKeysParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
