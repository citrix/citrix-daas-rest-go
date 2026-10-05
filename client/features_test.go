// Copyright © 2026. Citrix Systems, Inc.

package citrixclient

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// The client is feature agnostic, so its tests use placeholder names rather than real ones.
const testFeatureName = "SomeFeature"

func TestIsFeatureEnabled(t *testing.T) {
	testCases := []struct {
		name            string
		enabledFeatures []string
		expectedEnabled bool
	}{
		{
			name:            "Feature Is Listed",
			enabledFeatures: []string{"OtherFeature", testFeatureName},
			expectedEnabled: true,
		},
		{
			name:            "Feature Is Not Listed",
			enabledFeatures: []string{"OtherFeature"},
			expectedEnabled: false,
		},
		{
			name:            "Nothing Enabled",
			enabledFeatures: []string{},
			expectedEnabled: false,
		},
		{
			// The site decides the casing it reports, so matching must not depend on it.
			name:            "Casing Does Not Matter",
			enabledFeatures: []string{"somefeature"},
			expectedEnabled: true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			client := &CitrixDaasClient{
				ClientConfig: &ClientConfiguration{EnabledFeatures: testCase.enabledFeatures},
			}

			assert.Equal(t, testCase.expectedEnabled, client.IsFeatureEnabled(testFeatureName))
		})
	}
}

// The feature check runs during ModifyPlan, which can happen before the provider is
// configured, so an unconfigured client must report the feature as disabled rather than panic.
func TestIsFeatureEnabled_UnconfiguredClient(t *testing.T) {
	client := &CitrixDaasClient{}

	assert.False(t, client.IsFeatureEnabled(testFeatureName))
}
