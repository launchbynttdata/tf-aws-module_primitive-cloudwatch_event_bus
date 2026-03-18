package testimpl

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/eventbridge"
	eventbridgetypes "github.com/aws/aws-sdk-go-v2/service/eventbridge/types"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/launchbynttdata/lcaf-component-terratest/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func getAWSConfig(t *testing.T) aws.Config {
	cfg, err := config.LoadDefaultConfig(context.Background())
	require.NoError(t, err, "unable to load AWS SDK config")
	return cfg
}

func getEventBridgeClient(t *testing.T) *eventbridge.Client {
	return eventbridge.NewFromConfig(getAWSConfig(t))
}

func TestComposableComplete(t *testing.T, ctx types.TestContext) {
	eventBusName := terraform.Output(t, ctx.TerratestTerraformOptions(), "name")
	eventBusARN := terraform.Output(t, ctx.TerratestTerraformOptions(), "arn")
	eventBusID := terraform.Output(t, ctx.TerratestTerraformOptions(), "id")
	expectedKMSKeyARN := terraform.Output(t, ctx.TerratestTerraformOptions(), "kms_key_arn")

	assert.Equal(t, eventBusName, eventBusID, "id and name should be equal for EventBridge event bus")

	client := getEventBridgeClient(t)

	// Verify event bus exists and configuration via API
	result, err := client.DescribeEventBus(context.Background(), &eventbridge.DescribeEventBusInput{
		Name: aws.String(eventBusName),
	})
	require.NoError(t, err, "DescribeEventBus should succeed")
	require.NotNil(t, result, "DescribeEventBus result should not be nil")

	assert.Equal(t, eventBusName, aws.ToString(result.Name), "event bus name should match")
	assert.Equal(t, eventBusARN, aws.ToString(result.Arn), "event bus ARN should match")

	// Security verification: KMS encryption must be configured
	require.NotNil(t, result.KmsKeyIdentifier, "KMS key identifier must be present — encryption should be configured")
	assert.Equal(t, expectedKMSKeyARN, aws.ToString(result.KmsKeyIdentifier), "KMS key should match the key provisioned by Terraform")

	// Write operation: send a test event to the event bus
	detail, err := json.Marshal(map[string]string{"source": "terratest", "action": "verify"})
	require.NoError(t, err, "json.Marshal should succeed")
	putResult, err := client.PutEvents(context.Background(), &eventbridge.PutEventsInput{
		Entries: []eventbridgetypes.PutEventsRequestEntry{
			{
				Source:       aws.String("terratest.cloudwatch-event-bus"),
				DetailType:   aws.String("TerratestVerification"),
				Detail:       aws.String(string(detail)),
				EventBusName: aws.String(eventBusName),
			},
		},
	})
	require.NoError(t, err, "PutEvents should succeed")
	require.Len(t, putResult.Entries, 1, "PutEvents should return one entry")
	require.Empty(t, putResult.Entries[0].ErrorCode, "PutEvents entry should not have an error")
	assert.NotNil(t, putResult.Entries[0].EventId, "PutEvents should return an event ID")
}

func TestComposableCompleteReadOnly(t *testing.T, ctx types.TestContext) {
	eventBusName := terraform.Output(t, ctx.TerratestTerraformOptions(), "name")
	eventBusARN := terraform.Output(t, ctx.TerratestTerraformOptions(), "arn")
	eventBusID := terraform.Output(t, ctx.TerratestTerraformOptions(), "id")
	expectedKMSKeyARN := terraform.Output(t, ctx.TerratestTerraformOptions(), "kms_key_arn")

	assert.Equal(t, eventBusName, eventBusID, "id and name should be equal for EventBridge event bus")

	client := getEventBridgeClient(t)

	// Read-only: verify event bus exists and configuration via API
	result, err := client.DescribeEventBus(context.Background(), &eventbridge.DescribeEventBusInput{
		Name: aws.String(eventBusName),
	})
	require.NoError(t, err, "DescribeEventBus should succeed")
	require.NotNil(t, result, "DescribeEventBus result should not be nil")

	assert.Equal(t, eventBusName, aws.ToString(result.Name), "event bus name should match")
	assert.Equal(t, eventBusARN, aws.ToString(result.Arn), "event bus ARN should match")

	// Security verification: KMS encryption must be configured
	require.NotNil(t, result.KmsKeyIdentifier, "KMS key identifier must be present — encryption should be configured")
	assert.Equal(t, expectedKMSKeyARN, aws.ToString(result.KmsKeyIdentifier), "KMS key should match the key provisioned by Terraform")
}
