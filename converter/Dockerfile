FROM python:3.8-alpine

RUN pip install pyyaml
COPY yml2json.py /usr/local/bin
ENTRYPOINT ["python", "/usr/local/bin/yml2json.py", "-i", "/tmp/resume.yml", "-o", "/tmp/resume.json"]