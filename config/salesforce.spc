connection "salesforce" {
  plugin = "salesforce"

  # Your salesforce instance URL, for example, 'https://na01.salesforce.com/'
  # url       = "YOUR_SALESFORCE_URL"

  # Username of the Salesforce account
  # username  = "YOUR_SALESFORCE_USERNAME"

  # Password of the Salesforce account
  # password  = "YOUR_SALESFORCE_PASSWORD"

  # The Salesforce security token is only required If the client's IP address is not added to the organization's list of trusted IPs
  # https://help.salesforce.com/s/articleView?id=sf.security_networkaccess.htm&type=5
  # https://migration.trujay.com/help/how-to-add-an-ip-address-to-whitelist-on-salesforce/
  # token     = "YOUR_SALESFORCE_TOKEN"

  # Client Id of Connected App from Salesforce
  # client_id = "YOUR_SALESFORCE_CLIENT_ID"

  # List of salesforce tables to be generated. Salesforce Object API Name (refer https://developer.salesforce.com/docs/atlas.en-us.api.meta/api/sforce_api_objects_list.htm) should be used here."
  # Currently plugin supports a standard set of Salesforce tables by default in the plugin. If there is a need to query an standard object from, you can add standard object in the objects list. Steampipe will automatically generate respective table for the connection.
  # objects = ["Case", "Campaign", "CustomApp__c"]

  # Version of the salesforce API
  # api_version = "43.0"
}
