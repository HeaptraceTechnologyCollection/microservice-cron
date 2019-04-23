# Cron as a microservice
An OMG service for Cron, it is a time-based job scheduler in Unix-like computer operating systems.

[![Open Microservice Guide](https://img.shields.io/badge/OMG-enabled-brightgreen.svg?style=for-the-badge)](https://microservice.guide)


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
$  omg subscribe cron trigger -a interval=<INTERVAL> -e delay_interval=<DELAY_INTERVAL>
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
