#!/bin/sh

RESUME_FILE=/resume/resume.json

sed -i 's/###PHONE###/'$PHONE'/' $RESUME_FILE
sed -i 's/###EMAIL###/'$EMAIL'/' $RESUME_FILE

npm install -g jsonresume-theme-${THEME}

# hackmyresume NEW $RESUME_FILE
hackmyresume VALIDATE $RESUME_FILE
hackmyresume BUILD -t /usr/local/lib/node_modules/jsonresume-theme-${THEME} $RESUME_FILE TO CV.html