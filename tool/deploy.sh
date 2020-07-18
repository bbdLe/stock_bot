#!/bin/bash

time=`date +"%Y_%m_%d_%H_%M_%S"`
cd ../..

tar -czf stock_bot_$time.tar.gz stock_bot/cmd stock_bot/tool stock_bot/etc --exclude=stock_bot/cmd/main.go --exclude=stock_bot/tool/deploy.sh --exclude=nohup.out
mv stock_bot_$time.tar.gz stock_bot/tool