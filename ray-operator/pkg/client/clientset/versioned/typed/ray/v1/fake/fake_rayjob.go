// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"
	json "encoding/json"
	"fmt"

	v1 "github.com/ray-project/kuberay/ray-operator/apis/ray/v1"
	rayv1 "github.com/ray-project/kuberay/ray-operator/pkg/client/applyconfiguration/ray/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeRayJobs implements RayJobInterface
type FakeRayJobs struct {
	Fake *FakeRayV1
	ns   string
}

var rayjobsResource = v1.SchemeGroupVersion.WithResource("rayjobs")

var rayjobsKind = v1.SchemeGroupVersion.WithKind("RayJob")

// Get takes name of the rayJob, and returns the corresponding rayJob object, and an error if there is any.
func (c *FakeRayJobs) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.RayJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(rayjobsResource, c.ns, name), &v1.RayJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.RayJob), err
}

// List takes label and field selectors, and returns the list of RayJobs that match those selectors.
func (c *FakeRayJobs) List(ctx context.Context, opts metav1.ListOptions) (result *v1.RayJobList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(rayjobsResource, rayjobsKind, c.ns, opts), &v1.RayJobList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.RayJobList{ListMeta: obj.(*v1.RayJobList).ListMeta}
	for _, item := range obj.(*v1.RayJobList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested rayJobs.
func (c *FakeRayJobs) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(rayjobsResource, c.ns, opts))

}

// Create takes the representation of a rayJob and creates it.  Returns the server's representation of the rayJob, and an error, if there is any.
func (c *FakeRayJobs) Create(ctx context.Context, rayJob *v1.RayJob, opts metav1.CreateOptions) (result *v1.RayJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(rayjobsResource, c.ns, rayJob), &v1.RayJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.RayJob), err
}

// Update takes the representation of a rayJob and updates it. Returns the server's representation of the rayJob, and an error, if there is any.
func (c *FakeRayJobs) Update(ctx context.Context, rayJob *v1.RayJob, opts metav1.UpdateOptions) (result *v1.RayJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(rayjobsResource, c.ns, rayJob), &v1.RayJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.RayJob), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeRayJobs) UpdateStatus(ctx context.Context, rayJob *v1.RayJob, opts metav1.UpdateOptions) (*v1.RayJob, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(rayjobsResource, "status", c.ns, rayJob), &v1.RayJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.RayJob), err
}

// Delete takes name of the rayJob and deletes it. Returns an error if one occurs.
func (c *FakeRayJobs) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(rayjobsResource, c.ns, name, opts), &v1.RayJob{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeRayJobs) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(rayjobsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1.RayJobList{})
	return err
}

// Patch applies the patch and returns the patched rayJob.
func (c *FakeRayJobs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.RayJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(rayjobsResource, c.ns, name, pt, data, subresources...), &v1.RayJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.RayJob), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied rayJob.
func (c *FakeRayJobs) Apply(ctx context.Context, rayJob *rayv1.RayJobApplyConfiguration, opts metav1.ApplyOptions) (result *v1.RayJob, err error) {
	if rayJob == nil {
		return nil, fmt.Errorf("rayJob provided to Apply must not be nil")
	}
	data, err := json.Marshal(rayJob)
	if err != nil {
		return nil, err
	}
	name := rayJob.Name
	if name == nil {
		return nil, fmt.Errorf("rayJob.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(rayjobsResource, c.ns, *name, types.ApplyPatchType, data), &v1.RayJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.RayJob), err
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *FakeRayJobs) ApplyStatus(ctx context.Context, rayJob *rayv1.RayJobApplyConfiguration, opts metav1.ApplyOptions) (result *v1.RayJob, err error) {
	if rayJob == nil {
		return nil, fmt.Errorf("rayJob provided to Apply must not be nil")
	}
	data, err := json.Marshal(rayJob)
	if err != nil {
		return nil, err
	}
	name := rayJob.Name
	if name == nil {
		return nil, fmt.Errorf("rayJob.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(rayjobsResource, c.ns, *name, types.ApplyPatchType, data, "status"), &v1.RayJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.RayJob), err
}
