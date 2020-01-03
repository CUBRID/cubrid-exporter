// Copyright 2020 CUBRID Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Scrape CUBRID SpaceDB status data.

package collector

import (
	"context"
	"database/sql"
	"strconv"

	"github.com/prometheus/client_golang/prometheus"
)

const (
	spacedbStatus = "spacedb"

	spacedbQuery = "show spacedb demodb"
)

// Metric descriptors.
var (
	SpaceDbInfo = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "spacedb", "info"),
		"Information about CUBRID SpaceDB",
		[]string{"vol_no", "type", "purpose", "count", "used_pages", "free_pages"}, nil,
	)

	VolNoInfo = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "spacedb", "info"),
		"Information about CUBRID SpaceDB",
		[]string{"vol_no", "key"}, nil,
	)
)

// ScrapeSpaceDBStatus
type ScrapeSpaceDBStatus struct{}

// Name of the Scraper. Should be unique.
func (ScrapeSpaceDBStatus) Name() string {
	return spacedbStatus
}

// Help describes the role of the Scraper.
func (ScrapeSpaceDBStatus) Help() string {
	return "Scrape information from spacedbQuery"
}

// Version of CUBRID from which scraper is available.
func (ScrapeSpaceDBStatus) Version() float64 {
	return 10.2
}

// Scrape collects data from database connection and sends it over channel as prometheus metric.
func (ScrapeSpaceDBStatus) Scrape(ctx context.Context, db *sql.DB, ch chan<- prometheus.Metric) error {

	spaceDbRows, err := db.QueryContext(ctx, spacedbQuery)
	if err != nil {
		return err
	}

	defer spaceDbRows.Close()

	var vol_no string
	var _type string
	var purpose string
	var count string
	var used_pages string
	var free_pages string

	for spaceDbRows.Next() {

		err := spaceDbRows.Scan(&vol_no, &_type, &purpose, &count, &used_pages, &free_pages)
		if err != nil {
			return err
		}

		fValue, _ := strconv.ParseFloat(_type, 64)
		ch <- prometheus.MustNewConstMetric(VolNoInfo, prometheus.GaugeValue, fValue, vol_no, "_type")

		fValue, _ = strconv.ParseFloat(_type, 64)
		ch <- prometheus.MustNewConstMetric(VolNoInfo, prometheus.GaugeValue, fValue, vol_no, "purpose")

		fValue, _ = strconv.ParseFloat(count, 64)
		ch <- prometheus.MustNewConstMetric(VolNoInfo, prometheus.GaugeValue, fValue, vol_no, "count")

		fValue, _ = strconv.ParseFloat(used_pages, 64)
		fUsedPagesValue := fValue
		ch <- prometheus.MustNewConstMetric(VolNoInfo, prometheus.GaugeValue, fValue, vol_no, "used_pages")

		fValue, _ = strconv.ParseFloat(free_pages, 64)
		fFreePagesValue := fValue
		ch <- prometheus.MustNewConstMetric(VolNoInfo, prometheus.GaugeValue, fValue, vol_no, "free_pages")

		average := fUsedPagesValue / (fUsedPagesValue + fFreePagesValue) * 100
		if fUsedPagesValue == 0 {
			average = 0
		}
		ch <- prometheus.MustNewConstMetric(VolNoInfo, prometheus.GaugeValue, average, vol_no, "usedPercentage")

	}

	return nil
}

// check interface
var _ Scraper = ScrapeSpaceDBStatus{}
