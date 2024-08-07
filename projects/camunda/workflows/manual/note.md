export ZEEBE_ADDRESS="192.168.86.30:26500" && \
export ZEEBE_HOST="192.168.86.30" && \
export ZEEBE_PORT="26500" && \
export ZEEBE_CLIENT_ID="zeebe-worker" && \
export ZEEBE_CLIENT_SECRET="mVSsgWqmN7YcWoCkGQohFI7nAqPnvPPj" && \
export ZEEBE_AUTHORIZATION_SERVER_URL="http://192.168.86.30:18080/auth/realms/camunda-platform/protocol/openid-connect/token"


curl -X POST $ZEEBE_AUTHORIZATION_SERVER_URL -u $ZEEBE_CLIENT_ID:$ZEEBE_CLIENT_SECRET -d "grant_type=client_credentials"