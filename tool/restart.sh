#!/bin/bash

kill -TERM `cat ../etc/robot.pid`
cd ../cmd/
nohup ./cmd &