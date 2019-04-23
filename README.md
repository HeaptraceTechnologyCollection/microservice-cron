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

##### Create and save user
```sh
$ omg run createuser -a user_id=<USER_ID> -a email=<EMAIL_ADDRESS> -a phone=<PHONE_NUMBER> -a name=<USER_NAME> -a custom_attributes=<CUSTOM_ATTRIBUTES> -a companies=<COMPANIES_LIST> -e ACCESS_TOKEN=<ACCESS_TOKEN>
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
