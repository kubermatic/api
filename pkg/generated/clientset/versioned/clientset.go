// Code generated by client-gen. DO NOT EDIT.

package versioned

import (
	"fmt"
	"net/http"

	kubermaticappsv1 "k8c.io/api/v3/pkg/generated/clientset/versioned/typed/apps.kubermatic/v1"
	kubermaticenterpriseappsv1 "k8c.io/api/v3/pkg/generated/clientset/versioned/typed/ee.apps.kubermatic/v1"
	kubermaticenterprisev1 "k8c.io/api/v3/pkg/generated/clientset/versioned/typed/ee.kubermatic/v1"
	kubermaticv1 "k8c.io/api/v3/pkg/generated/clientset/versioned/typed/kubermatic/v1"
	discovery "k8s.io/client-go/discovery"
	rest "k8s.io/client-go/rest"
	flowcontrol "k8s.io/client-go/util/flowcontrol"
)

type Interface interface {
	Discovery() discovery.DiscoveryInterface
	KubermaticAppsV1() kubermaticappsv1.KubermaticAppsV1Interface
	KubermaticEnterpriseAppsV1() kubermaticenterpriseappsv1.KubermaticEnterpriseAppsV1Interface
	KubermaticEnterpriseV1() kubermaticenterprisev1.KubermaticEnterpriseV1Interface
	KubermaticV1() kubermaticv1.KubermaticV1Interface
}

// Clientset contains the clients for groups.
type Clientset struct {
	*discovery.DiscoveryClient
	kubermaticAppsV1           *kubermaticappsv1.KubermaticAppsV1Client
	kubermaticEnterpriseAppsV1 *kubermaticenterpriseappsv1.KubermaticEnterpriseAppsV1Client
	kubermaticEnterpriseV1     *kubermaticenterprisev1.KubermaticEnterpriseV1Client
	kubermaticV1               *kubermaticv1.KubermaticV1Client
}

// KubermaticAppsV1 retrieves the KubermaticAppsV1Client
func (c *Clientset) KubermaticAppsV1() kubermaticappsv1.KubermaticAppsV1Interface {
	return c.kubermaticAppsV1
}

// KubermaticEnterpriseAppsV1 retrieves the KubermaticEnterpriseAppsV1Client
func (c *Clientset) KubermaticEnterpriseAppsV1() kubermaticenterpriseappsv1.KubermaticEnterpriseAppsV1Interface {
	return c.kubermaticEnterpriseAppsV1
}

// KubermaticEnterpriseV1 retrieves the KubermaticEnterpriseV1Client
func (c *Clientset) KubermaticEnterpriseV1() kubermaticenterprisev1.KubermaticEnterpriseV1Interface {
	return c.kubermaticEnterpriseV1
}

// KubermaticV1 retrieves the KubermaticV1Client
func (c *Clientset) KubermaticV1() kubermaticv1.KubermaticV1Interface {
	return c.kubermaticV1
}

// Discovery retrieves the DiscoveryClient
func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	if c == nil {
		return nil
	}
	return c.DiscoveryClient
}

// NewForConfig creates a new Clientset for the given config.
// If config's RateLimiter is not set and QPS and Burst are acceptable,
// NewForConfig will generate a rate-limiter in configShallowCopy.
// NewForConfig is equivalent to NewForConfigAndClient(c, httpClient),
// where httpClient was generated with rest.HTTPClientFor(c).
func NewForConfig(c *rest.Config) (*Clientset, error) {
	configShallowCopy := *c

	if configShallowCopy.UserAgent == "" {
		configShallowCopy.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	// share the transport between all clients
	httpClient, err := rest.HTTPClientFor(&configShallowCopy)
	if err != nil {
		return nil, err
	}

	return NewForConfigAndClient(&configShallowCopy, httpClient)
}

// NewForConfigAndClient creates a new Clientset for the given config and http client.
// Note the http client provided takes precedence over the configured transport values.
// If config's RateLimiter is not set and QPS and Burst are acceptable,
// NewForConfigAndClient will generate a rate-limiter in configShallowCopy.
func NewForConfigAndClient(c *rest.Config, httpClient *http.Client) (*Clientset, error) {
	configShallowCopy := *c
	if configShallowCopy.RateLimiter == nil && configShallowCopy.QPS > 0 {
		if configShallowCopy.Burst <= 0 {
			return nil, fmt.Errorf("burst is required to be greater than 0 when RateLimiter is not set and QPS is set to greater than 0")
		}
		configShallowCopy.RateLimiter = flowcontrol.NewTokenBucketRateLimiter(configShallowCopy.QPS, configShallowCopy.Burst)
	}

	var cs Clientset
	var err error
	cs.kubermaticAppsV1, err = kubermaticappsv1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.kubermaticEnterpriseAppsV1, err = kubermaticenterpriseappsv1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.kubermaticEnterpriseV1, err = kubermaticenterprisev1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.kubermaticV1, err = kubermaticv1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}

	cs.DiscoveryClient, err = discovery.NewDiscoveryClientForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	return &cs, nil
}

// NewForConfigOrDie creates a new Clientset for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *Clientset {
	cs, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return cs
}

// New creates a new Clientset for the given RESTClient.
func New(c rest.Interface) *Clientset {
	var cs Clientset
	cs.kubermaticAppsV1 = kubermaticappsv1.New(c)
	cs.kubermaticEnterpriseAppsV1 = kubermaticenterpriseappsv1.New(c)
	cs.kubermaticEnterpriseV1 = kubermaticenterprisev1.New(c)
	cs.kubermaticV1 = kubermaticv1.New(c)

	cs.DiscoveryClient = discovery.NewDiscoveryClient(c)
	return &cs
}
