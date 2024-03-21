// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/scentregroup/terraform-provider-segment/internal/provider/types"
	"github.com/scentregroup/terraform-provider-segment/internal/sdk/models/operations"
	"github.com/scentregroup/terraform-provider-segment/internal/sdk/models/shared"
)

func (r *CreateWorkspaceRegulationV1InputResourceModel) ToSharedCreateWorkspaceRegulationV1Input() *shared.CreateWorkspaceRegulationV1Input {
	regulationType := shared.CreateWorkspaceRegulationV1InputRegulationType(r.RegulationType.ValueString())
	var subjectIds []string = nil
	for _, subjectIdsItem := range r.SubjectIds {
		subjectIds = append(subjectIds, subjectIdsItem.ValueString())
	}
	subjectType := shared.CreateWorkspaceRegulationV1InputSubjectType(r.SubjectType.ValueString())
	out := shared.CreateWorkspaceRegulationV1Input{
		RegulationType: regulationType,
		SubjectIds:     subjectIds,
		SubjectType:    subjectType,
	}
	return &out
}

func (r *CreateWorkspaceRegulationV1InputResourceModel) RefreshFromOperationsCreateWorkspaceRegulationResponseBody(resp *operations.CreateWorkspaceRegulationResponseBody) {
	if resp != nil {
		if resp.Data == nil {
			r.Data = nil
		} else {
			r.Data = &tfTypes.CreateCloudSourceRegulationV1Output{}
			r.Data.RegulateID = types.StringValue(resp.Data.RegulateID)
		}
	}
}
