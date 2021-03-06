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

	"github.com/crazedpeanut/go-luis-authoring-client/models"
)

// NewUpdateSublistParams creates a new UpdateSublistParams object
// with the default values initialized.
func NewUpdateSublistParams() *UpdateSublistParams {
	var ()
	return &UpdateSublistParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateSublistParamsWithTimeout creates a new UpdateSublistParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateSublistParamsWithTimeout(timeout time.Duration) *UpdateSublistParams {
	var ()
	return &UpdateSublistParams{

		timeout: timeout,
	}
}

// NewUpdateSublistParamsWithContext creates a new UpdateSublistParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateSublistParamsWithContext(ctx context.Context) *UpdateSublistParams {
	var ()
	return &UpdateSublistParams{

		Context: ctx,
	}
}

// NewUpdateSublistParamsWithHTTPClient creates a new UpdateSublistParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateSublistParamsWithHTTPClient(client *http.Client) *UpdateSublistParams {
	var ()
	return &UpdateSublistParams{
		HTTPClient: client,
	}
}

/*UpdateSublistParams contains all the parameters to send to the API endpoint
for the update sublist operation typically these are written to a http.Request
*/
type UpdateSublistParams struct {

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

// WithTimeout adds the timeout to the update sublist params
func (o *UpdateSublistParams) WithTimeout(timeout time.Duration) *UpdateSublistParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update sublist params
func (o *UpdateSublistParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update sublist params
func (o *UpdateSublistParams) WithContext(ctx context.Context) *UpdateSublistParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update sublist params
func (o *UpdateSublistParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update sublist params
func (o *UpdateSublistParams) WithHTTPClient(client *http.Client) *UpdateSublistParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update sublist params
func (o *UpdateSublistParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppID adds the appID to the update sublist params
func (o *UpdateSublistParams) WithAppID(appID string) *UpdateSublistParams {
	o.SetAppID(appID)
	return o
}

// SetAppID adds the appId to the update sublist params
func (o *UpdateSublistParams) SetAppID(appID string) {
	o.AppID = appID
}

// WithClEntityID adds the clEntityID to the update sublist params
func (o *UpdateSublistParams) WithClEntityID(clEntityID string) *UpdateSublistParams {
	o.SetClEntityID(clEntityID)
	return o
}

// SetClEntityID adds the clEntityId to the update sublist params
func (o *UpdateSublistParams) SetClEntityID(clEntityID string) {
	o.ClEntityID = clEntityID
}

// WithSubListID adds the subListID to the update sublist params
func (o *UpdateSublistParams) WithSubListID(subListID string) *UpdateSublistParams {
	o.SetSubListID(subListID)
	return o
}

// SetSubListID adds the subListId to the update sublist params
func (o *UpdateSublistParams) SetSubListID(subListID string) {
	o.SubListID = subListID
}

// WithVersionID adds the versionID to the update sublist params
func (o *UpdateSublistParams) WithVersionID(versionID string) *UpdateSublistParams {
	o.SetVersionID(versionID)
	return o
}

// SetVersionID adds the versionId to the update sublist params
func (o *UpdateSublistParams) SetVersionID(versionID string) {
	o.VersionID = versionID
}

// WithWordListBaseUpdateObject adds the wordListBaseUpdateObject to the update sublist params
func (o *UpdateSublistParams) WithWordListBaseUpdateObject(wordListBaseUpdateObject *models.WordListBaseUpdateObject) *UpdateSublistParams {
	o.SetWordListBaseUpdateObject(wordListBaseUpdateObject)
	return o
}

// SetWordListBaseUpdateObject adds the wordListBaseUpdateObject to the update sublist params
func (o *UpdateSublistParams) SetWordListBaseUpdateObject(wordListBaseUpdateObject *models.WordListBaseUpdateObject) {
	o.WordListBaseUpdateObject = wordListBaseUpdateObject
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateSublistParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
