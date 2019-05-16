/*
 * unified
 *
 * API generated from yang definitions: [fake,frinx-acl-extension,frinx-bfd,frinx-bfd-extension,frinx-bgp-extension,frinx-cdp,frinx-cisco-if-extension,frinx-cisco-mpls-te-extension,frinx-cisco-ospf-extension,frinx-cisco-pf-interfaces-extension,frinx-damping,frinx-dasan-vlan-extension,frinx-event-types,frinx-evpn,frinx-evpn-types,frinx-hsrp,frinx-if-aggregate-extension,frinx-isis-extension,frinx-juniper-if-aggregate-extension,frinx-juniper-if-extension,frinx-juniper-pf-interfaces-extension,frinx-juniper-probes-extension,frinx-l3ipvlan,frinx-lacp-lag-member,frinx-logging,frinx-mpls-ldp-extension,frinx-mpls-rsvp-extension,frinx-netflow,frinx-openconfig-acl,frinx-openconfig-bfd,frinx-openconfig-bgp,frinx-openconfig-bgp-policy,frinx-openconfig-bgp-types,frinx-openconfig-extensions,frinx-openconfig-if-aggregate,frinx-openconfig-if-ethernet,frinx-openconfig-if-ip,frinx-openconfig-if-ip-ext,frinx-openconfig-inet-types,frinx-openconfig-interfaces,frinx-openconfig-isis,frinx-openconfig-isis-lsdb-types,frinx-openconfig-isis-policy,frinx-openconfig-isis-types,frinx-openconfig-lacp,frinx-openconfig-lldp,frinx-openconfig-lldp-types,frinx-openconfig-local-routing,frinx-openconfig-mpls,frinx-openconfig-mpls-ldp,frinx-openconfig-mpls-rsvp,frinx-openconfig-mpls-types,frinx-openconfig-network-instance,frinx-openconfig-network-instance-l3,frinx-openconfig-network-instance-policy,frinx-openconfig-network-instance-types,frinx-openconfig-ospf-policy,frinx-openconfig-ospf-types,frinx-openconfig-ospfv2,frinx-openconfig-packet-match,frinx-openconfig-packet-match-types,frinx-openconfig-platform,frinx-openconfig-platform-linecard,frinx-openconfig-platform-port,frinx-openconfig-platform-transceiver,frinx-openconfig-platform-types,frinx-openconfig-policy-forwarding,frinx-openconfig-policy-types,frinx-openconfig-probes,frinx-openconfig-probes-types,frinx-openconfig-qos,frinx-openconfig-qos-types,frinx-openconfig-rib-bgp,frinx-openconfig-rib-bgp-ext,frinx-openconfig-rib-bgp-types,frinx-openconfig-routing-policy,frinx-openconfig-transport-types,frinx-openconfig-types,frinx-openconfig-vlan,frinx-openconfig-vlan-types,frinx-openconfig-yang-types,frinx-ospf-extension,frinx-qos-extension,frinx-snmp,frinx-uniconfig-topology,iana-if-type,ietf-inet-types,ietf-interfaces,ietf-yang-types,network-topology,network-topology,unified-topology,yang-ext]
 *
 * API version: 4.2.0.frinx
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger_unified

// Optional[Configuration data for the VRRP group] REF:Optional.empty
type FrinxOpenconfigIfIpIpvrrptopVrrpVrrpgroupConfig struct {
	// Optional[Set the delay the higher priority router waits before preempting] REF:Optional.empty
	PreemptDelay int32 `json:"preempt-delay,omitempty"`
	// Optional[Configure one or more virtual addresses for the VRRP group] REF:Optional.empty
	VirtualAddress []string `json:"virtual-address,omitempty"`
	// Optional[Sets the interval between successive VRRP advertisements -- RFC 5798 defines this as a 12-bit value expressed as 0.1 seconds, with default 100, i.e., 1 second.  Several implementation express this in units of seconds] REF:Optional.empty
	AdvertisementInterval int32 `json:"advertisement-interval,omitempty"`
	// Optional[Set the virtual router id for use by the VRRP group.  This usually also determines the virtual MAC address that is generated for the VRRP group] REF:Optional.empty
	VirtualRouterId int32 `json:"virtual-router-id,omitempty"`
	// Optional[For VRRP on IPv6 interfaces, sets the virtual link local address] REF:Optional.empty
	FrinxOpenconfigIfIpvirtualLinkLocal string `json:"frinx-openconfig-if-ip:virtual-link-local,omitempty"`
	// Optional[Specifies the sending VRRP interface's priority for the virtual router.  Higher values equal higher priority] REF:Optional.empty
	Priority int32 `json:"priority,omitempty"`
	// Optional[Configure whether packets destined for virtual addresses are accepted even when the virtual address is not owned by the router interface] REF:Optional.empty
	AcceptMode bool `json:"accept-mode,omitempty"`
	// Optional[When set to true, enables preemption by a higher priority backup router of a lower priority master router] REF:Optional.empty
	Preempt bool `json:"preempt,omitempty"`
}