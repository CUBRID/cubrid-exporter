# cubrid-exporter
A Prometheus interface for CUBRID Monitoring
# CUBRID_exporter - An interface providing CUBRID db metric to the prometheus
Overview
========
```
A Prometheus interface modules to provide CUBRID DB monitoring metrics to the prometheus.
Date: January, 2020
```

Abstract
========
Prometheus is an open source platform to collect various monitoring information from
multiple sources and keep it into the TSDB (Time Series DataBase).
CUBRID exporter is a gateway between prometheus and 
CUBRID MAS (Management Application Server, An CUBRID Monitoring Broker) running on CUBRID
Database nodes. It collects CUBRID database metrics such as 'database statistics', 'database volume info'
and 'cubrid broker status'.

Dependencies for CUBRID exporter
================================
```
  * CUBRID 9.3.9 or higher
  * OS: Windows (x86 and x86_64)
        Linux 64bit
  * Compiler: Go
```

How to Build
============
```
go build cubrid_exporter.go
```

Configure CUBRID Exporter
=========================
```
edit cubrid_exporter.ini
```
