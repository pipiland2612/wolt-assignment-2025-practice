package client

import (
	"context"
	"net/http"
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

	client := NewClient(&http.Client{})

	for _, v := range venues {
		t.Run(v, func(t *testing.T) {
			res, err := client.FetchApi(context.Background(), v)
			require.NoError(t, err, "fetch failed for %s", v)

			// check location
			require.NotNil(t, res.Location, "venue %s: location is nil", v)
			require.NotNil(t, res.Location.Coordinate, "venue %s: coordinates are empty", v)

			// check delivery specs
			require.NotNil(t, res.DeliverySpecs, "venue %s: delivery specs is nil", v)

			if res.DeliverySpecs.DeliveryPricing != nil {
				require.NotEmpty(t, res.DeliverySpecs.DeliveryPricing.DistanceRanges,
					"venue %s: delivery pricing distance ranges empty", v)
			}
		})
	}

}
