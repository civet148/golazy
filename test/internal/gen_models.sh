#!/bin/sh

OUT_DIR=.
PACK_NAME="models"
SUFFIX_NAME="do"
READ_ONLY=""
TABLE_NAME=""
WITH_OUT=""
TAGS=""
TINYINT_TO_BOOL="is_deleted"
DSN_URL="mysql://root:12345678@127.0.0.1:3306/test?charset=utf8"
JSON_PROPERTIES=""
SPEC_TYPES=""
IMPORT_MODELS="test/internal/models"
COMMON_TAGS=""

if [ $? -eq 0 ]; then
db2go --debug --url $DSN_URL --out $OUT_DIR --table "$TABLE_NAME" --json-properties $JSON_PROPERTIES --enable-decimal  --spec-type "$SPEC_TYPES" \
--suffix $SUFFIX_NAME --package $PACK_NAME --readonly "$READ_ONLY" --without "$WITH_OUT" --tinyint-as-bool "$TINYINT_TO_BOOL" \
--tag "$TAGS" --dao dao --import-models $IMPORT_MODELS

echo "generate go file ok, formatting..."
gofmt -w $OUT_DIR/$PACK_NAME

else
  echo "error: db2go execute failed"
fi

