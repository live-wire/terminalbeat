# Terminalbeat :pager:

Welcome to Terminalbeat.

This [beat](https://github.com/elastic/beats) runs a command (specified in the config file `terminalbeat.yml`) and forwards the corresponding [stdout](https://www.computerhope.com/jargon/s/stdout.htm) to Logstash or Elasticsearch.


Ensure that this folder is at the following location:
`${GOPATH}/src/github.com/live-wire/terminalbeat`

## Getting Started with Terminalbeat

### Requirements

* [Golang](https://golang.org/dl/) 1.7

### Init Project
To get running with Terminalbeat and also install the
dependencies, run the following command:

```
make setup
```

It will create a clean git history for each major step. Note that you can always rewrite the history if you wish before pushing your changes.

To push Terminalbeat in the git repository, run the following commands:

```
git remote set-url origin https://github.com/live-wire/terminalbeat
git push origin master
```

For further development, check out the [beat developer guide](https://www.elastic.co/guide/en/beats/libbeat/current/new-beat.html).

### Build

To build the binary for Terminalbeat run the command below. This will generate a binary
in the same directory with the name terminalbeat.

```
make
```


### Run

To run Terminalbeat with debugging output enabled, run:

```
./terminalbeat -c terminalbeat.yml -e -d "*"
```


### Test

To test Terminalbeat, run the following command:

```
make testsuite
```

alternatively:
```
make unit-tests
make system-tests
make integration-tests
make coverage-report
```

The test coverage is reported in the folder `./build/coverage/`

### Update

Each beat has a template for the mapping in elasticsearch and a documentation for the fields
which is automatically generated based on `fields.yml` by running the following command.

```
make update
```


### Cleanup

To clean  Terminalbeat source code, run the following command:

```
make fmt
```

To clean up the build directory and generated artifacts, run:

```
make clean
```


### Clone

To clone Terminalbeat from the git repository, run the following commands:

```
mkdir -p ${GOPATH}/src/github.com/live-wire/terminalbeat
git clone https://github.com/live-wire/terminalbeat ${GOPATH}/src/github.com/live-wire/terminalbeat
```


For further development, check out the [beat developer guide](https://www.elastic.co/guide/en/beats/libbeat/current/new-beat.html).


## Packaging

The beat frameworks provides tools to crosscompile and package your beat for different platforms. This requires [docker](https://www.docker.com/) and vendoring as described above. To build packages of your beat, run the following command:

```
make release
```

This will fetch and create all images required for the build process. The whole process to finish can take several minutes.
