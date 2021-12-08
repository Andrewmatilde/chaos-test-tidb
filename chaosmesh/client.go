package chaosmesh

import (
	"github.com/chaos-mesh/chaos-mesh/api/v1alpha1"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
)

type Client struct {
	restClient rest.RESTClient
}

func NewClientFor(config *rest.Config) (*Client, error) {
	err := v1alpha1.AddToScheme(scheme.Scheme)
	if err != nil {
		return nil, err
	}
	crdConfig := *config

	crdConfig.ContentConfig.GroupVersion = &v1alpha1.GroupVersion
	crdConfig.APIPath = "/apis"
	crdConfig.NegotiatedSerializer = serializer.NewCodecFactory(scheme.Scheme)
	crdConfig.UserAgent = rest.DefaultKubernetesUserAgent()

	restClient, err := rest.UnversionedRESTClientFor(&crdConfig)
	if err != nil {
		return nil, err
	}

	return &Client{
		*restClient,
	}, nil
}

func (c *Client) IoChaos(ns string) *iOChaos {
	return newIOChaoses(c, ns)
}

func (c *Client) HTTPChaos(ns string) *httpChaos {
	return newhttpChaoses(c, ns)
}
