package v1

// This file contains a collection of methods that can be used from go-restful to
// generate Swagger API documentation for its models. Please read this PR for more
// information on the implementation: https://github.com/emicklei/go-restful/pull/215
//
// TODOs are ignored from the parser (e.g. TODO(andronat):... || TODO:...) if and only if
// they are on one line! For multiple line or blocks that you want to ignore use ---.
// Any context after a --- is ignored.
//
// Those methods can be generated by using hack/update-generated-swagger-docs.sh

// AUTO-GENERATED FUNCTIONS START HERE
var map_AllowedFlexVolume = map[string]string{
	"":       "AllowedFlexVolume represents a single Flexvolume that is allowed to be used.",
	"driver": "Driver is the name of the Flexvolume driver.",
}

func (AllowedFlexVolume) SwaggerDoc() map[string]string {
	return map_AllowedFlexVolume
}

var map_FSGroupStrategyOptions = map[string]string{
	"":       "FSGroupStrategyOptions defines the strategy type and options used to create the strategy.",
	"type":   "Type is the strategy that will dictate what FSGroup is used in the SecurityContext.",
	"ranges": "Ranges are the allowed ranges of fs groups.  If you would like to force a single fs group then supply a single range with the same start and end.",
}

func (FSGroupStrategyOptions) SwaggerDoc() map[string]string {
	return map_FSGroupStrategyOptions
}

var map_IDRange = map[string]string{
	"":    "IDRange provides a min/max of an allowed range of IDs.",
	"min": "Min is the start of the range, inclusive.",
	"max": "Max is the end of the range, inclusive.",
}

func (IDRange) SwaggerDoc() map[string]string {
	return map_IDRange
}

var map_PodSecurityPolicyReview = map[string]string{
	"":       "PodSecurityPolicyReview checks which service accounts (not users, since that would be cluster-wide) can create the `PodTemplateSpec` in question.",
	"spec":   "spec is the PodSecurityPolicy to check.",
	"status": "status represents the current information/status for the PodSecurityPolicyReview.",
}

func (PodSecurityPolicyReview) SwaggerDoc() map[string]string {
	return map_PodSecurityPolicyReview
}

var map_PodSecurityPolicyReviewSpec = map[string]string{
	"":                    "PodSecurityPolicyReviewSpec defines specification for PodSecurityPolicyReview",
	"template":            "template is the PodTemplateSpec to check. The template.spec.serviceAccountName field is used if serviceAccountNames is empty, unless the template.spec.serviceAccountName is empty, in which case \"default\" is used. If serviceAccountNames is specified, template.spec.serviceAccountName is ignored.",
	"serviceAccountNames": "serviceAccountNames is an optional set of ServiceAccounts to run the check with. If serviceAccountNames is empty, the template.spec.serviceAccountName is used, unless it's empty, in which case \"default\" is used instead. If serviceAccountNames is specified, template.spec.serviceAccountName is ignored.",
}

func (PodSecurityPolicyReviewSpec) SwaggerDoc() map[string]string {
	return map_PodSecurityPolicyReviewSpec
}

var map_PodSecurityPolicyReviewStatus = map[string]string{
	"": "PodSecurityPolicyReviewStatus represents the status of PodSecurityPolicyReview.",
	"allowedServiceAccounts": "allowedServiceAccounts returns the list of service accounts in *this* namespace that have the power to create the PodTemplateSpec.",
}

func (PodSecurityPolicyReviewStatus) SwaggerDoc() map[string]string {
	return map_PodSecurityPolicyReviewStatus
}

var map_PodSecurityPolicySelfSubjectReview = map[string]string{
	"":       "PodSecurityPolicySelfSubjectReview checks whether this user/SA tuple can create the PodTemplateSpec",
	"spec":   "spec defines specification the PodSecurityPolicySelfSubjectReview.",
	"status": "status represents the current information/status for the PodSecurityPolicySelfSubjectReview.",
}

func (PodSecurityPolicySelfSubjectReview) SwaggerDoc() map[string]string {
	return map_PodSecurityPolicySelfSubjectReview
}

var map_PodSecurityPolicySelfSubjectReviewSpec = map[string]string{
	"":         "PodSecurityPolicySelfSubjectReviewSpec contains specification for PodSecurityPolicySelfSubjectReview.",
	"template": "template is the PodTemplateSpec to check.",
}

