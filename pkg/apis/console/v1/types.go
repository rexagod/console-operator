package v1

import (
	operatorsv1 "github.com/openshift/api/operator/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type Console struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ConsoleSpec   `json:"spec,omitempty"`
	Status ConsoleStatus `json:"status,omitempty"`
}

type ConsoleSpec struct {
	operatorsv1.OperatorSpec
	// Count is the number of Console replicas
	Count int32 `json:"count,omitempty"`
	// take a look @:
	// https://github.com/openshift/cluster-image-registry-operator/blob/master/pkg/apis/imageregistry/v1/types.go#L110-L115
	// DefaultRoute: T|F
	// additional routes config w/secrets
	// Route[]
}

type ConsoleStatus struct {
	operatorsv1.OperatorStatus
	// set once the router has a default host name
	DefaultHostName string `json:"defaultHostName"`
	OAuthSecret     string `json:"oauthSecret"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ConsoleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []Console `json:"items"`
}
