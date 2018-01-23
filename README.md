# Image

This image generates html / pdf file using npm package [hackmyresume](https://www.npmjs.com/package/hackmyresume).
[HackMyResume](http://please.hackmyresume.com/)  use standard format https://jsonresume.org/schema/ to generate the files.

[docker-hackmyresume-web](https://github.com/nouchka/docker-hackmyresume-web) extends this image with a webserver (nginx) for an all-in-one solution.

The entrypoint of the image:
* check if json file is missing, generates one
* update personal datas missing in json (###PHONE###, ###EMAIL### and ###ADDRESS###)
* validate json file
* install missing template
* generate output resume

# Use

Use from command line, put your resume.json in /tmp/resume/ and file will be generate in same directory:

	docker run -e 'OUTPUT_TEMPLATE=kendall' \
		-e 'PHONE=+33602030405' \
		-e 'EMAIL=docker@katagena.com' \
		-e 'ADDRESS=25 Rue Delphin Loche' \
		-e 'DISPLAY=unix$DISPLAY' \
		-v /tmp/resume/:/usr/share/nginx/html/ \
		-v /tmp/.X11-unix:/tmp/.X11-unix \
		nouchka/hackmyresume
or use with docker compose:
