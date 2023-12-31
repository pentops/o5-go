// Code generated by protoc-gen-go-messaging. DO NOT EDIT.

package dante_spb

// Service: DeadMessageCommandService
// Method: UpdateDeadMessage

func (msg *UpdateDeadMessageRequest) MessagingHeaders() map[string]string {
	headers := map[string]string{
		"grpc-service": "/o5.dante.v1.service.DeadMessageCommandService/UpdateDeadMessage",
		"grpc-message": "o5.dante.v1.service.UpdateDeadMessageRequest",
	}
	return headers
}

// Method: ReplayDeadMessage

func (msg *ReplayDeadMessageRequest) MessagingHeaders() map[string]string {
	headers := map[string]string{
		"grpc-service": "/o5.dante.v1.service.DeadMessageCommandService/ReplayDeadMessage",
		"grpc-message": "o5.dante.v1.service.ReplayDeadMessageRequest",
	}
	return headers
}

// Method: RejectDeadMessage

func (msg *RejectDeadMessageRequest) MessagingHeaders() map[string]string {
	headers := map[string]string{
		"grpc-service": "/o5.dante.v1.service.DeadMessageCommandService/RejectDeadMessage",
		"grpc-message": "o5.dante.v1.service.RejectDeadMessageRequest",
	}
	return headers
}
