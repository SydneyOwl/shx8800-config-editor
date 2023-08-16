@echo off
md tmp
cd lib
dotnet publish -c Release -r win-x86 /p:PublishSingleFile=true
move /Y bin\Release\net7.0\win-x86\publish\SHX8800.exe ..\tmp\SHX8800_x86.exe
dotnet publish -c Release -r win-x64 /p:PublishSingleFile=true
move bin\Release\net7.0\win-x64\publish\SHX8800.exe ..\tmp\SHX8800_x64.exe
rmdir /S /Q bin obj
cd ..
go mod tidy

REM 386
set GOARCH=386
go build -tags="dev"
@REM if "%1" neq "refs/heads/master" (
@REM     move /Y shx8800-config-editor.exe tmp\shx8800-config-editor-dev_%1_windows_386.exe
@REM )
copy /Y tmp\SHX8800_x86.exe pkg\filetools\SHX8800

go build -tags="release"
@REM if "%1" neq "refs/heads/master" (
@REM     move /Y shx8800-config-editor.exe tmp\shx8800-config-editor-release_%1_windows_386.exe
@REM )
del pkg\filetools\SHX8800

REM 64
set GOARCH=amd64
go build -tags="dev"
@REM if "%1" neq "refs/heads/master" (
@REM     move /Y shx8800-config-editor.exe tmp\shx8800-config-editor-dev_%1_windows_amd64.exe
@REM )
copy /Y tmp\SHX8800_x64.exe pkg\filetools\SHX8800

go build -tags="release"
@REM if "%1" neq "refs/heads/master" (
@REM     move /Y shx8800-config-editor.exe tmp\shx8800-config-editor-release_%1_windows_amd64.exe
@REM )