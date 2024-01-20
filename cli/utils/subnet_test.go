package utils

import (
	"net"
	"reflect"
	"testing"
)

func mustParseCIDR(cidrStr string) *net.IPNet {
	_, net, _ := net.ParseCIDR(cidrStr)
	return net
}

func TestDivideSubnet(t *testing.T) {
	tests := []struct {
		name           string
		network        string
		newPrefix      int
		expectedResult []*net.IPNet
		expectedError  bool
	}{
		{
			name:      "Valid split /24 into /25",
			network:   "192.168.0.0/24",
			newPrefix: 25,
			expectedResult: []*net.IPNet{
				mustParseCIDR("192.168.0.0/25"),
				mustParseCIDR("192.168.0.128/25"),
			},
			expectedError: false,
		},
		{
			name:           "Invalid split, new prefix smaller",
			network:        "192.168.0.0/24",
			newPrefix:      23,
			expectedResult: nil,
			expectedError:  true,
		},
		{
			name:      "Valid split /24 into /25",
			network:   "192.168.0.0/24",
			newPrefix: 25,
			expectedResult: []*net.IPNet{
				mustParseCIDR("192.168.0.0/25"),
				mustParseCIDR("192.168.0.128/25"),
			},
			expectedError: false,
		},
		{
			name:           "Invalid split, new prefix smaller",
			network:        "192.168.0.0/24",
			newPrefix:      23,
			expectedResult: nil,
			expectedError:  true,
		},
		{
			name:      "Valid split /24 into /26",
			network:   "192.168.0.0/24",
			newPrefix: 26,
			expectedResult: []*net.IPNet{
				mustParseCIDR("192.168.0.0/26"),
				mustParseCIDR("192.168.0.64/26"),
				mustParseCIDR("192.168.0.128/26"),
				mustParseCIDR("192.168.0.192/26"),
			},
			expectedError: false,
		},
		{
			name:           "Invalid network format",
			network:        "192.168.0.256/24", // Invalid IP
			newPrefix:      25,
			expectedResult: nil,
			expectedError:  true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			_, network, _ := net.ParseCIDR(tc.network)

			result, err := DivideSubnet4(network, tc.newPrefix)
			if (err != nil) != tc.expectedError {
				t.Errorf("Test '%s' failed: expected error status %v, got %v", tc.name, tc.expectedError, err != nil)
			}

			if !reflect.DeepEqual(result, tc.expectedResult) {
				t.Errorf("Test '%s' failed: expected result %v, got %v", tc.name, tc.expectedResult, result)
			}
		})
	}
}
