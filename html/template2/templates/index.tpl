{{ define "index" }}
{{- template "header" . -}}

Hello world! My name is {{ .Name | Upper }}

{{- template "footer" -}}
{{ end }}
