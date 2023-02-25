# Various things

## JWT

Keycloak clients are entities that are used to enable applications to authenticate and control authorization with Keycloak.
Let's say there are two applications: a frontend application for end-users, and an admin console. Frontend applications
should not receive the same permissions as the admin console.
Separate clients can be configured in Keycloak to handle the needs of both applications. We create two clients, frontend-client and admin-console-client. Frontend-client is restricted to the access-user-api role, which enables authenticated resources to access the user API.
Admin-console-client is restricted to the access-user-database role, which enables authenticated resources to directly access the user database.

When calls to a user database are requested, frontend applications request an access token from Keycloak before invoking the API request.
The frontend-client client_id and client_secret are passed to the Keycloak token endpoint. Notice this endpoint is not specific to
the frontend-client client.

```bash
CLIENT_ID=frontend-client
CLIENT_SECRET=IlMU8goi3qcDMA7HkDIY5Vo7ZBbK3Guj # client secret is created when authorization is enabled on the client
GRANT_TYPE=client_credentials
KEYCLOAK_TOKEN_ENDPOINT=http://localhost:30000/realms/master/protocol/openid-connect/token

curl -X POST $KEYCLOAK_TOKEN_ENDPOINT \
    -H "Content-Type: application/x-www-form-urlencoded" \
    -d "client_id=$CLIENT_ID&client_secret=$CLIENT_SECRET&grant_type=$GRANT_TYPE"
```

Before calls to the database are made from the admin console, a smiliar action occurs:

```bash
CLIENT_ID=admin-console-client
CLIENT_SECRET=dgb7ieXcuIwOwfjsxQapIuNIP75ysWFs
GRANT_TYPE=client_credentials
KEYCLOAK_TOKEN_ENDPOINT=http://localhost:30000/realms/master/protocol/openid-connect/token

curl -X POST $KEYCLOAK_TOKEN_ENDPOINT \
    -H "Content-Type: application/x-www-form-urlencoded" \
    -d "client_id=$CLIENT_ID&client_secret=$CLIENT_SECRET&grant_type=$GRANT_TYPE"
```

The frontend-client returns the following JSON object:

```json
{
    "access_token": "eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJSNGRKcGgyQ2pnbFc1VlhGZGdhSklQdzlfWlBWSVZHMnpfMGMwdkt5MHBRIn0.eyJleHAiOjE2NzczNTQyOTIsImlhdCI6MTY3NzM1NDIzMiwianRpIjoiZTQyYzJlOTEtODA2MC00YmY3LTljZGUtZjg3NjdmMzYzY2Q3IiwiaXNzIjoiaHR0cDovL2xvY2FsaG9zdDozMDAwMC9yZWFsbXMvbWFzdGVyIiwiYXVkIjoiYWNjb3VudCIsInN1YiI6IjQwYWMwZTQ2LTViNmYtNDNkNi05NzkwLTA4NGJkMzg3NmY0YyIsInR5cCI6IkJlYXJlciIsImF6cCI6ImZyb250ZW5kLWNsaWVudCIsImFjciI6IjEiLCJyZWFsbV9hY2Nlc3MiOnsicm9sZXMiOlsiZGVmYXVsdC1yb2xlcy1tYXN0ZXIiLCJhY2Nlc3MtdG8tdXNlci1hcGkiLCJvZmZsaW5lX2FjY2VzcyIsInVtYV9hdXRob3JpemF0aW9uIl19LCJyZXNvdXJjZV9hY2Nlc3MiOnsiYWNjb3VudCI6eyJyb2xlcyI6WyJtYW5hZ2UtYWNjb3VudCIsIm1hbmFnZS1hY2NvdW50LWxpbmtzIiwidmlldy1wcm9maWxlIl19fSwic2NvcGUiOiJwcm9maWxlIGVtYWlsIiwiY2xpZW50SWQiOiJmcm9udGVuZC1jbGllbnQiLCJlbWFpbF92ZXJpZmllZCI6ZmFsc2UsImNsaWVudEhvc3QiOiIxNzIuMTguMC4zIiwicHJlZmVycmVkX3VzZXJuYW1lIjoic2VydmljZS1hY2NvdW50LWZyb250ZW5kLWNsaWVudCIsImNsaWVudEFkZHJlc3MiOiIxNzIuMTguMC4zIn0.r6TD5KtPZto5986FPu0yKluGDU_LtUGae1-8CAy0UEKQVy3U3F6i8wWw2SFj8BuBdLpFZPXOQog-lNDr2dnWKAj7cZgDYX3xIYYxTShV40hwMklIQ9JGEdZ-tKE-0KZs3wKcbDR0SSHmaVB-ljG8_R3WlDkbT2L6h31OyZjXoEdXDLM3luWeMYmaPmh4Jjk1lrctII7UNij8e3I2jtKZ5hqZPA-EzH_XBV12kMGfK67KtRcBkTxSrrX3h2d2LfsvpH52zixVYIdttePG-zGkmByHCNfPSCs0INAkB0W7OfnwiZZ4lf2WrwHdZLoI39K2krpH7IDP1vh2Jya4NvSHZA",
    "expires_in": 60,
    "refresh_expires_in": 0,
    "token_type": "Bearer",
    "not-before-policy": 0,
    "scope": "profile email"
}
```

The admin-console-client returns the following JSON object:

