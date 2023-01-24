package config

import "github.com/upbound/upjet/pkg/config"

// ExternalNameNotTestedConfigs contains no-tested configurations for this
// provider.
var ExternalNameNotTestedConfigs = map[string]config.ExternalName{
	// appengine
	//
	// apps/{{project}}/domainMappings/{{domain_name}}
	"google_app_engine_domain_mapping": config.TemplatedStringAsIdentifier("domain_name", "apps/{{ .setup.configuration.project }}/domainMappings/{{ .external_name }}"),
	// apps/{{project}}/firewall/ingressRules/{{priority}}
	"google_app_engine_firewall_rule": config.TemplatedStringAsIdentifier("priority", "apps/{{ .setup.configuration.project }}/firewall/ingressRules/{{ .external_name }}"),
	// apps/{{project}}/services/{{service}}/versions/{{version_id}}
	"google_app_engine_flexible_app_version": config.TemplatedStringAsIdentifier("version_id", "apps/{{ .setup.configuration.project }}/services{{ .parameters.service }}/versions/{{ .external_name }}"),
	// apps/{{project}}/services/{{service}}
	"google_app_engine_service_split_traffic": config.TemplatedStringAsIdentifier("service", "apps/{{ .setup.configuration.project }}/services/{{ .external_name }}"),

	// assuredworkloads
	//
	// organizations/{{organization}}/locations/{{location}}/workloads/{{name}}
	"google_assured_workloads_workload": config.IdentifierFromProvider,

	// billing
	//
	// your-billing-account-id roles/billing.user
	"google_billing_account_iam_binding": config.TemplatedStringAsIdentifier("", "{{ .parameters.billing_account_id }} {{ .parameters.role }}"),
	// your-billing-account-id roles/billing.user user:jane@example.com
	"google_billing_account_iam_member": config.TemplatedStringAsIdentifier("", "{{ .parameters.billing_account_id }} {{ .parameters.role }} {{ .parameters.member }}"),
	// your-billing-account-id
	"google_billing_account_iam_policy": config.TemplatedStringAsIdentifier("", "{{ .parameters.billing_account_id }}"),
	// billingAccounts/{{billing_account}}/budgets/{{name}}
	// TODO: For now API is not normalized. While testing resource we can check the actual ID and normalize the API.
	"google_billing_budget": config.IdentifierFromProvider,

	// binaryauthorization
	//
	// projects/{{project}}/attestors/{{attestor}} roles/viewer
	"google_binary_authorization_attestor_iam_binding": config.TemplatedStringAsIdentifier("", "projects/{{ .setup.configuration.project }}/attestors/{{ .parameters.attestor }} {{ .parameters.role }}"),
	// projects/{{project}}/attestors/{{attestor}} roles/viewer user:jane@example.com
	"google_binary_authorization_attestor_iam_member": config.TemplatedStringAsIdentifier("", "projects/{{ .setup.configuration.project }}/attestors/{{ .parameters.attestor }} {{ .parameters.role }} {{ .parameters.member }}"),
	// projects/{{project}}/attestors/{{attestor}}
	"google_binary_authorization_attestor_iam_policy": config.TemplatedStringAsIdentifier("", "projects/{{ .setup.configuration.project }}/attestors/{{ .parameters.attestor }}"),

	// certificatemanager
	//
	// projects/{{project}}/locations/global/certificates/{{name}}
	"google_certificate_manager_certificate": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/locations/global/certificates/{{ .external_name }}"),
	// projects/{{project}}/locations/global/dnsAuthorizations/{{name}}
	"google_certificate_manager_dns_authorization": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/locations/global/dnsAuthorizations/{{ .external_name }}"),

	// cloudasset
	//
	// folders/{{folder_id}}/feeds/{{name}}
	"google_cloud_asset_folder_feed": config.IdentifierFromProvider,
	// organizations/{{org_id}}/feeds/{{name}}
	"google_cloud_asset_organization_feed": config.IdentifierFromProvider,
	// projects/{{project}}/feeds/{{name}}
	"google_cloud_asset_project_feed": config.IdentifierFromProvider,

	// clouddeploy
	//
	// projects/{{project}}/locations/{{location}}/deliveryPipelines/{{name}}
	"google_clouddeploy_delivery_pipeline": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/locations/{{ .parameters.location }}/deliveryPipelines/{{ .external_name }}"),
	// projects/{{project}}/locations/{{location}}/targets/{{name}}
	"google_clouddeploy_target": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/locations/{{ .parameters.location }}/targets/{{ .external_name }}"),

	// cloudidentity
	//
	// {{name}}
	"google_cloud_identity_group": config.IdentifierFromProvider,
	// {{name}}
	"google_cloud_identity_group_membership": config.IdentifierFromProvider,

	// datacatalog
	//
	// projects/{{project}}/locations/{{region}}/entryGroups/{{entry_group}} roles/viewer
	"google_data_catalog_entry_group_iam_binding": config.IdentifierFromProvider,
	// projects/{{project}}/locations/{{region}}/entryGroups/{{entry_group}} roles/viewer user:jane@example.com
	"google_data_catalog_entry_group_iam_member": config.IdentifierFromProvider,
	// projects/{{project}}/locations/{{region}}/entryGroups/{{entry_group}}
	"google_data_catalog_entry_group_iam_policy": config.IdentifierFromProvider,
	// projects/{{project}}/locations/{{region}}/tagTemplates/{{tag_template}} roles/viewer
	"google_data_catalog_tag_template_iam_binding": config.IdentifierFromProvider,
	// projects/{{project}}/locations/{{region}}/tagTemplates/{{tag_template}} roles/viewer user:jane@example.com
	"google_data_catalog_tag_template_iam_member": config.IdentifierFromProvider,
	// projects/{{project}}/locations/{{region}}/tagTemplates/{{tag_template}}
	"google_data_catalog_tag_template_iam_policy": config.IdentifierFromProvider,

	// accessapproval
	//
	// Imported by using the following format: folders/{{folder_id}}/accessApprovalSettings
	"google_folder_access_approval_settings": config.TemplatedStringAsIdentifier("", "folders/{{ .parameters.folder_id }}/accessApprovalSettings"),
	// Imported by using the following format: organizations/{{organization_id}}/accessApprovalSettings
	"google_organization_access_approval_settings": config.TemplatedStringAsIdentifier("", "organizations/{{ .parameters.organization_id }}/accessApprovalSettings"),
	// Imported by using the following format: projects/{{project_id}}/accessApprovalSettings
	"google_project_access_approval_settings": config.TemplatedStringAsIdentifier("", "projects/{{ .parameters.project_id }}/accessApprovalSettings"),

	// accesscontextmanager
	//
	//
	// Imported by using the following format: {{name}}: accessPolicies/{policy_id}/accessLevels/{short_name}
	"google_access_context_manager_access_level": config.IdentifierFromProvider,
	// No import
	"google_access_context_manager_access_level_condition": config.IdentifierFromProvider,
	// Imported by using {{name}}, but name doesn't exist in parameters. Try using IdentifierFromProvider
	"google_access_context_manager_access_policy": config.IdentifierFromProvider,
	// Imported by using the following format: accessPolicies/{{access_policy}} roles/accesscontextmanager.policyAdmin
	"google_access_context_manager_access_policy_iam_binding": config.IdentifierFromProvider,
	// Imported by using the following format: "accessPolicies/{{access_policy}} roles/accesscontextmanager.policyAdmin user:jane@example.com"
	"google_access_context_manager_access_policy_iam_member": config.IdentifierFromProvider,
	// Imported by using the following format: accessPolicies/{{access_policy}}
	"google_access_context_manager_access_policy_iam_policy": config.TemplatedStringAsIdentifier("name", "accessPolicies/{{ .external_name }}"),
	// Imported by using the following format: {{name}}
	"google_access_context_manager_gcp_user_access_binding": config.IdentifierFromProvider,
	// Imported by using the following format: {{name}}
	"google_access_context_manager_service_perimeter": config.IdentifierFromProvider,
	// Imported by using the following format: {{perimeter_name}}/{{resource}}
	"google_access_context_manager_service_perimeter_resource": config.TemplatedStringAsIdentifier("resource", "{{ .parameters.perimeter_name }}/{{ .external_name }}"),

	// activedirectory
	//
	// Imported by using the following format: projects/{{project}}/locations/global/domains/{{domain}}/{{target_domain_name}}
	"google_active_directory_domain_trust": config.TemplatedStringAsIdentifier("target_domain_name", "projects/{{ .setup.configuration.project }}/locations/global/domains/{{ .parameters.domain }}/{{ .external_name }}"),

	// apigee
	//
	// Imported by using the following format: {{org_id}}/endpointAttachments/{{endpoint_attachment_id}}
	"google_apigee_endpoint_attachment": config.TemplatedStringAsIdentifier("endpoint_attachment_id", "{{ .parameters.org_id }}/endpointAttachments/{{ .external_name }}"),
	// Imported by using the following format: {{envgroup_id}}/attachments/{{name}}. Name doesn't exist in parameters, try using IdentifierFromProvider
	"google_apigee_envgroup_attachment": config.IdentifierFromProvider,
	// Imported by using the following format: {{org_id}}/environments/{{environment}} roles/viewer
	"google_apigee_environment_iam_binding": config.IdentifierFromProvider,
	// Imported by using the following format: {{org_id}}/environments/{{environment}} roles/viewer user:jane@example.com
	"google_apigee_environment_iam_member": config.IdentifierFromProvider,
	// Imported by using the following format: {{org_id}}/environments/{{environment}}
	"google_apigee_environment_iam_policy": config.IdentifierFromProvider,
	// Imported by using the following format: {{instance_id}}/attachments/{{name}}. Name doesn't exist in parameters, try using IdentifierFromProvider
	"google_apigee_instance_attachment": config.IdentifierFromProvider,
	// Imported by using the following format: {{instance_id}}/natAddresses/{{name}}
	"google_apigee_nat_address": config.TemplatedStringAsIdentifier("name", "{{ .parameters.instance_id }}/natAddresses/{{ .external_name }}"),
	// Imported by using the following format: organizations/{{name}}/syncAuthorization
	"google_apigee_sync_authorization": config.TemplatedStringAsIdentifier("name", "organizations/{{ .external_name }}/syncAuthorization"),

	// apikeys
	//
	// Imported by using the following format: projects/{{project}}/locations/global/keys/{{name}}
	"google_apikeys_key": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/locations/global/keys/{{ .external_name }}"),
	// datalossprevention
	//
	// Imported by using the following format: {{parent}}/deidentifyTemplates/{{name}}
	"google_data_loss_prevention_deidentify_template": config.IdentifierFromProvider,
	// Imported by using the following format: {{parent}}/inspectTemplates/{{name}}
	"google_data_loss_prevention_inspect_template": config.IdentifierFromProvider,
	// Imported by using the following format: {{parent}}/jobTriggers/{{name}}
	"google_data_loss_prevention_job_trigger": config.IdentifierFromProvider,
	// Imported by using the following format: {{parent}}/storedInfoTypes/{{name}}
	"google_data_loss_prevention_stored_info_type": config.IdentifierFromProvider,

	// dataplex
	//
	// Imported by using the following format: projects/{{project}}/locations/{{location}}/lakes/{{name}}
	"google_dataplex_lake": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/locations/{{ .parameters.location }}/lakes/{{ .external_name }}"),

	// dataproc
	//
	// Imported by using the following format: projects/{project}/regions/{region}/clusters/{cluster} roles/editor
	"google_dataproc_cluster_iam_binding": config.TemplatedStringAsIdentifier("cluster", "projects/{ .setup.configuration.project }/regions/{{ .parameters.region }}/clusters/{{ .external_name }} {{ .parameters.role }}"),
	// Imported by using the following format: projects/{project}/regions/{region}/clusters/{cluster} roles/editor user:jane@example.com
	"google_dataproc_cluster_iam_member": config.TemplatedStringAsIdentifier("cluster", "projects/{ .setup.configuration.project }/regions/{{ .parameters.region }}/clusters/{{ .external_name }} {{ .parameters.role }} {{ .parameters.member }}"),
	// Imported by using the following format: projects/{project}/regions/{region}/clusters/{cluster}
	"google_dataproc_cluster_iam_policy": config.TemplatedStringAsIdentifier("cluster", "projects/{ .setup.configuration.project }/regions/{{ .parameters.region }}/clusters/{{ .external_name }}"),
	// Imported by using the following format: projects/{project}/regions/{region}/jobs/{job_id} roles/editor
	"google_dataproc_job_iam_binding": config.TemplatedStringAsIdentifier("job_id", "projects/{ .setup.configuration.project }/regions/{{ .parameters.region }}/jobs/{{ .external_name }} {{ .parameters.role }}"),
	// Imported by using the following format: projects/{project}/regions/{region}/jobs/{job_id} roles/editor user:jane@example.com
	"google_dataproc_job_iam_member": config.TemplatedStringAsIdentifier("job_id", "projects/{ .setup.configuration.project }/regions/{{ .parameters.region }}/jobs/{{ .external_name }} {{ .parameters.role }} {{ .parameters.member }}"),
	// Imported by using the following format: projects/{project}/regions/{region}/jobs/{job_id}
	"google_dataproc_job_iam_policy": config.TemplatedStringAsIdentifier("job_id", "projects/{ .setup.configuration.project }/regions/{{ .parameters.region }}/jobs/{{ .external_name }}"),

	// deploymentmanager
	//
	// Imported by using the following format: projects/{{project}}/deployments/{{name}}
	"google_deployment_manager_deployment": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/deployments/{{ .external_name }}"),

	// dialogflow
	//
	// Imported by using the following format: {{project}}
	"google_dialogflow_agent": config.ParameterAsIdentifier("project"),
	// Imported by using the following format: {{name}}
	"google_dialogflow_entity_type": config.IdentifierFromProvider,
	// Imported by using the following format: {{name}}
	"google_dialogflow_fulfillment": config.IdentifierFromProvider,
	// Imported by using the following format: {{name}}
	"google_dialogflow_intent": config.IdentifierFromProvider,

	// healthcare
	//
	// Imported by using the following format: {{dataset}}/consentStores/{{consent_store}} roles/viewer user:jane@example.com
	"google_healthcare_consent_store_iam_member": config.TemplatedStringAsIdentifier("consent_store_id", "{{ .parameters.dataset }}/consentStores/{{ .external_name }} {{ .parameters.role }} {{ .parameters.member }}"),
	// Imported by using the following format: your-project-id/location-name/dataset-name roles/viewer user:foo@example.com
	"google_healthcare_dataset_iam_member": config.TemplatedStringAsIdentifier("", "{{ .parameters.dataset_id }} {{ .parameters.role }} {{ .parameters.member }}"),
	// Imported by using the following format: {{dataset}}/dicomStores/{{name}}
	"google_healthcare_dicom_store": config.TemplatedStringAsIdentifier("name", "{{ .parameters.dataset }}/dicomStores/{{ .external_name }}"),
	// Imported by using the following format: your-project-id/location-name/dataset-name/dicom-store-name roles/viewer user:foo@example.com
	"google_healthcare_dicom_store_iam_member": config.TemplatedStringAsIdentifier("", "{{ .parameters.dicom_store_id }} {{ .parameters.role }} {{ .parameters.member }}"),
	// Imported by using the following format: {{dataset}}/fhirStores/{{name}}
	"google_healthcare_fhir_store": config.TemplatedStringAsIdentifier("name", "{{ .parameters.dataset }}/fhirStores/{{ .external_name }}"),
	// Imported by using the following format: your-project-id/location-name/dataset-name/fhir-store-name roles/viewer user:foo@example.com
	"google_healthcare_fhir_store_iam_member": config.TemplatedStringAsIdentifier("", "{{ .parameters.fhir_store_id }} {{ .parameters.role }} {{ .parameters.member }}"),
	// Imported by using the following format: {{dataset}}/hl7V2Stores/{{name}}
	"google_healthcare_hl7_v2_store": config.TemplatedStringAsIdentifier("name", "{{ .parameters.dataset }}/hl7V2Stores/{{ .external_name }}"),
	// Imported by using the following format: your-project-id/location-name/dataset-name/hl7-v2-store-name roles/viewer user:foo@example.com
	"google_healthcare_hl7_v2_store_iam_member": config.TemplatedStringAsIdentifier("", "{{ .parameters.hl7_v2_store_id }} {{ .parameters.role }} {{ .parameters.member }}"),

	// iap
	//
	// Imported by using the following format: projects/{{project_id}}/brands/{{brand_id}}
	"google_iap_brand": config.IdentifierFromProvider,
	// Imported by using the following format: {{brand}}/identityAwareProxyClients/{{client_id}}
	"google_iap_client": config.IdentifierFromProvider,
	// Imported by using the following format: projects/{{project}}/iap_web/appengine-{{appId}}/services/{{service}} roles/iap.httpsResourceAccessor user:jane@example.com
	"google_iap_app_engine_service_iam_member": config.TemplatedStringAsIdentifier("service", "projects/{{ .setup.configuration.project }}/iap_web/appengine-{{ .parameters.app_id }}/services/{{ .external_name }} {{ .parameters.role }} {{ .parameters.member }}"),
	// Imported by using the following format: projects/{{project}}/iap_web/appengine-{{appId}}/services/{{service}}/versions/{{versionId}} roles/iap.httpsResourceAccessor user:jane@example.com
	"google_iap_app_engine_version_iam_member": config.TemplatedStringAsIdentifier("version_id", "projects/{{ .setup.configuration.project }}/iap_web/appengine-{{ .parameters.app_id }}/services/{{ .parameters.service }}/versions/{{ .external_name }} {{ .parameters.role }} {{ .parameters.member }}"),
	// Imported by using the following format: projects/{{project}}/iap_web roles/iap.httpsResourceAccessor user:jane@example.com
	"google_iap_tunnel_iam_member": config.TemplatedStringAsIdentifier("", "/projects/{{ .setup.configuration.project }}/iap_web {{ .parameters.role }} {{ .parameters.member }}"),
	// logging
	//
	// Billing account logging exclusions can be imported using their URI
	// billingAccounts/my-billing_account/exclusions/my-exclusion
	"google_logging_billing_account_exclusion": config.TemplatedStringAsIdentifier("name", "billingAccounts/{{ .parameters.billing_account }}/exclusions/{{ .external_name }}"),
	// Billing account logging sinks can be imported using this format: billingAccounts/{{billing_account_id}}/sinks/{{sink_id}}
	"google_logging_billing_account_sink": config.IdentifierFromProvider,
	// This resource can be imported using the following format: folders/{{folder}}/locations/{{location}}/buckets/{{bucket_id}}
	"google_logging_folder_bucket_config": config.TemplatedStringAsIdentifier("", "folders/{{ .parameters.folder }}/locations/{{ .parameters.location }}/buckets/{{ .parameters.bucket_id }}"),
	// Folder-level logging exclusions can be imported using their URI
	// folders/my-folder/exclusions/my-exclusion
	"google_logging_folder_exclusion": config.TemplatedStringAsIdentifier("name", "folders/{{ .parameters.folder }}/exclusions/{{ .external_name }}"),
	// Folder-level logging sinks can be imported using this format: folders/{{folder_id}}/sinks/{{name}}
	"google_logging_folder_sink": config.TemplatedStringAsIdentifier("name", "folders/{{ .parameters.folder }}/sinks/{{ .external_name }}"),
	// LogView can be imported using any of these accepted formats: {{parent}}/locations/{{location}}/buckets/{{bucket}}/views/{{name}}
	// TODO: {{parent}} and {{location}} fields are optional in documentation. Check CRD if they are required and if not use config.IdentifierFromProvider
	"google_logging_log_view": config.TemplatedStringAsIdentifier("name", "{{ .parameters.parent }}/locations/{{ .parameters.location }}/buckets/{{ .parameters.bucket }}/views/{{ .external_name }}"),
	// Metric can be imported using Name
	"google_logging_metric": config.NameAsIdentifier,
	// This resource can be imported using the following format: organizations/{{organization}}/locations/{{location}}/buckets/{{bucket_id}}
	"google_logging_organization_bucket_config": config.TemplatedStringAsIdentifier("", "organizations/{{ .parameters.organization }}/locations/{{ .parameters.location }}/buckets/{{ .parameters.bucket_id }}"),
	// Organization-level logging exclusions can be imported using their URI
	// organizations/{{organization}}/exclusions/{{name}}
	"google_logging_organization_exclusion": config.TemplatedStringAsIdentifier("name", "organizations/{{ .parameters.org_id }}/exclusions/{{ .external_name }}"),
	// Organization-level logging sinks can be imported using this format: organizations/{{organization_id}}/sinks/{{sink_id}}
	"google_logging_organization_sink": config.IdentifierFromProvider,
	// This resource can be imported using the following format: projects/{{project}}/locations/{{location}}/buckets/{{bucket_id}}
	"google_logging_project_bucket_config": config.TemplatedStringAsIdentifier("", "projects/{{ .parameters.project }}/locations/{{ .parameters.location }}/buckets/{{ .parameters.bucket_id }}"),
	// Project-level logging exclusions can be imported using their URI
	// projects/my-project/exclusions/my-exclusion
	"google_logging_project_exclusion": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/exclusions/{{ .external_name }}"),
	// Project-level logging sinks can be imported using their URI
	// projects/my-project/sinks/my-sink
	"google_logging_project_sink": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/sinks/{{ .external_name }}"),

	// memcache
	//
	// nstance can be imported using Name
	"google_memcache_instance": config.NameAsIdentifier,

	// mlengine
	//
	// Model can be imported using Name
	"google_ml_engine_model": config.NameAsIdentifier,

	// networkservices
	//
	// EdgeCacheKeyset can be imported using Name
	"google_network_services_edge_cache_keyset": config.NameAsIdentifier,
	// EdgeCacheOrigin can be imported using Name
	"google_network_services_edge_cache_origin": config.NameAsIdentifier,
	// EdgeCacheService can be imported using Name
	"google_network_services_edge_cache_service": config.NameAsIdentifier,

	// notebooks
	//
	// IAM binding imports use space-delimited identifiers: the resource in question and the role
	// projects/{{project}}/locations/{{location}}/instances/{{instance_name}} roles/viewer
	"google_notebooks_instance_iam_binding": config.TemplatedStringAsIdentifier("instance_name", "projects/{{ .setup.configuration.project }}/locations/{{ .parameters.location }}/instances/{{ .external_name }} {{ .parameters.role }}"),
	// IAM policy imports use the identifier of the resource in question
	// projects/{{project}}/locations/{{location}}/instances/{{instance_name}}
	"google_notebooks_instance_iam_policy": config.TemplatedStringAsIdentifier("instance_name", "projects/{{ .setup.configuration.project }}/locations/{{ .parameters.location }}/instances/{{ .external_name }}"),
	// Location can be imported using Name
	"google_notebooks_location": config.NameAsIdentifier,
	// IAM binding imports use space-delimited identifiers: the resource in question and the role
	// projects/{{project}}/locations/{{location}}/runtimes/{{runtime_name}} roles/viewer
	"google_notebooks_runtime_iam_binding": config.TemplatedStringAsIdentifier("runtime_name", "projects/{{ .setup.configuration.project }}/locations/{{ .parameters.location }}/runtimes/{{ .external_name }} {{ .parameters.role }}"),
	// IAM policy imports use the identifier of the resource in question
	// projects/{{project}}/locations/{{location}}/runtimes/{{runtime_name}}
	"google_notebooks_runtime_iam_policy": config.TemplatedStringAsIdentifier("runtime_name", "projects/{{ .setup.configuration.project }}/locations/{{ .parameters.location }}/runtimes/{{ .external_name }}"),

	// artifactregistry
	//
	// Imported by using the following format: projects/{{project}}/locations/{{location}}/repositories/{{repository_id}}
	"google_artifact_registry_repository": config.TemplatedStringAsIdentifier("repository_id", "projects/{{ .setup.configuration.project }}/locations/{{ .parameters.location }}/repositories/{{ .external_name }}"),
	// Imported by using the following format: projects/{{project}}/locations/{{location}}/repositories/{{repository}} roles/artifactregistry.reader user:jane@example.com
	"google_artifact_registry_repository_iam_member": config.TemplatedStringAsIdentifier("repository", "projects/{{ .setup.configuration.project }}/locations/{{ .parameters.location }}/repositories/{{ .external_name }} {{ .parameters.role }} {{ .parameters.member }}"),

	// beyondcorp
	//
	// Imported by using the following projects/{{project}}/locations/{{region}}/appConnections/{{name}}
	"google_beyondcorp_app_connection": config.IdentifierFromProvider,
	// Imported by using the following projects/{{project}}/locations/{{region}}/appConnectors/{{name}}
	"google_beyondcorp_app_connector": config.IdentifierFromProvider,
	// Imported by using the following projects/{{project}}/locations/{{region}}/appGateways/{{name}}
	"google_beyondcorp_app_gateway": config.IdentifierFromProvider,

	// bigqueryanalyticshub
	//
	// Imported by using the following projects/{{project}}/locations/{{location}}/dataExchanges/{{data_exchange_id}}
	"google_bigquery_analytics_hub_data_exchange": config.TemplatedStringAsIdentifier("data_exchange_id", "projects/{{ .setup.configuration.project }}/locations/{{ .parameters.location }}/dataExchanges/{{ .external_name }}"),
	// Imported by using the following projects/{{project}}/locations/{{location}}/dataExchanges/{{data_exchange_id}} roles/viewer user:jane@example.com
	"google_bigquery_analytics_hub_data_exchange_iam_member": config.TemplatedStringAsIdentifier("data_exchange_id", "projects/{{ .setup.configuration.project }}/locations/{{ .parameters.location }}/dataExchanges/{{ .external_name }} {{ .parameters.role }} {{ .parameters.member }}"),
	// Imported by using the following projects/{{project}}/locations/{{location}}/dataExchanges/{{data_exchange_id}}/listings/{{listing_id}}
	"google_bigquery_analytics_hub_listing": config.TemplatedStringAsIdentifier("listing_id", "projects/{{ .setup.configuration.project }}/locations/{{ .parameters.location }}/dataExchanges/{{ .parameters.data_exchange_id }}/listings/{{ .external_name }}"),
	// Imported by using the following projects/{{project}}/locations/{{location}}/dataExchanges/{{data_exchange_id}}/listings/{{listing_id}} roles/viewer user:jane@example.com
	"google_bigquery_analytics_hub_listing_iam_member": config.TemplatedStringAsIdentifier("listing_id", "projects/{{ .setup.configuration.project }}/locations/{{ .parameters.location }}/dataExchanges/{{ .parameters.data_exchange_id }}/listings/{{ .external_name }} {{ .parameters.role }} {{ .parameters.member }}"),

	// bigqueryconnection
	//
	// Imported by using the following projects/{{project}}/locations/{{location}}/connections/{{connection_id}} roles/viewer user:jane@example.com
	"google_bigquery_connection_iam_member": config.TemplatedStringAsIdentifier("connection_id", "projects/{{ .setup.configuration.project }}/locations/{{ .parameters.location }}/connections/{{ .external_name }} {{ .parameters.role }} {{ .parameters.member }}"),

	// certificatemanager
	//
	// Imported by using the following projects/{{project}}/locations/global/certificateMaps/{{name}}
	"google_certificate_manager_certificate_map": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/locations/global/certificateMaps/{{ .external_name }}"),
	// Imported by using the following projects/{{project}}/locations/global/certificateMaps/{{map}}/certificateMapEntries/{{name}}
	"google_certificate_manager_certificate_map_entry": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/locations/global/certificateMaps/{{ .parameters.map }}/certificateMapEntries/{{ .external_name }}"),

	// ids
	//
	// Imported by using the following projects/{{project}}/locations/{{location}}/endpoints/{{name}}
	"google_cloud_ids_endpoint": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/locations/{{ .parameters.location }}/endpoints/{{ .exyernal_name }}"),

	// cloudrun
	//
	// Imported by using the following projects/{{project}}/locations/{{location}}/jobs/{{name}}
	"google_cloud_run_v2_job": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/locations/{{ .parameters.location }}/jobs/{{ .external_name }}"),
	// Imported by using the following projects/{{project}}/locations/{{location}}/services/{{name}}
	"google_cloud_run_v2_service": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/locations/{{ .parameters.location }}/services/{{ .external_name }}"),

	// cloudtasks
	//
	// Imported by using the following projects/{{project}}/locations/{{location}}/queues/{{queue}} roles/viewer user:jane@example.com
	"google_cloud_tasks_queue_iam_member": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/locations/{{ .parameters.location }}/queues/{{ .external_name }} {{ .parameters.role }} {{ .parameters.member }}"),

	// cloudfunctions2
	//
	// Imported by using the following projects/{{project}}/locations/{{location}}/functions/{{name}}
	"google_cloudfunctions2_function": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/locations/{{ .parameters.location }}/functions/{{ .external_name }}"),
	// Imported by using the following projects/{{project}}/locations/{{location}}/functions/{{cloud_function}} roles/viewer user:jane@example.com
	"google_cloudfunctions2_function_iam_member": config.TemplatedStringAsIdentifier("cloud_function", "projects/{{ .setup.configuration.project }}/locations/{{ .parameters.location }}/functions/{{ .external_name }} {{ .parameters.role }} {{ .parameters.member }}"),

	// cloudiot
	//
	// Imported by using the following projects/{{project}}/locations/{{location}}/registries/{{device_registry}} roles/viewer user:jane@example.com
	"google_cloudiot_registry_iam_member": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/locations/{{ .parameters.location }}/registries/{{ .external_name }} {{ .parameters.role }} {{ .parameters.member }}"),
}
