#!/bin/bash

CONFIG_FILE=~/.hcp/config
TARGET=$1

start=0
end=0
cnt=1
while IFS="=" read key value
do
    if [[ $key == \[*] ]]; then
    if [ $key == "[$1]" ]; then
        start=$cnt
        section=$(echo $key | tr -d "[]")
    fi
    fi

    if [[ $value ]] && [[ $section == 'default' ]]; then
        end=$cnt
#         if [[ $key == 'max_cluster_cpu' ]]; then
# IOP=$(cat <<EOF
# $IOP
#         ,"maxClusterCpu" : $value
# EOF
# )
#         elif [[ $key == 'max_cluster_mem' ]]; then
# IOP=$(cat <<EOF
# $IOP
#        , "maxClusterMem" : $value
# EOF
# )
#         elif [[ $key == 'default_node_option' ]]; then
# IOP=$(cat <<EOF
# $IOP
#         ,"defaultNodeOption" : $value
# EOF
# )
#         elif [[ $key == 'extra' ]]; then
# IOP=$(cat <<EOF
# $IOP
#         ,"extra" : $value
# EOF
# )
#         fi
    fi
    cnt=$(( cnt + 1 ))
done < $CONFIG_FILE

if [[ ! -z $start ]]; then
    exit 0
else
    sed "${start},${end}d" $CONFIG_FILE > $CONFIG_FILE
fi