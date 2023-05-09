# resumeHack

apt update
apt install latexmk
apt install texlive-fonts-extra

pdflatex -aux-directory=<DIRECTORY OF LOG & AUX FILES> -output-directory=<DIRECTORY OF OUTPUT PDF FILE> FILENAME.tex

go run cmd/main.go -email=my_email@email.com -phone=5647382910

docker build . test:v1
docker run -v $PWD/resume:/resume -it test:v1 sh

./main -email=my_email@email.com -phone=5647382910

pdflatex -output-directory=./resume ./template/modern/1.0/cv.tex

pdflatex --output-directory=./template/modern/1.0 ./template/modern/1.0/cv.tex