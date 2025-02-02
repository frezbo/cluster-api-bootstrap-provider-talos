// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package v1alpha3

import (
	"fmt"

	apierrors "k8s.io/apimachinery/pkg/api/errors"
	runtime "k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/validation/field"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

func (r *TalosConfig) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

//+kubebuilder:webhook:verbs=create;update;delete,path=/validate-bootstrap-cluster-x-k8s-io-v1alpha3-talosconfig,mutating=false,failurePolicy=fail,groups=bootstrap.cluster.x-k8s.io,resources=talosconfigs,versions=v1alpha3,name=vtalosconfig.cluster.x-k8s.io,sideEffects=None,admissionReviewVersions=v1

var _ webhook.Validator = &TalosConfig{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *TalosConfig) ValidateCreate() error {
	return r.validate()
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *TalosConfig) ValidateUpdate(old runtime.Object) error {
	return r.validate()
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *TalosConfig) ValidateDelete() error {
	return nil
}

func (r *TalosConfig) validate() error {
	var allErrs field.ErrorList

	switch r.Spec.Hostname.Source {
	case "":
	case HostnameSourceMachineName:
	default:
		allErrs = append(allErrs,
			field.Invalid(field.NewPath("spec").Child("hostname").Child("source"), r.Spec.Hostname.Source,
				fmt.Sprintf("valid values are: %q", []HostnameSource{HostnameSourceMachineName}),
			),
		)
	}

	if len(allErrs) == 0 {
		return nil
	}

	return apierrors.NewInvalid(
		schema.GroupKind{Group: GroupVersion.Group, Kind: "TalosConfig"},
		r.Name, allErrs)
}
