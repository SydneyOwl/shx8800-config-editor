@echo off
FOR /F %%v IN ('git describe --tags') DO SET BUILDVERSION=%%v
FOR /F %%c IN ('git rev-parse HEAD') DO SET COMMIT=%%c
SET BUILDTIME=%DATE%T%TIME%

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
go build -tags="dev" -ldflags "-X 'github.com/sydneyowl/shx8800-config-editor/config.VER=%BUILDVERSION%' -X 'github.com/sydneyowl/shx8800-config-editor/config.COMMIT=%COMMIT%' -X 'github.com/sydneyowl/shx8800-config-editor/config.BUILDTIME=%BUILDTIME%'"
move /Y shx8800-config-editor.exe tmp\shx8800-config-editor-dev_windows_386.exe 2>nul
copy /Y tmp\SHX8800_x86.exe pkg\filetools\SHX8800

go build -tags="release" -ldflags "-X 'github.com/sydneyowl/shx8800-config-editor/config.VER=%BUILDVERSION%' -X 'github.com/sydneyowl/shx8800-config-editor/config.COMMIT=%COMMIT%' -X 'github.com/sydneyowl/shx8800-config-editor/config.BUILDTIME=%BUILDTIME%'"
move /Y shx8800-config-editor.exe tmp\shx8800-config-editor-release_windows_386.exe 2>nul
del pkg\filetools\SHX8800

REM 64
set GOARCH=amd64
go build -tags="dev" -ldflags "-X 'github.com/sydneyowl/shx8800-config-editor/config.VER=%BUILDVERSION%' -X 'github.com/sydneyowl/shx8800-config-editor/config.COMMIT=%COMMIT%' -X 'github.com/sydneyowl/shx8800-config-editor/config.BUILDTIME=%BUILDTIME%'"
move /Y shx8800-config-editor.exe tmp\shx8800-config-editor-dev_windows_amd64.exe 2>nul
copy /Y tmp\SHX8800_x64.exe pkg\filetools\SHX8800

go build -tags="release" -ldflags "-X 'github.com/sydneyowl/shx8800-config-editor/config.VER=%BUILDVERSION%' -X 'github.com/sydneyowl/shx8800-config-editor/config.COMMIT=%COMMIT%' -X 'github.com/sydneyowl/shx8800-config-editor/config.BUILDTIME=%BUILDTIME%'"
move /Y shx8800-config-editor.exe tmp\shx8800-config-editor-release_windows_amd64.exe 2>nul