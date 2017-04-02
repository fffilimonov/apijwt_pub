#!/bin/bash


if [ $# -eq 0 ]; then
  IPPORT="likeuser.com";
else
  IPPORT=$1;
fi;

echo '<!DOCTYPE html>' > result.html;
echo '<html>' >> result.html;
echo '<head>' >> result.html;
echo '<title>Results</title>' >> result.html;
echo '</head>' >> result.html;
echo '<body>' >> result.html;

curl -s -H "Content-Type: application/json" -H "Accept: application/json" https://$IPPORT/api/signup -d '{"Username": "fffilimonov@yandex.ru", "Password": "12345"}';
curl -s -L http://$IPPORT/api/activate/ZmZmaWxpbW9ub3ZAeWFuZGV4LnJ1;
TOKEN=$(curl -s -H "Content-Type: application/json" -H "Accept: application/json" https://$IPPORT/api/login -d '{"Username": "fffilimonov@yandex.ru", "Password": "12345"}' | awk -F "\"" '{print $4}');
curl -s -H "Content-Type: application/json" -H "Accept: application/json" -H "Authorization: Bearer $TOKEN" -d '{"Name": "Test 1", "Text": "I open \"https://google.com\"\nI fill 1 input with \"my browser info\"\nI will see \"What'\''s My Browser\"\nI click on the text \"What'\''s My Browser\"\n"}' https://$IPPORT/api/add;
curl -s -L -H "Content-Type: application/json" -H "Accept: application/json" -H "Authorization: Bearer $TOKEN" http://$IPPORT/api/dashboard;
curl -s -H "Content-Type: application/json" -H "Accept: application/json" -H "Authorization: Bearer $TOKEN" -d '{"Scen": "Test 1","Browser": "chrome"}' https://$IPPORT/api/run | \
  while read LINE; do
    img=$(echo $LINE | grep -c data:image);
    if [ $img -gt 0 ]; then
      echo "<img style=\"display:block; width:100%;height:100%;\"src=\"$LINE\"/>" >> result.html;
    else
      echo $LINE >> result.html;
    fi;
  done
echo '</body>' >> result.html;
echo '</html>' >> result.html;

