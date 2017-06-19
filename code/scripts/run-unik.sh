#!/bin/bash
for i in $(seq 0 1 20)
do
   echo "time unik run --instanceName argo$i --imageName argo"
   time unik run --instanceName argo$i --imageName argo | grep 0m0
   sleep 4
   time unik delete-instance --force --instance argo$i
   sleep 2
done