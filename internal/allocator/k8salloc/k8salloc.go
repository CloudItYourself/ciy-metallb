// SPDX-License-Identifier:Apache-2.0

package k8salloc

import (
	"os"

	"go.universe.tf/metallb/internal/allocator"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/labels"
)

// Ports turns a service definition into a set of allocator ports.
func Ports(svc *v1.Service) []allocator.Port {
	var ret []allocator.Port
	for _, port := range svc.Spec.Ports {
		ret = append(ret, allocator.Port{
			Proto: string(port.Protocol),
			Port:  int(port.Port),
		})
	}
	return ret
}

// SharingKey extracts the sharing key for a service.
func SharingKey(svc *v1.Service) string {
	// check if CIY env variable is set
	svcAnnotation := svc.Annotations["metallb.universe.tf/allow-shared-ip"]
	if svcAnnotation == "" {
		if _, ok := os.LookupEnv("CIY"); ok {
			return "ciy-shared"
		}
	}
	return svcAnnotation
}

// BackendKey extracts the backend key for a service.
func BackendKey(svc *v1.Service) string {
	if svc.Spec.ExternalTrafficPolicy == v1.ServiceExternalTrafficPolicyTypeLocal {
		return labels.Set(svc.Spec.Selector).String()
	}
	// Cluster traffic policy can share services regardless of backends.
	return ""
}
