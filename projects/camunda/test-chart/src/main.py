from urllib.request import urlopen
import json

def main():
    print('Sending request...')
    url = "http://camunda-keycloak.camunda.svc.cluster.local/auth/realms/camunda-platform/protocol/openid-connect/certs"
  
    # store the response of URL
    response = urlopen(url)
    
    # storing the JSON response 
    # from url in data
    data_json = json.loads(response.read())
    
    # print the json response
    print(data_json)
    exit(0)


if __name__ == '__main__':
    main()