#!/bin/sh

RESUME_FILE=/resume/resume.json
THEME=jsonresume-theme-cora

sed -i 's/###PHONE###/'$PHONE'/' $RESUME_FILE
sed -i 's/###EMAIL###/'$EMAIL'/' $RESUME_FILE
sed -i "s/###CITY###/$CITY/" $RESUME_FILE
sed -i "s/###REGION###/$REGION/" $RESUME_FILE

npm install -g $THEME

# hackmyresume NEW $RESUME_FILE
hackmyresume VALIDATE $RESUME_FILE
hackmyresume BUILD -t /usr/local/lib/node_modules/$THEME $RESUME_FILE TO CV.html