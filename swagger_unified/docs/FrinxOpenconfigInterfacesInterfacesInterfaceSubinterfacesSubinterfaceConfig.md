# FrinxOpenconfigInterfacesInterfacesInterfaceSubinterfacesSubinterfaceConfig

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Optional[[adapted from IETF interfaces model (RFC 7223)]  The name of the interface.  A device MAY restrict the allowed values for this leaf, possibly depending on the type of the interface. For system-controlled interfaces, this leaf is the device-specific name of the interface.  The &#39;config false&#39; list interfaces/interface[name]/state contains the currently existing interfaces on the device.  If a client tries to create configuration for a system-controlled interface that is not present in the corresponding state list, the server MAY reject the request if the implementation does not support pre-provisioning of interfaces or if the name refers to an interface that can never exist in the system.  A NETCONF server MUST reply with an rpc-error with the error-tag &#39;invalid-value&#39; in this case.  The IETF model in RFC 7223 provides YANG features for the following (i.e., pre-provisioning and arbitrary-names), however they are omitted here:   If the device supports pre-provisioning of interface  configuration, the &#39;pre-provisioning&#39; feature is  advertised.   If the device allows arbitrarily named user-controlled  interfaces, the &#39;arbitrary-names&#39; feature is advertised.  When a configured user-controlled interface is created by the system, it is instantiated with the same name in the /interfaces/interface[name]/state list.] REF:Optional[RFC 7223: A YANG Data Model for Interface Management] | [optional] [default to null]
**Description** | **string** | Optional[[adapted from IETF interfaces model (RFC 7223)]  A textual description of the interface.  A server implementation MAY map this leaf to the ifAlias MIB object.  Such an implementation needs to use some mechanism to handle the differences in size and characters allowed between this leaf and ifAlias.  The definition of such a mechanism is outside the scope of this document.  Since ifAlias is defined to be stored in non-volatile storage, the MIB implementation MUST map ifAlias to the value of &#39;description&#39; in the persistently stored datastore.  Specifically, if the device supports &#39;:startup&#39;, when ifAlias is read the device MUST return the value of &#39;description&#39; in the &#39;startup&#39; datastore, and when it is written, it MUST be written to the &#39;running&#39; and &#39;startup&#39; datastores.  Note that it is up to the implementation to  decide whether to modify this single leaf in &#39;startup&#39; or perform an implicit copy-config from &#39;running&#39; to &#39;startup&#39;.  If the device does not support &#39;:startup&#39;, ifAlias MUST be mapped to the &#39;description&#39; leaf in the &#39;running&#39; datastore.] REF:Optional[RFC 2863: The Interfaces Group MIB - ifAlias] | [optional] [default to null]
**Index** | **int64** | Optional[The index of the subinterface, or logical interface number. On systems with no support for subinterfaces, or not using subinterfaces, this value should default to 0, i.e., the default subinterface.] REF:Optional.empty | [optional] [default to null]
**Enabled** | **bool** | Optional[[adapted from IETF interfaces model (RFC 7223)]  This leaf contains the configured, desired state of the interface.  Systems that implement the IF-MIB use the value of this leaf in the &#39;running&#39; datastore to set IF-MIB.ifAdminStatus to &#39;up&#39; or &#39;down&#39; after an ifEntry has been initialized, as described in RFC 2863.  Changes in this leaf in the &#39;running&#39; datastore are reflected in ifAdminStatus, but if ifAdminStatus is changed over SNMP, this leaf is not affected.] REF:Optional[RFC 2863: The Interfaces Group MIB - ifAdminStatus] | [optional] [default to null]
**FrinxJuniperIfExtensionrpmType** | [***FrinxJuniperIfExtensionRpmTypes**](frinx.juniper.if.extension.RpmTypes.md) | Optional[The rpm type of the subinterface] REF:Optional.empty | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

