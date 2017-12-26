#!/usr/bin/env bash

APPLICATION=io_match
APPLICATION_PATH=/home/apps/working/
LOG_PATH=/home/apps/working/
LOG_PATH_FILE=${LOG_PATH}${APPLICATION}.log

if [ ! -d "$LOG_PATH" ]
then
    mkdir -p ${LOG_PATH}
    touch ${LOG_PATH_FILE}
    chown -R apps:apps ${LOG_PATH}
fi


start(){
	nohup ${APPLICATION_PATH}${APPLICATION} >> ${LOG_PATH_FILE} 2>&1 &
	sleep 2
	ps aux|grep ${APPLICATION}|grep base.conf|grep -v grep
	tail -n 18 ${LOG_PATH_FILE}
}

stop(){
        pid=$(ps aux|grep io_match |grep -v grep|awk '{print $2}')
        if [[ $pid != "" ]]
        then
            ps aux|grep io_match|grep -v grep|awk '{print $2}'|xargs kill -15
        fi
}

restart(){
	stop && start
}

status(){
	ps aux|grep ${APPLICATION_PATH}|grep -v grep
}

case $1 in
	"start")
		start
	;;

	"stop")
		stop
	;;

	"restart")
		restart
	;;

	"status")
		status
	;;

	*)
		echo "Usage:{start|stop|restart|status}"
esac