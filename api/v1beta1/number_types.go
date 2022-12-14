/*
Copyright 2022.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// NumberSpec defines the desired state of Number
type NumberSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	Value int64 `json:"value,omitempty"`
}

// NumberStatus defines the observed state of Number
type NumberStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	FizzBuzz string `json:"fizz_buzz,omitempty"`
	IsSquare bool   `json:"is_square,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
//+kubebuilder:printcolumn:name="Value",type=integer,JSONPath=`.spec.value`
//+kubebuilder:printcolumn:name="FizzBuzz",type=string,JSONPath=`.status.fizz_buzz`
//+kubebuilder:printcolumn:name="Square",type=boolean,JSONPath=`.status.is_square`

// Number is the Schema for the numbers API
type Number struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NumberSpec   `json:"spec,omitempty"`
	Status NumberStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// NumberList contains a list of Number
type NumberList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Number `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Number{}, &NumberList{})
}
