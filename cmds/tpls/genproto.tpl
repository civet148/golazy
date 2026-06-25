#!/bin/bash

PROTO=
SWAGGER=0
for arg in "$@"
do
	if [[ $arg == *".proto" ]]; then
		PROTO=$arg
	elif [[ $arg == "swagger" ]]; then
		SWAGGER=1
	fi
done

if [ -z "$PROTO" ]; then
	echo "Need to specify .proto file"
	exit 1
fi


PROTONAME=${PROTO%.proto}
uname=$(uname);
case "$uname" in
    (*Linux*) READLINK=readlink; ;;
    (*Darwin*) READLINK=readlink; ;;
esac;
ROOT=`pwd`
PROJECTROOT=${ROOT#"$GOPATH"/src/}
SCRIPTROOT=$(dirname $($READLINK -f $0))

PROTOFILEPATH="$ROOT/$PROTO"

GATEWAYOPT="--grpc-gateway_out=logtostderr=true:."
if [ -f "$PROTONAME.yaml" ]; then
    echo "Additional gateway definition found $PROTONAME.yaml"
    GATEWAYOPT="--grpc-gateway_out=logtostderr=true,grpc_api_configuration=${PROJECTROOT}/${PROTONAME}.yaml:."
fi

cd ${GOPATH}/src

protoc -I. \
    -I${ROOT} \
    -I${GOPATH}/src \
    -I${GOPATH}/src/github.com/googleapis/googleapis \
    -I${GOPATH}/src/github.com/gogo/protobuf \
	--go_out=plugins=grpc:. \
    --govalidators_out=. \
    ${GATEWAYOPT} \
    ${PROJECTROOT}/${PROTO}


if [[ $SWAGGER == 1 ]]; then
protoc -I. \
    -I${ROOT} \
    -I${GOPATH}/src \
    -I${GOPATH}/src/github.com/googleapis/googleapis \
    -I${GOPATH}/src/github.com/gogo/protobuf \
	--openapiv2_out=enums_as_ints=true:. \
   	${PROJECTROOT}/${PROTO}
fi