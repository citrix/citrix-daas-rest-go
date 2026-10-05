// Copyright © 2026. Citrix Systems, Inc.

package citrixclient

import "strings"

// IsFeatureEnabled reports whether the named feature is enabled for the site. The site
// reports its enabled features when the client is set up, so this is a local lookup and
// works the same for on-premises and cloud deployments. A feature is enabled only if the
// site lists it.
func (c *CitrixDaasClient) IsFeatureEnabled(featureName string) bool {
	if c.ClientConfig == nil {
		return false
	}

	for _, enabledFeature := range c.ClientConfig.EnabledFeatures {
		if strings.EqualFold(enabledFeature, featureName) {
			return true
		}
	}

	return false
}
