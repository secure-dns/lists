#!/bin/bash

mkdir -p res
FROM=$PWD
for folder in lists/*/
do
  cd $FROM/$folder
  tar -czf "${FROM}/res/$(echo $folder | sed -e "s/^lists\///" -e "s/\/$//").tar.gz" *
  cd $FROM
done
