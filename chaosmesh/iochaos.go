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

// IOChaosesGetter has a method to return a IOChaosInterface.
// A group's client should implement this interface.
type IOChaosesGetter interface {
	IOChaoses(namespace string) IOChaosInterface
}

// IOChaosInterface has methods to work with IOChaos resources.
type IOChaosInterface interface {
	Create(ctx context.Context, iOChaos *v1alpha1.IOChaos, opts metav1.CreateOptions) (*v1alpha1.IOChaos, error)
	Update(ctx context.Context, iOChaos *v1alpha1.IOChaos, opts metav1.UpdateOptions) (*v1alpha1.IOChaos, error)
	UpdateStatus(ctx context.Context, iOChaos *v1alpha1.IOChaos, opts metav1.UpdateOptions) (*v1alpha1.IOChaos, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1alpha1.IOChaos, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1alpha1.IOChaosList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1alpha1.IOChaos, err error)
}

// iOChaoses implements IOChaosInterface
type iOChaos struct {
	client *rest.RESTClient
	ns     string
}

// newIOChaoses returns a IOChaoses
func newIOChaoses(c *Client, namespace string) *iOChaos {
	return &iOChaos{
		client: &c.restClient,
		ns:     namespace,
	}
}

// Get takes name of the iOChaos, and returns the corresponding iOChaos object, and an error if there is any.
func (c *iOChaos) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1alpha1.IOChaos, err error) {
	result = &v1alpha1.IOChaos{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("iochaos").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of IOChaoses that match those selectors.
func (c *iOChaos) List(ctx context.Context, opts metav1.ListOptions) (result *v1alpha1.IOChaosList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.IOChaosList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("iochaos").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested iOChaoses.
func (c *iOChaos) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("iochaos").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a iOChaos and creates it.  Returns the server's representation of the iOChaos, and an error, if there is any.
func (c *iOChaos) Create(ctx context.Context, iOChaos *v1alpha1.IOChaos, opts metav1.CreateOptions) (result *v1alpha1.IOChaos, err error) {
	result = &v1alpha1.IOChaos{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("iochaos").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(iOChaos).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a iOChaos and updates it. Returns the server's representation of the iOChaos, and an error, if there is any.
func (c *iOChaos) Update(ctx context.Context, iOChaos *v1alpha1.IOChaos, opts metav1.UpdateOptions) (result *v1alpha1.IOChaos, err error) {
	result = &v1alpha1.IOChaos{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("iochaos").
		Name(iOChaos.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(iOChaos).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *iOChaos) UpdateStatus(ctx context.Context, iOChaos *v1alpha1.IOChaos, opts metav1.UpdateOptions) (result *v1alpha1.IOChaos, err error) {
	result = &v1alpha1.IOChaos{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("iochaos").
		Name(iOChaos.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(iOChaos).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the iOChaos and deletes it. Returns an error if one occurs.
func (c *iOChaos) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("iochaos").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *iOChaos) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("iochaos").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched iOChaos.
func (c *iOChaos) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1alpha1.IOChaos, err error) {
	result = &v1alpha1.IOChaos{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("iochaos").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

func IoChaosForTikv(name string, ns string, tikvPod string) v1alpha1.IOChaos {
	return v1alpha1.IOChaos{
		TypeMeta: metav1.TypeMeta{
			Kind:       "IOChaos",
			APIVersion: "chaos-mesh.org/v1alpha1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:         name,
			GenerateName: "",
			Namespace:    "tidb-c0",
		},
		Spec: v1alpha1.IOChaosSpec{
			ContainerSelector: v1alpha1.ContainerSelector{
				PodSelector: v1alpha1.PodSelector{
					Selector: v1alpha1.PodSelectorSpec{
						GenericSelectorSpec: v1alpha1.GenericSelectorSpec{
							Namespaces: []string{ns},
						},
						Pods:              map[string][]string{ns: {tikvPod}},
						NodeSelectors:     nil,
						PodPhaseSelectors: nil,
					},
					Mode:  "all",
					Value: "",
				},
				ContainerNames: []string{"tikv"},
			},
			Action:     "latency",
			Delay:      "100ms",
			Errno:      0,
			Attr:       nil,
			Mistake:    nil,
			Path:       "",
			Methods:    []v1alpha1.IoMethod{v1alpha1.Write},
			Percent:    100,
			VolumePath: "/var/lib/tikv",
			Duration:   nil,
		},
		Status: v1alpha1.IOChaosStatus{},
	}
}
