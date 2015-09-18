FROM golang:1.4-onbuild
MAINTAINER Werner Gillmer <werner.gillmer@gmail.com>

### Settings
# Your fleet api service IP/hostname with port.
# Look at <insert link to README on github> on how to enable the fleet API
ENV FH_FLEET_API 54.78.171.187:49153

# Interval in seconds to run the checks
ENV FH_CHECK_INTERVAL 30
