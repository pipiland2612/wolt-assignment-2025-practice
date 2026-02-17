package calculator

import (
	"testing"

	"golang-api-practice/internal/model"

	"github.com/stretchr/testify/require"
)

func createVenue() model.Venue {
	return model.Venue{
		Location: &model.Location{
			Coordinate: []float64{60.0, 25.0},
		},
		DeliverySpecs: &model.DeliverySpecs{
			OrderMin: 200,
			DeliveryPricing: &model.DeliveryPricing{
				BasePrice: 100,
				DistanceRanges: []model.DistanceRange{
					{
						Min: 0,
						Max: 1000,
						A:   10,
						B:   1,
					},
				},
			},
		},
	}
}

func TestTotalFee_Success(t *testing.T) {
	venue := createVenue()

	userCoords := model.Location{
		Coordinate: []float64{60.0, 25.0},
	}

	tests := []struct {
		name      string
		cartValue float64
		expected  float64
	}{
		{
			name:      "cart above minimum",
			cartValue: 300,
			// delivery = round(100 + 10 + 0)
			// surcharge = max(300 - 200, 0) = 100
			// total = 300 + 100 + 110
			expected: 510,
		},
		{
			name:      "cart below minimum",
			cartValue: 150,
			// surcharge = max(150 - 200, 0) = 0
			// delivery = 110
			// total = 150 + 110
			expected: 260,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			total, err := TotalFee(tt.cartValue, &userCoords, &venue)
			require.NoError(t, err)
			require.Equal(t, tt.expected, total)
		})
	}
}

func TestTotalFee_OutOfRange(t *testing.T) {
	venue := createVenue()
	userCoords := model.Location{
		Coordinate: []float64{60.0, 200.0},
	}

	_, err := TotalFee(300, &userCoords, &venue)
	require.Error(t, err)
}
