connection "salesforce" {
  plugin = "salesforce"

  # Salesforce instance URL, e.g., "https://na01.salesforce.com/"
  # url = "https://na01.salesforce.com/"

  # Salesforce account name
  # username = "user@example.com"

  # Salesforce account password
  # password = "Dummy@~Password"

  # The Salesforce security token is only required If the client's IP address is not added to the organization's list of trusted IPs
  # https://help.salesforce.com/s/articleView?id=sf.security_networkaccess.htm&type=5
  # token = "ABO5C3PNqOP0BHsPFakeToken"

  # Salesforce client ID of the connected app
  # client_id = "3MVG99E3Ry5mh4z_FakeID"

  # List of Salesforce object names to generate additional tables for
  # This argument only accepts exact Salesforce standard and custom object names, e.g., AccountBrand, OpportunityStage, CustomApp__c
  # For a full list of standard object names, please see https://developer.salesforce.com/docs/atlas.en-us.api.meta/api/sforce_api_objects_list.htm)
  # All custom object names should end in "__c", following Salesforce object naming standards
  # objects = ["AccountBrand", "OpportunityStage", "CustomApp__c"]

  # Salesforce API version to connect to
  # api_version = "43.0"

  # The naming_convention allows users to control the naming format for tables and columns in the plugin. Below are the supported values:
  # api_native: If set to this value, the plugin will use the native format for table names, meaning there will be no salesforce_ prefix, and the table and column names will remain as they are in Salesforce.
  # snake_case: If the user does not specify any value or sets it to this value, the plugin will use snake case for table and column names and table names will have a salesforce_ prefix.
  # naming_convention = "snake_case"
}
