package server

import "fmt"

func FormatAlertText(alert Alert) string {
	statusIcon := "🔴"
	if alert.Status == "resolved" {
		statusIcon = "🟢"
	}
	return fmt.Sprintf("🚨 ALERT\n\n%s %s: %s", statusIcon, alert.AlertName, alert.Summary)
}
