/*
 * unified
 *
 * API generated from yang definitions: [fake,frinx-acl-extension,frinx-bfd,frinx-bfd-extension,frinx-bgp-extension,frinx-cdp,frinx-cisco-if-extension,frinx-cisco-mpls-te-extension,frinx-cisco-ospf-extension,frinx-cisco-pf-interfaces-extension,frinx-damping,frinx-dasan-vlan-extension,frinx-event-types,frinx-evpn,frinx-evpn-types,frinx-hsrp,frinx-if-aggregate-extension,frinx-isis-extension,frinx-juniper-if-aggregate-extension,frinx-juniper-if-extension,frinx-juniper-pf-interfaces-extension,frinx-juniper-probes-extension,frinx-l3ipvlan,frinx-lacp-lag-member,frinx-logging,frinx-mpls-ldp-extension,frinx-mpls-rsvp-extension,frinx-netflow,frinx-openconfig-acl,frinx-openconfig-bfd,frinx-openconfig-bgp,frinx-openconfig-bgp-policy,frinx-openconfig-bgp-types,frinx-openconfig-extensions,frinx-openconfig-if-aggregate,frinx-openconfig-if-ethernet,frinx-openconfig-if-ip,frinx-openconfig-if-ip-ext,frinx-openconfig-inet-types,frinx-openconfig-interfaces,frinx-openconfig-isis,frinx-openconfig-isis-lsdb-types,frinx-openconfig-isis-policy,frinx-openconfig-isis-types,frinx-openconfig-lacp,frinx-openconfig-lldp,frinx-openconfig-lldp-types,frinx-openconfig-local-routing,frinx-openconfig-mpls,frinx-openconfig-mpls-ldp,frinx-openconfig-mpls-rsvp,frinx-openconfig-mpls-types,frinx-openconfig-network-instance,frinx-openconfig-network-instance-l3,frinx-openconfig-network-instance-policy,frinx-openconfig-network-instance-types,frinx-openconfig-ospf-policy,frinx-openconfig-ospf-types,frinx-openconfig-ospfv2,frinx-openconfig-packet-match,frinx-openconfig-packet-match-types,frinx-openconfig-platform,frinx-openconfig-platform-linecard,frinx-openconfig-platform-port,frinx-openconfig-platform-transceiver,frinx-openconfig-platform-types,frinx-openconfig-policy-forwarding,frinx-openconfig-policy-types,frinx-openconfig-probes,frinx-openconfig-probes-types,frinx-openconfig-qos,frinx-openconfig-qos-types,frinx-openconfig-rib-bgp,frinx-openconfig-rib-bgp-ext,frinx-openconfig-rib-bgp-types,frinx-openconfig-routing-policy,frinx-openconfig-transport-types,frinx-openconfig-types,frinx-openconfig-vlan,frinx-openconfig-vlan-types,frinx-openconfig-yang-types,frinx-ospf-extension,frinx-qos-extension,frinx-snmp,frinx-uniconfig-topology,iana-if-type,ietf-inet-types,ietf-interfaces,ietf-yang-types,network-topology,network-topology,unified-topology,yang-ext]
 *
 * API version: 4.2.0.frinx
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger_unified

// Optional[Configuration parameters relating to graceful-restart] REF:Optional.empty
type FrinxOpenconfigBgpBgpglobalbaseGracefulrestartConfig struct {
	// Optional[Estimated time (in seconds) for the local BGP speaker to restart a session. This value is advertise in the graceful restart BGP capability.  This is a 12-bit value, referred to as Restart Time in RFC4724.  Per RFC4724, the suggested default value is <= the hold-time value.] REF:Optional.empty
	RestartTime int32 `json:"restart-time,omitempty"`
	// Optional[Enable graceful-restart in helper mode only. When this leaf is set, the local system does not retain forwarding its own state during a restart, but supports procedures for the receiving speaker, as defined in RFC4724.] REF:Optional.empty
	HelperOnly bool `json:"helper-only,omitempty"`
	// Optional[An upper-bound on the time thate stale routes will be retained by a router after a session is restarted. If an End-of-RIB (EOR) marker is received prior to this timer expiring stale-routes will be flushed upon its receipt - if no EOR is received, then when this timer expires stale paths will be purged. This timer is referred to as the Selection_Deferral_Timer in RFC4724] REF:Optional.empty
	StaleRoutesTime int32 `json:"stale-routes-time,omitempty"`
	// Optional[Enable or disable the graceful-restart capability.] REF:Optional.empty
	Enabled bool `json:"enabled,omitempty"`
}
