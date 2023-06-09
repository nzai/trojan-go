@echo off

REM 获取当前目录
set CURRENT_DIR=%CD%

REM 注册服务
sc create tg binPath= "%CURRENT_DIR%\winsvc.exe -c %CURRENT_DIR%\config.json" start= auto

REM 启动服务
sc start tg