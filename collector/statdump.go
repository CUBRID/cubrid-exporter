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

// Scrape CUBRID Statdump data.

package collector

import (
	"context"
	"database/sql"

	"strconv"

	"github.com/prometheus/client_golang/prometheus"
)

const (
	statdump = "statdump"

	statdumpQuery = "show statdump demodb"
)

// Metric descriptors.
var (
	StatdumpInfo = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "statdump", "info"),
		"Information about CUBRID Statdump", []string{"key"}, nil,
	)
)

// ScrapeStatdump
type ScrapeStatdump struct{}

// Name of the Scraper. Should be unique.
func (ScrapeStatdump) Name() string {
	return statdump
}

// Help describes the role of the Scraper.
func (ScrapeStatdump) Help() string {
	return "Scrape information from statdump Query"
}

// Version of CUBRID from which scraper is available.
func (ScrapeStatdump) Version() float64 {
	return 10.2
}

// Scrape collects data from database connection and sends it over channel as prometheus metric.
func (ScrapeStatdump) Scrape(ctx context.Context, db *sql.DB, ch chan<- prometheus.Metric) error {

	statdumpRows, err := db.QueryContext(ctx, statdumpQuery)
	if err != nil {
		return err
	}

	var key string
	var value string

	for statdumpRows.Next() {

		err := statdumpRows.Scan(&key, &value)
		if err != nil {
			return err
		}

		floatValue, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return err
		}

		ch <- prometheus.MustNewConstMetric(StatdumpInfo, prometheus.GaugeValue, floatValue, key)
	}

	return nil
}

// check interface
var _ Scraper = ScrapeStatdump{}
