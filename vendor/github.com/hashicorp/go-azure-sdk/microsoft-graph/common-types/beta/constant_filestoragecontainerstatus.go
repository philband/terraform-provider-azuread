package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FileStorageContainerStatus string

const (
	FileStorageContainerStatus_Active   FileStorageContainerStatus = "active"
	FileStorageContainerStatus_Inactive FileStorageContainerStatus = "inactive"
)

func PossibleValuesForFileStorageContainerStatus() []string {
	return []string{
		string(FileStorageContainerStatus_Active),
		string(FileStorageContainerStatus_Inactive),
	}
}

func (s *FileStorageContainerStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseFileStorageContainerStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseFileStorageContainerStatus(input string) (*FileStorageContainerStatus, error) {
	vals := map[string]FileStorageContainerStatus{
		"active":   FileStorageContainerStatus_Active,
		"inactive": FileStorageContainerStatus_Inactive,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := FileStorageContainerStatus(input)
	return &out, nil
}