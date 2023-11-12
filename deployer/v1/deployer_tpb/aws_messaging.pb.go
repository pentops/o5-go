// Code generated by protoc-gen-go-messaging. DO NOT EDIT.

package deployer_tpb

// Service: AWSCommandTopic
// Method: CreateNewStack

func (msg *CreateNewStackMessage) MessagingHeaders() map[string]string {
	headers := map[string]string{
		"grpc-service": "/o5.deployer.v1.topic.AWSCommandTopic/CreateNewStack",
		"grpc-message": "o5.deployer.v1.topic.CreateNewStackMessage",
	}
	return headers
}

// Method: UpdateStack

func (msg *UpdateStackMessage) MessagingHeaders() map[string]string {
	headers := map[string]string{
		"grpc-service": "/o5.deployer.v1.topic.AWSCommandTopic/UpdateStack",
		"grpc-message": "o5.deployer.v1.topic.UpdateStackMessage",
	}
	return headers
}

// Method: DeleteStack

func (msg *DeleteStackMessage) MessagingHeaders() map[string]string {
	headers := map[string]string{
		"grpc-service": "/o5.deployer.v1.topic.AWSCommandTopic/DeleteStack",
		"grpc-message": "o5.deployer.v1.topic.DeleteStackMessage",
	}
	return headers
}

// Method: ScaleStack

func (msg *ScaleStackMessage) MessagingHeaders() map[string]string {
	headers := map[string]string{
		"grpc-service": "/o5.deployer.v1.topic.AWSCommandTopic/ScaleStack",
		"grpc-message": "o5.deployer.v1.topic.ScaleStackMessage",
	}
	return headers
}

// Method: CancelStackUpdate

func (msg *CancelStackUpdateMessage) MessagingHeaders() map[string]string {
	headers := map[string]string{
		"grpc-service": "/o5.deployer.v1.topic.AWSCommandTopic/CancelStackUpdate",
		"grpc-message": "o5.deployer.v1.topic.CancelStackUpdateMessage",
	}
	return headers
}

// Method: UpsertSNSTopics

func (msg *UpsertSNSTopicsMessage) MessagingHeaders() map[string]string {
	headers := map[string]string{
		"grpc-service": "/o5.deployer.v1.topic.AWSCommandTopic/UpsertSNSTopics",
		"grpc-message": "o5.deployer.v1.topic.UpsertSNSTopicsMessage",
	}
	return headers
}

// Method: RunDatabaseMigration

func (msg *RunDatabaseMigrationMessage) MessagingHeaders() map[string]string {
	headers := map[string]string{
		"grpc-service": "/o5.deployer.v1.topic.AWSCommandTopic/RunDatabaseMigration",
		"grpc-message": "o5.deployer.v1.topic.RunDatabaseMigrationMessage",
	}
	return headers
}

// Service: AWSStatusTopic
// Method: StackStatusChanged

func (msg *StackStatusChangedMessage) MessagingHeaders() map[string]string {
	headers := map[string]string{
		"grpc-service": "/o5.deployer.v1.topic.AWSStatusTopic/StackStatusChanged",
		"grpc-message": "o5.deployer.v1.topic.StackStatusChangedMessage",
	}
	return headers
}

// Method: MigrationStatusChanged

func (msg *MigrationStatusChangedMessage) MessagingHeaders() map[string]string {
	headers := map[string]string{
		"grpc-service": "/o5.deployer.v1.topic.AWSStatusTopic/MigrationStatusChanged",
		"grpc-message": "o5.deployer.v1.topic.MigrationStatusChangedMessage",
	}
	return headers
}
