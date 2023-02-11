package main

import (
    "droute/winroute"
    "flag"
    "net"
    "sort"
    "strings"

    log "github.com/sirupsen/logrus"
)

func LocalAddr() ([]string, error) {
    address, err := net.InterfaceAddrs()
    if err != nil {
        return nil, err
    }
    IPs := make([]string, 0)
    for _, a := range address {
        if ipNet, ok := a.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
            if ipNet.IP.To4() != nil {
                IPs = append(IPs, ipNet.IP.To4().String())
            }
        }
    }
    return IPs, nil
}

func fromGatewayGetAddr(gateway string) (index int, ipAddr string) {
    networks, _ := net.Interfaces()
    for _, network := range networks {
        address, _ := network.Addrs()
        for _, addr := range address {
            ipNet, isIpNet := addr.(*net.IPNet)
            if isIpNet && !ipNet.IP.IsLoopback() {
                if ipNet.IP.To4() != nil {
                    ipAddr = ipNet.IP.String()
                    ipTmp1 := strings.Join(strings.Split(ipAddr, ".")[0:3], ".")
                    ipTmp2 := strings.Join(strings.Split(gateway, ".")[0:3], ".")
                    if ipTmp1 == ipTmp2 {
                        index = network.Index
                        return
                    }
                }
            }
        }
    }

    return
}

func main() {
    reverse := flag.Bool("r", false, "enable debug logging")
    flag.Parse()
    var gateway []string
    for _, q := range winroute.WbemQuery() {
        if len(q.DefaultIPGateway) == 0 {
            continue
        }
        gateway = append(gateway, q.DefaultIPGateway[0])
    }

    address, err := LocalAddr()
    if err != nil {
        log.Fatal(err)
    }
    for _, addr := range address {
        if strings.HasPrefix(addr, "169.254") {
            log.Fatalf("DHCP server is not reachable: %s", addr)
        }
    }

    if len(address) != len(gateway) {
        log.Fatalf("The number of addresses and gateways is different. try reboot")
    }
    if len(gateway) == 1 {
        log.Fatalf("one gateway and address are found. no need set")
    }

    r := winroute.NewNetRoute()
    defer r.Close()
    _ = r.ResetRoute()

    if *reverse {
        Reverse(gateway)
    }
    for gwIndex, gw := range gateway {
        index, _ := fromGatewayGetAddr(gw)
        i, err := r.GetInterfaceByIndex(uint32(index))
        if err != nil {
            log.Fatal(err)
        }
        //fmt.Println(i.InterfaceIndex, gw, i.Metric)
        r1 := &winroute.IPForwardRow{
            ForwardDest:    winroute.Inet_aton("0.0.0.0", false),
            ForwardMask:    winroute.Inet_aton("0.0.0.0", false),
            ForwardNextHop: winroute.Inet_aton(gw, false),
            ForwardIfIndex: i.InterfaceIndex,
            ForwardType:    4,
            ForwardProto:   3,
            ForwardMetric1: i.Metric + uint32(gwIndex),
        }

        if err = r.AddRoute(r1); err == nil {
            log.Infof("Added route: %s", gw)
        } else {
            log.Warnf("Error adding route: %s", gw)
        }
    }

}

func printRoutes(r *winroute.NetRoute) {
    routes, err := r.GetRoutes()
    if err != nil {
        log.Error("Error getting routes")
    }
    for _, route := range routes {
        dest := winroute.Inet_ntoa(route.ForwardDest, false)
        mask := winroute.Inet_ntoa(route.ForwardMask, false)
        gate := winroute.Inet_ntoa(route.ForwardNextHop, false)
        log.WithFields(log.Fields{
            "dest":    dest,
            "mask":    mask,
            "gate":    gate,
            "Proto":   route.ForwardProto,
            "Type":    route.ForwardType,
            "metric":  route.ForwardMetric1,
            "ifIndex": route.ForwardIfIndex,
        }).Info("Route")
    }
}

func Reverse(s interface{}) {
    sort.SliceStable(s, func(i, j int) bool {
        return true
    })
}
