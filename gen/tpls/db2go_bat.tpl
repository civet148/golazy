@echo off

set OUT_DIR=.
set PACK_NAME="models"
set SUFFIX_NAME=""
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

rem 判断本地系统是否已安装db2go工具，没有则进行安装
where db2go.exe

IF "%errorlevel%" == "0" (
    echo db2go already installed.
) ELSE (
    echo db2go not found in system %%PATH%%, installing...
    go install github.com/civet148/db2go@latest
    If "%errorlevel%" == "0" (
        echo db2go install succeeded
    ) ELSE (
        rem 安装失败，Linux/Mac请安装gcc工具链，Windows系统可以通过链接直接下载二进制(https://github.com/civet148/release/tree/master/db2go/v3)
        echo error: db2go install failed, Linux/Mac please install gcc tool-chain, Windows download from https://github.com/civet148/release/tree/master/db2go/v3
    )
)

If "%errorlevel%" == "0" (
db2go.exe --url "%DSN_URL%" --out "%OUT_DIR%" --table "%TABLE_NAME""% --json-properties "%JSON_PROPERTIES%" --enable-decimal  --spec-type "%SPEC_TYPES%" ^
--suffix "%SUFFIX_NAME%" --package "%PACK_NAME%" --readonly "%READ_ONLY%" --without "%WITH_OUT%" --tinyint-as-bool "%TINYINT_TO_BOOL%" ^
--tag "%TAGS%" --import-models "%IMPORT_MODELS%"

echo generate go file ok, formatting...
gofmt -w %OUT_DIR%/%PACK_NAME%
db2go.exe -v
)
pause

