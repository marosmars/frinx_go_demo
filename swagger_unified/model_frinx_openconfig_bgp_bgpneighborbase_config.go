/*
 * unified
 *
 * API generated from yang definitions: [fake,frinx-acl-extension,frinx-bfd,frinx-bfd-extension,frinx-bgp-extension,frinx-cdp,frinx-cisco-if-extension,frinx-cisco-mpls-te-extension,frinx-cisco-ospf-extension,frinx-cisco-pf-interfaces-extension,frinx-damping,frinx-dasan-vlan-extension,frinx-event-types,frinx-evpn,frinx-evpn-types,frinx-hsrp,frinx-if-aggregate-extension,frinx-isis-extension,frinx-juniper-if-aggregate-extension,frinx-juniper-if-extension,frinx-juniper-pf-interfaces-extension,frinx-juniper-probes-extension,frinx-l3ipvlan,frinx-lacp-lag-member,frinx-logging,frinx-mpls-ldp-extension,frinx-mpls-rsvp-extension,frinx-netflow,frinx-openconfig-acl,frinx-openconfig-bfd,frinx-openconfig-bgp,frinx-openconfig-bgp-policy,frinx-openconfig-bgp-types,frinx-openconfig-extensions,frinx-openconfig-if-aggregate,frinx-openconfig-if-ethernet,frinx-openconfig-if-ip,frinx-openconfig-if-ip-ext,frinx-openconfig-inet-types,frinx-openconfig-interfaces,frinx-openconfig-isis,frinx-openconfig-isis-lsdb-types,frinx-openconfig-isis-policy,frinx-openconfig-isis-types,frinx-openconfig-lacp,frinx-openconfig-lldp,frinx-openconfig-lldp-types,frinx-openconfig-local-routing,frinx-openconfig-mpls,frinx-openconfig-mpls-ldp,frinx-openconfig-mpls-rsvp,frinx-openconfig-mpls-types,frinx-openconfig-network-instance,frinx-openconfig-network-instance-l3,frinx-openconfig-network-instance-policy,frinx-openconfig-network-instance-types,frinx-openconfig-ospf-policy,frinx-openconfig-ospf-types,frinx-openconfig-ospfv2,frinx-openconfig-packet-match,frinx-openconfig-packet-match-types,frinx-openconfig-platform,frinx-openconfig-platform-linecard,frinx-openconfig-platform-port,frinx-openconfig-platform-transceiver,frinx-openconfig-platform-types,frinx-openconfig-policy-forwarding,frinx-openconfig-policy-types,frinx-openconfig-probes,frinx-openconfig-probes-types,frinx-openconfig-qos,frinx-openconfig-qos-types,frinx-openconfig-rib-bgp,frinx-openconfig-rib-bgp-ext,frinx-openconfig-rib-bgp-types,frinx-openconfig-routing-policy,frinx-openconfig-transport-types,frinx-openconfig-types,frinx-openconfig-vlan,frinx-openconfig-vlan-types,frinx-openconfig-yang-types,frinx-ospf-extension,frinx-qos-extension,frinx-snmp,frinx-uniconfig-topology,iana-if-type,ietf-inet-types,ietf-interfaces,ietf-yang-types,network-topology,network-topology,unified-topology,yang-ext]
 *
 * API version: 4.2.0.frinx
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger_unified

// Optional[Configuration parameters relating to the BGP neighbor or group] REF:Optional.empty
type FrinxOpenconfigBgpBgpneighborbaseConfig struct {
	// Optional[Explicitly designate the peer or peer group as internal (iBGP) or external (eBGP).] REF:Optional.empty
	PeerType *FrinxOpenconfigBgpTypesPeerType `json:"peer-type,omitempty"`
	// Optional[Remove private AS numbers from updates sent to peers - when this leaf is not specified, the AS_PATH attribute should be sent to the peer unchanged] REF:Optional.empty
	RemovePrivateAs *FrinxOpenconfigBgpTypesRemovePrivateAsOption `json:"remove-private-as,omitempty"`
	// Optional[Specify which types of community should be sent to the neighbor or group. The default is to not send the community attribute] REF:Optional.empty
	SendCommunity *FrinxOpenconfigBgpTypesCommunityType `json:"send-community,omitempty"`
	// Optional[An optional textual description (intended primarily for use with a peer or group] REF:Optional.empty
	Description string `json:"description,omitempty"`
	// Optional[AS number of the peer.] REF:Optional.empty
	PeerAs int64 `json:"peer-as,omitempty"`
	// Optional[Address of the BGP peer, either in IPv4 or IPv6] REF:Optional.empty
	NeighborAddress string `json:"neighbor-address,omitempty"`
	// Optional[Enable route flap damping.] REF:Optional.empty
	RouteFlapDamping bool `json:"route-flap-damping,omitempty"`
	// Optional[The peer-group with which this neighbor is associated] REF:Optional.empty
	PeerGroup string `json:"peer-group,omitempty"`
	// Optional[Configures an MD5 authentication password for use with neighboring devices.] REF:Optional.empty
	AuthPassword string `json:"auth-password,omitempty"`
	// Optional[Whether the BGP peer is enabled. In cases where the enabled leaf is set to false, the local system should not initiate connections to the neighbor, and should not respond to TCP connections attempts from the neighbor. If the state of the BGP session is ESTABLISHED at the time that this leaf is set to false, the BGP session should be ceased.] REF:Optional.empty
	Enabled bool `json:"enabled,omitempty"`
	// Optional[The local autonomous system number that is to be used when establishing sessions with the remote peer or peer group, if this differs from the global BGP router autonomous system number.] REF:Optional.empty
	LocalAs int64 `json:"local-as,omitempty"`
}
