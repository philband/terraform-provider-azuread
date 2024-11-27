package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ItemFacet = ItemEmail{}

type ItemEmail struct {
	// The email address itself.
	Address *string `json:"address,omitempty"`

	// The name or label a user has associated with a particular email address.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	Type *EmailType `json:"type,omitempty"`

	// Fields inherited from ItemFacet

	// The audiences that are able to see the values contained within the associated entity. Possible values are: me,
	// family, contacts, groupMembers, organization, federatedOrganizations, everyone, unknownFutureValue.
	AllowedAudiences *AllowedAudiences `json:"allowedAudiences,omitempty"`

	CreatedBy IdentitySet `json:"createdBy"`

	// Provides the dateTimeOffset for when the entity was created.
	CreatedDateTime *string `json:"createdDateTime,omitempty"`

	// Contains inference detail if the entity is inferred by the creating or modifying application.
	Inference *InferenceData `json:"inference,omitempty"`

	IsSearchable   nullable.Type[bool] `json:"isSearchable,omitempty"`
	LastModifiedBy IdentitySet         `json:"lastModifiedBy"`

	// Provides the dateTimeOffset for when the entity was created.
	LastModifiedDateTime *string `json:"lastModifiedDateTime,omitempty"`

	// Where the values within an entity originated if synced from another service.
	Source *PersonDataSources `json:"source,omitempty"`

	// Where the values within an entity originated if synced from another source.
	Sources *[]ProfileSourceAnnotation `json:"sources,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s ItemEmail) ItemFacet() BaseItemFacetImpl {
	return BaseItemFacetImpl{
		AllowedAudiences:     s.AllowedAudiences,
		CreatedBy:            s.CreatedBy,
		CreatedDateTime:      s.CreatedDateTime,
		Inference:            s.Inference,
		IsSearchable:         s.IsSearchable,
		LastModifiedBy:       s.LastModifiedBy,
		LastModifiedDateTime: s.LastModifiedDateTime,
		Source:               s.Source,
		Sources:              s.Sources,
		Id:                   s.Id,
		ODataId:              s.ODataId,
		ODataType:            s.ODataType,
	}
}

func (s ItemEmail) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ItemEmail{}

func (s ItemEmail) MarshalJSON() ([]byte, error) {
	type wrapper ItemEmail
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ItemEmail: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ItemEmail: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.itemEmail"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ItemEmail: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &ItemEmail{}

func (s *ItemEmail) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Address              *string                    `json:"address,omitempty"`
		DisplayName          nullable.Type[string]      `json:"displayName,omitempty"`
		Type                 *EmailType                 `json:"type,omitempty"`
		AllowedAudiences     *AllowedAudiences          `json:"allowedAudiences,omitempty"`
		CreatedDateTime      *string                    `json:"createdDateTime,omitempty"`
		Inference            *InferenceData             `json:"inference,omitempty"`
		IsSearchable         nullable.Type[bool]        `json:"isSearchable,omitempty"`
		LastModifiedDateTime *string                    `json:"lastModifiedDateTime,omitempty"`
		Source               *PersonDataSources         `json:"source,omitempty"`
		Sources              *[]ProfileSourceAnnotation `json:"sources,omitempty"`
		Id                   *string                    `json:"id,omitempty"`
		ODataId              *string                    `json:"@odata.id,omitempty"`
		ODataType            *string                    `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Address = decoded.Address
	s.DisplayName = decoded.DisplayName
	s.Type = decoded.Type
	s.AllowedAudiences = decoded.AllowedAudiences
	s.CreatedDateTime = decoded.CreatedDateTime
	s.Id = decoded.Id
	s.Inference = decoded.Inference
	s.IsSearchable = decoded.IsSearchable
	s.LastModifiedDateTime = decoded.LastModifiedDateTime
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.Source = decoded.Source
	s.Sources = decoded.Sources

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling ItemEmail into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["createdBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'CreatedBy' for 'ItemEmail': %+v", err)
		}
		s.CreatedBy = impl
	}

	if v, ok := temp["lastModifiedBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'LastModifiedBy' for 'ItemEmail': %+v", err)
		}
		s.LastModifiedBy = impl
	}

	return nil
}