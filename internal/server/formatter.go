package server

import "fmt"

func FormatAlertText(item AlertItem) string {
	statusIcon := "🔴"
	if item.Status == "resolved" {
		statusIcon = "🟢"
	}
	alertName := item.Labels["alertname"]
	summary := item.Annotations["summary"]
	return fmt.Sprintf("🚨 ALERT\n\n%s %s: %s", statusIcon, alertName, summary)
}
