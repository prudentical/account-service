package discovery

import (
	"account-service/configuration"
	"fmt"
	"os"

	"github.com/hashicorp/consul/api"
)

type service struct {
	name        string
	host        string
	healthcheck string
	port        int
}

type ServiceDiscovery interface {
	Register() error
	Discover(name string) (string, error)
}

type consulServiceDiscovery struct {
	service service
	client  *api.Client
}

func NewServiceDiscovery(config configuration.Config) ServiceDiscovery {
	consulConfig := api.DefaultConfig()
	address := fmt.Sprintf("%s:%d", config.Discovery.Server.Host, config.Discovery.Server.Port)
	consulConfig.Address = address

	client, err := api.NewClient(consulConfig)

	if err != nil {
		panic(err)
	}
	return consulServiceDiscovery{
		service: service{
			name:        config.App.Name,
			host:        config.Server.Host,
			healthcheck: config.Server.Healthcheck,
			port:        config.Server.Port,
		},
		client: client,
	}
}

func (c consulServiceDiscovery) Register() error {
	host := c.service.host
	if host == "" {
		hostname, err := os.Hostname()
		if err != nil {
			return err
		}
		host = hostname
	}
	id := fmt.Sprintf("%s::%s:%d", c.service.name, host, c.service.port)

	service := &api.AgentServiceRegistration{
		ID:      id,
		Name:    c.service.name,
		Address: host,
		Port:    c.service.port,
		Check: &api.AgentServiceCheck{
			HTTP:     fmt.Sprintf("http://%s:%v/health", c.service.healthcheck, c.service.port),
			Interval: "10s",
			Timeout:  "3s",
		},
	}
	return c.client.Agent().ServiceRegister(service)
}

func (c consulServiceDiscovery) Discover(name string) (string, error) {
	services, _, err := c.client.Catalog().Service(name, "", nil)
	if err != nil {
		return "", err
	}
	result := fmt.Sprintf("%s:%d", services[0].ServiceAddress, services[0].ServicePort)
	return result, nil
}
