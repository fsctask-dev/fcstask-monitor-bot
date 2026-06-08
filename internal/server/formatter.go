package server

import (
	"fmt"
	"strings"
)

func FormatAlertText(payload AlertmanagerPayload) string {
	var sb strings.Builder
	for _, alert := range payload.Alerts {
		statusIcon := "🔴"
		if alert.Status == "resolved" {
			statusIcon = "🟢"
		}
		alertName := alert.Labels["alertname"]
		summary := alert.Annotations["summary"]
		sb.WriteString(fmt.Sprintf("🚨 ALERT\n\n%s %s: %s", statusIcon, alertName, summary))
	}
	return sb.String()
}
