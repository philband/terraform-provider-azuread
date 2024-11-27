package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsAutopilotDeploymentProfile interface {
	Entity
	WindowsAutopilotDeploymentProfile() BaseWindowsAutopilotDeploymentProfileImpl
}

var _ WindowsAutopilotDeploymentProfile = BaseWindowsAutopilotDeploymentProfileImpl{}

type BaseWindowsAutopilotDeploymentProfileImpl struct {
	// The list of assigned devices for the profile.
	AssignedDevices *[]WindowsAutopilotDeviceIdentity `json:"assignedDevices,omitempty"`

	// The list of group assignments for the profile.
	Assignments *[]WindowsAutopilotDeploymentProfileAssignment `json:"assignments,omitempty"`

	// The date and time of when the deployment profile was created. The value cannot be modified and is automatically
	// populated when the profile was created. The timestamp type represents date and time information using ISO 8601 format
	// and is always in UTC time. For example, midnight UTC on Jan 1, 2014 would look like this: '2014-01-01T00:00:00Z'.
	// Supports: $select, $top, $skip. $Search, $orderBy and $filter are not supported. Read-Only.
	CreatedDateTime *string `json:"createdDateTime,omitempty"`

	// A description of the deployment profile. Max allowed length is 1500 chars. Supports: $select, $top, $skip, $orderBy.
	// $Search and $filter are not supported.
	Description nullable.Type[string] `json:"description,omitempty"`

	// The template used to name the Autopilot device. This can be a custom text and can also contain either the serial
	// number of the device, or a randomly generated number. The total length of the text generated by the template can be
	// no more than 15 characters. Supports: $select, $top, $skip. $Search, $orderBy and $filter are not supported.
	DeviceNameTemplate nullable.Type[string] `json:"deviceNameTemplate,omitempty"`

	DeviceType *WindowsAutopilotDeviceType `json:"deviceType,omitempty"`

	// The display name of the deployment profile. Max allowed length is 200 chars. Returned by default. Supports: $select,
	// $top, $skip, $orderby. $Search and $filter are not supported.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// Indicates whether the user is allowed to use Windows Autopilot for pre-provisioned deployment mode during Out of Box
	// experience (OOBE). When TRUE, indicates that Windows Autopilot for pre-provisioned deployment mode is allowed. When
	// false, Windows Autopilot for pre-provisioned deployment mode is not allowed. The default is FALSE. Read-Only.
	// Starting from May 2024 this property will no longer be supported and will be marked as deprecated. Use
	// preprovisioningAllowed instead.
	EnableWhiteGlove *bool `json:"enableWhiteGlove,omitempty"`

	// Enrollment status screen setting
	EnrollmentStatusScreenSettings *WindowsEnrollmentStatusScreenSettings `json:"enrollmentStatusScreenSettings,omitempty"`

	// Indicates whether the profile supports the extraction of hardware hash values and registration of the device into
	// Windows Autopilot. When TRUE, indicates if hardware extraction and Windows Autopilot registration will happen on the
	// next successful check-in. When FALSE, hardware hash extraction and Windows Autopilot registration will not happen.
	// Default value is FALSE. Supports: $select, $top, $skip. $Search, $orderBy and $filter are not supported. Read-Only.
	// Starting from May 2024 this property will no longer be supported and will be marked as deprecated. Use
	// hardwareHashExtractionEnabled instead.
	ExtractHardwareHash *bool `json:"extractHardwareHash,omitempty"`

	// Indicates whether the profile supports the extraction of hardware hash values and registration of the device into
	// Windows Autopilot. When TRUE, indicates if hardware extraction and Windows Autopilot registration will happen on the
	// next successful check-in. When FALSE, hardware hash extraction and Windows Autopilot registration will not happen.
	// Default value is FALSE. Supports: $select, $top, $skip. $Search, $orderBy and $filter are not supported.
	HardwareHashExtractionEnabled nullable.Type[bool] `json:"hardwareHashExtractionEnabled,omitempty"`

	// The language code to be used when configuring the device. E.g. en-US. The default value is os-default. Supports:
	// $select, $top, $skip. $Search, $orderBy and $filter are not supported. Read-Only. Starting from May 2024 this
	// property will no longer be supported and will be marked as deprecated. Use locale instead.
	Language nullable.Type[string] `json:"language,omitempty"`

	// The date and time of when the deployment profile was last modified. The value cannot be updated manually and is
	// automatically populated when any changes are made to the profile. The Timestamp type represents date and time
	// information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 would look like
	// this: '2014-01-01T00:00:00Z'. Supports: $select, $top, $skip. $Search, $orderBy and $filter are not supported
	// Read-Only.
	LastModifiedDateTime *string `json:"lastModifiedDateTime,omitempty"`

	// The locale (language) to be used when configuring the device. E.g. en-US. The default value is os-default. Supports:
	// $select, $top, $skip. $Search, $orderBy and $filter are not supported.
	Locale nullable.Type[string] `json:"locale,omitempty"`

	// The Entra management service App ID which gets used during client device-based enrollment discovery. Supports:
	// $select, $top, $skip. $Search, $orderBy and $filter are not supported.
	ManagementServiceAppId nullable.Type[string] `json:"managementServiceAppId,omitempty"`

	// The Windows Autopilot Deployment Profile settings used by the device for the out-of-box experience. Supports:
	// $select, $top, $skip. $Search, $orderBy and $filter are not supported.
	OutOfBoxExperienceSetting *OutOfBoxExperienceSetting `json:"outOfBoxExperienceSetting,omitempty"`

	// The Windows Autopilot Deployment Profile settings used by the Autopilot device for out-of-box experience. Supports:
	// $select, $top, $skip. $Search, $orderBy and $filter are not supported. Read-Only. Starting from May 2024 this
	// property will no longer be supported and will be marked as deprecated. Use outOfBoxExperienceSetting instead.
	OutOfBoxExperienceSettings *OutOfBoxExperienceSettings `json:"outOfBoxExperienceSettings,omitempty"`

	// Indicates whether the user is allowed to use Windows Autopilot for pre-provisioned deployment mode during Out of Box
	// experience (OOBE). When TRUE, indicates that Windows Autopilot for pre-provisioned deployment mode for OOBE is
	// allowed to be used. When false, Windows Autopilot for pre-provisioned deployment mode for OOBE is not allowed. The
	// default is FALSE.
	PreprovisioningAllowed nullable.Type[bool] `json:"preprovisioningAllowed,omitempty"`

	// List of role scope tags for the deployment profile.
	RoleScopeTagIds *[]string `json:"roleScopeTagIds,omitempty"`

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

func (s BaseWindowsAutopilotDeploymentProfileImpl) WindowsAutopilotDeploymentProfile() BaseWindowsAutopilotDeploymentProfileImpl {
	return s
}

func (s BaseWindowsAutopilotDeploymentProfileImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ WindowsAutopilotDeploymentProfile = RawWindowsAutopilotDeploymentProfileImpl{}

// RawWindowsAutopilotDeploymentProfileImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawWindowsAutopilotDeploymentProfileImpl struct {
	windowsAutopilotDeploymentProfile BaseWindowsAutopilotDeploymentProfileImpl
	Type                              string
	Values                            map[string]interface{}
}

func (s RawWindowsAutopilotDeploymentProfileImpl) WindowsAutopilotDeploymentProfile() BaseWindowsAutopilotDeploymentProfileImpl {
	return s.windowsAutopilotDeploymentProfile
}

func (s RawWindowsAutopilotDeploymentProfileImpl) Entity() BaseEntityImpl {
	return s.windowsAutopilotDeploymentProfile.Entity()
}

var _ json.Marshaler = BaseWindowsAutopilotDeploymentProfileImpl{}

func (s BaseWindowsAutopilotDeploymentProfileImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseWindowsAutopilotDeploymentProfileImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseWindowsAutopilotDeploymentProfileImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseWindowsAutopilotDeploymentProfileImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.windowsAutopilotDeploymentProfile"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseWindowsAutopilotDeploymentProfileImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalWindowsAutopilotDeploymentProfileImplementation(input []byte) (WindowsAutopilotDeploymentProfile, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling WindowsAutopilotDeploymentProfile into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.activeDirectoryWindowsAutopilotDeploymentProfile") {
		var out ActiveDirectoryWindowsAutopilotDeploymentProfile
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ActiveDirectoryWindowsAutopilotDeploymentProfile: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.azureADWindowsAutopilotDeploymentProfile") {
		var out AzureADWindowsAutopilotDeploymentProfile
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AzureADWindowsAutopilotDeploymentProfile: %+v", err)
		}
		return out, nil
	}

	var parent BaseWindowsAutopilotDeploymentProfileImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseWindowsAutopilotDeploymentProfileImpl: %+v", err)
	}

	return RawWindowsAutopilotDeploymentProfileImpl{
		windowsAutopilotDeploymentProfile: parent,
		Type:                              value,
		Values:                            temp,
	}, nil

}