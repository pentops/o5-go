// Code generated by protoc-gen-go-messaging. DO NOT EDIT.

package dante_pb

// Service: DanteQueryService
// Method: ListMessages

func (msg *ListMessagesRequest) MessagingHeaders() map[string]string {
	headers := map[string]string{
		"grpc-service": "/o5.dante.v1.service.DanteQueryService/ListMessages",
		"grpc-message": "o5.dante.v1.service.ListMessagesRequest",
	}
	return headers
}

// Method: GetMessage

func (msg *GetMessageRequest) MessagingHeaders() map[string]string {
	headers := map[string]string{
		"grpc-service": "/o5.dante.v1.service.DanteQueryService/GetMessage",
		"grpc-message": "o5.dante.v1.service.GetMessageRequest",
	}
	return headers
}

// Method: MessagesAction

func (msg *MessagesActionRequest) MessagingHeaders() map[string]string {
	headers := map[string]string{
		"grpc-service": "/o5.dante.v1.service.DanteQueryService/MessagesAction",
		"grpc-message": "o5.dante.v1.service.MessagesActionRequest",
	}
	return headers
}
