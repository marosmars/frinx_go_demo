/*
 * unified
 *
 * API generated from yang definitions: [fake,frinx-acl-extension,frinx-bfd,frinx-bfd-extension,frinx-bgp-extension,frinx-cdp,frinx-cisco-if-extension,frinx-cisco-mpls-te-extension,frinx-cisco-ospf-extension,frinx-cisco-pf-interfaces-extension,frinx-damping,frinx-dasan-vlan-extension,frinx-event-types,frinx-evpn,frinx-evpn-types,frinx-hsrp,frinx-if-aggregate-extension,frinx-isis-extension,frinx-juniper-if-aggregate-extension,frinx-juniper-if-extension,frinx-juniper-pf-interfaces-extension,frinx-juniper-probes-extension,frinx-l3ipvlan,frinx-lacp-lag-member,frinx-logging,frinx-mpls-ldp-extension,frinx-mpls-rsvp-extension,frinx-netflow,frinx-openconfig-acl,frinx-openconfig-bfd,frinx-openconfig-bgp,frinx-openconfig-bgp-policy,frinx-openconfig-bgp-types,frinx-openconfig-extensions,frinx-openconfig-if-aggregate,frinx-openconfig-if-ethernet,frinx-openconfig-if-ip,frinx-openconfig-if-ip-ext,frinx-openconfig-inet-types,frinx-openconfig-interfaces,frinx-openconfig-isis,frinx-openconfig-isis-lsdb-types,frinx-openconfig-isis-policy,frinx-openconfig-isis-types,frinx-openconfig-lacp,frinx-openconfig-lldp,frinx-openconfig-lldp-types,frinx-openconfig-local-routing,frinx-openconfig-mpls,frinx-openconfig-mpls-ldp,frinx-openconfig-mpls-rsvp,frinx-openconfig-mpls-types,frinx-openconfig-network-instance,frinx-openconfig-network-instance-l3,frinx-openconfig-network-instance-policy,frinx-openconfig-network-instance-types,frinx-openconfig-ospf-policy,frinx-openconfig-ospf-types,frinx-openconfig-ospfv2,frinx-openconfig-packet-match,frinx-openconfig-packet-match-types,frinx-openconfig-platform,frinx-openconfig-platform-linecard,frinx-openconfig-platform-port,frinx-openconfig-platform-transceiver,frinx-openconfig-platform-types,frinx-openconfig-policy-forwarding,frinx-openconfig-policy-types,frinx-openconfig-probes,frinx-openconfig-probes-types,frinx-openconfig-qos,frinx-openconfig-qos-types,frinx-openconfig-rib-bgp,frinx-openconfig-rib-bgp-ext,frinx-openconfig-rib-bgp-types,frinx-openconfig-routing-policy,frinx-openconfig-transport-types,frinx-openconfig-types,frinx-openconfig-vlan,frinx-openconfig-vlan-types,frinx-openconfig-yang-types,frinx-ospf-extension,frinx-qos-extension,frinx-snmp,frinx-uniconfig-topology,iana-if-type,ietf-inet-types,ietf-interfaces,ietf-yang-types,network-topology,network-topology,unified-topology,yang-ext]
 *
 * API version: 4.2.0.frinx
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger_unified

// Optional[list of per-interface RSVP configurations] REF:Optional.empty
type FrinxOpenconfigNetworkInstanceNetworkinstancesNetworkinstanceMplsSignalingprotocolsRsvpteInterfaceattributesInterface struct {
	// Optional[Enclosing container for bandwidth reservation] REF:Optional.empty
	BandwidthReservations *FrinxOpenconfigMplsRsvpMplsrsvpinterfacereservationsBandwidthReservations `json:"bandwidth-reservations,omitempty"`
	// Optional[link-protection (NHOP) related configuration] REF:Optional.empty
	Protection *FrinxOpenconfigMplsRsvpMplsrsvplinkprotectionProtection `json:"protection,omitempty"`
	// Optional[Bandwidth percentage reservable by RSVP on an interface] REF:Optional.empty
	Subscription *FrinxOpenconfigNetworkInstanceNetworkinstancesNetworkinstanceMplsSignalingprotocolsRsvpteInterfaceattributesInterfaceSubscription `json:"subscription,omitempty"`
	// Optional[Top level container for RSVP hello parameters] REF:Optional.empty
	Hellos *FrinxOpenconfigMplsRsvpMplsrsvphellosHellos `json:"hellos,omitempty"`
	// Optional[Configuration of per-interface RSVP parameters] REF:Optional.empty
	Config *FrinxOpenconfigMplsRsvpRsvpglobalRsvpteInterfaceattributesInterfaceConfig `json:"config,omitempty"`
	// Optional[reference to the interface-id data] REF:Optional.empty
	InterfaceId string `json:"interface-id,omitempty"`
	// Optional[Reference to an interface or subinterface] REF:Optional.empty
	InterfaceRef *FrinxOpenconfigInterfacesInterfacerefInterfaceRef `json:"interface-ref,omitempty"`
	// Optional[Configuration and state parameters relating to RSVP authentication as per RFC2747] REF:Optional.empty
	Authentication *FrinxOpenconfigMplsRsvpMplsrsvpauthenticationAuthentication `json:"authentication,omitempty"`
}
