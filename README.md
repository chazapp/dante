#Tigbra

An ELK installation for class notes index and searching.

## Objectives

This project intends to provide Elastic Search and Kibana visualization for books citations.
Book text quotes are written and tagged in TXT, processed by the `pkg/parser` and serialized to JSON.
These documents are then saved to a MongoDB instance to be indexed by ElasticSearch. 

## Dataset

The dataset are not to be commited in the repository.  
The input dataset TXT format is as follows (WIP):
```txt
1.A category name
A theme
-"A citation" p.42
-"A citation" p.43
-"A citation" p.44
Another theme
-"Another citation" p.45
-"Another citation" p.46
2.Another category
A theme
-"A citation" p.47
...
```

## Usage

```bash
$ GO111MODULE=on go mod vendor
$ go build cmd/cli/main.go
$ ./main parse -f dataset.txt 
```

Example output: 
```json
  {"category":"1. A category name", "theme":"Condition féminine","quote":"A citation","page":"11"}
  {"category":"1. A category name", "theme":"Condition féminine","quote":"A citation","page":"11"}
  {"category":"1. A category name", "theme":"Condition féminine","quote":"A citation","page":"11"}
``` 