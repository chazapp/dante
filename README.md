# Project Dante

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

Print result to the standard output.
```bash
$ GO111MODULE=on go mod vendor
$ cd /go/src/github.com/chazapp/dante/cmd/cli && go build -o dante-cli 
$ ./dante-cli print -f dataset.txt 
```

Example output: 
```json
  {"category":"1. A category name", "theme":"A theme","quote":"A citation","page":40}
  {"category":"1. A category name", "theme":"A theme","quote":"A citation","page":41}
  {"category":"1. A category name", "theme":"A theme","quote":"A citation","page":42}
``` 

Store the results in a running mongodb instance.
```bash
./main mongo --db mongodb://localhost:27017 --file dataset.txt --name "BookName"
```
Store the results in a running ElasticSearch instance.
```bash
./main elastic --db http://localhost:9200 --file dataset.txt --name "BookName"
```

## Docker Orchestration
The aim of the project is to provide visualization to data gathered in book quotes. These visualisations
are provided by ElasticSearch and Kibana. To this effect, a `docker-compose.yml` is available.
To start the infrastructure:    
```bash
$ docker-compose up
Creating network "dante_default" with the default driver
Creating dante_kibana_1         ... done
Creating dante_elasticsearch_1  ... done
Creating dante_cli_1            ... done
'Compose: docker-compose.yml' has been deployed successfully.
```
The DanteCLI will create an index in ElasticSearch, then upload the processed dataset provided in the Dockerfile.
Kibana is then available at http://localhost:5061 to create visualizations.  
Warning: The current `docker-compose.yml` does not store the work written in Kibana yet. The visualizations have to be
exported before killing the process. 

## Inspiration

Tigbra
![TigbraProject](https://i.redd.it/n11syzm0v6x21.jpg)
