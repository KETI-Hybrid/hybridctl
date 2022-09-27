#!/bin/bash

CONFIG_FILE=~/.hcp/config

IOP=$(cat <<EOF
[
EOF
)


while IFS="=" read key value
do
    if [[ $key == \[*] ]]; then
    if [ -z $section ]; then
        section=$(echo $key | tr -d "[]")
IOP=$(cat <<EOF
$IOP
        {
            "section" : "$section"
EOF
)
    else
        section=$(echo $key | tr -d "[]")
IOP=$(cat <<EOF
$IOP
        },
        {
            "section" : "$section"
EOF
)
    fi
    fi

    if [[ $value ]] && [[ $section == 'default' ]]; then
        if [[ $key == 'max_cluster_cpu' ]]; then
IOP=$(cat <<EOF
$IOP
        ,"maxClusterCpu" : $value
EOF
)
        elif [[ $key == 'max_cluster_mem' ]]; then
IOP=$(cat <<EOF
$IOP
       ,"maxClusterMem" : $value
EOF
)
        elif [[ $key == 'default_node_option' ]]; then
IOP=$(cat <<EOF
$IOP
        ,"defaultNodeOption" : "$value"
EOF
)
        elif [[ $key == 'extra' ]]; then
IOP=$(cat <<EOF
$IOP
        ,"extra" : $value
EOF
)
        fi
    fi
done < $CONFIG_FILE

echo $IOP
if [[ $IOP == "[" ]]; then
    echo "[]" > tmp.json
    exit 0
else
IOP=$(cat <<EOF
$IOP
    }
]
EOF
)
fi

# tmp.json에 저장
echo $IOP > tmp.json