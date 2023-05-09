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

echo '--------Starting generation process'
EXIST=$(docker images -q generator:latest)
if [ -z $EXIST ]; then
echo '--------Building Image generator:latest'
	docker build . -t generator:latest
fi
docker run --rm -v $PWD/resume:/resume -e "PHONE=${PHONE}" -e "EMAIL=${EMAIL}" generator:latest
echo '--------Ending generation process'