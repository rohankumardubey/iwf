/*
Workflow APIs

This APIs for iwf SDKs to operate workflows

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package iwfidl

import (
	"encoding/json"
)

// WorkflowStateStartResponse struct for WorkflowStateStartResponse
type WorkflowStateStartResponse struct {
	UpsertSearchAttributes []SearchAttribute `json:"upsertSearchAttributes,omitempty"`
	UpsertDataObjects []KeyValue `json:"upsertDataObjects,omitempty"`
	CommandRequest *CommandRequest `json:"commandRequest,omitempty"`
	UpsertStateLocals []KeyValue `json:"upsertStateLocals,omitempty"`
	RecordEvents []KeyValue `json:"recordEvents,omitempty"`
	PublishToInterStateChannel []InterStateChannelPublishing `json:"publishToInterStateChannel,omitempty"`
}

// NewWorkflowStateStartResponse instantiates a new WorkflowStateStartResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowStateStartResponse() *WorkflowStateStartResponse {
	this := WorkflowStateStartResponse{}
	return &this
}

// NewWorkflowStateStartResponseWithDefaults instantiates a new WorkflowStateStartResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowStateStartResponseWithDefaults() *WorkflowStateStartResponse {
	this := WorkflowStateStartResponse{}
	return &this
}

// GetUpsertSearchAttributes returns the UpsertSearchAttributes field value if set, zero value otherwise.
func (o *WorkflowStateStartResponse) GetUpsertSearchAttributes() []SearchAttribute {
	if o == nil || isNil(o.UpsertSearchAttributes) {
		var ret []SearchAttribute
		return ret
	}
	return o.UpsertSearchAttributes
}

// GetUpsertSearchAttributesOk returns a tuple with the UpsertSearchAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowStateStartResponse) GetUpsertSearchAttributesOk() ([]SearchAttribute, bool) {
	if o == nil || isNil(o.UpsertSearchAttributes) {
    return nil, false
	}
	return o.UpsertSearchAttributes, true
}

// HasUpsertSearchAttributes returns a boolean if a field has been set.
func (o *WorkflowStateStartResponse) HasUpsertSearchAttributes() bool {
	if o != nil && !isNil(o.UpsertSearchAttributes) {
		return true
	}

	return false
}

// SetUpsertSearchAttributes gets a reference to the given []SearchAttribute and assigns it to the UpsertSearchAttributes field.
func (o *WorkflowStateStartResponse) SetUpsertSearchAttributes(v []SearchAttribute) {
	o.UpsertSearchAttributes = v
}

// GetUpsertDataObjects returns the UpsertDataObjects field value if set, zero value otherwise.
func (o *WorkflowStateStartResponse) GetUpsertDataObjects() []KeyValue {
	if o == nil || isNil(o.UpsertDataObjects) {
		var ret []KeyValue
		return ret
	}
	return o.UpsertDataObjects
}

// GetUpsertDataObjectsOk returns a tuple with the UpsertDataObjects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowStateStartResponse) GetUpsertDataObjectsOk() ([]KeyValue, bool) {
	if o == nil || isNil(o.UpsertDataObjects) {
    return nil, false
	}
	return o.UpsertDataObjects, true
}

// HasUpsertDataObjects returns a boolean if a field has been set.
func (o *WorkflowStateStartResponse) HasUpsertDataObjects() bool {
	if o != nil && !isNil(o.UpsertDataObjects) {
		return true
	}

	return false
}

// SetUpsertDataObjects gets a reference to the given []KeyValue and assigns it to the UpsertDataObjects field.
func (o *WorkflowStateStartResponse) SetUpsertDataObjects(v []KeyValue) {
	o.UpsertDataObjects = v
}

// GetCommandRequest returns the CommandRequest field value if set, zero value otherwise.
func (o *WorkflowStateStartResponse) GetCommandRequest() CommandRequest {
	if o == nil || isNil(o.CommandRequest) {
		var ret CommandRequest
		return ret
	}
	return *o.CommandRequest
}

// GetCommandRequestOk returns a tuple with the CommandRequest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowStateStartResponse) GetCommandRequestOk() (*CommandRequest, bool) {
	if o == nil || isNil(o.CommandRequest) {
    return nil, false
	}
	return o.CommandRequest, true
}

// HasCommandRequest returns a boolean if a field has been set.
func (o *WorkflowStateStartResponse) HasCommandRequest() bool {
	if o != nil && !isNil(o.CommandRequest) {
		return true
	}

	return false
}

// SetCommandRequest gets a reference to the given CommandRequest and assigns it to the CommandRequest field.
func (o *WorkflowStateStartResponse) SetCommandRequest(v CommandRequest) {
	o.CommandRequest = &v
}

// GetUpsertStateLocals returns the UpsertStateLocals field value if set, zero value otherwise.
func (o *WorkflowStateStartResponse) GetUpsertStateLocals() []KeyValue {
	if o == nil || isNil(o.UpsertStateLocals) {
		var ret []KeyValue
		return ret
	}
	return o.UpsertStateLocals
}

// GetUpsertStateLocalsOk returns a tuple with the UpsertStateLocals field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowStateStartResponse) GetUpsertStateLocalsOk() ([]KeyValue, bool) {
	if o == nil || isNil(o.UpsertStateLocals) {
    return nil, false
	}
	return o.UpsertStateLocals, true
}

// HasUpsertStateLocals returns a boolean if a field has been set.
func (o *WorkflowStateStartResponse) HasUpsertStateLocals() bool {
	if o != nil && !isNil(o.UpsertStateLocals) {
		return true
	}

	return false
}

// SetUpsertStateLocals gets a reference to the given []KeyValue and assigns it to the UpsertStateLocals field.
func (o *WorkflowStateStartResponse) SetUpsertStateLocals(v []KeyValue) {
	o.UpsertStateLocals = v
}

// GetRecordEvents returns the RecordEvents field value if set, zero value otherwise.
func (o *WorkflowStateStartResponse) GetRecordEvents() []KeyValue {
	if o == nil || isNil(o.RecordEvents) {
		var ret []KeyValue
		return ret
	}
	return o.RecordEvents
}

// GetRecordEventsOk returns a tuple with the RecordEvents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowStateStartResponse) GetRecordEventsOk() ([]KeyValue, bool) {
	if o == nil || isNil(o.RecordEvents) {
    return nil, false
	}
	return o.RecordEvents, true
}

// HasRecordEvents returns a boolean if a field has been set.
func (o *WorkflowStateStartResponse) HasRecordEvents() bool {
	if o != nil && !isNil(o.RecordEvents) {
		return true
	}

	return false
}

// SetRecordEvents gets a reference to the given []KeyValue and assigns it to the RecordEvents field.
func (o *WorkflowStateStartResponse) SetRecordEvents(v []KeyValue) {
	o.RecordEvents = v
}

// GetPublishToInterStateChannel returns the PublishToInterStateChannel field value if set, zero value otherwise.
func (o *WorkflowStateStartResponse) GetPublishToInterStateChannel() []InterStateChannelPublishing {
	if o == nil || isNil(o.PublishToInterStateChannel) {
		var ret []InterStateChannelPublishing
		return ret
	}
	return o.PublishToInterStateChannel
}

// GetPublishToInterStateChannelOk returns a tuple with the PublishToInterStateChannel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowStateStartResponse) GetPublishToInterStateChannelOk() ([]InterStateChannelPublishing, bool) {
	if o == nil || isNil(o.PublishToInterStateChannel) {
    return nil, false
	}
	return o.PublishToInterStateChannel, true
}

// HasPublishToInterStateChannel returns a boolean if a field has been set.
func (o *WorkflowStateStartResponse) HasPublishToInterStateChannel() bool {
	if o != nil && !isNil(o.PublishToInterStateChannel) {
		return true
	}

	return false
}

// SetPublishToInterStateChannel gets a reference to the given []InterStateChannelPublishing and assigns it to the PublishToInterStateChannel field.
func (o *WorkflowStateStartResponse) SetPublishToInterStateChannel(v []InterStateChannelPublishing) {
	o.PublishToInterStateChannel = v
}

func (o WorkflowStateStartResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.UpsertSearchAttributes) {
		toSerialize["upsertSearchAttributes"] = o.UpsertSearchAttributes
	}
	if !isNil(o.UpsertDataObjects) {
		toSerialize["upsertDataObjects"] = o.UpsertDataObjects
	}
	if !isNil(o.CommandRequest) {
		toSerialize["commandRequest"] = o.CommandRequest
	}
	if !isNil(o.UpsertStateLocals) {
		toSerialize["upsertStateLocals"] = o.UpsertStateLocals
	}
	if !isNil(o.RecordEvents) {
		toSerialize["recordEvents"] = o.RecordEvents
	}
	if !isNil(o.PublishToInterStateChannel) {
		toSerialize["publishToInterStateChannel"] = o.PublishToInterStateChannel
	}
	return json.Marshal(toSerialize)
}

type NullableWorkflowStateStartResponse struct {
	value *WorkflowStateStartResponse
	isSet bool
}

func (v NullableWorkflowStateStartResponse) Get() *WorkflowStateStartResponse {
	return v.value
}

func (v *NullableWorkflowStateStartResponse) Set(val *WorkflowStateStartResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowStateStartResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowStateStartResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowStateStartResponse(val *WorkflowStateStartResponse) *NullableWorkflowStateStartResponse {
	return &NullableWorkflowStateStartResponse{value: val, isSet: true}
}

func (v NullableWorkflowStateStartResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowStateStartResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


