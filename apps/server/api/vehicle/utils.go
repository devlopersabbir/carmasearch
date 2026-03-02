package vehicle

import esCore "github.com/carmasearch/carma-server/internal/elastic/core"

func intPtr(v int) *int {
	if v == 0 {
		return nil
	}
	return &v
}
func buildVehicleSearchAndCompare(
	body *esCore.CompareRequest,
	query *esCore.CompareRequestQuery,
) *esCore.VehicleSearchAndCompare {
	search := &esCore.VehicleSearchAndCompare{
		CompareRequestQuery: *query,
	}

	// Map body URL into search structure
	if body.Url != "" {
		search.ListingURL = &body.Url
	}

	return search
}

func float64Ptr(v float64) *float64 {
	if v == 0 {
		return nil
	}
	return &v
}
