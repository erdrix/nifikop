package config

// Template for nifi's authorizer.xml file, used by the initial cluster's nodes
var UsersGroupsTemplate = `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<tenants>
    <groups/>
    <users>
      <user identifier="{{ .InitialAdminUserId }}" identity="{{ .InitialAdminUser }}"/>
{{- range $uuID, $host := .NodesHost }}
      <user identifier="{{ $host }}" identity="{{ $uuID }}"/>
{{- end }}
      </users>
</tenants>
`
// <tenants>
//  <groups>
//    <group identifier="{{ .AdminsGroupId }}" name="admins">
//      <user identifier="{{ .ControllerUserId }}"/>
//      <user identifier="{{ .InitialAdminUserId }}"/>
//    </group>
//      <group identifier="{{ .NodesGroupId }}" name="{{ .NodesGroupName }}">
//{{- range $uuID, $host := .NodesHost }}
//        <user identifier="{{ $uuID }}"/>
//{{- end }}
//      </group>
//      <group identifier="{{ .ReadersGroupId }}" name="readers"/>
//  </groups>
//  <users>
//    <user identifier="{{ .ControllerUserId }}" identity="{{ .ControllerUser }}"/>
//    <user identifier="{{ .InitialAdminUserId }}" identity="{{ .InitialAdminUser }}"/>
//{{- range $uuID, $host := .NodesHost }}
//    <user identifier="{{ $host }}" identity="{{ $uuID }}"/>
//{{- end }}
//  </users>
//</tenants>