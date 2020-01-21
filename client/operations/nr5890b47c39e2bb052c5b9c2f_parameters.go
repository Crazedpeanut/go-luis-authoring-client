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

	models "github.com/crazedpeanut/luis/models"
)

// NewNr5890b47c39e2bb052c5b9c2fParams creates a new Nr5890b47c39e2bb052c5b9c2fParams object
// with the default values initialized.
func NewNr5890b47c39e2bb052c5b9c2fParams() *Nr5890b47c39e2bb052c5b9c2fParams {
	var ()
	return &Nr5890b47c39e2bb052c5b9c2fParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewNr5890b47c39e2bb052c5b9c2fParamsWithTimeout creates a new Nr5890b47c39e2bb052c5b9c2fParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewNr5890b47c39e2bb052c5b9c2fParamsWithTimeout(timeout time.Duration) *Nr5890b47c39e2bb052c5b9c2fParams {
	var ()
	return &Nr5890b47c39e2bb052c5b9c2fParams{

		timeout: timeout,
	}
}

// NewNr5890b47c39e2bb052c5b9c2fParamsWithContext creates a new Nr5890b47c39e2bb052c5b9c2fParams object
// with the default values initialized, and the ability to set a context for a request
func NewNr5890b47c39e2bb052c5b9c2fParamsWithContext(ctx context.Context) *Nr5890b47c39e2bb052c5b9c2fParams {
	var ()
	return &Nr5890b47c39e2bb052c5b9c2fParams{

		Context: ctx,
	}
}

// NewNr5890b47c39e2bb052c5b9c2fParamsWithHTTPClient creates a new Nr5890b47c39e2bb052c5b9c2fParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewNr5890b47c39e2bb052c5b9c2fParamsWithHTTPClient(client *http.Client) *Nr5890b47c39e2bb052c5b9c2fParams {
	var ()
	return &Nr5890b47c39e2bb052c5b9c2fParams{
		HTTPClient: client,
	}
}

/*Nr5890b47c39e2bb052c5b9c2fParams contains all the parameters to send to the API endpoint
for the 5890b47c39e2bb052c5b9c2f operation typically these are written to a http.Request
*/
type Nr5890b47c39e2bb052c5b9c2fParams struct {

	/*ApplicationCreateObject
	  Application name length has to be less than 50 characters.

	The culture and tokenizer version of an app cannot be changed after the app is created.

	The default version is "0.1"

	*/
	ApplicationCreateObject *models.ApplicationCreateObject

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the 5890b47c39e2bb052c5b9c2f params
func (o *Nr5890b47c39e2bb052c5b9c2fParams) WithTimeout(timeout time.Duration) *Nr5890b47c39e2bb052c5b9c2fParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the 5890b47c39e2bb052c5b9c2f params
func (o *Nr5890b47c39e2bb052c5b9c2fParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the 5890b47c39e2bb052c5b9c2f params
func (o *Nr5890b47c39e2bb052c5b9c2fParams) WithContext(ctx context.Context) *Nr5890b47c39e2bb052c5b9c2fParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the 5890b47c39e2bb052c5b9c2f params
func (o *Nr5890b47c39e2bb052c5b9c2fParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the 5890b47c39e2bb052c5b9c2f params
func (o *Nr5890b47c39e2bb052c5b9c2fParams) WithHTTPClient(client *http.Client) *Nr5890b47c39e2bb052c5b9c2fParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the 5890b47c39e2bb052c5b9c2f params
func (o *Nr5890b47c39e2bb052c5b9c2fParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithApplicationCreateObject adds the applicationCreateObject to the 5890b47c39e2bb052c5b9c2f params
func (o *Nr5890b47c39e2bb052c5b9c2fParams) WithApplicationCreateObject(applicationCreateObject *models.ApplicationCreateObject) *Nr5890b47c39e2bb052c5b9c2fParams {
	o.SetApplicationCreateObject(applicationCreateObject)
	return o
}

// SetApplicationCreateObject adds the applicationCreateObject to the 5890b47c39e2bb052c5b9c2f params
func (o *Nr5890b47c39e2bb052c5b9c2fParams) SetApplicationCreateObject(applicationCreateObject *models.ApplicationCreateObject) {
	o.ApplicationCreateObject = applicationCreateObject
}

// WriteToRequest writes these params to a swagger request
func (o *Nr5890b47c39e2bb052c5b9c2fParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ApplicationCreateObject != nil {
		if err := r.SetBodyParam(o.ApplicationCreateObject); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
