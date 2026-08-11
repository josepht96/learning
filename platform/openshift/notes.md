oc create serviceaccount logging-admin -n openshift-logging
oc adm policy add-cluster-role-to-user collect-application-logs -z logging-admin -n openshift-logging
oc adm policy add-cluster-role-to-user collect-infrastructure-logs -z logging-admin -n openshift-logging
oc adm policy add-cluster-role-to-user collect-audit-logs -z logging-admin -n openshift-logging


oc delete clusterlogging instance -n openshift-logging
oc delete elasticsearch elasticsearch -n openshift-logging
oc delete subscription elasticsearch-operator -n openshift-operators-redhat
oc delete csv elasticsearch-operator.v5.8.21 -n openshift-operators-redhat