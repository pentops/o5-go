// Code generated by protoc-gen-go-messaging. DO NOT EDIT.

package dempe_pb

// Service: DempeService
// Method: ListMessages

func (msg *ListMessagesRequest) MessagingHeaders() map[string]string {
	headers := map[string]string{
		"grpc-service": "/o5.dempe.v1.DempeService/ListMessages",
		"grpc-message": "o5.dempe.v1.ListMessagesRequest",
	}
	return headers
}

// Method: GetMessage

func (msg *GetMessageRequest) MessagingHeaders() map[string]string {
	headers := map[string]string{
		"grpc-service": "/o5.dempe.v1.DempeService/GetMessage",
		"grpc-message": "o5.dempe.v1.GetMessageRequest",
	}
	return headers
}

// Method: MessagesAction

func (msg *MessagesActionRequest) MessagingHeaders() map[string]string {
	headers := map[string]string{
		"grpc-service": "/o5.dempe.v1.DempeService/MessagesAction",
		"grpc-message": "o5.dempe.v1.MessagesActionRequest",
	}
	return headers
}
