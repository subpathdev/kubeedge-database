# KubeEdge Database

## Table of Content
* [Description](#description)
* [Setup](#setup)

## Description
This extension to [KubeEdge](https://kubeedge.io) will append KubeEdge device/crd with the
with a database connection. This tool will run on the kubernetes cluster an not in on the edge.
In this first version we will use
[PostgreSQL](https://www.postgresql.org) as database with a opportunity to use automatically the
[Timescale](https://www.timescale.com) plugin. To view the data you can connect to the database via the commandline
tool `psql` or the graphical tool `pgadmin` but we recommend the usage of [Grafana](https://grafana.com/). If you are
using grafana, you have to convert the value of the to a number. The next example shows you an example grafana database
query to use by time series views. We will specify the data to one device in one namespace.
```SQL
SELECT $_time(time), value::numeric FROM devices WHERE $_timeFilter(time) AND device="test" AND namespace="default"
```
[TODO Beispiel Bild einfügen von Grafana]: <>

## Setup
In this readme we will not explain how to setup Grafana, PostgreSQL, Timescale or KubeEdge. Before you are trying to setting
this up, you have to deploy PostgreSQL and KubeEdge with the device/crd in your kubernetes cluster. We are also not describing
how to create a database role, a database or a schema in PostgreSQL. But you need a database role and a database to set up
this connector. With this description you
should be able to set up the kubeedge-database in your own kubernetes cluster. The next to chapters will show you two
different ways to install this kubeedge database connector.

### Yaml
In this case you have to update the YAML file which are provided in the root directory of this repository.
The following has to be change to your own specification in you kubernetes cluster:
- address; this has to be set to your database; specified in the kubeedge-database deployment
- user; this has to be set to a database user; specified in the kubeedge-database secret
- password; this has to be set to the database user password; specified in the kubeedge-database secret

To apply the pod you can execute this command this:
```
kubectl apply -f kubernetes.yaml
```
