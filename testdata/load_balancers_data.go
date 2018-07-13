package testdata

import "github.com/ingrammicro/concerto/api/types"

// GetLoadBalancerData loads test data
func GetLoadBalancerData() *[]types.LoadBalancer {

	testLoadBalancers := []types.LoadBalancer{
		{
			ID:                       "fakeID0",
			Name:                     "fakeName0",
			Fqdn:                     "fakeFqdn0",
			Protocol:                 "fakeProtocol0",
			Port:                     1234,
			Algorithm:                "fakeAlgorithm0",
			SSLCertificate:           "fakeSSLCertificate0",
			SSLCertificatePrivateKey: "fakeSSLCertificatePrivateKey0",
			DomainID:                 "fakeDomainID0",
			CloudProviderID:          "fakeCloudProviderID0",
			TrafficIn:                1024,
			TrafficOut:               2048,
		},
		{
			ID:                       "fakeID1",
			Name:                     "fakeName1",
			Fqdn:                     "fakeFqdn1",
			Protocol:                 "fakeProtocol1",
			Port:                     1235,
			Algorithm:                "fakeAlgorithm1",
			SSLCertificate:           "fakeSSLCertificate1",
			SSLCertificatePrivateKey: "fakeSSLCertificatePrivateKey1",
			DomainID:                 "fakeDomainID1",
			CloudProviderID:          "fakeCloudProviderID1",
			TrafficIn:                10240,
			TrafficOut:               20480,
		},
	}

	return &testLoadBalancers
}

// GetLoadBalancerRecordData loads test data
func GetLBNodeData() *[]types.LBNode {

	testLBNode := []types.LBNode{
		{
			ID:       "fakeID0.0",
			Name:     "fakeName0.0",
			PublicIP: "fakePublicIP0",
			State:    "fakeState0",
			ServerID: "fakeServerID0",
			Port:     1234,
		},
		{
			ID:       "fakeID0.1",
			Name:     "fakeName0.1",
			PublicIP: "fakePublicIP1",
			State:    "fakeState1",
			ServerID: "fakeServerID1",
			Port:     1235,
		},
		{
			ID:       "fakeID0.2",
			Name:     "fakeName0.2",
			PublicIP: "fakePublicIP2",
			State:    "fakeState2",
			ServerID: "fakeServerID2",
			Port:     1236,
		},
	}

	return &testLBNode
}
