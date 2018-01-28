#!/bin/bash

#Convert YAML to JSON
python ./src/python/yml2json.py -i ./resume/resume.yml -o ./resume/resume.json

#Run Docker Command 
docker run -e 'OUTPUT_TEMPLATE=positive' \
		-e 'PHONE=0123456789' \
		-e 'EMAIL=email@email.com' \
		-e 'ADDRESS=myAddress' \
		-e 'CITY=City' \
		-e 'CODE=Code' \
		-e 'DISPLAY=unix$DISPLAY' \
		-v $(pwd)/resume/:/usr/share/nginx/html/ \
		-v /tmp/.X11-unix:/tmp/.X11-unix \
		nouchka/hackmyresume
