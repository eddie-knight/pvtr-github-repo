package vuln_management

import (
	"github.com/revanite-io/sci/pkg/layer4"

	"github.com/revanite-io/pvtr-github-repo/evaluation_plans/reusable_steps"
)

func hasSecContact(payloadData interface{}, _ map[string]*layer4.Change) (result layer4.Result, message string) {
	data, message := reusable_steps.VerifyPayload(payloadData)
	if message != "" {
		return layer4.Unknown, message
	}

	if data.Insights.Project.Vulnerability.Contact.Email != "" {
		return layer4.Passed, "Security contacts were specified in Security Insights data"
	}
	for _, champion := range data.Insights.Repository.Security.Champions {
		if champion.Email != "" {
			return layer4.Passed, "Security contacts were specified in Security Insights data"
		}
	}

	return layer4.Failed, "Security contacts were NOT specified in Security Insights data"
}
