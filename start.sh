#!/bin/bash

RESUME=$PWD/resume

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

#Run python container to convert yaml to json
echo '--------Starting conversion process'
EXIST=$(docker images -q converter:latest)
if [ -z $EXIST ]; then
echo '--------Building Image converter:latest'
	docker build ./converter/ -t converter:latest
fi
docker run -v $RESUME:/tmp converter:latest
echo '--------Ending conversion process'

#Run node container to generate resume
echo '--------Starting generation process'
EXIST=$(docker images -q generator:latest)
if [ -z $EXIST ]; then
echo '--------Building Image generator:latest'
	docker build ./generator/ -t generator:latest
fi
docker run -e "PHONE=${PHONE}" \
			-e "EMAIL=${EMAIL}" \
			-e "CITY=${CITY}" \
			-e "REGION=${REGION}" \
			-v $RESUME:/resume \
			generator:latest
echo '--------Ending generation process'