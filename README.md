# _Cron_ OMG Microservice

[![Open Microservice Guide](https://img.shields.io/badge/OMG%20Enabled-👍-green.svg?)](https://microservice.guide)
[![Build Status](https://travis-ci.org/heaptracetechnology/microservice-cron.svg?branch=master)](https://travis-ci.org/heaptracetechnology/microservice-cron)
[![codecov](https://codecov.io/gh/heaptracetechnology/microservice-cron/branch/master/graph/badge.svg)](https://codecov.io/gh/heaptracetechnology/microservice-cron)

An OMG service for Cron, it is a time-based job scheduler in Unix-like computer operating systems.

## Direct usage in [Storyscript](https://storyscript.io/):

```coffee
>>> cron subscribe

```

Curious to [learn more](https://docs.storyscript.io/)?

✨🍰✨

## Usage with [OMG CLI](https://www.npmjs.com/package/omg)
##### Cron Trigger
```sh
$  omg subscribe event triggers -a interval=<INTERVAL> -a initial_delay=<INITIAL_DELAY>
```

**Note**: the OMG CLI requires [Docker](https://docs.docker.com/install/) to be installed.

## License
[MIT License](https://github.com/omg-services/cron/blob/master/LICENSE).
