package client

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFetchApiReal(t *testing.T) {
	venues := []string{
		"home-assignment-venue-helsinki",
		"home-assignment-venue-stockholm",
		"home-assignment-venue-berlin",
		"home-assignment-venue-tokyo",
	}

	for _, v := range venues {
		t.Run(v, func(t *testing.T) {
			result, err := FetchApi(v)
			require.NoError(t, err, "fetch failed for %s", v)

			// check location
			require.NotNil(t, result.Location, "venue %s: location is nil", v)
			require.NotNil(t, result.Location.Coordinate, "venue %s: coordinates are empty", v)

			// check delivery specs
			require.NotNil(t, result.DeliverySpecs, "venue %s: delivery specs is nil", v)

			if result.DeliverySpecs.DeliveryPricing != nil {
				require.NotEmpty(t, result.DeliverySpecs.DeliveryPricing.DistanceRanges,
					"venue %s: delivery pricing distance ranges empty", v)
			}
		})
	}

}
