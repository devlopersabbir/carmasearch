package elastic

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/carmasearch/carma-server/api/vehicle/core"
	db "github.com/carmasearch/carma-server/internal/database"
)

// BulkIndexVehicles sends a batch of vehicles to Elasticsearch using the Bulk API.
// It builds a newline-delimited NDJSON body (action meta + source) and issues a
// single _bulk request. Each document is upserted (index action) keyed by vehicle ID.
func BulkIndexVehicles(ctx context.Context, vehicles []*core.Vehicle) error {
	if db.ESClient == nil {
		return fmt.Errorf("elasticsearch client is not initialised")
	}
	if len(vehicles) == 0 {
		return nil
	}

	var buf bytes.Buffer

	for _, v := range vehicles {
		// Action meta line
		meta := map[string]interface{}{
			"index": map[string]interface{}{
				"_index": db.ESIndexName,
				"_id":    strconv.Itoa(int(v.ID)),
			},
		}
		metaLine, err := json.Marshal(meta)
		if err != nil {
			return fmt.Errorf("marshal meta for vehicle %d: %w", v.ID, err)
		}
		buf.Write(metaLine)
		buf.WriteByte('\n')

		// Source line
		srcLine, err := json.Marshal(v)
		if err != nil {
			return fmt.Errorf("marshal vehicle %d: %w", v.ID, err)
		}
		buf.Write(srcLine)
		buf.WriteByte('\n')
	}

	res, err := db.ESClient.Bulk(
		bytes.NewReader(buf.Bytes()),
		db.ESClient.Bulk.WithContext(ctx),
		db.ESClient.Bulk.WithIndex(db.ESIndexName),
		db.ESClient.Bulk.WithRefresh("false"), // avoid refresh overhead during bulk import
	)
	if err != nil {
		return fmt.Errorf("bulk request failed: %w", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		return fmt.Errorf("bulk response error: %s", res.String())
	}

	// Parse response to detect per-document errors
	var bulkResp struct {
		Errors bool `json:"errors"`
		Items  []map[string]struct {
			ID     string `json:"_id"`
			Status int    `json:"status"`
			Error  *struct {
				Type   string `json:"type"`
				Reason string `json:"reason"`
			} `json:"error,omitempty"`
		} `json:"items"`
	}

	if err := json.NewDecoder(res.Body).Decode(&bulkResp); err != nil {
		return fmt.Errorf("decode bulk response: %w", err)
	}

	if bulkResp.Errors {
		var errs []string
		for _, item := range bulkResp.Items {
			if action, ok := item["index"]; ok && action.Error != nil {
				errs = append(errs, fmt.Sprintf("id=%s: %s – %s", action.ID, action.Error.Type, action.Error.Reason))
			}
		}
		log.Printf("⚠️  bulk index partial errors (%d): %v", len(errs), errs)
	}

	log.Printf("✅ bulk indexed %d vehicles", len(vehicles))
	return nil
}
