/*
 * unified
 *
 * API generated from yang definitions: [fake,frinx-acl-extension,frinx-bfd,frinx-bfd-extension,frinx-bgp-extension,frinx-cdp,frinx-cisco-if-extension,frinx-cisco-mpls-te-extension,frinx-cisco-ospf-extension,frinx-cisco-pf-interfaces-extension,frinx-damping,frinx-dasan-vlan-extension,frinx-event-types,frinx-evpn,frinx-evpn-types,frinx-hsrp,frinx-if-aggregate-extension,frinx-isis-extension,frinx-juniper-if-aggregate-extension,frinx-juniper-if-extension,frinx-juniper-pf-interfaces-extension,frinx-juniper-probes-extension,frinx-l3ipvlan,frinx-lacp-lag-member,frinx-logging,frinx-mpls-ldp-extension,frinx-mpls-rsvp-extension,frinx-netflow,frinx-openconfig-acl,frinx-openconfig-bfd,frinx-openconfig-bgp,frinx-openconfig-bgp-policy,frinx-openconfig-bgp-types,frinx-openconfig-extensions,frinx-openconfig-if-aggregate,frinx-openconfig-if-ethernet,frinx-openconfig-if-ip,frinx-openconfig-if-ip-ext,frinx-openconfig-inet-types,frinx-openconfig-interfaces,frinx-openconfig-isis,frinx-openconfig-isis-lsdb-types,frinx-openconfig-isis-policy,frinx-openconfig-isis-types,frinx-openconfig-lacp,frinx-openconfig-lldp,frinx-openconfig-lldp-types,frinx-openconfig-local-routing,frinx-openconfig-mpls,frinx-openconfig-mpls-ldp,frinx-openconfig-mpls-rsvp,frinx-openconfig-mpls-types,frinx-openconfig-network-instance,frinx-openconfig-network-instance-l3,frinx-openconfig-network-instance-policy,frinx-openconfig-network-instance-types,frinx-openconfig-ospf-policy,frinx-openconfig-ospf-types,frinx-openconfig-ospfv2,frinx-openconfig-packet-match,frinx-openconfig-packet-match-types,frinx-openconfig-platform,frinx-openconfig-platform-linecard,frinx-openconfig-platform-port,frinx-openconfig-platform-transceiver,frinx-openconfig-platform-types,frinx-openconfig-policy-forwarding,frinx-openconfig-policy-types,frinx-openconfig-probes,frinx-openconfig-probes-types,frinx-openconfig-qos,frinx-openconfig-qos-types,frinx-openconfig-rib-bgp,frinx-openconfig-rib-bgp-ext,frinx-openconfig-rib-bgp-types,frinx-openconfig-routing-policy,frinx-openconfig-transport-types,frinx-openconfig-types,frinx-openconfig-vlan,frinx-openconfig-vlan-types,frinx-openconfig-yang-types,frinx-ospf-extension,frinx-qos-extension,frinx-snmp,frinx-uniconfig-topology,iana-if-type,ietf-inet-types,ietf-interfaces,ietf-yang-types,network-topology,network-topology,unified-topology,yang-ext]
 *
 * API version: 4.2.0.frinx
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger_unified

// Optional[List of BGP neighbors configured on the local system, uniquely identified by peer IPv[46] address] REF:Optional.empty
type FrinxOpenconfigNetworkInstanceNetworkinstancesNetworkinstanceProtocolsProtocolBgpNeighborsNeighbor struct {
	// Optional[Timers related to a BGP neighbor] REF:Optional.empty
	Timers *FrinxOpenconfigBgpBgpneighborbaseTimers `json:"timers,omitempty"`
	// Optional[Anchor point for routing policies in the model. Import and export policies are with respect to the local routing table, i.e., export (send) and import (receive), depending on the context.] REF:Optional.empty
	ApplyPolicy *FrinxOpenconfigRoutingPolicyApplypolicygroupApplyPolicy `json:"apply-policy,omitempty"`
	// Optional[Route reflector parameters for the BGPgroup] REF:Optional.empty
	RouteReflector *FrinxOpenconfigBgpBgpcommonstructureneighborgrouproutereflectorRouteReflector `json:"route-reflector,omitempty"`
	// Optional[Enable BFD for liveliness detection to the next-hop or neighbour.] REF:Optional.empty
	FrinxOpenconfigBfdenableBfd *FrinxOpenconfigBfdBfdenableEnableBfd `json:"frinx-openconfig-bfd:enable-bfd,omitempty"`
	// Optional[Transport session parameters for the BGP neighbor] REF:Optional.empty
	Transport *FrinxOpenconfigBgpBgpneighborbaseTransport `json:"transport,omitempty"`
	// Optional[AS_PATH manipulation parameters for the BGP neighbor or group] REF:Optional.empty
	AsPathOptions *FrinxOpenconfigBgpBgpcommonstructureneighborgroupaspathoptionsAsPathOptions `json:"as-path-options,omitempty"`
	// Optional[eBGP multi-hop parameters for the BGPgroup] REF:Optional.empty
	EbgpMultihop *FrinxOpenconfigBgpBgpcommonstructureneighborgroupebgpmultihopEbgpMultihop `json:"ebgp-multihop,omitempty"`
	// Optional[Error handling parameters used for the BGP neighbor or group] REF:Optional.empty
	ErrorHandling *FrinxOpenconfigBgpBgpneighborbaseErrorHandling `json:"error-handling,omitempty"`
	// Optional[Parameters related to the use of multiple-paths for the same NLRI when they are received only from this neighbor] REF:Optional.empty
	UseMultiplePaths *FrinxOpenconfigBgpBgpneighborusemultiplepathsUseMultiplePaths `json:"use-multiple-paths,omitempty"`
	// Optional[Parameters relating to the advertisement and receipt of multiple paths for a single NLRI (add-paths)] REF:Optional.empty
	AddPaths *FrinxOpenconfigBgpBgpcommonstructureneighborgroupaddpathsAddPaths `json:"add-paths,omitempty"`
	// Optional[Per-address-family configuration parameters associated with the neighbor] REF:Optional.empty
	AfiSafis *FrinxOpenconfigNetworkInstanceNetworkinstancesNetworkinstanceProtocolsProtocolBgpNeighborsNeighborAfiSafis `json:"afi-safis,omitempty"`
	// Optional[Parameters relating the graceful restart mechanism for BGP] REF:Optional.empty
	GracefulRestart *FrinxOpenconfigBgpBgpneighborbaseGracefulRestart `json:"graceful-restart,omitempty"`
	// Optional[Reference to the address of the BGP neighbor used as a key in the neighbor list] REF:Optional.empty
	NeighborAddress string `json:"neighbor-address,omitempty"`
	// Optional[Logging options for events related to the BGP neighbor or group] REF:Optional.empty
	LoggingOptions *FrinxOpenconfigBgpBgpcommonstructureneighborgrouploggingoptionsLoggingOptions `json:"logging-options,omitempty"`
	// Optional[Configuration parameters relating to the BGP neighbor or group] REF:Optional.empty
	Config *FrinxOpenconfigBgpBgpneighborbaseConfig `json:"config,omitempty"`
}