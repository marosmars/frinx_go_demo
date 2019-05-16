/*
 * southbound
 *
 * API generated from yang definitions: [aaa-encrypt-service-config,cli-topology,cli-translate-registry,cluster-singleton-service-impl,cluster-singleton-service-spi,config,fake,general-entity,ietf-inet-types,ietf-netconf,ietf-netconf-monitoring,ietf-netconf-monitoring-extension,ietf-netconf-notifications,ietf-yang-library,ietf-yang-types,journal,nc-notifications,netconf-keystore,netconf-node-inventory,netconf-node-topology,network-topology,network-topology,notifications,opendaylight-config-dom-datastore,opendaylight-entity-ownership-service,opendaylight-inmemory-datastore-provider,opendaylight-inventory,opendaylight-legacy-entity-ownership-service-provider,opendaylight-md-sal-common,opendaylight-md-sal-dom,opendaylight-operational-dom-datastore,rpc-context,yang-ext]
 *
 * API version: 4.2.0.frinx
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger_southbound

// Optional.empty REF:Optional.empty
type IetfNetconfMonitoringGetschemaInput struct {
	// Optional[The data modeling language of the schema.  If this parameter is not present, and more than one formats of the schema exists on the server, a 'data-not-unique' error is returned, as described above.] REF:Optional.empty
	Format *IetfNetconfMonitoringFormatIdentityref `json:"format,omitempty"`
	// Optional[Identifier for the schema list entry.] REF:Optional.empty
	Identifier string `json:"identifier,omitempty"`
	// Optional[Version of the schema requested.  If this parameter is not present, and more than one version of the schema exists on the server, a 'data-not-unique' error is returned, as described above.] REF:Optional.empty
	Version string `json:"version,omitempty"`
}