```json
{
    "access_token": "eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJSNGRKcGgyQ2pnbFc1VlhGZGdhSklQdzlfWlBWSVZHMnpfMGMwdkt5MHBRIn0.eyJleHAiOjE2NzczNTQyNDIsImlhdCI6MTY3NzM1NDE4MiwianRpIjoiNzQ5OWMwMzAtYTM4OC00MzUwLWE2ZTYtZmMyODZiN2YyMzBkIiwiaXNzIjoiaHR0cDovL2xvY2FsaG9zdDozMDAwMC9yZWFsbXMvbWFzdGVyIiwiYXVkIjoiYWNjb3VudCIsInN1YiI6ImFkMTQxZDk2LWUzNjUtNDhiOS05ODY0LTNkYzc3NTIzNmEwNSIsInR5cCI6IkJlYXJlciIsImF6cCI6ImFkbWluLWNvbnNvbGUtY2xpZW50IiwiYWNyIjoiMSIsInJlYWxtX2FjY2VzcyI6eyJyb2xlcyI6WyJkZWZhdWx0LXJvbGVzLW1hc3RlciIsIm9mZmxpbmVfYWNjZXNzIiwidW1hX2F1dGhvcml6YXRpb24iLCJhY2Nlc3MtdG8tdXNlci1kYXRhYmFzZSJdfSwicmVzb3VyY2VfYWNjZXNzIjp7ImFjY291bnQiOnsicm9sZXMiOlsibWFuYWdlLWFjY291bnQiLCJtYW5hZ2UtYWNjb3VudC1saW5rcyIsInZpZXctcHJvZmlsZSJdfX0sInNjb3BlIjoicHJvZmlsZSBlbWFpbCIsImNsaWVudEhvc3QiOiIxNzIuMTguMC4zIiwiZW1haWxfdmVyaWZpZWQiOmZhbHNlLCJjbGllbnRJZCI6ImFkbWluLWNvbnNvbGUtY2xpZW50IiwicHJlZmVycmVkX3VzZXJuYW1lIjoic2VydmljZS1hY2NvdW50LWFkbWluLWNvbnNvbGUtY2xpZW50IiwiY2xpZW50QWRkcmVzcyI6IjE3Mi4xOC4wLjMifQ.uXVAwEw7hJjCW_zmkIL-YGJwptA8qN050uvgfVfqkboaJ1xYUpqQhkiztM4WwPXDGsbSR3M_U-FD_8QF5MSWivBzFb-8s9qt9MdXNFqNRBAVjR7fl9vq9PyXNchRYN91Y3eVtqyufEVC5tut79nkWQ65uRgTMsMYKTeN2k1jXpR3FHXohu6_tukMbMqA4gewSFoq1iQ9Qgr7-fRROo_mO31aBZhSvtK6q6iOGLbsauE1UeDgisqtGh06T8lm4IOGxEkzqyY6m-SyZo1EwbLjfNXP1O9IwqQfmMpb2rnjoes_OrpZrBPcX8WsXg7TgGvIdgw9beqZW1S_y6y3ExUqAg",
    "expires_in": 60,
    "refresh_expires_in": 0,
    "token_type": "Bearer",
    "not-before-policy": 0,
    "scope": "profile email"
}
```

The access_token can subsequently be passed to API/database calls. For information on the structure of a JWT, see [here](https://jwt.io/introduction).
If you take the access_token and paste it into [here](https://jwt.io/), you will quickly see the token contains information about who requested and provisioned the token, what permissions to grants, when it expires, etc...
Receiving applications are responsible for validating and mapping permissions to application access (what operations does the role "access-to-user-database" actually enable).

## Routing in Istio

In the simplest implementation, Istio ingress traffic is handled with the following resources: gateway, virtual service, service, pod.
A Gateway resource is used to configure an istio ingressgateway, which is generally istio into the istio-system (or istio-system-common) namespace.
A Virtual service attaches itself to a specific gateway and handles routes for a list of host names.
A service is a standard kubernetes resource that directs traffic to a set of pods based on a label.
A pod is a pod. If `strict` mTLS is enforced on the cluster or namespace, the destination pod will need to be running with an istio sidecar.

```yaml
apiVersion: networking.istio.io/v1alpha3
kind: Gateway
metadata:
  name: istio-gateway
spec:
  selector:
    app: my-istio-ingressgateway
  servers:
  - port:
      number: 80
      name: http
      protocol: HTTP
    hosts:
    - "*.my-domain.com"
---
apiVersion: networking.istio.io/v1alpha3
kind: VirtualService
metadata:
  name: nginx-vs
spec:
  gateways:
  - istio-gateway
  hosts:
  - nginx.my-domain.com
  http:
  - route:
    - destination:
        host: nginx
        port: 80
---
apiVersion: v1
kind: Service
metadata:
  name: nginx-svc
spec:
  selector:
    app.kubernetes.io/name: nginx
  ports:
    - protocol: TCP
      port: 80
      targetPort: 80
---
apiVersion: v1
kind: Pod
metadata:
  name: nginx
  labels:
    app.kubernetes.io/name: nginx
spec:
  containers:
  - name: nginx
    image: nginx:stable
    ports:
      - containerPort: 80
```

## Istio sidecar injection

Istio uses a mutating webhook to watch for pod creation filtered by namespace or pod annotations for OCP.
Essentially, when the Kubernetes API receives a request to create a pod, the Isito mutating webhook modifies the pod defintion
to add a InitContainer and an Envoy proxy container. The istio InitContainer modifies the pod's iptables (containers within a pod share a network namespace, so changes made by any container to iptables affects all other containers). Once the InitContainer compmletes, the primary container (nginx) and the Envoy proxy container are started. Now when traffic is direct to the pod (say nginx:80), the traffic is "intercepted" by the Envoy proxy. The same happens when the primary container sends egress traffic (kubectl exec -it nginx -- curl google.com).
Generally speaking, when Istio is enabled for a pod or namespace, istio injection is completely automated.

Read [more](https://istio.io/latest/blog/2019/data-plane-setup/)