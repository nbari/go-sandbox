{{ define "indexPage" }}
<!DOCTYPE html>
<html>
{{- template "header" }}
<body>
Hello world! My name is {{ .Name }}
</body>
{{- template "footer" }}
{{ end }}
