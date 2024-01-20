package utils

import (
	"fmt"
	"net"

	"github.com/apparentlymart/go-cidr/cidr"
)

func DivideSubnet4(network *net.IPNet, newPrefix int) ([]*net.IPNet, error) {
	if network == nil {
		return nil, fmt.Errorf("undefined network")
	}

	if network.IP.To4() == nil {
		return nil, fmt.Errorf("network is not an IPv4 subnet")
	}

	if newPrefix < 0 || newPrefix > 32 {
		return nil, fmt.Errorf("new prefix must be between 0 and 32")
	}

	oldPrefixSize, _ := network.Mask.Size()

	if newPrefix <= oldPrefixSize {
		return nil, fmt.Errorf("new prefix must be larger than the original prefix")
	}

	subnetCount := 1 << uint(newPrefix-oldPrefixSize) // 2^(newPrefix−oldPrefixSize)

	var subnets []*net.IPNet

	for i := 0; i < subnetCount; i++ {
		subnet, err := cidr.Subnet(network, newPrefix-oldPrefixSize, i)
		if err != nil {
			return nil, err
		}

		subnets = append(subnets, subnet)
	}

	return subnets, nil
}
