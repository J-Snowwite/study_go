package winroute

import (
    "github.com/StackExchange/wmi"
    log "github.com/sirupsen/logrus"
    "time"
)

func WbemQuery() []Win32_NetworkAdapterConfiguration {
    s, err := wmi.InitializeSWbemServices(wmi.DefaultClient)
    if err != nil {
        log.Fatalf("InitializeSWbemServices: %s", err)
    }

    var dst []Win32_NetworkAdapterConfiguration
    q := wmi.CreateQuery(&dst, "WHERE IPEnabled=True")
    errQuery := s.Query(q, &dst)
    if errQuery != nil {
        log.Fatalf("Query1: %s", errQuery)
    }
    return dst
}

type Win32_NetworkAdapterConfiguration struct {
    Caption                      string
    Description                  string
    SettingID                    string
    ArpAlwaysSourceRoute         bool
    ArpUseEtherSNAP              bool
    DatabasePath                 string
    DeadGWDetectEnabled          bool
    DefaultIPGateway             []string
    DefaultTOS                   uint8
    DefaultTTL                   uint8
    DHCPEnabled                  bool
    DHCPLeaseExpires             *time.Time
    DHCPLeaseObtained            *time.Time
    DHCPServer                   string
    DNSDomain                    string
    DNSDomainSuffixSearchOrder   []string
    DNSEnabledForWINSResolution  bool
    DNSHostName                  string
    DNSServerSearchOrder         []string
    DomainDNSRegistrationEnabled bool
    ForwardBufferMemory          uint32
    FullDNSRegistrationEnabled   bool
    GatewayCostMetric            []int32
    IGMPLevel                    uint8
    Index                        uint32
    InterfaceIndex               uint32
    IPAddress                    []string
    IPConnectionMetric           uint32
    IPEnabled                    bool
    IPFilterSecurityEnabled      bool
    IPPortSecurityEnabled        bool
    IPSecPermitIPProtocols       []string
    IPSecPermitTCPPorts          []string
    IPSecPermitUDPPorts          []string
    IPSubnet                     []string
    IPUseZeroBroadcast           bool
    IPXAddress                   string
    IPXEnabled                   bool
    IPXFrameType                 []uint32
    IPXMediaType                 uint32
    IPXNetworkNumber             []string
    IPXVirtualNetNumber          string
    KeepAliveInterval            uint32
    KeepAliveTime                uint32
    MACAddress                   string
    MTU                          uint32
    NumForwardPackets            uint32
    PMTUBHDetectEnabled          bool
    PMTUDiscoveryEnabled         bool
    ServiceName                  string
    TcpipNetbiosOptions          uint32
    TcpMaxConnectRetransmissions uint32
    TcpMaxDataRetransmissions    uint32
    TcpNumConnections            uint32
    TcpUseRFC1122UrgentPointer   bool
    TcpWindowSize                uint16
    WINSEnableLMHostsLookup      bool
    WINSHostLookupFile           string
    WINSPrimaryServer            string
    WINSScopeID                  string
    WINSSecondaryServer          string
}
