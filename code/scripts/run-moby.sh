#!/bin/bash
for i in $(seq 0 1 2)
do
   echo "time docker run --name argo$i gunjan5/argo:v1 "
   sleep 2
   time docker run --name argo$i gunjan5/argo:v1
done