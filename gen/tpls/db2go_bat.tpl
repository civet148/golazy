@echo off

set OUT_DIR=.
set PACK_NAME="models"
set SUFFIX_NAME="do"
set READ_ONLY=""
set TABLE_NAME=""
set WITH_OUT=""
set TAGS="bson"
set TINYINT_TO_BOOL="is_deleted"
set DSN_URL="mysql://root:12345678@127.0.0.1:3306/test?charset=utf8"
set JSON_PROPERTIES=""
set SPEC_TYPES=""
set IMPORT_MODELS="{{ .importModel}}/models"
set COMMON_TAGS=""

If "%errorlevel%" == "0" (
db2go.exe --url %DSN_URL% --out %OUT_DIR% --table %TABLE_NAME% --json-properties %JSON_PROPERTIES% --enable-decimal  --spec-type %SPEC_TYPES% ^
--suffix %SUFFIX_NAME% --package %PACK_NAME% --readonly %READ_ONLY% --without %WITH_OUT% --tinyint-as-bool %TINYINT_TO_BOOL% ^
--tag %TAGS% --dao dao --import-models %IMPORT_MODELS%

echo generate go file ok, formatting...
gofmt -w %OUT_DIR%/%PACK_NAME%
)
pause

