package service_discovery

import (
	"fmt"
	"log"
	"net/http"

	consulapi "github.com/hashicorp/consul/api"
)

type ServiceName string

const (
	ServiceName_USER_SERVICE         ServiceName = "USER_SERVICE"
	ServiceName_PARCEL_SERVICE       ServiceName = "PARCEL_SERVICE"
	ServiceName_NOTIFICATION_SERVICE ServiceName = "NOTIFICATION_SERVICE"
)

type ServiceDiscoveryClient struct {
	Client *consulapi.Client
}

func check(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Consul check")
}

func GetServiceDiscoveryClient() *ServiceDiscoveryClient {
	config := consulapi.DefaultConfig()
	consul, err := consulapi.NewClient(config)
	if err != nil {
		log.Println(err)
	}
	return &ServiceDiscoveryClient{Client: consul}
}

func (c *ServiceDiscoveryClient) ServiceRegistry(serviceID string, address string, port int, hbPort int) {

	http.HandleFunc("/check", check)
	hbPortString := fmt.Sprintf("%v", hbPort)
	http.ListenAndServe(string(hbPortString), nil)

	registration := &consulapi.AgentServiceRegistration{
		ID:      serviceID,
		Name:    serviceID,
		Port:    port,
		Address: address,
		Check: &consulapi.AgentServiceCheck{
			HTTP:     fmt.Sprintf("http://%s:%v/check", address, hbPort),
			Interval: "10s",
			Timeout:  "30s",
		},
	}

	regisErr := c.Client.Agent().ServiceRegister(registration)

	if regisErr != nil {
		log.Printf("Failed to register service: %s:%v ", address, port)
	} else {
		log.Printf("successfully register service: %s:%v", address, port)
	}

}

func (c *ServiceDiscoveryClient) ServiceDiscovery(serviceID string) string {
	services, error := c.Client.Agent().Services()
	if error != nil {
		fmt.Println(error)
	}

	service := services[serviceID]
	address := service.Address
	port := service.Port

	url := fmt.Sprintf("%s:%v", address, port)
	log.Println("Discovered", serviceID, url)

	return url
}
