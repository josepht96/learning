oc create serviceaccount logging-admin -n openshift-logging
oc adm policy add-cluster-role-to-user collect-application-logs -z logging-admin -n openshift-logging
oc adm policy add-cluster-role-to-user collect-infrastructure-logs -z logging-admin -n openshift-logging
oc adm policy add-cluster-role-to-user collect-audit-logs -z logging-admin -n openshift-logging


oc delete clusterlogging instance -n openshift-logging
oc delete elasticsearch elasticsearch -n openshift-logging
oc delete subscription elasticsearch-operator -n openshift-operators-redhat
oc delete csv elasticsearch-operator.v5.8.21 -n openshift-operators-redhat


oc get secret <cluster-name>-es-http-certs-public -n <es-namespace> -o jsonpath='{.data.ca\.crt}' | base64 -d > ca.crt
oc create secret generic eck-es-ca -n openshift-logging --from-file=ca-bundle.crt=ca.crt

oc get secret <cluster-name>-es-elastic-user -n <es-namespace> -o jsonpath='{.data.elastic}' | base64 -d