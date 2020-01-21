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

// NewNr59087d2d5aca2f03943a889aParams creates a new Nr59087d2d5aca2f03943a889aParams object
// with the default values initialized.
func NewNr59087d2d5aca2f03943a889aParams() *Nr59087d2d5aca2f03943a889aParams {
	var ()
	return &Nr59087d2d5aca2f03943a889aParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewNr59087d2d5aca2f03943a889aParamsWithTimeout creates a new Nr59087d2d5aca2f03943a889aParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewNr59087d2d5aca2f03943a889aParamsWithTimeout(timeout time.Duration) *Nr59087d2d5aca2f03943a889aParams {
	var ()
	return &Nr59087d2d5aca2f03943a889aParams{

		timeout: timeout,
	}
}

// NewNr59087d2d5aca2f03943a889aParamsWithContext creates a new Nr59087d2d5aca2f03943a889aParams object
// with the default values initialized, and the ability to set a context for a request
func NewNr59087d2d5aca2f03943a889aParamsWithContext(ctx context.Context) *Nr59087d2d5aca2f03943a889aParams {
	var ()
	return &Nr59087d2d5aca2f03943a889aParams{

		Context: ctx,
	}
}

// NewNr59087d2d5aca2f03943a889aParamsWithHTTPClient creates a new Nr59087d2d5aca2f03943a889aParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewNr59087d2d5aca2f03943a889aParamsWithHTTPClient(client *http.Client) *Nr59087d2d5aca2f03943a889aParams {
	var ()
	return &Nr59087d2d5aca2f03943a889aParams{
		HTTPClient: client,
	}
}

/*Nr59087d2d5aca2f03943a889aParams contains all the parameters to send to the API endpoint
for the 59087d2d5aca2f03943a889a operation typically these are written to a http.Request
*/
type Nr59087d2d5aca2f03943a889aParams struct {

	/*AppID
	  Format - guid. The application ID.

	*/
	AppID string
	/*ClEntityID
	  Format - guid. The closedlist ID.

	*/
	ClEntityID string
	/*SubListID
	  Format - long. The sublist ID

	*/
	SubListID string
	/*VersionID
	  The task version ID.

	*/
	VersionID string
	/*WordListBaseUpdateObject
	  A sublist update object containing the new canonical form and the list of words.

	*/
	WordListBaseUpdateObject *models.WordListBaseUpdateObject

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the 59087d2d5aca2f03943a889a params
func (o *Nr59087d2d5aca2f03943a889aParams) WithTimeout(timeout time.Duration) *Nr59087d2d5aca2f03943a889aParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the 59087d2d5aca2f03943a889a params
func (o *Nr59087d2d5aca2f03943a889aParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the 59087d2d5aca2f03943a889a params
func (o *Nr59087d2d5aca2f03943a889aParams) WithContext(ctx context.Context) *Nr59087d2d5aca2f03943a889aParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the 59087d2d5aca2f03943a889a params
func (o *Nr59087d2d5aca2f03943a889aParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the 59087d2d5aca2f03943a889a params
func (o *Nr59087d2d5aca2f03943a889aParams) WithHTTPClient(client *http.Client) *Nr59087d2d5aca2f03943a889aParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the 59087d2d5aca2f03943a889a params
func (o *Nr59087d2d5aca2f03943a889aParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppID adds the appID to the 59087d2d5aca2f03943a889a params
func (o *Nr59087d2d5aca2f03943a889aParams) WithAppID(appID string) *Nr59087d2d5aca2f03943a889aParams {
	o.SetAppID(appID)
	return o
}

// SetAppID adds the appId to the 59087d2d5aca2f03943a889a params
func (o *Nr59087d2d5aca2f03943a889aParams) SetAppID(appID string) {
	o.AppID = appID
}

// WithClEntityID adds the clEntityID to the 59087d2d5aca2f03943a889a params
func (o *Nr59087d2d5aca2f03943a889aParams) WithClEntityID(clEntityID string) *Nr59087d2d5aca2f03943a889aParams {
	o.SetClEntityID(clEntityID)
	return o
}

// SetClEntityID adds the clEntityId to the 59087d2d5aca2f03943a889a params
func (o *Nr59087d2d5aca2f03943a889aParams) SetClEntityID(clEntityID string) {
	o.ClEntityID = clEntityID
}

// WithSubListID adds the subListID to the 59087d2d5aca2f03943a889a params
func (o *Nr59087d2d5aca2f03943a889aParams) WithSubListID(subListID string) *Nr59087d2d5aca2f03943a889aParams {
	o.SetSubListID(subListID)
	return o
}

// SetSubListID adds the subListId to the 59087d2d5aca2f03943a889a params
func (o *Nr59087d2d5aca2f03943a889aParams) SetSubListID(subListID string) {
	o.SubListID = subListID
}

// WithVersionID adds the versionID to the 59087d2d5aca2f03943a889a params
func (o *Nr59087d2d5aca2f03943a889aParams) WithVersionID(versionID string) *Nr59087d2d5aca2f03943a889aParams {
	o.SetVersionID(versionID)
	return o
}

// SetVersionID adds the versionId to the 59087d2d5aca2f03943a889a params
func (o *Nr59087d2d5aca2f03943a889aParams) SetVersionID(versionID string) {
	o.VersionID = versionID
}

// WithWordListBaseUpdateObject adds the wordListBaseUpdateObject to the 59087d2d5aca2f03943a889a params
func (o *Nr59087d2d5aca2f03943a889aParams) WithWordListBaseUpdateObject(wordListBaseUpdateObject *models.WordListBaseUpdateObject) *Nr59087d2d5aca2f03943a889aParams {
	o.SetWordListBaseUpdateObject(wordListBaseUpdateObject)
	return o
}

// SetWordListBaseUpdateObject adds the wordListBaseUpdateObject to the 59087d2d5aca2f03943a889a params
func (o *Nr59087d2d5aca2f03943a889aParams) SetWordListBaseUpdateObject(wordListBaseUpdateObject *models.WordListBaseUpdateObject) {
	o.WordListBaseUpdateObject = wordListBaseUpdateObject
}

// WriteToRequest writes these params to a swagger request
func (o *Nr59087d2d5aca2f03943a889aParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param appId
	if err := r.SetPathParam("appId", o.AppID); err != nil {
		return err
	}

	// path param clEntityId
	if err := r.SetPathParam("clEntityId", o.ClEntityID); err != nil {
		return err
	}

	// path param subListId
	if err := r.SetPathParam("subListId", o.SubListID); err != nil {
		return err
	}

	// path param versionId
	if err := r.SetPathParam("versionId", o.VersionID); err != nil {
		return err
	}

	if o.WordListBaseUpdateObject != nil {
		if err := r.SetBodyParam(o.WordListBaseUpdateObject); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
