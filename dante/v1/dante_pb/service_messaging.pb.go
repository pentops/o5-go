// Code generated by protoc-gen-go-messaging. DO NOT EDIT.

package dante_pb

// Service: DempeService
// Method: ListMessages

func (msg *ListMessagesRequest) MessagingHeaders() map[string]string {
	headers := map[string]string{
		"grpc-service": "/o5.dante.v1.service.DempeService/ListMessages",
		"grpc-message": "o5.dante.v1.service.ListMessagesRequest",
	}
	return headers
}

// Method: GetMessage

func (msg *GetMessageRequest) MessagingHeaders() map[string]string {
	headers := map[string]string{
		"grpc-service": "/o5.dante.v1.service.DempeService/GetMessage",
		"grpc-message": "o5.dante.v1.service.GetMessageRequest",
	}
	return headers
}

// Method: MessagesAction

func (msg *MessagesActionRequest) MessagingHeaders() map[string]string {
	headers := map[string]string{
		"grpc-service": "/o5.dante.v1.service.DempeService/MessagesAction",
		"grpc-message": "o5.dante.v1.service.MessagesActionRequest",
	}
	return headers
}
