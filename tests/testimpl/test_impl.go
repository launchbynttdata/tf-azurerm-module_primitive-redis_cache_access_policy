package testimpl

import (
	"context"
	"os"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	armredis "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/redis/armredis/v3"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/launchbynttdata/lcaf-component-terratest/types"
	"github.com/stretchr/testify/assert"
)

func TestComposableComplete(t *testing.T, ctx types.TestContext) {
	subscriptionId := os.Getenv("ARM_SUBSCRIPTION_ID")

	if len(subscriptionId) == 0 {
		t.Fatal("ARM_SUBSCRIPTION_ID environment variable is not set")
	}

	// Create Azure credential using default credential chain
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	assert.NoError(t, err, "Failed to create Azure credential")


	resourceGroupName := terraform.Output(t, ctx.TerratestTerraformOptions(), "resource_group_name")
	redisCacheName := terraform.Output(t, ctx.TerratestTerraformOptions(), "redis_cache_name")
	accessPolicyName := terraform.Output(t, ctx.TerratestTerraformOptions(), "access_policy_name")

	t.Run("EnsureAccessPolicyExists", func(t *testing.T) {
		client, err := armredis.NewAccessPolicyClient(subscriptionId, cred, nil)
		assert.NoError(t, err, "Failed to create Redis access policy client")

		ctx := context.Background()
		accessPolicy, err := client.Get(ctx, resourceGroupName, redisCacheName, accessPolicyName, nil)
		assert.NoError(t, err, "Access policy should exist but was not found")
		assert.NotNil(t, accessPolicy.CacheAccessPolicy, "Access policy response should not be nil")
	})
}
