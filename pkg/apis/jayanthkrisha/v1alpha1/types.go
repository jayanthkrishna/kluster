package v1aplha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Kluster struct {
	metav1.TypeMeta

	metav1.ObjectMeta

	Spec KlusterSpec
}

type KlusterSpec struct {
	Name      string
	Region    string
	Version   string
	NodePools []Nodepool
}

type Nodepool struct {
	Size  string
	Name  string
	Count int
}
