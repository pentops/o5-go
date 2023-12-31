// Code generated by protoc-gen-go-messaging. DO NOT EDIT.

package dante_spb

// Service: DeadMessageQueryService
// Method: GetDeadMessage

func (msg *GetDeadMessageRequest) MessagingHeaders() map[string]string {
	headers := map[string]string{
		"grpc-service": "/o5.dante.v1.service.DeadMessageQueryService/GetDeadMessage",
		"grpc-message": "o5.dante.v1.service.GetDeadMessageRequest",
	}
	return headers
}

// Method: ListDeadMessages

func (msg *ListDeadMessagesRequest) MessagingHeaders() map[string]string {
	headers := map[string]string{
		"grpc-service": "/o5.dante.v1.service.DeadMessageQueryService/ListDeadMessages",
		"grpc-message": "o5.dante.v1.service.ListDeadMessagesRequest",
	}
	return headers
}

// Method: ListDeadMessageEvents

func (msg *ListDeadMessageEventsRequest) MessagingHeaders() map[string]string {
	headers := map[string]string{
		"grpc-service": "/o5.dante.v1.service.DeadMessageQueryService/ListDeadMessageEvents",
		"grpc-message": "o5.dante.v1.service.ListDeadMessageEventsRequest",
	}
	return headers
}
