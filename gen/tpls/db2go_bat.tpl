@echo off

set OUT_DIR=.
set PACK_NAME="models"
set READ_ONLY=""
set TABLE_NAME=""
set WITH_OUT=""
set TAGS="gorm"
set DSN_URL="mysql://root:123456@127.0.0.1:3306/test?charset=utf8"
set SPEC_TYPES=""

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
db2go.exe --url "%DSN_URL%" --out "%OUT_DIR%" --table "%TABLE_NAME""% --enable-decimal --spec-type "%SPEC_TYPES%" ^
          --package "%PACK_NAME%" --readonly "%READ_ONLY%" --without "%WITH_OUT%" --tag "%TAGS%"
)
pause

