package main

import (
	"fmt"
	"subnetCalculation/subnet"
)

func main() {
	sub := subnet.SubnetCalculator("10.196.199.0", 23)
	fmt.Println(sub)

	fmt.Println(sub.GetNumberIPAddresses())      // 512
	fmt.Println(sub.GetNumberAddressableHosts()) // 510
	fmt.Println(sub.GetIPAddressRange())         // [ 192.168.112.0, 192.168.113.255 ]
	fmt.Println(sub.GetNetworkSize())            // 23
	fmt.Println(sub.GetBroadcastAddress())       // 192.168.113.255

	fmt.Println(sub.GetIPAddress())       // 192.168.112.203
	fmt.Println(sub.GetIPAddressQuads())  // [ 192, 168, 112, 203 ]
	fmt.Println(sub.GetIPAddressHex())    // C0A870CB
	fmt.Println(sub.GetIPAddressBinary()) // 11000000101010000111000011001011

	fmt.Println(sub.GetSubnetMask())       // 255.255.254.0
	fmt.Println(sub.GetSubnetMaskQuards()) // [ 255, 255, 254, 0 ]
	fmt.Println(sub.GetSubnetMaskHex())    // FFFFFE00
	fmt.Println(sub.GetSubnetMaskBinary()) // 11111111111111111111111000000000

	fmt.Println(sub.GetNetworkPortion())       // 192.168.112.0
	fmt.Println(sub.GetNetworkPortionQuards()) // [ 192, 168, 112, 0 ]
	fmt.Println(sub.GetNetworkPortionHex())    // C0A87000
	fmt.Println(sub.GetNetworkPortionBinary()) // 11000000101010000111000000000000

	fmt.Println(sub.GetHostPortion())       // 0.0.0.203
	fmt.Println(sub.GetHostPortionQuards()) // [ 0, 0, 0, 203 ]
	fmt.Println(sub.GetHostPortionHex())    // 000000CB
	fmt.Println(sub.GetHostPortionBinary()) // 00000000000000000000000011001011
}
