package chaosmesh

import (
	"context"
	"github.com/chaos-mesh/chaos-mesh/api/v1alpha1"
	"k8s.io/client-go/kubernetes/scheme"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// HTTPChaosesGetter has a method to return a HTTPChaosInterface.
// A group's client should implement this interface.
type HTTPChaosesGetter interface {
	HTTPChaoses(namespace string) HTTPChaosInterface
}

// HTTPChaosInterface has methods to work with HTTPChaos resources.
type HTTPChaosInterface interface {
	Create(ctx context.Context, httpChaos *v1alpha1.HTTPChaos, opts metav1.CreateOptions) (*v1alpha1.HTTPChaos, error)
	Update(ctx context.Context, httpChaos *v1alpha1.HTTPChaos, opts metav1.UpdateOptions) (*v1alpha1.HTTPChaos, error)
	UpdateStatus(ctx context.Context, httpChaos *v1alpha1.HTTPChaos, opts metav1.UpdateOptions) (*v1alpha1.HTTPChaos, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1alpha1.HTTPChaos, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1alpha1.HTTPChaosList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1alpha1.HTTPChaos, err error)
}

// httpChaos implements IOChaosInterface
type httpChaos struct {
	client *rest.RESTClient
	ns     string
}

// newIOChaoses returns a IOChaoses
func newhttpChaoses(c *Client, namespace string) *httpChaos {
	return &httpChaos{
		client: &c.restClient,
		ns:     namespace,
	}
}

// Get takes name of the iOChaos, and returns the corresponding iOChaos object, and an error if there is any.
func (c *httpChaos) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1alpha1.HTTPChaos, err error) {
	result = &v1alpha1.HTTPChaos{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("httpchaos").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of IOChaoses that match those selectors.
func (c *httpChaos) List(ctx context.Context, opts metav1.ListOptions) (result *v1alpha1.HTTPChaosList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.HTTPChaosList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("httpchaos").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested iOChaoses.
func (c *httpChaos) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("httpchaos").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a iOChaos and creates it.  Returns the server's representation of the iOChaos, and an error, if there is any.
func (c *httpChaos) Create(ctx context.Context, httpChaos *v1alpha1.HTTPChaos, opts metav1.CreateOptions) (result *v1alpha1.HTTPChaos, err error) {
	result = &v1alpha1.HTTPChaos{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("httpchaos").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(httpChaos).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a iOChaos and updates it. Returns the server's representation of the iOChaos, and an error, if there is any.
func (c *httpChaos) Update(ctx context.Context, httpChaos *v1alpha1.HTTPChaos, opts metav1.UpdateOptions) (result *v1alpha1.HTTPChaos, err error) {
	result = &v1alpha1.HTTPChaos{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("httpchaos").
		Name(httpChaos.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(httpChaos).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *httpChaos) UpdateStatus(ctx context.Context, httpChaos *v1alpha1.HTTPChaos, opts metav1.UpdateOptions) (result *v1alpha1.HTTPChaos, err error) {
	result = &v1alpha1.HTTPChaos{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("httpchaos").
		Name(httpChaos.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(httpChaos).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the iOChaos and deletes it. Returns an error if one occurs.
func (c *httpChaos) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("httpchaos").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *httpChaos) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("httpchaos").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched iOChaos.
func (c *httpChaos) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1alpha1.HTTPChaos, err error) {
	result = &v1alpha1.HTTPChaos{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("httpchaos").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

func Hhhchaos(name string, ns string) v1alpha1.HTTPChaos {
	return v1alpha1.HTTPChaos{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: ns,
		},
		Spec: v1alpha1.HTTPChaosSpec{
			PodSelector: v1alpha1.PodSelector{
				Selector: v1alpha1.PodSelectorSpec{
					GenericSelectorSpec: v1alpha1.GenericSelectorSpec{
						Namespaces:     []string{ns},
						LabelSelectors: map[string]string{"app": "http"},
					},
				},
				Mode: v1alpha1.OneMode,
			},
			Port:   8080,
			Target: "Request",
			PodHttpChaosActions: v1alpha1.PodHttpChaosActions{
				Replace: &v1alpha1.PodHttpChaosReplaceActions{
					Body: []byte("assasadsadsa"),
				},
			},
		},
	}
}
