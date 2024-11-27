package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdEventIdInstanceIdExceptionOccurrenceIdAttachmentId{}

// UserIdEventIdInstanceIdExceptionOccurrenceIdAttachmentId is a struct representing the Resource ID for a User Id Event Id Instance Id Exception Occurrence Id Attachment
type UserIdEventIdInstanceIdExceptionOccurrenceIdAttachmentId struct {
	UserId       string
	EventId      string
	EventId1     string
	EventId2     string
	AttachmentId string
}

// NewUserIdEventIdInstanceIdExceptionOccurrenceIdAttachmentID returns a new UserIdEventIdInstanceIdExceptionOccurrenceIdAttachmentId struct
func NewUserIdEventIdInstanceIdExceptionOccurrenceIdAttachmentID(userId string, eventId string, eventId1 string, eventId2 string, attachmentId string) UserIdEventIdInstanceIdExceptionOccurrenceIdAttachmentId {
	return UserIdEventIdInstanceIdExceptionOccurrenceIdAttachmentId{
		UserId:       userId,
		EventId:      eventId,
		EventId1:     eventId1,
		EventId2:     eventId2,
		AttachmentId: attachmentId,
	}
}

// ParseUserIdEventIdInstanceIdExceptionOccurrenceIdAttachmentID parses 'input' into a UserIdEventIdInstanceIdExceptionOccurrenceIdAttachmentId
func ParseUserIdEventIdInstanceIdExceptionOccurrenceIdAttachmentID(input string) (*UserIdEventIdInstanceIdExceptionOccurrenceIdAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdEventIdInstanceIdExceptionOccurrenceIdAttachmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdEventIdInstanceIdExceptionOccurrenceIdAttachmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdEventIdInstanceIdExceptionOccurrenceIdAttachmentIDInsensitively parses 'input' case-insensitively into a UserIdEventIdInstanceIdExceptionOccurrenceIdAttachmentId
// note: this method should only be used for API response data and not user input
func ParseUserIdEventIdInstanceIdExceptionOccurrenceIdAttachmentIDInsensitively(input string) (*UserIdEventIdInstanceIdExceptionOccurrenceIdAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdEventIdInstanceIdExceptionOccurrenceIdAttachmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdEventIdInstanceIdExceptionOccurrenceIdAttachmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdEventIdInstanceIdExceptionOccurrenceIdAttachmentId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.EventId, ok = input.Parsed["eventId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "eventId", input)
	}

	if id.EventId1, ok = input.Parsed["eventId1"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "eventId1", input)
	}

	if id.EventId2, ok = input.Parsed["eventId2"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "eventId2", input)
	}

	if id.AttachmentId, ok = input.Parsed["attachmentId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "attachmentId", input)
	}

	return nil
}

// ValidateUserIdEventIdInstanceIdExceptionOccurrenceIdAttachmentID checks that 'input' can be parsed as a User Id Event Id Instance Id Exception Occurrence Id Attachment ID
func ValidateUserIdEventIdInstanceIdExceptionOccurrenceIdAttachmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdEventIdInstanceIdExceptionOccurrenceIdAttachmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Event Id Instance Id Exception Occurrence Id Attachment ID
func (id UserIdEventIdInstanceIdExceptionOccurrenceIdAttachmentId) ID() string {
	fmtString := "/users/%s/events/%s/instances/%s/exceptionOccurrences/%s/attachments/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.EventId, id.EventId1, id.EventId2, id.AttachmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Event Id Instance Id Exception Occurrence Id Attachment ID
func (id UserIdEventIdInstanceIdExceptionOccurrenceIdAttachmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("events", "events", "events"),
		resourceids.UserSpecifiedSegment("eventId", "eventId"),
		resourceids.StaticSegment("instances", "instances", "instances"),
		resourceids.UserSpecifiedSegment("eventId1", "eventId1"),
		resourceids.StaticSegment("exceptionOccurrences", "exceptionOccurrences", "exceptionOccurrences"),
		resourceids.UserSpecifiedSegment("eventId2", "eventId2"),
		resourceids.StaticSegment("attachments", "attachments", "attachments"),
		resourceids.UserSpecifiedSegment("attachmentId", "attachmentId"),
	}
}

// String returns a human-readable description of this User Id Event Id Instance Id Exception Occurrence Id Attachment ID
func (id UserIdEventIdInstanceIdExceptionOccurrenceIdAttachmentId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Event Id 1: %q", id.EventId1),
		fmt.Sprintf("Event Id 2: %q", id.EventId2),
		fmt.Sprintf("Attachment: %q", id.AttachmentId),
	}
	return fmt.Sprintf("User Id Event Id Instance Id Exception Occurrence Id Attachment (%s)", strings.Join(components, "\n"))
}