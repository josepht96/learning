provider "azurerm" {
  features {}
}

terraform {
#   backend "azurerm" {
#     resource_group_name  = "test-business-services-shared"
#     storage_account_name = "cmtestnonprodtf"
#     container_name       = "cm-test-nonprod-tfstate"
#     key                  = "cm-test-app-svc-test.tfstate"
#   }
}

############################ East US  ##########################################

module "appsvc-eastus" {
  source                 = "./modules"
  resource_group_name    = "test-resource-group"
  plan_name              = "plan-bdservices-test-eastus-001"
  env                    = "dev"
  location               = "eastus"
  kind                   = "Windows"
  per_site_scaling       = false
  size                   = "P1V2"
  tier                   = "Premium"
  hybrid_host_name       = "FHIRDB01"
  hybrid_port_number     = 1433
  hybrid_connection_name = "FHIRDB01"
  app_names = [
    "app-ConditionAPI-test-eastus-001",
    "app-EncounterAPI-test-eastus-001",
    "app-PatientAPI-test-eastus-001",
    "app-DocumentReferenceAPI-test-eastus-001",
    "app-AllergyIntoleranceAPI-test-eastus-001",
    "app-CarePlanAPI-test-eastus-001",
    "app-DeviceAPI-test-eastus-001",
    "app-GoalAPI-test-eastus-001",
    "app-ImmunizationAPI-test-eastus-001",
    "app-LocationAPI-test-eastus-001",
    "app-MedicationRequestAPI-test-eastus-001",
    "app-ObservationAPI-test-eastus-001",
    "app-OrganizationAPI-test-eastus-001",
    "app-PractitionerAPI-test-eastus-001",
    "app-ProcedureAPI-test-eastus-001",
    # "app-GraphqlAPI-test-eastus-001",
    # "app-CareTeam-test-eastus-001",
    # "app-RelatedPerson-test-eastus-001",
    # "app-UserPreference-test-eastus-001",
    # "app-FhirInfrastructure-test-eastus-001",
    # "app-PractitionerRole-test-eastus-001",
    # "app-DiagnosticReportAPI-test-eastus-001"
  ]

  service_bus_name  = "test-test-servicebus-eastus"
  app_insights_name = "appi-test-test-eastus-001"

  api_management_name = "apim-test-bdservices-test-eastus"
  api_management_sku  = "Developer_1"
  publisher_name      = "Joe"
  publisher_email     = "joe.thomas@allscripts.com"

  api_revision = [
    "v1",
    "v1",
    "v1",
    "v1",
    "v1",
    "v1",
    "v1",
    "v1",
    "v1",
    "v1",
    "v1",
    "v1",
    "v1",
    "v1",
    "v1",
    # "v1",
    # "v1",
    # "v1",
    # "v1",
    # "v1",
    # "v1"
  ]
  api_display_names = [
    "Condition API",
    "Encounter API",
    "Patient API",
    "Document Reference API",
    "Allergy Intolerance API",
    "Care Plan API",
    "Device API",
    "Goal API",
    "Immunization API",
    "Location API",
    "Medication Request API",
    "Observation API",
    "Organization API",
    "Practitioner API",
    "Procedure API",
    # "DiagnosticReport API",
    # "CareTeam API",
    # "RelatedPerson API",
    # "UserPreference API",
    # "FhirInfrastructure API",
    # "PractitionerRole API"
  ]
  api_path = [
    "condition",
    "encounter",
    "patient",
    "document-reference",
    "allergy-intolerance",
    "care-plan",
    "device",
    "goal",
    "immunization",
    "location",
    "medication-request",
    "observation",
    "organization",
    "practitioner",
    "procedure",
    # "DiagnosticReport",
    # "careteam",
    # "relatedperson",
    # "userpreference",
    # "fhirinfrastructure",
    # "practitionerrole"
  ]

  api_content_value = [
    # "https://app-ConditionAPI-test-eastus-001.azurewebsites.net/swagger/v1/swagger.json",
    # "https://app-EncounterAPI-test-eastus-001.azurewebsites.net/swagger/v1/swagger.json",
    # "https://app-PatientAPI-test-eastus-001.azurewebsites.net/swagger/v1/swagger.json",
    # "https://app-DocumentReferenceAPI-test-eastus-001.azurewebsites.net/swagger/DocumentReferenceController/swagger.json",
    # "https://app-AllergyIntoleranceAPI-test-eastus-001.azurewebsites.net/swagger/v1/swagger.json",
    # "https://app-CarePlanAPI-test-eastus-001.azurewebsites.net/swagger/v1/swagger.json",
    # "https://app-DeviceAPI-test-eastus-001.azurewebsites.net/swagger/v1/swagger.json",
    # "https://app-GoalAPI-test-eastus-001.azurewebsites.net/swagger/v1/swagger.json",
    # "https://app-ImmunizationAPI-test-eastus-001.azurewebsites.net/swagger/v1/swagger.json",
    # "https://app-LocationAPI-test-eastus-001.azurewebsites.net/swagger/v1/swagger.json",
    # "https://app-MedicationRequestAPI-test-eastus-001.azurewebsites.net/swagger/v1/swagger.json",
    # "https://app-ObservationAPI-test-eastus-001.azurewebsites.net/swagger/v1/swagger.json",
    # "https://app-OrganizationAPI-test-eastus-001.azurewebsites.net/swagger/v1/swagger.json",
    # "https://app-PractitionerAPI-test-eastus-001.azurewebsites.net/swagger/v1/swagger.json",
    # "https://app-ProcedureAPI-test-eastus-001.azurewebsites.net/swagger/v1/swagger.json",
    # # "https://app-DiagnosticReportAPI-test-eastus-001.azurewebsites.net/swagger/v1/swagger.json",
    # # "https://app-CareAPI-test-eastus-001.azurewebsites.net/swagger/v1/swagger.json",
    # # "https://app-RelatedPersonAPI-test-eastus-001.azurewebsites.net/swagger/v1/swagger.json",
    # # "https://app-UserPreferenceAPI-test-eastus-001.azurewebsites.net/swagger/v1/swagger.json",
    # # "https://app-FhirInfrastructureAPI-test-eastus-001.azurewebsites.net/swagger/v1/swagger.json",
    # # "https://app-PractitionerRoleAPI-test-eastus-001.azurewebsites.net/swagger/v1/swagger.json",
    # # "https://app-GraphqlAPI-test-eastus-001.azurewebsites.net/swagger/v1/swagger.json"
  ]

  tags = {
    "organization" = "",
    "sub-org"      = "",
    "department"   = "test",
    "project"      = "platform-of-health",
  }

}
