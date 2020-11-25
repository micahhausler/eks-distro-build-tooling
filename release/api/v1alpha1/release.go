/*
Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License"). You may
not use this file except in compliance with the License. A copy of the
License is located at

    http://aws.amazon.com/apache2.0/

or in the "license" file accompanying this file. This file is distributed
on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
express or implied. See the License for the specific language governing
permissions and limitations under the License.
*/

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ReleaseSpec defines the desired state of Release
type ReleaseSpec struct {

	// +kubebuilder:validation:Required
	Channel string `json:"channel,omitempty"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Minimum=1
	// Monotonically increasing release number
	Number int `json:"number,omitempty"`

	// +kubebuilder:validation:Required
	BuildRepoCommit string `json:"buildRepoCommit,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:printcolumn:name="Release Channel",type=string,JSONPath=`.spec.releaseChannel`,description="The release channel"
// +kubebuilder:printcolumn:name="Release",type=integer,JSONPath=`.spec.release`,description="Release number"
// +kubebuilder:printcolumn:name="Release Date",type=string,format=date-time,JSONPath=`.spec.releaseDate`,description="The date the release was published"
// +kubebuilder:resource:singular="release",path="releases",shortName={"rel"}

// Release is the Schema for the releases API
type Release struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ReleaseSpec   `json:"spec,omitempty"`
	Status ReleaseStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ReleaseList contains a list of Release
type ReleaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Release `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Release{}, &ReleaseList{})
}