#!/bin/bash

# 1. select configuration

# Create a new configuration
# $1 - CONFIGURATION_NAME
EXIST=0
for instance in $(gcloud config configurations list --format="value(name)")
do
    INSTANCE_NO_WHITESPACE="$(echo ${instance} | tr -d ' ')"
    echo $INSTANCE_NO_WHITESPACE
    if [ $1 == $INSTANCE_NO_WHITESPACE ]; then
        EXIST=1
    fi
done


if [ ! $EXIST ]
then
   echo "[step 1] create a new configuration"
   gcloud config configurations create $1
else
   echo "[step 1] Switch to and re-initialize existing configuration: [$1]"
   gcloud config configurations activate $1
fi


# 2. set up credentials 
# $2 - KEY_FILE
echo "[step 2] set up credentials"
gcloud auth activate-service-account --key-file="/root/hcp-key.json"

# 3. set project
# $3 - PROJECT_ID
echo "[step 3] set project"
gcloud config set project $2

# [optional] 4. set default GCE Zone
# $4 - ZONE
if [ ! -z $3]
then 
    echo "[step 4] set default zone"
    gcloud config set compute/zone $3
fi

# [optional] 5. set default GCE region
# $5 - REGION
if [ ! -z $4]
then 
    echo "[step 5] set default region"
    gcloud config set compute/region $4
fi


# # Reference : https://stackoverflow.com/questions/42379685/can-i-automate-google-cloud-sdk-gcloud-init-interactive-command