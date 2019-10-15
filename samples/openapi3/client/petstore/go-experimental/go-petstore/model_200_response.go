/*
 * OpenAPI Petstore
 *
 * This spec is mainly for testing Petstore server and contains fake endpoints, models. Please do not use this for any other purpose. Special characters: \" \\
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
import (
	"encoding/json"
)
// Model200Response Model for testing model name starting with number
type Model200Response struct {
	Name *int32 `json:"name,omitempty"`

	Class *string `json:"class,omitempty"`

}

// GetName returns the Name field if non-nil, zero value otherwise.
func (o *Model200Response) GetName() int32 {
	if o == nil || o.Name == nil {
		var ret int32
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Model200Response) GetNameOk() (int32, bool) {
	if o == nil || o.Name == nil {
		var ret int32
		return ret, false
	}
	return *o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Model200Response) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given int32 and assigns it to the Name field.
func (o *Model200Response) SetName(v int32) {
	o.Name = &v
}

// GetClass returns the Class field if non-nil, zero value otherwise.
func (o *Model200Response) GetClass() string {
	if o == nil || o.Class == nil {
		var ret string
		return ret
	}
	return *o.Class
}

// GetClassOk returns a tuple with the Class field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Model200Response) GetClassOk() (string, bool) {
	if o == nil || o.Class == nil {
		var ret string
		return ret, false
	}
	return *o.Class, true
}

// HasClass returns a boolean if a field has been set.
func (o *Model200Response) HasClass() bool {
	if o != nil && o.Class != nil {
		return true
	}

	return false
}

// SetClass gets a reference to the given string and assigns it to the Class field.
func (o *Model200Response) SetClass(v string) {
	o.Class = &v
}


// MarshalJSON returns the JSON representation of the model.
func (o Model200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Class != nil {
		toSerialize["class"] = o.Class
	}
	return json.Marshal(toSerialize)
}

