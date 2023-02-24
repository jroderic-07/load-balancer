### Load Balancer
Load balancer tool developed using go.

## Installation
Clone the repository and build using the make file.
```bash
git clone https://github.com/jroderic-07/load-balancer.git 
make
``` 

## Usage
Default values are:
- Port: :8080
- config-file: data/config.json
- lb-type: roundRobin

Specify port:
```bash
./bin/loadbalancer -port :8000
```

Specify configuration file:
```bash
./bin/loadbalancer -config-file /tmp/test.json
```

Specify load balancing type:
```bash
./bin/loadbalancer -lb-type roundRobin
```

For help, run:
```bash
./bin/loadbalancer --help
```

## Requirements
- UNIX-based system
- Go
