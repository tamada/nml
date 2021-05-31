# :alarm_clock: nml

[![build](https://github.com/tamada/nml/actions/workflows/build.yml/badge.svg)](https://github.com/tamada/nml/actions/workflows/build.yml)
[![Coverage Status](https://coveralls.io/repos/github/tamada/nml/badge.svg?branch=main)](https://coveralls.io/github/tamada/nml?branch=main)
[![Go Report Card](https://goreportcard.com/badge/github.com/tamada/nml)](https://goreportcard.com/report/github.com/tamada/nml)
[![codebeat badge](https://codebeat.co/badges/ad1c690d-9ba3-46f5-87b1-4f68fbdfb882)](https://codebeat.co/projects/github-com-tamada-nml-main)

[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg?logo=spdx)](https://github.com/tamada/nml/blob/main/LICENSE)

Notify me later!

## :speaking_head: Description



## :runner: Usage

```sh
nml [OPTIONS] <NUMBER> [-- <COMMANDS...>]
nml means 'notify me, later!'
OPTIONS
    -b, --background      runs nml on background.  Same as 'nml ... &'.
    -H, --with-header     shows the header on notification.
    -u, --unit <UNIT>     specifies the time unit. Default is min.
                          Available units are: nsec, msec, sec, min, and hour.
    -h, --help            prints this message.
NUMBER
    specifies the number for timer.
COMMANDS
    specifies the commands execute after timer.
    If no commands are specified, nml notifies by printing message
    "MESSAGE FROM NML ARRIVE!! (<NUMBER> <UNIT> left)" to STDOUT.
```

## :muscle: Compile

```
make
```

## :smile: About

### :man_office_worker: Authors :woman_office_worker:

* Haruaki Tamada ([tamada](https://github.com/tamada))

### :scroll: License

[Apache 2.0](https://github.com/tamada/nml/blob/main/LICENSE)

### :jack_o_lantern: Icon

![nml](https://github.com/tamada/nml/blob/main/docs/static/images/nml.svg)

This image is obtained from [freesvg.org](https://freesvg.org/zz-bell-silver)
