connection "salesforce" {
  plugin = "salesforce"

  # Your salesforce instance URL, for example, 'https://na01.salesforce.com/'
  # url       = "YOUR_SALESFORCE_URL"

  # Username of the Salesforce account
  # user      = "YOUR_SALESFORCE_USERNAME"

  # Password of the Salesforce account
  # password  = "YOUR_SALESFORCE_PASSWORD"

  # Security token, could be omitted if Trusted IP is configured
  # token     = "YOUR_SALESFORCE_TOKEN"

  # Client Id of Salesforce from Connected App"
  # client_id = "YOUR_SALESFORCE_CLIENT_ID"

  # List of salesforce tables to be generated. Salesforce Object API Name (refer https://developer.salesforce.com/docs/atlas.en-us.api.meta/api/sforce_api_objects_list.htm) should be used here."
  # Currently plugin supports a standard set of Salesforce tables by default in the plugin. If there is a need to query an standard object from, you can add standard object in the tables list. Steampipe will automatically generate respective table for the connection.
  # tables = ["Case", "Campaign", "CustomApp__c"]

  # Version of the salesforce API
  # api_version = "43.0"
}
