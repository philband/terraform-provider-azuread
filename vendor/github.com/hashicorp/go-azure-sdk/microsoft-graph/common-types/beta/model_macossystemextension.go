package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSSystemExtension struct {
	// Gets or sets the bundle identifier of the system extension.
	BundleId *string `json:"bundleId,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Gets or sets the team identifier that was used to sign the system extension.
	TeamIdentifier nullable.Type[string] `json:"teamIdentifier,omitempty"`
}