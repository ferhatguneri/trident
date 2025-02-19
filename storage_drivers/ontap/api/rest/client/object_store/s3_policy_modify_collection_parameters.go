// Code generated by go-swagger; DO NOT EDIT.

package object_store

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewS3PolicyModifyCollectionParams creates a new S3PolicyModifyCollectionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewS3PolicyModifyCollectionParams() *S3PolicyModifyCollectionParams {
	return &S3PolicyModifyCollectionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewS3PolicyModifyCollectionParamsWithTimeout creates a new S3PolicyModifyCollectionParams object
// with the ability to set a timeout on a request.
func NewS3PolicyModifyCollectionParamsWithTimeout(timeout time.Duration) *S3PolicyModifyCollectionParams {
	return &S3PolicyModifyCollectionParams{
		timeout: timeout,
	}
}

// NewS3PolicyModifyCollectionParamsWithContext creates a new S3PolicyModifyCollectionParams object
// with the ability to set a context for a request.
func NewS3PolicyModifyCollectionParamsWithContext(ctx context.Context) *S3PolicyModifyCollectionParams {
	return &S3PolicyModifyCollectionParams{
		Context: ctx,
	}
}

// NewS3PolicyModifyCollectionParamsWithHTTPClient creates a new S3PolicyModifyCollectionParams object
// with the ability to set a custom HTTPClient for a request.
func NewS3PolicyModifyCollectionParamsWithHTTPClient(client *http.Client) *S3PolicyModifyCollectionParams {
	return &S3PolicyModifyCollectionParams{
		HTTPClient: client,
	}
}

/*
S3PolicyModifyCollectionParams contains all the parameters to send to the API endpoint

	for the s3 policy modify collection operation.

	Typically these are written to a http.Request.
*/
type S3PolicyModifyCollectionParams struct {

	/* Comment.

	   Filter by comment
	*/
	Comment *string

	/* ContinueOnFailure.

	   Continue even when the operation fails on one of the records.
	*/
	ContinueOnFailure *bool

	/* Info.

	   Info specification
	*/
	Info S3PolicyModifyCollectionBody

	/* Name.

	   Filter by name
	*/
	Name *string

	/* ReadOnly.

	   Filter by read-only
	*/
	ReadOnly *bool

	/* ReturnRecords.

	   The default is true for GET calls.  When set to false, only the number of records is returned.

	   Default: true
	*/
	ReturnRecords *bool

	/* ReturnTimeout.

	   The number of seconds to allow the call to execute before returning.  When iterating over a collection, the default is 15 seconds.  ONTAP returns earlier if either max records or the end of the collection is reached.

	   Default: 15
	*/
	ReturnTimeout *int64

	/* SerialRecords.

	   Perform the operation on the records synchronously.
	*/
	SerialRecords *bool

	/* StatementsActions.

	   Filter by statements.actions
	*/
	StatementsActions *string

	/* StatementsEffect.

	   Filter by statements.effect
	*/
	StatementsEffect *string

	/* StatementsIndex.

	   Filter by statements.index
	*/
	StatementsIndex *int64

	/* StatementsResources.

	   Filter by statements.resources
	*/
	StatementsResources *string

	/* StatementsSid.

	   Filter by statements.sid
	*/
	StatementsSid *string

	/* SvmName.

	   Filter by svm.name
	*/
	SvmName *string

	/* SvmUUID.

	   UUID of the SVM to which this object belongs.
	*/
	SvmUUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the s3 policy modify collection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *S3PolicyModifyCollectionParams) WithDefaults() *S3PolicyModifyCollectionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the s3 policy modify collection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *S3PolicyModifyCollectionParams) SetDefaults() {
	var (
		continueOnFailureDefault = bool(false)

		returnRecordsDefault = bool(true)

		returnTimeoutDefault = int64(15)

		serialRecordsDefault = bool(false)
	)

	val := S3PolicyModifyCollectionParams{
		ContinueOnFailure: &continueOnFailureDefault,
		ReturnRecords:     &returnRecordsDefault,
		ReturnTimeout:     &returnTimeoutDefault,
		SerialRecords:     &serialRecordsDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the s3 policy modify collection params
func (o *S3PolicyModifyCollectionParams) WithTimeout(timeout time.Duration) *S3PolicyModifyCollectionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the s3 policy modify collection params
func (o *S3PolicyModifyCollectionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the s3 policy modify collection params
func (o *S3PolicyModifyCollectionParams) WithContext(ctx context.Context) *S3PolicyModifyCollectionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the s3 policy modify collection params
func (o *S3PolicyModifyCollectionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the s3 policy modify collection params
func (o *S3PolicyModifyCollectionParams) WithHTTPClient(client *http.Client) *S3PolicyModifyCollectionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the s3 policy modify collection params
func (o *S3PolicyModifyCollectionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithComment adds the comment to the s3 policy modify collection params
func (o *S3PolicyModifyCollectionParams) WithComment(comment *string) *S3PolicyModifyCollectionParams {
	o.SetComment(comment)
	return o
}

// SetComment adds the comment to the s3 policy modify collection params
func (o *S3PolicyModifyCollectionParams) SetComment(comment *string) {
	o.Comment = comment
}

// WithContinueOnFailure adds the continueOnFailure to the s3 policy modify collection params
func (o *S3PolicyModifyCollectionParams) WithContinueOnFailure(continueOnFailure *bool) *S3PolicyModifyCollectionParams {
	o.SetContinueOnFailure(continueOnFailure)
	return o
}

// SetContinueOnFailure adds the continueOnFailure to the s3 policy modify collection params
func (o *S3PolicyModifyCollectionParams) SetContinueOnFailure(continueOnFailure *bool) {
	o.ContinueOnFailure = continueOnFailure
}

// WithInfo adds the info to the s3 policy modify collection params
func (o *S3PolicyModifyCollectionParams) WithInfo(info S3PolicyModifyCollectionBody) *S3PolicyModifyCollectionParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the s3 policy modify collection params
func (o *S3PolicyModifyCollectionParams) SetInfo(info S3PolicyModifyCollectionBody) {
	o.Info = info
}

// WithName adds the name to the s3 policy modify collection params
func (o *S3PolicyModifyCollectionParams) WithName(name *string) *S3PolicyModifyCollectionParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the s3 policy modify collection params
func (o *S3PolicyModifyCollectionParams) SetName(name *string) {
	o.Name = name
}

// WithReadOnly adds the readOnly to the s3 policy modify collection params
func (o *S3PolicyModifyCollectionParams) WithReadOnly(readOnly *bool) *S3PolicyModifyCollectionParams {
	o.SetReadOnly(readOnly)
	return o
}

// SetReadOnly adds the readOnly to the s3 policy modify collection params
func (o *S3PolicyModifyCollectionParams) SetReadOnly(readOnly *bool) {
	o.ReadOnly = readOnly
}

// WithReturnRecords adds the returnRecords to the s3 policy modify collection params
func (o *S3PolicyModifyCollectionParams) WithReturnRecords(returnRecords *bool) *S3PolicyModifyCollectionParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the s3 policy modify collection params
func (o *S3PolicyModifyCollectionParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WithReturnTimeout adds the returnTimeout to the s3 policy modify collection params
func (o *S3PolicyModifyCollectionParams) WithReturnTimeout(returnTimeout *int64) *S3PolicyModifyCollectionParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the s3 policy modify collection params
func (o *S3PolicyModifyCollectionParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WithSerialRecords adds the serialRecords to the s3 policy modify collection params
func (o *S3PolicyModifyCollectionParams) WithSerialRecords(serialRecords *bool) *S3PolicyModifyCollectionParams {
	o.SetSerialRecords(serialRecords)
	return o
}

// SetSerialRecords adds the serialRecords to the s3 policy modify collection params
func (o *S3PolicyModifyCollectionParams) SetSerialRecords(serialRecords *bool) {
	o.SerialRecords = serialRecords
}

// WithStatementsActions adds the statementsActions to the s3 policy modify collection params
func (o *S3PolicyModifyCollectionParams) WithStatementsActions(statementsActions *string) *S3PolicyModifyCollectionParams {
	o.SetStatementsActions(statementsActions)
	return o
}

// SetStatementsActions adds the statementsActions to the s3 policy modify collection params
func (o *S3PolicyModifyCollectionParams) SetStatementsActions(statementsActions *string) {
	o.StatementsActions = statementsActions
}

// WithStatementsEffect adds the statementsEffect to the s3 policy modify collection params
func (o *S3PolicyModifyCollectionParams) WithStatementsEffect(statementsEffect *string) *S3PolicyModifyCollectionParams {
	o.SetStatementsEffect(statementsEffect)
	return o
}

// SetStatementsEffect adds the statementsEffect to the s3 policy modify collection params
func (o *S3PolicyModifyCollectionParams) SetStatementsEffect(statementsEffect *string) {
	o.StatementsEffect = statementsEffect
}

// WithStatementsIndex adds the statementsIndex to the s3 policy modify collection params
func (o *S3PolicyModifyCollectionParams) WithStatementsIndex(statementsIndex *int64) *S3PolicyModifyCollectionParams {
	o.SetStatementsIndex(statementsIndex)
	return o
}

// SetStatementsIndex adds the statementsIndex to the s3 policy modify collection params
func (o *S3PolicyModifyCollectionParams) SetStatementsIndex(statementsIndex *int64) {
	o.StatementsIndex = statementsIndex
}

// WithStatementsResources adds the statementsResources to the s3 policy modify collection params
func (o *S3PolicyModifyCollectionParams) WithStatementsResources(statementsResources *string) *S3PolicyModifyCollectionParams {
	o.SetStatementsResources(statementsResources)
	return o
}

// SetStatementsResources adds the statementsResources to the s3 policy modify collection params
func (o *S3PolicyModifyCollectionParams) SetStatementsResources(statementsResources *string) {
	o.StatementsResources = statementsResources
}

// WithStatementsSid adds the statementsSid to the s3 policy modify collection params
func (o *S3PolicyModifyCollectionParams) WithStatementsSid(statementsSid *string) *S3PolicyModifyCollectionParams {
	o.SetStatementsSid(statementsSid)
	return o
}

// SetStatementsSid adds the statementsSid to the s3 policy modify collection params
func (o *S3PolicyModifyCollectionParams) SetStatementsSid(statementsSid *string) {
	o.StatementsSid = statementsSid
}

// WithSvmName adds the svmName to the s3 policy modify collection params
func (o *S3PolicyModifyCollectionParams) WithSvmName(svmName *string) *S3PolicyModifyCollectionParams {
	o.SetSvmName(svmName)
	return o
}

// SetSvmName adds the svmName to the s3 policy modify collection params
func (o *S3PolicyModifyCollectionParams) SetSvmName(svmName *string) {
	o.SvmName = svmName
}

// WithSvmUUID adds the svmUUID to the s3 policy modify collection params
func (o *S3PolicyModifyCollectionParams) WithSvmUUID(svmUUID string) *S3PolicyModifyCollectionParams {
	o.SetSvmUUID(svmUUID)
	return o
}

// SetSvmUUID adds the svmUuid to the s3 policy modify collection params
func (o *S3PolicyModifyCollectionParams) SetSvmUUID(svmUUID string) {
	o.SvmUUID = svmUUID
}

// WriteToRequest writes these params to a swagger request
func (o *S3PolicyModifyCollectionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Comment != nil {

		// query param comment
		var qrComment string

		if o.Comment != nil {
			qrComment = *o.Comment
		}
		qComment := qrComment
		if qComment != "" {

			if err := r.SetQueryParam("comment", qComment); err != nil {
				return err
			}
		}
	}

	if o.ContinueOnFailure != nil {

		// query param continue_on_failure
		var qrContinueOnFailure bool

		if o.ContinueOnFailure != nil {
			qrContinueOnFailure = *o.ContinueOnFailure
		}
		qContinueOnFailure := swag.FormatBool(qrContinueOnFailure)
		if qContinueOnFailure != "" {

			if err := r.SetQueryParam("continue_on_failure", qContinueOnFailure); err != nil {
				return err
			}
		}
	}
	if err := r.SetBodyParam(o.Info); err != nil {
		return err
	}

	if o.Name != nil {

		// query param name
		var qrName string

		if o.Name != nil {
			qrName = *o.Name
		}
		qName := qrName
		if qName != "" {

			if err := r.SetQueryParam("name", qName); err != nil {
				return err
			}
		}
	}

	if o.ReadOnly != nil {

		// query param read-only
		var qrReadOnly bool

		if o.ReadOnly != nil {
			qrReadOnly = *o.ReadOnly
		}
		qReadOnly := swag.FormatBool(qrReadOnly)
		if qReadOnly != "" {

			if err := r.SetQueryParam("read-only", qReadOnly); err != nil {
				return err
			}
		}
	}

	if o.ReturnRecords != nil {

		// query param return_records
		var qrReturnRecords bool

		if o.ReturnRecords != nil {
			qrReturnRecords = *o.ReturnRecords
		}
		qReturnRecords := swag.FormatBool(qrReturnRecords)
		if qReturnRecords != "" {

			if err := r.SetQueryParam("return_records", qReturnRecords); err != nil {
				return err
			}
		}
	}

	if o.ReturnTimeout != nil {

		// query param return_timeout
		var qrReturnTimeout int64

		if o.ReturnTimeout != nil {
			qrReturnTimeout = *o.ReturnTimeout
		}
		qReturnTimeout := swag.FormatInt64(qrReturnTimeout)
		if qReturnTimeout != "" {

			if err := r.SetQueryParam("return_timeout", qReturnTimeout); err != nil {
				return err
			}
		}
	}

	if o.SerialRecords != nil {

		// query param serial_records
		var qrSerialRecords bool

		if o.SerialRecords != nil {
			qrSerialRecords = *o.SerialRecords
		}
		qSerialRecords := swag.FormatBool(qrSerialRecords)
		if qSerialRecords != "" {

			if err := r.SetQueryParam("serial_records", qSerialRecords); err != nil {
				return err
			}
		}
	}

	if o.StatementsActions != nil {

		// query param statements.actions
		var qrStatementsActions string

		if o.StatementsActions != nil {
			qrStatementsActions = *o.StatementsActions
		}
		qStatementsActions := qrStatementsActions
		if qStatementsActions != "" {

			if err := r.SetQueryParam("statements.actions", qStatementsActions); err != nil {
				return err
			}
		}
	}

	if o.StatementsEffect != nil {

		// query param statements.effect
		var qrStatementsEffect string

		if o.StatementsEffect != nil {
			qrStatementsEffect = *o.StatementsEffect
		}
		qStatementsEffect := qrStatementsEffect
		if qStatementsEffect != "" {

			if err := r.SetQueryParam("statements.effect", qStatementsEffect); err != nil {
				return err
			}
		}
	}

	if o.StatementsIndex != nil {

		// query param statements.index
		var qrStatementsIndex int64

		if o.StatementsIndex != nil {
			qrStatementsIndex = *o.StatementsIndex
		}
		qStatementsIndex := swag.FormatInt64(qrStatementsIndex)
		if qStatementsIndex != "" {

			if err := r.SetQueryParam("statements.index", qStatementsIndex); err != nil {
				return err
			}
		}
	}

	if o.StatementsResources != nil {

		// query param statements.resources
		var qrStatementsResources string

		if o.StatementsResources != nil {
			qrStatementsResources = *o.StatementsResources
		}
		qStatementsResources := qrStatementsResources
		if qStatementsResources != "" {

			if err := r.SetQueryParam("statements.resources", qStatementsResources); err != nil {
				return err
			}
		}
	}

	if o.StatementsSid != nil {

		// query param statements.sid
		var qrStatementsSid string

		if o.StatementsSid != nil {
			qrStatementsSid = *o.StatementsSid
		}
		qStatementsSid := qrStatementsSid
		if qStatementsSid != "" {

			if err := r.SetQueryParam("statements.sid", qStatementsSid); err != nil {
				return err
			}
		}
	}

	if o.SvmName != nil {

		// query param svm.name
		var qrSvmName string

		if o.SvmName != nil {
			qrSvmName = *o.SvmName
		}
		qSvmName := qrSvmName
		if qSvmName != "" {

			if err := r.SetQueryParam("svm.name", qSvmName); err != nil {
				return err
			}
		}
	}

	// path param svm.uuid
	if err := r.SetPathParam("svm.uuid", o.SvmUUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
