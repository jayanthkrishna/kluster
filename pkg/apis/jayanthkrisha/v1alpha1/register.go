package v1aplha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

var SchemeGroupVersion = schema.GroupVersion{
	Group:   "jayanthkrishna ",
	Version: "v1alpha1",
}

var (
	SchemeBuilder runtime.SchemeBuilder
)

func init() {

	SchemeBuilder.Register(addKnownTypes)

}

func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(SchemeGroupVersion, &Kluster{}, &KlusterList{})
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)

	return nil
}
