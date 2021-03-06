/*
Copyright 2019 The Kubernetes Authors.

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

package v1alpha3

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	clusterv1 "sigs.k8s.io/cluster-api/api/v1alpha3"
	"sigs.k8s.io/cluster-api/errors"
)

const (
	// MachineFinalizer allows ReconcileVSphereMachine to clean up VSphere
	// resources associated with VSphereMachine before removing it from the
	// API Server.
	MachineFinalizer = "vspheremachine.infrastructure.cluster.x-k8s.io"
)

// VSphereMachineSpec defines the desired state of VSphereMachine
type VSphereMachineSpec struct {
	VirtualMachineCloneSpec `json:",inline"`

	// ProviderID is the virtual machine's BIOS UUID formated as
	// vsphere://12345678-1234-1234-1234-123456789abc
	// +optional
	ProviderID *string `json:"providerID,omitempty"`
}

// VSphereMachineStatus defines the observed state of VSphereMachine
type VSphereMachineStatus struct {
	// Ready is true when the provider resource is ready.
	// +optional
	Ready bool `json:"ready"`

	// Addresses contains the VSphere instance associated addresses.
	Addresses []clusterv1.MachineAddress `json:"addresses,omitempty"`

	// Network returns the network status for each of the machine's configured
	// network interfaces.
	// +optional
	Network []NetworkStatus `json:"network,omitempty"`

	// ErrorReason will be set in the event that there is a terminal problem
	// reconciling the Machine and will contain a succinct value suitable
	// for machine interpretation.
	//
	// This field should not be set for transitive errors that a controller
	// faces that are expected to be fixed automatically over
	// time (like service outages), but instead indicate that something is
	// fundamentally wrong with the Machine's spec or the configuration of
	// the controller, and that manual intervention is required. Examples
	// of terminal errors would be invalid combinations of settings in the
	// spec, values that are unsupported by the controller, or the
	// responsible controller itself being critically misconfigured.
	//
	// Any transient errors that occur during the reconciliation of Machines
	// can be added as events to the Machine object and/or logged in the
	// controller's output.
	// +optional
	ErrorReason *errors.MachineStatusError `json:"errorReason,omitempty"`

	// ErrorMessage will be set in the event that there is a terminal problem
	// reconciling the Machine and will contain a more verbose string suitable
	// for logging and human consumption.
	//
	// This field should not be set for transitive errors that a controller
	// faces that are expected to be fixed automatically over
	// time (like service outages), but instead indicate that something is
	// fundamentally wrong with the Machine's spec or the configuration of
	// the controller, and that manual intervention is required. Examples
	// of terminal errors would be invalid combinations of settings in the
	// spec, values that are unsupported by the controller, or the
	// responsible controller itself being critically misconfigured.
	//
	// Any transient errors that occur during the reconciliation of Machines
	// can be added as events to the Machine object and/or logged in the
	// controller's output.
	// +optional
	ErrorMessage *string `json:"errorMessage,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=vspheremachines,scope=Namespaced,categories=cluster-api
// +kubebuilder:subresource:status

// VSphereMachine is the Schema for the vspheremachines API
type VSphereMachine struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VSphereMachineSpec   `json:"spec,omitempty"`
	Status VSphereMachineStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VSphereMachineList contains a list of VSphereMachine
type VSphereMachineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VSphereMachine `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VSphereMachine{}, &VSphereMachineList{})
}