func (PodSecurityPolicySelfSubjectReviewSpec) SwaggerDoc() map[string]string {
	return map_PodSecurityPolicySelfSubjectReviewSpec
}

var map_PodSecurityPolicySubjectReview = map[string]string{
	"":       "PodSecurityPolicySubjectReview checks whether a particular user/SA tuple can create the PodTemplateSpec.",
	"spec":   "spec defines specification for the PodSecurityPolicySubjectReview.",
	"status": "status represents the current information/status for the PodSecurityPolicySubjectReview.",
}

func (PodSecurityPolicySubjectReview) SwaggerDoc() map[string]string {
	return map_PodSecurityPolicySubjectReview
}

var map_PodSecurityPolicySubjectReviewSpec = map[string]string{
	"":         "PodSecurityPolicySubjectReviewSpec defines specification for PodSecurityPolicySubjectReview",
	"template": "template is the PodTemplateSpec to check. If template.spec.serviceAccountName is empty it will not be defaulted. If its non-empty, it will be checked.",
	"user":     "user is the user you're testing for. If you specify \"user\" but not \"group\", then is it interpreted as \"What if user were not a member of any groups. If user and groups are empty, then the check is performed using *only* the serviceAccountName in the template.",
	"groups":   "groups is the groups you're testing for.",
}

func (PodSecurityPolicySubjectReviewSpec) SwaggerDoc() map[string]string {
	return map_PodSecurityPolicySubjectReviewSpec
}

var map_PodSecurityPolicySubjectReviewStatus = map[string]string{
	"":          "PodSecurityPolicySubjectReviewStatus contains information/status for PodSecurityPolicySubjectReview.",
	"allowedBy": "allowedBy is a reference to the rule that allows the PodTemplateSpec. A rule can be a SecurityContextConstraint or a PodSecurityPolicy A `nil`, indicates that it was denied.",
	"reason":    "A machine-readable description of why this operation is in the \"Failure\" status. If this value is empty there is no information available.",
	"template":  "template is the PodTemplateSpec after the defaulting is applied.",
}

func (PodSecurityPolicySubjectReviewStatus) SwaggerDoc() map[string]string {
	return map_PodSecurityPolicySubjectReviewStatus
}

var map_RunAsUserStrategyOptions = map[string]string{
	"":            "RunAsUserStrategyOptions defines the strategy type and any options used to create the strategy.",
	"type":        "Type is the strategy that will dictate what RunAsUser is used in the SecurityContext.",
	"uid":         "UID is the user id that containers must run as.  Required for the MustRunAs strategy if not using namespace/service account allocated uids.",
	"uidRangeMin": "UIDRangeMin defines the min value for a strategy that allocates by range.",
	"uidRangeMax": "UIDRangeMax defines the max value for a strategy that allocates by range.",
}

func (RunAsUserStrategyOptions) SwaggerDoc() map[string]string {
	return map_RunAsUserStrategyOptions
}

var map_SELinuxContextStrategyOptions = map[string]string{
	"":               "SELinuxContextStrategyOptions defines the strategy type and any options used to create the strategy.",
	"type":           "Type is the strategy that will dictate what SELinux context is used in the SecurityContext.",
	"seLinuxOptions": "seLinuxOptions required to run as; required for MustRunAs",
}

func (SELinuxContextStrategyOptions) SwaggerDoc() map[string]string {
	return map_SELinuxContextStrategyOptions
}

