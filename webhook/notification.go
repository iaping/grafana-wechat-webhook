package webhook

import "fmt"

type Notification struct {
	Receiver        string   `json:"receiver"`
	Status          string   `json:"status"`
	OrgId           int      `json:"orgId"`
	Alerts          []*Alert `json:"alerts"`
	ExternalURL     string   `json:"externalURL"`
	Version         string   `json:"version"`
	GroupKey        string   `json:"groupKey"`
	TruncatedAlerts int      `json:"truncatedAlerts"`
	Title           string   `json:"title"`
	State           string   `json:"state"`
	Message         string   `json:"message"`
}

func (n *Notification) BuildMarkdownTemplate() string {
	template := fmt.Sprintf("# %s\n", n.Status)

	for _, alert := range n.Alerts {
		template += fmt.Sprintf("> %s\n", alert.Labels.Alertname)
		template += fmt.Sprintf("> **desc:**<font color=\"comment\">%s</font>\n", alert.Annotations.Description)
		template += fmt.Sprintf("> **summary:**<font color=\"comment\">%s</font>\n", alert.Annotations.Summary)
		template += fmt.Sprintf("> **time:**<font color=\"comment\">%s</font>\n", alert.StartsAt)
		template += fmt.Sprintf("<font color=\"comment\">%s</font>", alert.ValueString)
		template += "\n--------------------------\n"
	}

	template += fmt.Sprintf("[查看详情](%s)", n.ExternalURL)

	return template
}

type Alert struct {
	Status string `json:"status"`
	Labels struct {
		Alertname string `json:"alertname"`
		Team      string `json:"team"`
		Zone      string `json:"zone"`
	} `json:"labels"`
	Annotations struct {
		Description string `json:"description"`
		RunbookUrl  string `json:"runbook_url"`
		Summary     string `json:"summary"`
	} `json:"annotations"`
	StartsAt     string `json:"startsAt"`
	EndsAt       string `json:"endsAt"`
	GeneratorURL string `json:"generatorURL"`
	Fingerprint  string `json:"fingerprint"`
	SilenceURL   string `json:"silenceURL"`
	DashboardURL string `json:"dashboardURL"`
	PanelURL     string `json:"panelURL"`
	ValueString  string `json:"valueString"`
}
