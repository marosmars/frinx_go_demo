/*
 * unified
 *
 * API generated from yang definitions: [fake,frinx-acl-extension,frinx-bfd,frinx-bfd-extension,frinx-bgp-extension,frinx-cdp,frinx-cisco-if-extension,frinx-cisco-mpls-te-extension,frinx-cisco-ospf-extension,frinx-cisco-pf-interfaces-extension,frinx-damping,frinx-dasan-vlan-extension,frinx-event-types,frinx-evpn,frinx-evpn-types,frinx-hsrp,frinx-if-aggregate-extension,frinx-isis-extension,frinx-juniper-if-aggregate-extension,frinx-juniper-if-extension,frinx-juniper-pf-interfaces-extension,frinx-juniper-probes-extension,frinx-l3ipvlan,frinx-lacp-lag-member,frinx-logging,frinx-mpls-ldp-extension,frinx-mpls-rsvp-extension,frinx-netflow,frinx-openconfig-acl,frinx-openconfig-bfd,frinx-openconfig-bgp,frinx-openconfig-bgp-policy,frinx-openconfig-bgp-types,frinx-openconfig-extensions,frinx-openconfig-if-aggregate,frinx-openconfig-if-ethernet,frinx-openconfig-if-ip,frinx-openconfig-if-ip-ext,frinx-openconfig-inet-types,frinx-openconfig-interfaces,frinx-openconfig-isis,frinx-openconfig-isis-lsdb-types,frinx-openconfig-isis-policy,frinx-openconfig-isis-types,frinx-openconfig-lacp,frinx-openconfig-lldp,frinx-openconfig-lldp-types,frinx-openconfig-local-routing,frinx-openconfig-mpls,frinx-openconfig-mpls-ldp,frinx-openconfig-mpls-rsvp,frinx-openconfig-mpls-types,frinx-openconfig-network-instance,frinx-openconfig-network-instance-l3,frinx-openconfig-network-instance-policy,frinx-openconfig-network-instance-types,frinx-openconfig-ospf-policy,frinx-openconfig-ospf-types,frinx-openconfig-ospfv2,frinx-openconfig-packet-match,frinx-openconfig-packet-match-types,frinx-openconfig-platform,frinx-openconfig-platform-linecard,frinx-openconfig-platform-port,frinx-openconfig-platform-transceiver,frinx-openconfig-platform-types,frinx-openconfig-policy-forwarding,frinx-openconfig-policy-types,frinx-openconfig-probes,frinx-openconfig-probes-types,frinx-openconfig-qos,frinx-openconfig-qos-types,frinx-openconfig-rib-bgp,frinx-openconfig-rib-bgp-ext,frinx-openconfig-rib-bgp-types,frinx-openconfig-routing-policy,frinx-openconfig-transport-types,frinx-openconfig-types,frinx-openconfig-vlan,frinx-openconfig-vlan-types,frinx-openconfig-yang-types,frinx-ospf-extension,frinx-qos-extension,frinx-snmp,frinx-uniconfig-topology,iana-if-type,ietf-inet-types,ietf-interfaces,ietf-yang-types,network-topology,network-topology,unified-topology,yang-ext]
 *
 * API version: 4.2.0.frinx
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger_unified

// Optional[Configuration parameters relating to the forwarding rule's action.] REF:Optional.empty
type FrinxOpenconfigPolicyForwardingPfforwardingpolicystructuralPoliciesPolicyRulesRuleActionConfig struct {
	// Optional[When this leaf is set, packets matching the match criteria for the forwarding rule should be looked up in the network-instance that is referenced rather than the network-instance with which the interface is associated. Such configuration allows policy-routing into multiple sub-topologies from a single ingress access interface, or different send and receive contexts for a particular interface (sometimes referred to as half-duplex VRF).] REF:Optional.empty
	NetworkInstance string `json:"network-instance,omitempty"`
	// Optional[When this leaf is set to true, the local system should remove the GRE header from the packet matching the rule. Following the decapsulation it should subsequently forward the encapsulated packet according to the relevant lookup (e.g., if the encapsulated packet is IP, the packet should be routed according to the IP destination).] REF:Optional.empty
	DecapsulateGre bool `json:"decapsulate-gre,omitempty"`
	// Optional[When this leaf is set to true, the local system should drop packets that match the rule.] REF:Optional.empty
	Discard bool `json:"discard,omitempty"`
	// Optional[When an IP next-hop is specified in the next-hop field, packets matching the match criteria for the forwarding rule should be forwarded to the next-hop IP address, bypassing any lookup on the local system.] REF:Optional.empty
	NextHop string `json:"next-hop,omitempty"`
	// Optional[When path-selection-group is set, packets matching the match criteria for the forwarding rule should be forwarded only via one of the paths that is specified within the referenced path-selection-group. The next-hop of the packet within the routing context should be used to determine between multiple paths that are specified within the group.] REF:Optional.empty
	PathSelectionGroup string `json:"path-selection-group,omitempty"`
}