{{- define "mock-worker.labels" -}}
{{ include "mock-worker.selectorLabels" . }}
{{- end -}}

{{- define "mock-worker.selectorLabels" -}}
app: mock-worker
gardener.cloud/role: worker
role: worker
gardener.cloud/pool: {{ .Values.worker.pool }}
{{- end -}}
