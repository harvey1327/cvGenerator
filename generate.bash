#!/bin/bash

while getopts ":p:e:" arg; do
	case "${arg}" in
		p)
		  PHONE=${OPTARG}
		  ;;
		e)
		  EMAIL=${OPTARG}
		  ;;
	esac
done

#Convert YAML to JSON
python ./src/python/yml2json.py -i ./resume/resume.yml -o ./resume/resume.json

#Run Docker Command 
docker run -e 'OUTPUT_TEMPLATE=cora' \
               -e "PHONE=${PHONE}" \
               -e "EMAIL=${EMAIL}" \
               -e 'ADDRESS=myAddress' \
               -e 'CITY=City' \
               -e 'CODE=Code' \
               -e 'DISPLAY=unix$DISPLAY' \
               -v $(pwd)/resume/:/usr/share/nginx/html/ \
               -v /tmp/.X11-unix:/tmp/.X11-unix \
               nouchka/hackmyresume
