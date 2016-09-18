# chronosctl
A tool to interact with [Mesosphere's Chronos](https://mesos.github.io/chronos/docs/) via the commandline.

## Features
For now, you can perform the following actions:
* List jobs and see an overview of their status
* Inspect individual jobs
* Launch jobs at will
* Create jobs
* Create docker jobs
* Delete jobs
* Kill tasks for a given job

## Compiling
Just enter `make` in a terminal, it will work out just fine.

## Using
Help is provided. Just enter `./chronosctl -h`

## Configuring
You must use a configuration file named `chronosctl.yml`

It should look like:
```yaml
url: https://yourchronosendpoint
# Those two are optional, they are the basic auth
# credentials required :)
username: chronos
password: ieatzeus
```
## License
GPL v3

## Contributing
I'll more than gladly accept contributions, whatever
they are, pull requests, suggestions, tests, fixes or
what ever do not hesitate to submit amelioration.

For any question, ask [me](https://twitter.com/thomas_maurice)
