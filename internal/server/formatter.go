package server

import (
	"fmt"
	"html"
	"sort"
	"strings"
)

// Лейблы которые показываются отдельно и не дублируются в блоке контекста
var skipLabels = map[string]bool{
	"alertname": true,
	"severity":  true,
}

// Приоритет отображения лейблов в блоке контекста
var labelOrder = []string{"instance", "job", "route", "role", "replica_index"}

func FormatAlertText(payload AlertmanagerPayload) string {
	var sb strings.Builder
	for i, alert := range payload.Alerts {
		if i > 0 {
			sb.WriteString("\n\n")
		}
		sb.WriteString(formatAlert(alert))
	}
	return sb.String()
}

func formatAlert(alert AlertItem) string {
	var sb strings.Builder

	if alert.Status == "resolved" {
		sb.WriteString("✅ <b>RESOLVED</b>\n\n")
	} else {
		sb.WriteString("🚨 <b>ALERT — Firing</b>\n\n")
	}

	if name := alert.Labels["alertname"]; name != "" {
		sb.WriteString(fmt.Sprintf("<b>%s</b>\n", html.EscapeString(name)))
	}

	if severity := alert.Labels["severity"]; severity != "" {
		sb.WriteString(fmt.Sprintf("%s %s\n", severityIcon(severity), html.EscapeString(severity)))
	}

	// Сначала выводим лейблы в приоритетном порядке, потом остальные
	shown := make(map[string]bool)
	for _, key := range labelOrder {
		if val, ok := alert.Labels[key]; ok {
			sb.WriteString(fmt.Sprintf("\n%s: <code>%s</code>", key, html.EscapeString(val)))
			shown[key] = true
		}
	}
	// Оставшиеся лейблы в алфавитном порядке
	rest := make([]string, 0)
	for key := range alert.Labels {
		if !skipLabels[key] && !shown[key] {
			rest = append(rest, key)
		}
	}
	sort.Strings(rest)
	for _, key := range rest {
		sb.WriteString(fmt.Sprintf("\n<code>%s</code>: %s", html.EscapeString(key), html.EscapeString(alert.Labels[key])))
	}

	if summary := alert.Annotations["summary"]; summary != "" {
		sb.WriteString(fmt.Sprintf("\n\n📝 %s", html.EscapeString(summary)))
	}

	if desc := alert.Annotations["description"]; desc != "" && desc != alert.Annotations["summary"] {
		sb.WriteString(fmt.Sprintf("\nℹ️ %s", html.EscapeString(desc)))
	}

	sb.WriteString(fmt.Sprintf("\n\n⏰ %s", alert.StartsAt.UTC().Format("02.01.2006 15:04 UTC")))

	if alert.GeneratorURL != "" {
		sb.WriteString(fmt.Sprintf("\n🔗 <code>%s</code>", html.EscapeString(alert.GeneratorURL)))
	}

	return sb.String()
}

func severityIcon(severity string) string {
	switch strings.ToLower(severity) {
	case "critical":
		return "🔴"
	case "warning":
		return "🟡"
	case "info":
		return "🔵"
	default:
		return "⚪"
	}
}
