#!/bin/sh

OUT_DIR=.
PACK_NAME="models"
SUFFIX_NAME=""
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

# 检查 db2go 是否已安装
if ! which db2go >/dev/null 2>&1; then
    # 安装最新版 db2go
    go install github.com/civet148/db2go@latest

    # 检查是否安装成功
    if which db2go >/dev/null 2>&1; then
        echo "✅  db2go install success, $(which db2go)"
    else
        echo "❌  db2go install failed, please check go env and gcc tool-chain"
        exit 1
    fi
fi

db2go --url "$DSN_URL" --out "$OUT_DIR" --table "$TABLE_NAME" --json-properties "$JSON_PROPERTIES" --enable-decimal  --spec-type "$SPEC_TYPES" \
--suffix "$SUFFIX_NAME" --package "$PACK_NAME" --readonly "$READ_ONLY" --without "$WITH_OUT" --tinyint-as-bool "$TINYINT_TO_BOOL" \
--tag "$TAGS" --import-models "$IMPORT_MODELS"

echo "generate go file ok, formatting..."
gofmt -w $OUT_DIR/$PACK_NAME
db2go -v
