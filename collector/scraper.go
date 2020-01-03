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

package collector

import (
	"context"
	"database/sql"

	_ "github.com/cubrid/cubrid-go"
	"github.com/prometheus/client_golang/prometheus"
)

type Scraper interface {
	// Name of the Scraper. Should be unique.
	Name() string

	// Help describes the role of the Scraper.
	// Example: "Collect from SHOW ENGINE INNODB STATUS"
	Help() string

	// Version of CUBRID from which scraper is available.
	Version() float64

	// Scrape collects data from database connection and sends it over channel as prometheus metric.
	Scrape(ctx context.Context, db *sql.DB, ch chan<- prometheus.Metric) error
}
