#!/bin/sh

if test -n "$(echo $1|sed -n "/^[0-9]\+$/p")"  ;then
	times=$1
else
	times=1
fi

if test -n "$(echo $2|sed -n "/^[0-9]\+$/p")"  ;then
	sec=$2
else
	sec=1
fi



if test  -n "$(echo $3|sed -n "/^[0-9]\+$/p")"  ;then
	localip=$3
else
	localip="xxx"
fi

#echo ${var1}
#echo ${var2}

vari=0

#for j in $(seq 1 $times)
#echo "=========xxxxxxxx"
#for (( i=0;i < $times; i++ ))
while (($vari < $times))
do
	echo "==================================================================="
	echo "第一个参数循环次数(默认为1)"
	echo "第二个参数为间隔时间(默认为1)"
	echo "第三个参数为本机IP(默认为空)"

	vari=$(($vari+1))




	if test `uname -a|awk '{print $1}'` == "Linux" ;then

		echo "Linux"
		echo -e "\033[41;37mConnections:\033[0m"
		printf "%-26s%-26s%s\n" "Source IP:port" "Dest IP:port" "PID/CMD"
		printf "%-26s%-26s%s\n" "==============" "============" "================"
		netstat -antp|grep ESTAB|grep -v fff|awk '{printf("%-26s%-26s%s\n"), $4,$5,$7}'

		#/bin/netstat -antp |grep ESTAB |sed 's/[0-9]:/=/'|awk -F '=' '{ printf("localhost:%s\n",$2) }'|awk 'BEGIN{print "Source IP:port\t\t\tDest IP:port\t\tPID/CMD";print "==============\t\t\t============\t\t======="}{printf("%-28s%-28s%-28s\n",$1,$2,$4)}'|grep -v "127.0.0.1"|grep -v ffff|grep -v $localip

		netstat -antp |grep ESTAB |grep -v fff|awk '{print $5}'|awk -F ':' '{print $1}' > sdf_x_tmp.5as45fas4d5f

		#for j in `netstat -antp |grep ESTAB |awk '{print $5}'|awk -F ':' '{for(i=1;i < NF-1;i++)$i="";print}'|sed 's/:/ /'|awk '{print $1}'|sort -u`
		for j in `cat sdf_x_tmp.5as45fas4d5f |sort -u`
		do
			count=0
			for x in `cat sdf_x_tmp.5as45fas4d5f`
			do
				if test $j == $x ;then
					count=$(($count+1))
				fi
			done
			echo -e "本机与 "$j" 的连接个数: "$count"\t"
		done

		rm -f ./sdf_x_tmp.5as45fas4d5f

		echo "==================================================================="

		#echo "duiduan:"
		#/bin/netstat -antp|grep ESTABLISHED|/bin/awk '{print $4"\t\t"$5"\t\t"$7}'

		echo -e "\033[41;37mListening Port:\033[0m"

		/bin/netstat -antp|grep LISTEN |/bin/awk 'BEGIN{print "Listening:port\t\tPID/CMD";print "============\t\t======="}{printf("%-24s%-24s\n",$4,$7)}'

	else
		echo "Unix"
		echo "\033[41;37mConnections:\033[0m"
		printf "%-26s%-26s%s\n" "Source IP:port" "Dest IP:port" "PID/CMD"
		printf "%-26s%-26s%s\n" "==============" "============" "================"
		for ux in `netstat -Aan |grep ESTABLISHED|awk '{print $1"="$5"="$6}'`
		do
			#echo $ux
			i=`echo $ux|awk -F '=' '{print $1}'`
			j=`rmsock $i tcpcb|awk '{print $9}'`
			k=`ps -ef|grep -v grep|awk '{print $2" "$8" "$9" "$10}'|grep $j`
			#k2=`ps -ef|grep -v grep|grep $j|awk '{print $9}'`
			u=`echo $ux|awk -F '=' '{print $2}'|awk -F '.' '{print $1"."$2"."$3"."$4":"$5}'`
			v=`echo $ux|awk -F '=' '{print $3}'|awk -F '.' '{print $1"."$2"."$3"."$4":"$5}'`
			#echo $u"\t"$v"\t"$k
			#echo "x"|awk -v vu=$u -v vv=$v -v vk=$k '{printf("%-28s%-28s%-28s\n",vu,vv,vk)}'
			printf "%-26s%-26s" $u $v
			echo $k
			#echo "====3"
		done

		netstat -ant |grep ESTAB |grep -v fff|awk '{print $5}'|awk -F '.' '{print $1"."$2"."$3"."$4}' > sdf_x_tmp.5as45fas4d5f
		for j in `cat sdf_x_tmp.5as45fas4d5f |sort -u`
		do
			count=0
			for x in `cat sdf_x_tmp.5as45fas4d5f`
			do
				if test $j == $x ;then
					count=$(($count+1))
				fi
			done
			echo  "本机与 "$j" 的连接个数: "$count"\t"
		done

		rm -f ./sdf_x_tmp.5as45fas4d5f

		echo "==================================================================="
		echo "\033[41;37mListening Port:\033[0m"
		printf "%-26s%s\n" "Listening:port" "PID/CMD"
		printf "%-26s%s\n" "==============" "================"
		for u2x in `netstat -Aan |grep LISTEN|awk '{print $1"="$5}'`
		do
			i=`echo $u2x|awk -F '=' '{print $1}'`
			j=`rmsock $i tcpcb|awk '{print $9}'`
			k=`ps -ef|grep -v grep|awk '{print $2" "$8" "$9" "$10}'|grep $j`
			#k2=`ps -ef|grep -v grep|grep $j|awk '{print $9}'`
			u=`echo $u2x|awk -F '=' '{print $2}'`
			printf "%-26s" $u
			echo $k
		done

	fi

	echo "==================================================================="
	sleep $sec

done
