/*
 * unified
 *
 * API generated from yang definitions: [fake,frinx-acl-extension,frinx-bfd,frinx-bfd-extension,frinx-bgp-extension,frinx-cdp,frinx-cisco-if-extension,frinx-cisco-mpls-te-extension,frinx-cisco-ospf-extension,frinx-cisco-pf-interfaces-extension,frinx-damping,frinx-dasan-vlan-extension,frinx-event-types,frinx-evpn,frinx-evpn-types,frinx-hsrp,frinx-if-aggregate-extension,frinx-isis-extension,frinx-juniper-if-aggregate-extension,frinx-juniper-if-extension,frinx-juniper-pf-interfaces-extension,frinx-juniper-probes-extension,frinx-l3ipvlan,frinx-lacp-lag-member,frinx-logging,frinx-mpls-ldp-extension,frinx-mpls-rsvp-extension,frinx-netflow,frinx-openconfig-acl,frinx-openconfig-bfd,frinx-openconfig-bgp,frinx-openconfig-bgp-policy,frinx-openconfig-bgp-types,frinx-openconfig-extensions,frinx-openconfig-if-aggregate,frinx-openconfig-if-ethernet,frinx-openconfig-if-ip,frinx-openconfig-if-ip-ext,frinx-openconfig-inet-types,frinx-openconfig-interfaces,frinx-openconfig-isis,frinx-openconfig-isis-lsdb-types,frinx-openconfig-isis-policy,frinx-openconfig-isis-types,frinx-openconfig-lacp,frinx-openconfig-lldp,frinx-openconfig-lldp-types,frinx-openconfig-local-routing,frinx-openconfig-mpls,frinx-openconfig-mpls-ldp,frinx-openconfig-mpls-rsvp,frinx-openconfig-mpls-types,frinx-openconfig-network-instance,frinx-openconfig-network-instance-l3,frinx-openconfig-network-instance-policy,frinx-openconfig-network-instance-types,frinx-openconfig-ospf-policy,frinx-openconfig-ospf-types,frinx-openconfig-ospfv2,frinx-openconfig-packet-match,frinx-openconfig-packet-match-types,frinx-openconfig-platform,frinx-openconfig-platform-linecard,frinx-openconfig-platform-port,frinx-openconfig-platform-transceiver,frinx-openconfig-platform-types,frinx-openconfig-policy-forwarding,frinx-openconfig-policy-types,frinx-openconfig-probes,frinx-openconfig-probes-types,frinx-openconfig-qos,frinx-openconfig-qos-types,frinx-openconfig-rib-bgp,frinx-openconfig-rib-bgp-ext,frinx-openconfig-rib-bgp-types,frinx-openconfig-routing-policy,frinx-openconfig-transport-types,frinx-openconfig-types,frinx-openconfig-vlan,frinx-openconfig-vlan-types,frinx-openconfig-yang-types,frinx-ospf-extension,frinx-qos-extension,frinx-snmp,frinx-uniconfig-topology,iana-if-type,ietf-inet-types,ietf-interfaces,ietf-yang-types,network-topology,network-topology,unified-topology,yang-ext]
 *
 * API version: 4.2.0.frinx
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger_unified

// Optional[Configuration data for ethernet interfaces] REF:Optional.empty
type FrinxOpenconfigIfEthernetEthernettopEthernetConfig struct {
	// Optional[Member port of LACP has key value. All member ports in one aggregator have same key values. To make the aggregator consisted of specified member ports, configure the different key value with the key value of another port.] REF:Optional.empty
	FrinxIfAggregateExtensionadminKey int32 `json:"frinx-if-aggregate-extension:admin-key,omitempty"`
	// Optional[Assigns a MAC address to the Ethernet interface.  If not specified, the corresponding operational state leaf is expected to show the system-assigned MAC address.] REF:Optional.empty
	MacAddress string `json:"mac-address,omitempty"`
	// Optional[When auto-negotiate is TRUE, this optionally sets the port-speed mode that will be advertised to the peer for negotiation.  If unspecified, it is expected that the interface will select the highest speed available based on negotiation.  When auto-negotiate is set to FALSE, sets the link speed to a fixed value -- supported values are defined by ETHERNET_SPEED identities] REF:Optional.empty
	PortSpeed *FrinxOpenconfigIfEthernetPortSpeedIdentityref `json:"port-speed,omitempty"`
	// Optional[Enable or disable flow control for this interface. Ethernet flow control is a mechanism by which a receiver may send PAUSE frames to a sender to stop transmission for a specified time.  This setting should override auto-negotiated flow control settings.  If left unspecified, and auto-negotiate is TRUE, flow control mode is negotiated with the peer interface.] REF:Optional[IEEE 802.3x]
	EnableFlowControl bool `json:"enable-flow-control,omitempty"`
	// Optional[Set the period between LACP messages -- uses the lacp-period-type enumeration.] REF:Optional.empty
	FrinxLacpLagMemberinterval *FrinxOpenconfigLacpLacpPeriodType `json:"frinx-lacp-lag-member:interval,omitempty"`
	// Optional[ACTIVE is to initiate the transmission of LACP packets. PASSIVE is to wait for peer to initiate the transmission of LACP packets.] REF:Optional.empty
	FrinxLacpLagMemberlacpMode *FrinxOpenconfigLacpLacpActivityType `json:"frinx-lacp-lag-member:lacp-mode,omitempty"`
	// Optional[Specify the logical aggregate interface to which this interface belongs] REF:Optional.empty
	FrinxOpenconfigIfAggregateaggregateId string `json:"frinx-openconfig-if-aggregate:aggregate-id,omitempty"`
	// Optional[Set to TRUE to request the interface to auto-negotiate transmission parameters with its peer interface.  When set to FALSE, the transmission parameters are specified manually.] REF:Optional[IEEE 802.3-2012 auto-negotiation transmission parameters]
	AutoNegotiate bool `json:"auto-negotiate,omitempty"`
	// Optional[When auto-negotiate is TRUE, this optionally sets the duplex mode that will be advertised to the peer.  If unspecified, the interface should negotiate the duplex mode directly (typically full-duplex).  When auto-negotiate is FALSE, this sets the duplex mode on the interface directly.] REF:Optional.empty
	DuplexMode *FrinxOpenconfigIfEthernetDuplexModeEnumeration `json:"duplex-mode,omitempty"`
}