# bleve-explorer

An example app providing an HTTP/REST/JSON front-end to bleve.  It provides a REST API and an HTML interface to:

* create/list/delete indexes
* view index document count
* view index mapping
* index/delete documents
* query indexes
* monitor system performance

## Building

        go build -tags full

The ```-tags full``` is optional, but includes all the optional components of bleve.

## Running

		mkdir data
        ./bleve-explorer

This will use the default "data" dir for storing indexes.  Once started you can access the web UI at http://localhost:8095/

## Screenshots

Tabs showing operations available on an index

![](docs/index.png)

The monitoring capabilities

![](docs/monitor.png)