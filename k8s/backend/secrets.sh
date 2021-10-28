#!/bin/bash

VAULT_CRED_NAME=`echo -n "vault_db" | base64`
VAULT_ROLE_ID=`echo -n "550693ef-d956-a588-6af7-720583c20a5d" | base64`

cat <<EOF
apiVersion: v1
kind: Secret
metadata:
  namespace: queue-system
  name: backend-secret
data:
  VAULT_CRED_NAME: $VAULT_CRED_NAME
  VAULT_ROLE_ID: $VAULT_ROLE_ID
EOF

# ./secrets.sh | kubectl apply -f -