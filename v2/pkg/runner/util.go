package runner

import (
	"fmt"
	"os"
	"runtime"

	"github.com/projectdiscovery/gologger"
	"github.com/projectdiscovery/iputil"
)

func isRoot() bool {
	return os.Geteuid() == 0
}

func (r *Runner) host2ips(target string) (targetIPs []string, err error) {
	// If the host is a Domain, then perform resolution and discover all IP
	// addresses for a given host. Else use that host for port scanning
	if !iputil.IsIP(target) {
		var ips []string
		ips, err = r.dnsclient.Lookup(target)
		if err != nil {
			gologger.Warning().Msgf("Could not get IP for host: %s\n", target)
			return
		}
		for _, ip := range ips {
			if iputil.IsIPv4(ip) {
				targetIPs = append(targetIPs, ip)
			}
		}

		if len(targetIPs) == 0 {
			return targetIPs, fmt.Errorf("no IP addresses found for host: %s", target)
		}
	} else {
		targetIPs = append(targetIPs, target)
		gologger.Debug().Msgf("Found %d addresses for %s\n", len(targetIPs), target)
	}

	return
}

func isOSSupported() bool {
	return runtime.GOOS == "linux" || runtime.GOOS == "darwin"
}

func mapKeysToSliceInt(m map[int]struct{}) (s []int) {
	for k := range m {
		s = append(s, k)
	}
	return
}
