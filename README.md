# nml

Notify me later!

## Description



## Usage

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

## Compile

```
make
```

## About

### Authors

* Haruaki Tamada ([tamada](https://github.com/tamada))

### License

[Apache 2.0](https://github.com/tamada/nml/blob/main/LICENSE)

### Icon

![nml](https://github.com/tamada/nml/blob/main/docs/static/images/nml.svg)
