[![Build Status](https://travis-ci.org/rck/simplepush.svg?branch=master)](https://travis-ci.org/rck/simplepush)
[![GO Report](https://goreportcard.com/badge/github.com/rck/simplepush)](https://goreportcard.com/report/github.com/rck/simplepush)
[![Docker Automated build](https://img.shields.io/docker/automated/rck/simplepush.svg)](https://hub.docker.com/r/rck81/simplepush/)

# simplepush
Golang client and library for simplepush.io

# Synopsis
```
simplepush
Usage: simplepush -k key -m message [-t title] [-e event] [-p password]
  -e event
    	Set event
  -k key
    	Set simplepush.io key
  -m message
    	Set message
  -s salt
    	Set custom salt
  -p password
    	Set password, if set send message encrypted
  -t title
    	Set title
```

# Releases
Pre-built binaries are provided [here](https://github.com/rck/simplepush/releases/latest). Please note that these
binaries are automatically built by [Travis-CI](https://travis-ci.org). Your decision if you trust them.

# Docker
```
docker pull rck81/simplepush
docker run -it --rm rck81/simplepush -k YOURKEY -m "Test"
```
