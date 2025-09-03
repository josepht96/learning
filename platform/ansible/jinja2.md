# jinja2

{% for name in names -%}
{{name}}
{% endfor %}

{% for host in hosts -%}
{{ host.name }} {{ host.ip_address }}
{% endfor %}

```yaml
hosts:
task:
- name: test
  template:
    src: index.html.j2
    dest: /var/www/index.html
```