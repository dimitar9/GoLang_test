
/* ResolveIP
 */

package main

import (
	"net"
	"os"
	"fmt"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s hostname\n", os.Args[0])
		fmt.Println("Usage: ", os.Args[0], "hostname")
		os.Exit(1)
	}
	name := os.Args[1]

	addr, err := net.ResolveIPAddr("ip6", name)// #the first argument can be ip, ip4, ip6
	if err != nil {
		fmt.Println("Resolution error", err.Error())
		os.Exit(1)
	}
	fmt.Println("Resolved address is ", addr.String())
	os.Exit(0)
}
