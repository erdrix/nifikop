package config

// Template for nifi's authorizer.xml file, used by the initial cluster's nodes
var AuthorizationsTemplate = `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<authorizations>
<policies>
<policy identifier="f99bccd1-a30e-3e4a-98a2-dbc708edc67f" resource="/flow" action="R">
<user identifier="{{ .InitialAdminUserId }}"/>
</policy>
<policy identifier="b8775bd4-704a-34c6-987b-84f2daf7a515" resource="/restricted-components" action="W">
<user identifier="{{ .InitialAdminUserId }}"/>
</policy>
<policy identifier="627410be-1717-35b4-a06f-e9362b89e0b7" resource="/tenants" action="R">
<user identifier="{{ .InitialAdminUserId }}"/>
</policy>
<policy identifier="15e4e0bd-cb28-34fd-8587-f8d15162cba5" resource="/tenants" action="W">
<user identifier="{{ .InitialAdminUserId }}"/>
</policy>
<policy identifier="ff96062a-fa99-36dc-9942-0f6442ae7212" resource="/policies" action="R">
<user identifier="{{ .InitialAdminUserId }}"/>
</policy>
<policy identifier="ad99ea98-3af6-3561-ae27-5bf09e1d969d" resource="/policies" action="W">
<user identifier="{{ .InitialAdminUserId }}"/>
</policy>
<policy identifier="2e1015cb-0fed-3005-8e0d-722311f21a03" resource="/controller" action="R">
<user identifier="{{ .InitialAdminUserId }}"/>
</policy>
<policy identifier="c6322e6c-4cc1-3bcc-91b3-2ed2111674cf" resource="/controller" action="W">
<user identifier="{{ .InitialAdminUserId }}"/>
</policy>
<policy identifier="287edf48-da72-359b-8f61-da5d4c45a270" resource="/proxy" action="W">
{{- range $uuID, $host := .NodesHost }}
<user identifier="{{ $uuID }}"/>
{{- end }}
</policy>
</policies>
</authorizations>
`



// <?xml version="1.0" encoding="UTF-8" standalone="yes"?>
//<authorizations>
//    <policies>
//        <policy identifier="f99bccd1-a30e-3e4a-98a2-dbc708edc67f" resource="/flow" action="R">
//            <group identifier="{{ .ReadersGroupId }}"/>
//            <group identifier="{{ .AdminsGroupId }}"/>
//        </policy>
//        <policy identifier="b8775bd4-704a-34c6-987b-84f2daf7a515" resource="/restricted-components" action="W">
//            <group identifier="{{ .AdminsGroupId }}"/>
//        </policy>
//        <policy identifier="627410be-1717-35b4-a06f-e9362b89e0b7" resource="/tenants" action="R">
//            <group identifier="{{ .ReadersGroupId }}"/>
//            <group identifier="{{ .AdminsGroupId }}"/>
//        </policy>
//        <policy identifier="15e4e0bd-cb28-34fd-8587-f8d15162cba5" resource="/tenants" action="W">
//            <group identifier="{{ .AdminsGroupId }}"/>
//        </policy>
//        <policy identifier="ff96062a-fa99-36dc-9942-0f6442ae7212" resource="/policies" action="R">
//            <group identifier="{{ .ReadersGroupId }}"/>
//            <group identifier="{{ .AdminsGroupId }}"/>
//        </policy>
//        <policy identifier="ad99ea98-3af6-3561-ae27-5bf09e1d969d" resource="/policies" action="W">
//            <group identifier="{{ .AdminsGroupId }}"/>
//        </policy>
//        <policy identifier="2e1015cb-0fed-3005-8e0d-722311f21a03" resource="/controller" action="R">
//            <group identifier="{{ .ReadersGroupId }}"/>
//            <group identifier="{{ .AdminsGroupId }}"/>
//        </policy>
//        <policy identifier="c6322e6c-4cc1-3bcc-91b3-2ed2111674cf" resource="/controller" action="W">
//            <group identifier="{{ .AdminsGroupId }}"/>
//        </policy>
//        <policy identifier="287edf48-da72-359b-8f61-da5d4c45a270" resource="/proxy" action="W">
//            <group identifier="{{ .NodesGroupId }}"/>
//{{- range $uuID, $host := .NodesHost }}
//            <user identifier="{{ $uuID }}"/>
//{{- end }}
//        </policy>
//        <policy identifier="2c2d93df-0175-1000-ffff-fffff14d91f6" resource="/provenance" action="R">
//            <group identifier="{{ .AdminsGroupId }}"/>
//        </policy>
//        <policy identifier="2c2e74c7-0175-1000-0000-0000442a8c8c" resource="/site-to-site" action="R">
//            <group identifier="{{ .AdminsGroupId }}"/>
//        </policy>
//        <policy identifier="2c2e906f-0175-1000-ffff-fffff599cb12" resource="/system" action="R">
//            <group identifier="{{ .AdminsGroupId }}"/>
//        </policy>
//        <policy identifier="2c2ed24c-0175-1000-0000-00007ee8c4c3" resource="/counters" action="R">
//            <group identifier="{{ .ReadersGroupId }}"/>
//            <group identifier="{{ .AdminsGroupId }}"/>
//        </policy>
//        <policy identifier="2c2efe29-0175-1000-0000-000009663dda" resource="/counters" action="W">
//            <group identifier="{{ .AdminsGroupId }}"/>
//        </policy>
//    </policies>
//</authorizations>	