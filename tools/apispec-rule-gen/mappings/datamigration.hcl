mapping "azurerm_database_migration_project" {
  import_path = "azure-rest-api-specs/specification/datamigration/resource-manager/Microsoft.DataMigration/stable/2018-04-19/definitions/Projects.json"

  source_platform = ProjectSourcePlatform
  target_platform = ProjectTargetPlatform
}

mapping "azurerm_database_migration_service" {
  import_path = "azure-rest-api-specs/specification/datamigration/resource-manager/Microsoft.DataMigration/stable/2018-04-19/definitions/Services.json"

  subnet_id = DataMigrationServiceProperties.virtualSubnetId
  sku_name  = ServiceSku.name
}
