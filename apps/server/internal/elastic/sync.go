package elastic

import (
	"context"
	"log"

	"github.com/carmasearch/carma-server/api/vehicle/core"
	db "github.com/carmasearch/carma-server/internal/database"
)

// SyncOptions controls the batching behaviour of SyncAllVehicles.
type SyncOptions struct {
	// BatchSize is the number of PostgreSQL rows fetched and indexed per iteration.
	// Defaults to 1000 if zero or negative.
	BatchSize int
}

// defaultBatchSize is the number of records processed per round-trip.
const defaultBatchSize = 1000

// VehiclePageFetcher is a function that fetches a page of vehicles from the
// persistent store. offset is the number of rows to skip; limit is the page
// size. It returns the rows and any error.
//
// Accepting a function rather than a concrete repository lets callers inject
// any data source (Postgres, CSV, etc.) and keeps this package free of direct
// GORM / domain dependencies.
type VehiclePageFetcher func(ctx context.Context, limit, offset int) ([]*core.Vehicle, error)

// SyncAllVehicles pages through all vehicles returned by fetcher, indexing
// them into Elasticsearch in batches. It respects the ctx cancellation so it
// can be interrupted gracefully.
//
// Returns the total number of vehicles synced and any terminal error.
func SyncAllVehicles(ctx context.Context, fetcher VehiclePageFetcher, opts SyncOptions) (int, error) {
	if db.ESClient == nil {
		return 0, nil
	}

	batchSize := opts.BatchSize
	if batchSize <= 0 {
		batchSize = defaultBatchSize
	}

	total := 0
	offset := 0

	for {
		// Stop if the caller cancelled the context
		select {
		case <-ctx.Done():
			return total, ctx.Err()
		default:
		}

		vehicles, err := fetcher(ctx, batchSize, offset)
		if err != nil {
			return total, err
		}

		if len(vehicles) == 0 {
			// No more data
			break
		}

		log.Printf("🔄 syncing batch offset=%d size=%d", offset, len(vehicles))

		if err := BulkIndexVehicles(ctx, vehicles); err != nil {
			return total, err
		}

		total += len(vehicles)
		offset += len(vehicles)

		if len(vehicles) < batchSize {
			// Last (possibly partial) page – we are done
			break
		}
	}

	log.Printf("✅ sync complete: %d vehicles indexed", total)
	return total, nil
}
