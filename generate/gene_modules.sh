#!/bin/sh
#go install github.com/ZenLiuCN/engine/gene@v0.6.2

pkg="github.com/ZenLiuCN/engineX/modules/"
checkEmpty() {
    if [[ -n "$(ls -A $1|grep '.go')" ]]; then
        return 1
      else
        return 0
    fi
}
function generate() {
    find $1 -type d |  while read -r file; do
      if [[ $file == $1 ]] ;then
          echo "ignore root $file"
      elif [[ "$file" == *sample* ]] ;then
          echo "ignore sample $file"
      elif [[ "$file" == */.* ]] ;then
          echo "ignore dot $file"
      elif [[ "$file" == */doc ]] ;then
          echo "ignore doc $file"
      elif [[ "$file" == *"/ws" ]]; then
         gene -rt $pkg -i=github.com/gogo/protobuf/proto -ro ../modules $file
      else
        checkEmpty $file
        if [ $? -eq 0 ] ;then
           echo "ignore empty $file"
        else
          gene -rt $pkg -ro ../modules $file
        fi
      fi
    done
}

generate "D:/Dev/tmp/go/pkg/mod/github.com/larksuite/base-sdk-go/v3@v3.0.2/"
generate "D:/Dev/tmp/go/pkg/mod/github.com/larksuite/oapi-sdk-go/v3@v3.3.0/"
