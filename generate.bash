#!/bin/bash

while getopts ":p:e:r:c:" arg; do
	case "${arg}" in
		p)
		  PHONE=${OPTARG}
		  ;;
		e)
		  EMAIL=${OPTARG}
		  ;;
		r)
		  REGION=${OPTARG}
		  ;;
		c)
		  CITY=${OPTARG}
		  ;;
	esac
done

#Convert YAML to JSON
python ./src/python/yml2json.py -i ./resume/resume.yml -o ./resume/resume.json

#Build Docker Image if non-existant

EXIST=$(docker images -q resume:latest)
if [ -z $EXIST ]; then
	docker build ./src/docker/ -t resume:latest
fi

#Run Docker Command 
docker run -e 'OUTPUT_TEMPLATE=cora' \
               -e "PHONE=${PHONE}" \
               -e "EMAIL=${EMAIL}" \
               -e "CITY=${CITY}" \
               -e "REGION=${REGION}" \
               -e 'DISPLAY=unix$DISPLAY' \
               -v $(pwd)/resume/:/usr/share/nginx/html/ \
               -v /tmp/.X11-unix:/tmp/.X11-unix \
               resume:latest
