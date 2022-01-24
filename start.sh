#!/bin/bash

RESUME=$PWD/resume

while getopts ":p:e:t:" arg; do
	case "${arg}" in
		p)
		  PHONE=${OPTARG}
		  ;;
		e)
		  EMAIL=${OPTARG}
		  ;;
		t)
		  THEME=${OPTARG}
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
docker run --name converter -v $RESUME:/tmp converter:latest
echo '--------Ending conversion process'

#Run node container to generate resume
echo '--------Starting generation process'
EXIST=$(docker images -q generator:latest)
if [ -z $EXIST ]; then
echo '--------Building Image generator:latest'
	docker build ./generator/ -t generator:latest
fi
docker run --name generator \
			-e "PHONE=${PHONE}" \
			-e "EMAIL=${EMAIL}" \
			-e "THEME=${THEME}" \
			-v $RESUME:/resume \
			generator:latest
echo '--------Ending generation process'

echo '--------Clearing containers'
docker container rm converter generator