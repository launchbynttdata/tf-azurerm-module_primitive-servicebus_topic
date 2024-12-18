package testimpl

import (
	"context"
	"net/url"
	"os"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/launchbynttdata/lcaf-component-terratest/types"
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

	client, err := azservicebus.NewClient(busName, credential, nil)
	if err != nil {
		t.Fatalf("Unable to create service bus admin client: %e\n", err)
	}

	topicName := terraform.Output(t, ctx.TerratestTerraformOptions(), "name")

	t.Run("CanClientSendMessage", func(t *testing.T) {
		sender, err := client.NewSender(topicName, nil)
		if err != nil {
			t.Fatalf("Unable to create a Sender: %e\n", err)
		}
		defer sender.Close(context.TODO())

		err = sender.SendMessage(context.TODO(), &azservicebus.Message{Body: []byte("Hello, World!")}, nil)
		if err != nil {
			t.Fatalf("Unable to send message: %e\n", err)
		}
	})
}
