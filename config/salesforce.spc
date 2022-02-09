connection "salesforce" {
  plugin = "salesforce"

  # Salesforce instance URL, e.g., "https://na01.salesforce.com/"
  # url = "YOUR_SALESFORCE_URL"

  # Salesforce account name
  # user = "YOUR_SALESFORCE_USERNAME"

  # Salesforce account password
  # password = "YOUR_SALESFORCE_PASSWORD"

  # The Salesforce security token is only required If the client's IP address is not added to the organization's list of trusted IPs
  # https://help.salesforce.com/s/articleView?id=sf.security_networkaccess.htm&type=5
  # https://migration.trujay.com/help/how-to-add-an-ip-address-to-whitelist-on-salesforce/
  # token     = "YOUR_SALESFORCE_TOKEN"

  # Client Id of Connected App from Salesforce
  # Client id is optional. https://developer.salesforce.com/docs/atlas.en-us.214.0.api.meta/api/sforce_api_calls_login.htm
  # client_id = "YOUR_SALESFORCE_CLIENT_ID"

  # List of Salesforce object names to generate additional tables for
  # This argument only accepts exact Salesforce standard and custom object names, e.g., AccountBrand, OpportunityStage, CustomApp__c
  # For a full list of standard object names, please see https://developer.salesforce.com/docs/atlas.en-us.api.meta/api/sforce_api_objects_list.htm)
  # All custom object names should end in "__c", following Salesforce object naming
  # objects = ["AccountBrand", "OpportunityStage", "CustomApp__c"]

  # Salesforce API version to connect to.
  # api_version = "43.0"
}
