#!/bin/sh

if [ $1 -gt 0 ] 2>/dev/null ;then
	times=$1
else
	times=1
fi


if [ $2 -gt 0 ] 2>/dev/null ;then
	sec=$2
else
	sec=1
fi

vari=0
while (($vari < $times))
do
	echo "==================================================================="
	echo "第一个参数循环次数(默认为1)"
	echo "第二个参数为间隔时间(默认为1)"
	
	vari=$(($vari+1))
	sleep $sec
	
	if test `uname -a|awk '{print $1}'` == "Linux" ;then
		echo "Linux"
		
		netstat -ant|grep LISTEN|grep -v fff|grep -v "::"|awk '{print $4"-"}'|awk -F ':' '{print $2}'|sort -u > linux899x8iitmpd.xx1
		netstat -ant|grep ESTAB|grep -v fff|awk '{print $4"-"}'|awk -F ':' '{print $2}'|sort -u > linux899x8iitmpd.xx2
		sort linux899x8iitmpd.xx1 > linux899x8iitmpd.xx3
		sort linux899x8iitmpd.xx2 > linux899x8iitmpd.xx4
		comm  linux899x8iitmpd.xx3  linux899x8iitmpd.xx4 -1 -2 > linux899x8iitmpd.xx5
		netstat -antp|grep ESTAB|grep -v fff|awk '{print $4"- "$5" "$7}'|sort  > linux899x8iitmpd.xx6

		for i in $(cat linux899x8iitmpd.xx5)
		do
		grep  $i linux899x8iitmpd.xx6 >> linux899x8iitmpd.xx7
		done
		
		echo -e "\033[41;37mOther IP to localhost:\033[0m"
		
		printf "%-26s" "对端IP"
		echo -e "\033[47;30m本机IP:本地监听端口\033[0m"
		printf "%-26s%s\n" "===================" "=============="
		cat linux899x8iitmpd.xx7 |grep -v "127.0.0.1"|awk -F ':' '{printf("%s:%s\n",$1,$2)}'|awk -F '- ' '{printf("%-26s%s\n",$2,$1)}'|sort -u

		sort linux899x8iitmpd.xx6 > linux899x8iitmpd.xx8
		sort linux899x8iitmpd.xx7 > linux899x8iitmpd.xx9
		:> linux899x8iitmpd.xx7
		comm linux899x8iitmpd.xx8 linux899x8iitmpd.xx9 -2 -3 > linux899x8iitmpd.xx10
#==================================================================================================================================
		echo -e "\033[41;37mConnections TO other IP:\033[0m"
		
		echo -e "\033[47;30m本地IP\033[0m\c"
		printf "%45s\n" "对端IP:对端监听端口"
		printf "%-24s%s\n" "==============" "==================="
		cat linux899x8iitmpd.xx10 |grep -v "127.0.0.1"|awk '{print $1":"$2":"$3}'|awk -F ':' '{printf("%-24s%s:%s\n",$1,$3,$4)}'|sort -u
		
		echo "-----------------------------------------------"
		netstat -antp |grep ESTAB |grep -v fff|awk '{print $5}'|awk -F ':' '{print $1}' > linux899x8iitmpd.xx11
		for j in `cat linux899x8iitmpd.xx11 |sort -u`
		do
			count=0
			for x in `cat linux899x8iitmpd.xx11`
			do
				if test $j == $x ;then
					count=$(($count+1))
				fi
			done
			echo -e "本机与 "$j" 的连接个数: "$count"\t"
		done
		
	else
		echo "Unix"

		netstat -Aan |grep LISTEN|grep -v tcp6|awk '{print $5}'|awk -F '.' '{print $2"-"}'|sort -u > jj827348907.tmpy1

		netstat -Aan |grep ESTABLISHED|awk '{print $5}'|awk -F '.' '{print $5"-"}'|sort -u > jj827348907.tmpy2

		sort jj827348907.tmpy1 > jj827348907.tmpy3
		sort jj827348907.tmpy2 > jj827348907.tmpy4

		comm -1 -2 jj827348907.tmpy3 jj827348907.tmpy4 > jj827348907.tmpy5

		netstat -Aan |grep ESTABLISHED |awk '{print $5"- "$6}'|sort > jj827348907.tmpy6

		for i in $(cat jj827348907.tmpy5)
		do 
			grep  $i jj827348907.tmpy6 >> jj827348907.tmpy7
		done
		
		echo "\033[41;37mOther IP to localhost:\033[0m"
#		printf "%-26s%s\n" "Source IP" "Dest IP:port"
#		printf "%-26s%s\n" "=========" "=============="
		printf "%-26s" "对端IP"
		echo  "\033[47;30m本机IP:本地监听端口\033[0m"
		printf "%-26s%s\n" "===================" "=============="
		cat jj827348907.tmpy7|awk -F '- ' '{print $1"."$2}'|awk -F '.' '{print $1"."$2"."$3"."$4":"$5" "$6"."$7"."$8"."$9}'|awk '{printf("%-26s%s\n",$2,$1)}'|sort

		sort jj827348907.tmpy6 > jj827348907.tmpy8
		sort jj827348907.tmpy7 > jj827348907.tmpy9
		:> jj827348907.tmpy7
		comm -2 -3 jj827348907.tmpy8 jj827348907.tmpy9 > jj827348907.tmpy10
		
		
		echo  "\033[41;37mConnections TO other IP:\033[0m"
#		printf "%-26s%s\n" "Source IP" "Dest IP:port"
#		printf "%-26s%s\n" "=========" "==============="
		echo  "\033[47;30m本地IP\033[0m\c"
		printf "%45s\n" "对端IP:对端监听端口"
		printf "%-24s%s\n" "==============" "==================="
		cat jj827348907.tmpy10|awk -F '- ' '{print $1"."$2}'|awk -F '.' '{print $1"."$2"."$3"."$4" "$6"."$7"."$8"."$9":"$10}'|awk '{printf("%-26s%s\n",$1,$2)}'|sort
		
		echo "-----------------------------------------------"
		netstat -ant |grep ESTAB |grep -v fff|awk '{print $5}'|awk -F '.' '{print $1"."$2"."$3"."$4}' > jj827348907.tmpy11
		for j in `cat jj827348907.tmpy11 |sort -u`
		do
			count=0
			for x in `cat jj827348907.tmpy11`
			do
				if test $j == $x ;then
					count=$(($count+1))
				fi
			done
			echo  "本机与 "$j" 的连接个数: "$count"\t"
		done
	fi
done
rm -f ./jj827348907.tmpy*
rm -f  ./linux899x8iitmpd.xx*