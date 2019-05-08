# Cron as a microservice
An OMG service for Cron, it is a time-based job scheduler in Unix-like computer operating systems.

[![Open Microservice Guide](https://img.shields.io/badge/OMG-enabled-brightgreen.svg?style=for-the-badge)](https://microservice.guide)
[![Build Status](https://travis-ci.org/heaptracetechnology/microservice-cron.svg?branch=master)](https://travis-ci.org/heaptracetechnology/microservice-cron)
[![codecov](https://codecov.io/gh/heaptracetechnology/microservice-cron/branch/master/graph/badge.svg)](https://codecov.io/gh/heaptracetechnology/microservice-cron)


## [OMG](hhttps://microservice.guide) CLI

### OMG

* omg validate
```
omg validate
```
* omg build
```
omg build
```
### Test Service

* Test the service by following OMG commands

### CLI

##### Cron Trigger
```sh
$  omg subscribe event triggers -a interval=<INTERVAL> -a initial_delay=<INITIAL_DELAY>
```
## License
### [MIT](https://choosealicense.com/licenses/mit/)

## Docker
### Build
```
docker build -t microservice-cron .
```
### RUN
```
docker run -p 3000:3000 microservice-cron
```
