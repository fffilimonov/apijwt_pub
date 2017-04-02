# apijwt
Golang API with JWT auth

git status | awk '{if($0~"modified"&&$0~".go"){print $2}}' | while read FILE; do go fmt $FILE;done

docker build --build-arg CACHE_DATE=$(date +%Y-%m-%d:%H:%M:%S) -t fffilimonov/apijwt .;docker push fffilimonov/apijwt;hyper pull fffilimonov/apijwt;hyper images
