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

// Scrape CUBRID broker status data.

package collector

import (
	"context"
	"database/sql"

	"strconv"

	"github.com/prometheus/client_golang/prometheus"
)

const (
	brokerStatus = "broker_status"

	brokerStatusQuery = "show brokers"
)

// Metric descriptors.
var (
	BrokersInfo = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "broker_status", "info"),
		"Information about CUBRID Broker Status",
		[]string{"broker_name", "num_as", "pid", "port", "qsize", "num_select", "num_insert", "num_update", "num_delete", "num_trans", "num_conns", "tps", "qps"}, nil,
	)

	BrokerInfo = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "broker_status", "info"),
		"Information about CUBRID Broker Status",
		[]string{"broker_name", "key"}, nil,
	)
)

// ScrapeBrokerStatus
type ScrapeBrokerStatus struct{}

// Name of the Scraper. Should be unique.
func (ScrapeBrokerStatus) Name() string {
	return brokerStatus
}

// Help describes the role of the Scraper.
func (ScrapeBrokerStatus) Help() string {
	return "Scrape information from brokerStatusQuery"
}

// Version of CUBRID from which scraper is available.
func (ScrapeBrokerStatus) Version() float64 {
	return 10.2
}

// Scrape collects data from database connection and sends it over channel as prometheus metric.
func (ScrapeBrokerStatus) Scrape(ctx context.Context, db *sql.DB, ch chan<- prometheus.Metric) error {

	brokerStatusRows, err := db.QueryContext(ctx, brokerStatusQuery)
	if err != nil {
		return err
	}

	defer brokerStatusRows.Close()

	var broker_name string
	var num_as string
	var pid string
	var port string
	var qsize string
	var num_select string
	var num_insert string
	var num_update string
	var num_delete string
	var num_trans string
	var num_query string
	var num_conns string
	var num_long_query string
	var num_error_query string
	var num_uniq_error string

	for brokerStatusRows.Next() {

		err := brokerStatusRows.Scan(&broker_name, &num_as, &pid, &port, &qsize, &num_select, &num_insert, &num_update, &num_delete, &num_trans, &num_query, &num_conns, &num_long_query, &num_error_query, &num_uniq_error)
		if err != nil {
			return err
		}

		count, _ := strconv.ParseFloat(num_as, 64)
		ch <- prometheus.MustNewConstMetric(BrokerInfo, prometheus.GaugeValue, count, broker_name, "num_as")

		count, _ = strconv.ParseFloat(pid, 64)
		ch <- prometheus.MustNewConstMetric(BrokerInfo, prometheus.GaugeValue, count, broker_name, "pid")

		count, _ = strconv.ParseFloat(port, 64)
		ch <- prometheus.MustNewConstMetric(BrokerInfo, prometheus.GaugeValue, count, broker_name, "port")

		count, _ = strconv.ParseFloat(qsize, 64)
		ch <- prometheus.MustNewConstMetric(BrokerInfo, prometheus.GaugeValue, count, broker_name, "qsize")

		count, _ = strconv.ParseFloat(num_select, 64)
		ch <- prometheus.MustNewConstMetric(BrokerInfo, prometheus.GaugeValue, count, broker_name, "num_select")

		count, _ = strconv.ParseFloat(num_insert, 64)
		ch <- prometheus.MustNewConstMetric(BrokerInfo, prometheus.GaugeValue, count, broker_name, "num_insert")

		count, _ = strconv.ParseFloat(num_update, 64)
		ch <- prometheus.MustNewConstMetric(BrokerInfo, prometheus.GaugeValue, count, broker_name, "num_update")

		count, _ = strconv.ParseFloat(num_delete, 64)
		ch <- prometheus.MustNewConstMetric(BrokerInfo, prometheus.GaugeValue, count, broker_name, "num_delete")

		count, _ = strconv.ParseFloat(num_trans, 64)
		ch <- prometheus.MustNewConstMetric(BrokerInfo, prometheus.GaugeValue, count, broker_name, "num_trans")

		count, _ = strconv.ParseFloat(num_query, 64)
		ch <- prometheus.MustNewConstMetric(BrokerInfo, prometheus.GaugeValue, count, broker_name, "num_query")

		count, _ = strconv.ParseFloat(num_conns, 64)
		ch <- prometheus.MustNewConstMetric(BrokerInfo, prometheus.GaugeValue, count, broker_name, "num_conns")

		count, _ = strconv.ParseFloat(num_long_query, 64)
		ch <- prometheus.MustNewConstMetric(BrokerInfo, prometheus.GaugeValue, count, broker_name, "num_long_query")

		count, _ = strconv.ParseFloat(num_error_query, 64)
		ch <- prometheus.MustNewConstMetric(BrokerInfo, prometheus.GaugeValue, count, broker_name, "num_error_query")

		count, _ = strconv.ParseFloat(num_uniq_error, 64)
		ch <- prometheus.MustNewConstMetric(BrokerInfo, prometheus.GaugeValue, count, broker_name, "num_uniq_error")
	}

	return nil
}

// check interface
var _ Scraper = ScrapeBrokerStatus{}
