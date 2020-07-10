package v1alpha1

import (
	cloudbuild "google.golang.org/api/cloudbuild/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient:noStatus
// +genreconciler
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Build represents a Google Cloud Build.
//
// +k8s:openapi-gen=true
type Build struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +optional
	Spec BuildSpec `json:"spec"`
}

type gcbBuild cloudbuild.Build

func (in *gcbBuild) DeepCopyInto(out *gcbBuild) {
	*out = *in
	return
}

type BuildSpec struct {
	gcbBuild `json:",inline"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// BuildList contains a list of Build
type BuildList struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Build `json:"items"`
}
