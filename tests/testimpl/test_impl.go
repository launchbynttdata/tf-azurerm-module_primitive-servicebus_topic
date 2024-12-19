package testimpl

import (
	"context"
	"net/url"
	"os"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus/admin"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/launchbynttdata/lcaf-component-terratest/types"
	"github.com/stretchr/testify/assert"
)

func TestComposableComplete(t *testing.T, ctx types.TestContext) {
	subscriptionId := os.Getenv("ARM_SUBSCRIPTION_ID")
	if len(subscriptionId) == 0 {
		t.Fatal("ARM_SUBSCRIPTION_ID environment variable is not set")
	}

	credential, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		t.Fatalf("Unable to get credentials: %e\n", err)
	}

	// The client requires the full hostname of the service bus
	busEndpoint := terraform.Output(t, ctx.TerratestTerraformOptions(), "endpoint")
	u, err := url.Parse(busEndpoint)
	if err != nil {
		t.Fatalf("Unable to parse service bus endpoint: %e\n", err)
	}
	busName := u.Hostname()

	adminClient, err := admin.NewClient(busName, credential, nil)
	if err != nil {
		t.Fatalf("Unable to create service bus admin client: %e\n", err)
	}

	t.Run("DoesTopicExist", func(t *testing.T) {
		topicName := terraform.Output(t, ctx.TerratestTerraformOptions(), "name")
		resp, err := adminClient.GetTopic(context.TODO(), topicName, nil)
		if err != nil {
			t.Fatalf("Unable to retrieve topic: %e\n", err)
		}

		assert.Equal(t, topicName, resp.TopicName, "Expected topic name to be %s, got %s", topicName, resp.TopicName)
		assert.Equal(t, "Active", string(*resp.Status), "Expected topic status to be Active, got %s", string(*resp.Status))
	})
}
