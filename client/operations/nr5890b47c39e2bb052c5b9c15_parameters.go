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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewNr5890b47c39e2bb052c5b9c15Params creates a new Nr5890b47c39e2bb052c5b9c15Params object
// with the default values initialized.
func NewNr5890b47c39e2bb052c5b9c15Params() *Nr5890b47c39e2bb052c5b9c15Params {
	var (
		skipDefault = int64(0)
		takeDefault = int64(100)
	)
	return &Nr5890b47c39e2bb052c5b9c15Params{
		Skip: &skipDefault,
		Take: &takeDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewNr5890b47c39e2bb052c5b9c15ParamsWithTimeout creates a new Nr5890b47c39e2bb052c5b9c15Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewNr5890b47c39e2bb052c5b9c15ParamsWithTimeout(timeout time.Duration) *Nr5890b47c39e2bb052c5b9c15Params {
	var (
		skipDefault = int64(0)
		takeDefault = int64(100)
	)
	return &Nr5890b47c39e2bb052c5b9c15Params{
		Skip: &skipDefault,
		Take: &takeDefault,

		timeout: timeout,
	}
}

// NewNr5890b47c39e2bb052c5b9c15ParamsWithContext creates a new Nr5890b47c39e2bb052c5b9c15Params object
// with the default values initialized, and the ability to set a context for a request
func NewNr5890b47c39e2bb052c5b9c15ParamsWithContext(ctx context.Context) *Nr5890b47c39e2bb052c5b9c15Params {
	var (
		skipDefault = int64(0)
		takeDefault = int64(100)
	)
	return &Nr5890b47c39e2bb052c5b9c15Params{
		Skip: &skipDefault,
		Take: &takeDefault,

		Context: ctx,
	}
}

// NewNr5890b47c39e2bb052c5b9c15ParamsWithHTTPClient creates a new Nr5890b47c39e2bb052c5b9c15Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewNr5890b47c39e2bb052c5b9c15ParamsWithHTTPClient(client *http.Client) *Nr5890b47c39e2bb052c5b9c15Params {
	var (
		skipDefault = int64(0)
		takeDefault = int64(100)
	)
	return &Nr5890b47c39e2bb052c5b9c15Params{
		Skip:       &skipDefault,
		Take:       &takeDefault,
		HTTPClient: client,
	}
}

/*Nr5890b47c39e2bb052c5b9c15Params contains all the parameters to send to the API endpoint
for the 5890b47c39e2bb052c5b9c15 operation typically these are written to a http.Request
*/
type Nr5890b47c39e2bb052c5b9c15Params struct {

	/*AppID
	  Format - guid. The application ID.

	*/
	AppID string
	/*Skip
	  The number of entries to skip. Default value is 0.

	*/
	Skip *int64
	/*Take
	  The number of entries to return. Maximum page size is 500. Default is 100.

	*/
	Take *int64
	/*VersionID
	  The task version ID.

	*/
	VersionID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the 5890b47c39e2bb052c5b9c15 params
func (o *Nr5890b47c39e2bb052c5b9c15Params) WithTimeout(timeout time.Duration) *Nr5890b47c39e2bb052c5b9c15Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the 5890b47c39e2bb052c5b9c15 params
func (o *Nr5890b47c39e2bb052c5b9c15Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the 5890b47c39e2bb052c5b9c15 params
func (o *Nr5890b47c39e2bb052c5b9c15Params) WithContext(ctx context.Context) *Nr5890b47c39e2bb052c5b9c15Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the 5890b47c39e2bb052c5b9c15 params
func (o *Nr5890b47c39e2bb052c5b9c15Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the 5890b47c39e2bb052c5b9c15 params
func (o *Nr5890b47c39e2bb052c5b9c15Params) WithHTTPClient(client *http.Client) *Nr5890b47c39e2bb052c5b9c15Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the 5890b47c39e2bb052c5b9c15 params
func (o *Nr5890b47c39e2bb052c5b9c15Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppID adds the appID to the 5890b47c39e2bb052c5b9c15 params
func (o *Nr5890b47c39e2bb052c5b9c15Params) WithAppID(appID string) *Nr5890b47c39e2bb052c5b9c15Params {
	o.SetAppID(appID)
	return o
}

// SetAppID adds the appId to the 5890b47c39e2bb052c5b9c15 params
func (o *Nr5890b47c39e2bb052c5b9c15Params) SetAppID(appID string) {
	o.AppID = appID
}

// WithSkip adds the skip to the 5890b47c39e2bb052c5b9c15 params
func (o *Nr5890b47c39e2bb052c5b9c15Params) WithSkip(skip *int64) *Nr5890b47c39e2bb052c5b9c15Params {
	o.SetSkip(skip)
	return o
}

// SetSkip adds the skip to the 5890b47c39e2bb052c5b9c15 params
func (o *Nr5890b47c39e2bb052c5b9c15Params) SetSkip(skip *int64) {
	o.Skip = skip
}

// WithTake adds the take to the 5890b47c39e2bb052c5b9c15 params
func (o *Nr5890b47c39e2bb052c5b9c15Params) WithTake(take *int64) *Nr5890b47c39e2bb052c5b9c15Params {
	o.SetTake(take)
	return o
}

// SetTake adds the take to the 5890b47c39e2bb052c5b9c15 params
func (o *Nr5890b47c39e2bb052c5b9c15Params) SetTake(take *int64) {
	o.Take = take
}

// WithVersionID adds the versionID to the 5890b47c39e2bb052c5b9c15 params
func (o *Nr5890b47c39e2bb052c5b9c15Params) WithVersionID(versionID string) *Nr5890b47c39e2bb052c5b9c15Params {
	o.SetVersionID(versionID)
	return o
}

// SetVersionID adds the versionId to the 5890b47c39e2bb052c5b9c15 params
func (o *Nr5890b47c39e2bb052c5b9c15Params) SetVersionID(versionID string) {
	o.VersionID = versionID
}

// WriteToRequest writes these params to a swagger request
func (o *Nr5890b47c39e2bb052c5b9c15Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param appId
	if err := r.SetPathParam("appId", o.AppID); err != nil {
		return err
	}

	if o.Skip != nil {

		// query param skip
		var qrSkip int64
		if o.Skip != nil {
			qrSkip = *o.Skip
		}
		qSkip := swag.FormatInt64(qrSkip)
		if qSkip != "" {
			if err := r.SetQueryParam("skip", qSkip); err != nil {
				return err
			}
		}

	}

	if o.Take != nil {

		// query param take
		var qrTake int64
		if o.Take != nil {
			qrTake = *o.Take
		}
		qTake := swag.FormatInt64(qrTake)
		if qTake != "" {
			if err := r.SetQueryParam("take", qTake); err != nil {
				return err
			}
		}

	}

	// path param versionId
	if err := r.SetPathParam("versionId", o.VersionID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
