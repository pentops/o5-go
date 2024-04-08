// Code generated by protoc-gen-go-psm. DO NOT EDIT.

package dante_spb

import (
	psm "github.com/pentops/protostate/psm"
)

// State Query Service for %sMessage
// QuerySet is the query set for the Message service.

type MessagePSMQuerySet = psm.StateQuerySet[
	*GetDeadMessageRequest,
	*GetDeadMessageResponse,
	*ListDeadMessagesRequest,
	*ListDeadMessagesResponse,
	*ListDeadMessageEventsRequest,
	*ListDeadMessageEventsResponse,
]

func NewMessagePSMQuerySet(
	smSpec psm.QuerySpec[
		*GetDeadMessageRequest,
		*GetDeadMessageResponse,
		*ListDeadMessagesRequest,
		*ListDeadMessagesResponse,
		*ListDeadMessageEventsRequest,
		*ListDeadMessageEventsResponse,
	],
	options psm.StateQueryOptions,
) (*MessagePSMQuerySet, error) {
	return psm.BuildStateQuerySet[
		*GetDeadMessageRequest,
		*GetDeadMessageResponse,
		*ListDeadMessagesRequest,
		*ListDeadMessagesResponse,
		*ListDeadMessageEventsRequest,
		*ListDeadMessageEventsResponse,
	](smSpec, options)
}

type MessagePSMQuerySpec = psm.QuerySpec[
	*GetDeadMessageRequest,
	*GetDeadMessageResponse,
	*ListDeadMessagesRequest,
	*ListDeadMessagesResponse,
	*ListDeadMessageEventsRequest,
	*ListDeadMessageEventsResponse,
]

func DefaultMessagePSMQuerySpec(tableSpec psm.QueryTableSpec) MessagePSMQuerySpec {
	return psm.QuerySpec[
		*GetDeadMessageRequest,
		*GetDeadMessageResponse,
		*ListDeadMessagesRequest,
		*ListDeadMessagesResponse,
		*ListDeadMessageEventsRequest,
		*ListDeadMessageEventsResponse,
	]{
		QueryTableSpec: tableSpec,
		ListRequestFilter: func(req *ListDeadMessagesRequest) (map[string]interface{}, error) {
			filter := map[string]interface{}{}
			return filter, nil
		},
		ListEventsRequestFilter: func(req *ListDeadMessageEventsRequest) (map[string]interface{}, error) {
			filter := map[string]interface{}{}
			filter["message_id"] = req.MessageId
			return filter, nil
		},
	}
}