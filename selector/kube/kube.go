// Package kube is a selector which always returns the name specified
// (dots replaced with dashes) with a port-number appended.
// AN optional domain-name will also be added.
package kube

import (
	"fmt"
	"os"

	"github.com/micro/go-micro/cmd"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/selector"
)

const (
	ENV_KUBE_SELECTOR_DOMAIN_NAME = "KUBE_SELECTOR_DOMAIN_NAME"
	ENV_KUBE_SELECTOR_PORT_NUMBER = "KUBE_SELECTOR_PORT_NUMBER"
	DEFAULT_PORT_NUMBER             = "8080"
)

type kubeSelector struct {
	addressSuffix string
	envDomainName string
	envPortNumber string
}

func init() {
	cmd.DefaultSelectors["kube"] = NewSelector
}

func (s *kubeSelector) Init(opts ...selector.Option) error {
	return nil
}

func (s *kubeSelector) Options() selector.Options {
	return selector.Options{}
}

func (s *kubeSelector) Select(service string, opts ...selector.SelectOption) (selector.Next, error) {
	// If your go-micro service is micro.service.greeter, your kubernetes service
	// name should be micro-service-greeter
	service = strings.Replace(service, ".", "-", -1)
	node := &registry.Node{
		Id:      service,
		Address: fmt.Sprintf("%v%v", service, s.addressSuffix),
	}

	return func() (*registry.Node, error) {
		return node, nil
	}, nil
}

func (s *kubeSelector) Mark(service string, node *registry.Node, err error) {
	return
}

func (s *kubeSelector) Reset(service string) {
	return
}

func (s *kubeSelector) Close() error {
	return nil
}

func (s *kubeSelector) String() string {
	return "kube"
}

func NewSelector(opts ...selector.Option) selector.Selector {

	// Build a new
	s := &kubeSelector{
		addressSuffix: "",
		envDomainName: os.Getenv(ENV_KUBE_SELECTOR_DOMAIN_NAME),
		envPortNumber: os.Getenv(ENV_KUBE_SELECTOR_PORT_NUMBER),
	}

	// Add the dns domain-name (if one was specified by an env-var):
	if s.envDomainName != "" {
		s.addressSuffix += fmt.Sprintf(".%v", s.envDomainName)
	}

	// Either add the default port-number, or override with one specified by an env-var:
	if s.envPortNumber == "" {
		s.addressSuffix += fmt.Sprintf(":%v", DEFAULT_PORT_NUMBER)
	} else {
		s.addressSuffix += fmt.Sprintf(":%v", s.envPortNumber)
	}

	return s
}
