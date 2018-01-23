# Image

Forked over from https://github.com/nouchka/docker-hackmyresume

Use from command line, put your resume.json in /tmp/resume/ and file will be generate in same directory:

	docker run -e 'OUTPUT_TEMPLATE=<template>' \
		-e 'PHONE=<number>' \
		-e 'EMAIL=<email>' \
		-e 'ADDRESS=<address>' \
		-e 'DISPLAY=unix$DISPLAY' \
		-v /tmp/resume/:/usr/share/nginx/html/ \
		-v /tmp/.X11-unix:/tmp/.X11-unix \
		nouchka/hackmyresume
		
resume.json format is based on https://github.com/fresh-standard/FRESCA and https://jsonresume.org/schema/
