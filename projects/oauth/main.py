import requests

# Replace these with your actual values
oauth_url = "https://auth-service.com/oauth/token"  # OAuth token URL
client_id = "your-client-id"
client_secret = "your-client-secret"
username = "your-username"
password = "your-password"

# Step 1: Retrieve the OAuth token
def get_oauth_token():
    payload = {
        'grant_type': 'password',
        'client_id': client_id,
        'client_secret': client_secret,
        'username': username,
        'password': password
    }
    
    response = requests.post(oauth_url, data=payload)
    
    if response.status_code == 200:
        token_data = response.json()
        return token_data['access_token']
    else:
        print(f"Failed to retrieve token: {response.status_code}, {response.text}")
        return None

# Step 2: Send a GET request to the service
def make_authorized_request(token):
    headers = {
        'Authorization': f"Bearer {token}"
    }
    
    response = requests.get("https://my-service.com/test", headers=headers)
    print(f"Response Code: {response.status_code}")

# Main function
def main():
    token = get_oauth_token()
    if token:
        make_authorized_request(token)

if __name__ == "__main__":
    main()
