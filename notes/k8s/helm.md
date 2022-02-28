# Helm Charts
Package manager for kubernetes
Single file to declare all custom settings - values.yml

## cmds 
helm install <release name> <helm chart name>
helm upgrade <release name>
helm rollback <release name>
helm uninstall <release name>

## values.yml
Template deployment files by using values.yml file


# Releases
When a helm chart is installed on a cluster, a release is created


# Structure
A helm directory expects the following structure:
```
wordpress/
  Chart.yaml          # A YAML file containing information about the chart
  LICENSE             # OPTIONAL: A plain text file containing the license for the chart
  README.md           # OPTIONAL: A human-readable README file
  values.yaml         # The default configuration values for this chart
  values.schema.json  # OPTIONAL: A JSON Schema for imposing a structure on the values.yaml file
  charts/             # A directory containing any charts upon which this chart depends.
  crds/               # Custom Resource Definitions
  templates/          # A directory of templates that, when combined with values,
                      # will generate valid Kubernetes manifest files.
  templates/NOTES.txt # OPTIONAL: A plain text file containing short usage notes
```