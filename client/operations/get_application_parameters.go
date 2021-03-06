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

// NewGetApplicationParams creates a new GetApplicationParams object
// with the default values initialized.
func NewGetApplicationParams() *GetApplicationParams {
	var ()
	return &GetApplicationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetApplicationParamsWithTimeout creates a new GetApplicationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetApplicationParamsWithTimeout(timeout time.Duration) *GetApplicationParams {
	var ()
	return &GetApplicationParams{

		timeout: timeout,
	}
}

// NewGetApplicationParamsWithContext creates a new GetApplicationParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetApplicationParamsWithContext(ctx context.Context) *GetApplicationParams {
	var ()
	return &GetApplicationParams{

		Context: ctx,
	}
}

// NewGetApplicationParamsWithHTTPClient creates a new GetApplicationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetApplicationParamsWithHTTPClient(client *http.Client) *GetApplicationParams {
	var ()
	return &GetApplicationParams{
		HTTPClient: client,
	}
}

/*GetApplicationParams contains all the parameters to send to the API endpoint
for the get application operation typically these are written to a http.Request
*/
type GetApplicationParams struct {

	/*AppID
	  Format - guid. The application ID.

	*/
	AppID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get application params
func (o *GetApplicationParams) WithTimeout(timeout time.Duration) *GetApplicationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get application params
func (o *GetApplicationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get application params
func (o *GetApplicationParams) WithContext(ctx context.Context) *GetApplicationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get application params
func (o *GetApplicationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get application params
func (o *GetApplicationParams) WithHTTPClient(client *http.Client) *GetApplicationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get application params
func (o *GetApplicationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppID adds the appID to the get application params
func (o *GetApplicationParams) WithAppID(appID string) *GetApplicationParams {
	o.SetAppID(appID)
	return o
}

// SetAppID adds the appId to the get application params
func (o *GetApplicationParams) SetAppID(appID string) {
	o.AppID = appID
}

// WriteToRequest writes these params to a swagger request
func (o *GetApplicationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param appId
	if err := r.SetPathParam("appId", o.AppID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
