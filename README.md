# filebeat.nsq.output
A Filebeat embedding a nsq output.

You can use the output in every beat you want. This repository offers a Filebeat "main" that embeds it.
You can compile and use it by following the Golang setup detailed in the CONTRIBUTE instructions of beats :
https://www.elastic.co/guide/en/beats/devguide/current/beats-contributing.html#setting-up-dev-environment

And compiling this repository using :

install go version 1.13.10  
update go.mod with beats/go.mod

```Go
git clone github.com/elastic/beats
git clone github.com/chennqqi/filebeat.nsq.output
cd filebeat.nsq.output
ln -s ../beats/vendor .
GO111MODULE=off go build
* go build

To build the project, and :

* ./filebeat.nsq.output

To launch it.

inspired by <https://github.com/dapicard/filebeat.mongodb.output>

## reference

* official develop documentation <https://www.elastic.co/guide/en/beats/devguide/current/beats-contributing.html#setting-up-dev-environment>
* an third party mongodb output <https://github.com/dapicard/filebeat.mongodb.output>
* official kafka output <https://github.com/elastic/beats/tree/master/libbeat/outputs/kafka>
