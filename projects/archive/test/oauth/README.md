# oauth probe

Project is to test oauth2.0 authenication against a backend application
Application should go through oauth2.0 flow and send token to backend application
Application should not care where oauth validation actually happens

Pass in oauth2.0/OIDC info as environment variables, go through flow, send request
where bearer token to target_url
