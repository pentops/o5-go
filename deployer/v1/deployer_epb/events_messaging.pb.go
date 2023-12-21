// Code generated by protoc-gen-go-messaging. DO NOT EDIT.

package deployer_epb

// Service: DeployerEventsTopic
// Method: DeploymentEvent

func (msg *DeploymentEventMessage) MessagingTopic() string {
	return "o5-deployer-event"
}
func (msg *DeploymentEventMessage) MessagingHeaders() map[string]string {
	headers := map[string]string{
		"grpc-service": "/o5.deployer.v1.events.DeployerEventsTopic/DeploymentEvent",
		"grpc-message": "o5.deployer.v1.events.DeploymentEventMessage",
	}
	return headers
}

// Method: StackEvent

func (msg *StackEventMessage) MessagingTopic() string {
	return "o5-deployer-event"
}
func (msg *StackEventMessage) MessagingHeaders() map[string]string {
	headers := map[string]string{
		"grpc-service": "/o5.deployer.v1.events.DeployerEventsTopic/StackEvent",
		"grpc-message": "o5.deployer.v1.events.StackEventMessage",
	}
	return headers
}
