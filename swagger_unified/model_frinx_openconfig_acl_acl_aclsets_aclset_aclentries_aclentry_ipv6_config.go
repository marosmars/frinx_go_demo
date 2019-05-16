/*
 * unified
 *
 * API generated from yang definitions: [fake,frinx-acl-extension,frinx-bfd,frinx-bfd-extension,frinx-bgp-extension,frinx-cdp,frinx-cisco-if-extension,frinx-cisco-mpls-te-extension,frinx-cisco-ospf-extension,frinx-cisco-pf-interfaces-extension,frinx-damping,frinx-dasan-vlan-extension,frinx-event-types,frinx-evpn,frinx-evpn-types,frinx-hsrp,frinx-if-aggregate-extension,frinx-isis-extension,frinx-juniper-if-aggregate-extension,frinx-juniper-if-extension,frinx-juniper-pf-interfaces-extension,frinx-juniper-probes-extension,frinx-l3ipvlan,frinx-lacp-lag-member,frinx-logging,frinx-mpls-ldp-extension,frinx-mpls-rsvp-extension,frinx-netflow,frinx-openconfig-acl,frinx-openconfig-bfd,frinx-openconfig-bgp,frinx-openconfig-bgp-policy,frinx-openconfig-bgp-types,frinx-openconfig-extensions,frinx-openconfig-if-aggregate,frinx-openconfig-if-ethernet,frinx-openconfig-if-ip,frinx-openconfig-if-ip-ext,frinx-openconfig-inet-types,frinx-openconfig-interfaces,frinx-openconfig-isis,frinx-openconfig-isis-lsdb-types,frinx-openconfig-isis-policy,frinx-openconfig-isis-types,frinx-openconfig-lacp,frinx-openconfig-lldp,frinx-openconfig-lldp-types,frinx-openconfig-local-routing,frinx-openconfig-mpls,frinx-openconfig-mpls-ldp,frinx-openconfig-mpls-rsvp,frinx-openconfig-mpls-types,frinx-openconfig-network-instance,frinx-openconfig-network-instance-l3,frinx-openconfig-network-instance-policy,frinx-openconfig-network-instance-types,frinx-openconfig-ospf-policy,frinx-openconfig-ospf-types,frinx-openconfig-ospfv2,frinx-openconfig-packet-match,frinx-openconfig-packet-match-types,frinx-openconfig-platform,frinx-openconfig-platform-linecard,frinx-openconfig-platform-port,frinx-openconfig-platform-transceiver,frinx-openconfig-platform-types,frinx-openconfig-policy-forwarding,frinx-openconfig-policy-types,frinx-openconfig-probes,frinx-openconfig-probes-types,frinx-openconfig-qos,frinx-openconfig-qos-types,frinx-openconfig-rib-bgp,frinx-openconfig-rib-bgp-ext,frinx-openconfig-rib-bgp-types,frinx-openconfig-routing-policy,frinx-openconfig-transport-types,frinx-openconfig-types,frinx-openconfig-vlan,frinx-openconfig-vlan-types,frinx-openconfig-yang-types,frinx-ospf-extension,frinx-qos-extension,frinx-snmp,frinx-uniconfig-topology,iana-if-type,ietf-inet-types,ietf-interfaces,ietf-yang-types,network-topology,network-topology,unified-topology,yang-ext]
 *
 * API version: 4.2.0.frinx
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger_unified

// Optional[Configuration data for IPv6 match fields] REF:Optional.empty
type FrinxOpenconfigAclAclAclsetsAclsetAclentriesAclentryIpv6Config struct {
	// Optional[The IP packet's hop limit -- known as TTL (in hops) in IPv4 packets, and hop limit in IPv6] REF:Optional.empty
	HopLimit int32 `json:"hop-limit,omitempty"`
	// Optional[Destination IPv6 Flow label.] REF:Optional.empty
	DestinationFlowLabel int64 `json:"destination-flow-label,omitempty"`
	// Optional[The protocol carried in the IP packet, expressed either as its IP protocol number, or by a defined identity.] REF:Optional.empty
	Protocol string `json:"protocol,omitempty"`
	// Optional[Value of diffserv codepoint.] REF:Optional.empty
	Dscp int32 `json:"dscp,omitempty"`
	// Optional.empty REF:Optional.empty
	FrinxAclExtensionsourceAddressWildcarded *FrinxAclExtensionSrcdstipv6addresswildcardedSourceAddressWildcarded `json:"frinx-acl-extension:source-address-wildcarded,omitempty"`
	// Optional[The IP packet's hop range (TTL)] REF:Optional.empty
	FrinxAclExtensionhopRange string `json:"frinx-acl-extension:hop-range,omitempty"`
	// Optional[Source IPv6 Flow label.] REF:Optional.empty
	SourceFlowLabel int64 `json:"source-flow-label,omitempty"`
	// Optional.empty REF:Optional.empty
	FrinxAclExtensiondestinationAddressWildcarded *FrinxAclExtensionSrcdstipv6addresswildcardedDestinationAddressWildcarded `json:"frinx-acl-extension:destination-address-wildcarded,omitempty"`
	// Optional[Destination IPv6 address prefix.] REF:Optional.empty
	DestinationAddress string `json:"destination-address,omitempty"`
	// Optional[Source IPv6 address prefix.] REF:Optional.empty
	SourceAddress string `json:"source-address,omitempty"`
}