var map_SecurityContextConstraints = map[string]string{
	"":                         "SecurityContextConstraints governs the ability to make requests that affect the SecurityContext that will be applied to a container.",
	"metadata":                 "Standard object's metadata. More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#metadata",
	"priority":                 "Priority influences the sort order of SCCs when evaluating which SCCs to try first for a given pod request based on access in the Users and Groups fields.  The higher the int, the higher priority. An unset value is considered a 0 priority. If scores for multiple SCCs are equal they will be sorted from most restrictive to least restrictive. If both priorities and restrictions are equal the SCCs will be sorted by name.",
	"allowPrivilegedContainer": "AllowPrivilegedContainer determines if a container can request to be run as privileged.",
	"defaultAddCapabilities":   "DefaultAddCapabilities is the default set of capabilities that will be added to the container unless the pod spec specifically drops the capability.  You may not list a capabiility in both DefaultAddCapabilities and RequiredDropCapabilities.",
	"requiredDropCapabilities": "RequiredDropCapabilities are the capabilities that will be dropped from the container.  These are required to be dropped and cannot be added.",
	"allowedCapabilities":      "AllowedCapabilities is a list of capabilities that can be requested to add to the container. Capabilities in this field maybe added at the pod author's discretion. You must not list a capability in both AllowedCapabilities and RequiredDropCapabilities. To allow all capabilities you may use '*'.",
	"allowHostDirVolumePlugin": "AllowHostDirVolumePlugin determines if the policy allow containers to use the HostDir volume plugin",
	"volumes":                  "Volumes is a white list of allowed volume plugins.  FSType corresponds directly with the field names of a VolumeSource (azureFile, configMap, emptyDir).  To allow all volumes you may use \"*\". To allow no volumes, set to [\"none\"].",
	"allowedFlexVolumes":       "AllowedFlexVolumes is a whitelist of allowed Flexvolumes.  Empty or nil indicates that all Flexvolumes may be used.  This parameter is effective only when the usage of the Flexvolumes is allowed in the \"Volumes\" field.",
	"allowHostNetwork":         "AllowHostNetwork determines if the policy allows the use of HostNetwork in the pod spec.",
	"allowHostPorts":           "AllowHostPorts determines if the policy allows host ports in the containers.",
	"allowHostPID":             "AllowHostPID determines if the policy allows host pid in the containers.",
	"allowHostIPC":             "AllowHostIPC determines if the policy allows host ipc in the containers.",
	"seLinuxContext":           "SELinuxContext is the strategy that will dictate what labels will be set in the SecurityContext.",
	"runAsUser":                "RunAsUser is the strategy that will dictate what RunAsUser is used in the SecurityContext.",
	"supplementalGroups":       "SupplementalGroups is the strategy that will dictate what supplemental groups are used by the SecurityContext.",
	"fsGroup":                  "FSGroup is the strategy that will dictate what fs group is used by the SecurityContext.",
	"readOnlyRootFilesystem":   "ReadOnlyRootFilesystem when set to true will force containers to run with a read only root file system.  If the container specifically requests to run with a non-read only root file system the SCC should deny the pod. If set to false the container may run with a read only root file system if it wishes but it will not be forced to.",
	"users":                    "The users who have permissions to use this security context constraints",
	"groups":                   "The groups that have permission to use this security context constraints",
	"seccompProfiles":          "SeccompProfiles lists the allowed profiles that may be set for the pod or container's seccomp annotations.  An unset (nil) or empty value means that no profiles may be specifid by the pod or container.\tThe wildcard '*' may be used to allow all profiles.  When used to generate a value for a pod the first non-wildcard profile will be used as the default.",
}

func (SecurityContextConstraints) SwaggerDoc() map[string]string {
	return map_SecurityContextConstraints
}

var map_SecurityContextConstraintsList = map[string]string{
	"":         "SecurityContextConstraintsList is a list of SecurityContextConstraints objects",
	"metadata": "More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#metadata",
	"items":    "List of security context constraints.",
}

func (SecurityContextConstraintsList) SwaggerDoc() map[string]string {
	return map_SecurityContextConstraintsList
}

var map_ServiceAccountPodSecurityPolicyReviewStatus = map[string]string{
	"":     "ServiceAccountPodSecurityPolicyReviewStatus represents ServiceAccount name and related review status",
	"name": "name contains the allowed and the denied ServiceAccount name",
}

func (ServiceAccountPodSecurityPolicyReviewStatus) SwaggerDoc() map[string]string {
	return map_ServiceAccountPodSecurityPolicyReviewStatus
}

var map_SupplementalGroupsStrategyOptions = map[string]string{
	"":       "SupplementalGroupsStrategyOptions defines the strategy type and options used to create the strategy.",
	"type":   "Type is the strategy that will dictate what supplemental groups is used in the SecurityContext.",
	"ranges": "Ranges are the allowed ranges of supplemental groups.  If you would like to force a single supplemental group then supply a single range with the same start and end.",
}

func (SupplementalGroupsStrategyOptions) SwaggerDoc() map[string]string {
	return map_SupplementalGroupsStrategyOptions
}

// AUTO-GENERATED FUNCTIONS END HERE
