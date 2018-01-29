#!/bin/bash

while getopts ":p:e:a:c:z:" arg; do
	case "${arg}" in
		p)
		  PHONE=${OPTARG}
		  ;;
		e)
		  EMAIL=${OPTARG}
		  ;;
		a)
		  ADDRESS=${OPTARG}
		  ;;
		c)
		  CITY=${OPTARG}
		  ;;
		z)
		  CODE=${OPTARG}
		  ;;
	esac
done

#Convert YAML to JSON
python ./src/python/yml2json.py -i ./resume/resume.yml -o ./resume/resume.json

#Build Docker Image
#docker build ./src/docker/Dockerfile -t personalresume

#Run Docker Command 
docker run -e 'OUTPUT_TEMPLATE=cora' \
               -e "PHONE=${PHONE}" \
               -e "EMAIL=${EMAIL}" \
               -e "ADDRESS=${ADDRESS}" \
               -e "CITY=${CITY}" \
               -e "CODE=${CODE}" \
               -e 'DISPLAY=unix$DISPLAY' \
               -v $(pwd)/resume/:/usr/share/nginx/html/ \
               -v /tmp/.X11-unix:/tmp/.X11-unix \
               personalresume
