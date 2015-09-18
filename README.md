### Fleet health

For Fleet Health to work your fleet server, needs to have it's
API enabled.
To do so on CoreOS look at :

https://coreos.com/fleet/docs/latest/deployment-and-configuration.html

Environment settings :

The ip/hostname and port of your fleet service. Default to 127.0.0.1 if empty

```
FH_FLEET_API=<ip>:<port>
```

Time is seconds to do a fleet unit status check. Defaults to 60 seconds if empty

```
FH_CHECK_INTERVAL=30
```

You can either run fleethealth as a standalone application or use docker(recommended)


To build docker container

```
docker build -t fleethealth .
```

To run

```
docker run -t fleethealth
```
