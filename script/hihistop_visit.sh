#!/bin/bash

fmtTime() {
	timeType=$1
	if [ $timeType == "perHour" ];then
		echo $(date "+%Y-%m-%d %H" -d "+0 hour")
	elif [ $timeType == "per6Hour" ];then
		echo $(date "+%Y-%m-%d %H" -d "+0 hour")
	elif [ $timeType == "perDay" ];then
		echo $(date "+%Y-%m-%d" -d "+0 hour")
	else
		echo '[Error] Param Error'
	fi
}
