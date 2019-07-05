# _Cron_ OMG Microservice

[![Open Microservice Guide](https://img.shields.io/badge/OMG%20Enabled-👍-green.svg?)](https://microservice.guide)
[![Build Status](https://travis-ci.com/omg-services/cron.svg?branch=master)](https://travis-ci.com/omg-services/cron)
[![codecov](https://codecov.io/gh/omg-services/cron/branch/master/graph/badge.svg)](https://codecov.io/gh/omg-services/cron)

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
$  omg subscribe event triggers -a interval=<INTERVAL> -a initialDelay=<INITIAL_DELAY>
```

**Note**: the OMG CLI requires [Docker](https://docs.docker.com/install/) to be installed.

## License
[MIT License](https://github.com/omg-services/cron/blob/master/LICENSE).
