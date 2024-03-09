mkdir ./build/
mkdir ./build/go-register/
mkdir ./build/go-register/static/

go build .

mv ./app.exe ./build/go-register/
cp auth.txt ./build/go-register/
cp provider.json ./build/go-register/

cp -r ./static/dependencies/ ./build/go-register/static/
cp -r ./static/js/ ./build/go-register/static/
cp -r ./static/view/ ./build/go-register/static/
cp -r ./static/index.html ./build/go-register/static/