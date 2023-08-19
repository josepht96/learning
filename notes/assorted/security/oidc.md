# OIDC

## keycloak to freeipa

Keycloak connects to FreeIPA, or any Idm provider, via the LDAP protocol. LDAP uses a bind user (a general admin account) to query the idm database. Via TCP, this information can be sent back to Keycloak, which it is stored in a keycloak database - user federation. Communcation between Keycloak and FreeIpa is not continuous. If updates are made to FreeIpa, Keycloak must synchronize this data (unsure if its possible to configure this to automatically happen). Users and group data can be viewed in Keycloak.

## kubernetes to keycloak

From the AWS console EKS users are able to see the OIDC configuration for a specific cluster. This configuration does not estalbish any communication between EKS and your OIDC provider (in this case Keycloak). This configuration directs the cluster on how to handle OIDC tokens attached to API requests. It also tells EKS how to parse an Id token. When a user authenticates to EKS, it uses this configuration to determine if the users id token is valid. User validation does not require communcation between EKS and the OIDC provider. 

## kubelogin

To access EKS over OIDC, a user must include a valid id token in their API request. This process requires uses to retrieve an id token on their own and manually configure a kubeconfig user with it. Kubelogin is an open-source kubectl plugin that handles this process for the user. Users must know their OIDC issuer url, client id, client secret, and have valid OIDC credentials. Kubelogin invoked with the authorization code flow operates similarly to how OIDC generally works for web applications. 
1. Kubelogin directs user to OIDC issuer login page
2. A local server is started which handle upcoming communication  
4. User credentials are provided, authorizing the request
5. An authorization token is returned to the users local server via OIDC callback URI (this needs to be configured prior to this process)
6. Kubelogin exchanges auth token for access token, refresh token, id token.
7. id token and refresh token are added to the kubectl cli
8. When invoked, kubectl will pass id token to EKS API server for validation
9. API Server, based on EKS OIDC configuration, will parse ID token to determine user permissions

1. Start a local server at the port.
2. Open a browser and navigate it to the local server.
3. Wait for the user authorization.
4. Receive a code via an authorization response (HTTP redirect).
5. Exchange the code and a token.
6. Return the code. 

When done via ROPC, service user credentials can be provided via enivironment variables and the authorization code stage is skipped. 
client id, client secret, scope, and user credentials can be provided directly to the OIDC provider in exchange for access, refresh, and id tokens.