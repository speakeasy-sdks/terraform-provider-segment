// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// CreateSourceRegulationV1InputRegulationType - The regulation type to create.
type CreateSourceRegulationV1InputRegulationType string

const (
	CreateSourceRegulationV1InputRegulationTypeDeleteInternal     CreateSourceRegulationV1InputRegulationType = "DELETE_INTERNAL"
	CreateSourceRegulationV1InputRegulationTypeDeleteOnly         CreateSourceRegulationV1InputRegulationType = "DELETE_ONLY"
	CreateSourceRegulationV1InputRegulationTypeSuppressOnly       CreateSourceRegulationV1InputRegulationType = "SUPPRESS_ONLY"
	CreateSourceRegulationV1InputRegulationTypeSuppressWithDelete CreateSourceRegulationV1InputRegulationType = "SUPPRESS_WITH_DELETE"
	CreateSourceRegulationV1InputRegulationTypeUnsuppress         CreateSourceRegulationV1InputRegulationType = "UNSUPPRESS"
)

func (e CreateSourceRegulationV1InputRegulationType) ToPointer() *CreateSourceRegulationV1InputRegulationType {
	return &e
}

func (e *CreateSourceRegulationV1InputRegulationType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "DELETE_INTERNAL":
		fallthrough
	case "DELETE_ONLY":
		fallthrough
	case "SUPPRESS_ONLY":
		fallthrough
	case "SUPPRESS_WITH_DELETE":
		fallthrough
	case "UNSUPPRESS":
		*e = CreateSourceRegulationV1InputRegulationType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateSourceRegulationV1InputRegulationType: %v", v)
	}
}

// CreateSourceRegulationV1InputSubjectType - The subject type.
type CreateSourceRegulationV1InputSubjectType string

const (
	CreateSourceRegulationV1InputSubjectTypeUserID CreateSourceRegulationV1InputSubjectType = "USER_ID"
)

func (e CreateSourceRegulationV1InputSubjectType) ToPointer() *CreateSourceRegulationV1InputSubjectType {
	return &e
}

func (e *CreateSourceRegulationV1InputSubjectType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "USER_ID":
		*e = CreateSourceRegulationV1InputSubjectType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateSourceRegulationV1InputSubjectType: %v", v)
	}
}

// CreateSourceRegulationV1Input - The input to create a Source-scoped regulation.
type CreateSourceRegulationV1Input struct {
	// The regulation type to create.
	RegulationType CreateSourceRegulationV1InputRegulationType `json:"regulationType"`
	// The list of `userId` or `objectId` values of the subjects to regulate.
	//
	// Config API note: equal to `parent` but allows an array.
	SubjectIds []string `json:"subjectIds"`
	// The subject type.
	SubjectType CreateSourceRegulationV1InputSubjectType `json:"subjectType"`
}
