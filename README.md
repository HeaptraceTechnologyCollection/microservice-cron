# _Cron_ Open Microservice

[![Open Microservices](https://img.shields.io/badge/OMS%20Enabled-👍-green.svg?)](https://openmicroservices.org)
[![Build Status](https://travis-ci.com/oms-services/cron.svg?branch=master)](https://travis-ci.com/oms-services/cron)
[![codecov](https://codecov.io/gh/omg-services/cron/branch/master/graph/badge.svg)](https://codecov.io/gh/oms-services/cron)

An OMG service for Cron, it is a time-based job scheduler in Unix-like computer operating systems.

## Direct usage in [Storyscript](https://storyscript.io/):

```coffee
>>> cron subscribe

```

Curious to [learn more](https://docs.storyscript.io/)?

✨🍰✨

## Usage with [OMS CLI](https://www.npmjs.com/package/@microservices/oms)
##### Cron Trigger
```sh
$  oms subscribe event triggers -a interval=<INTERVAL> -a initialDelay=<INITIAL_DELAY>
```

**Note**: the OMS CLI requires [Docker](https://docs.docker.com/install/) to be installed.

## License
[MIT License](https://github.com/oms-services/cron/blob/master/LICENSE).
