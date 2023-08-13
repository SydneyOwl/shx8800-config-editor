@echo off
cd lib
dotnet publish -c Release -r win-x86 /p:PublishSingleFile=true
move bin\Release\net7.0\win-x86\publish\SHX8800.exe ..\pkg\filetools\SHX8800_32
dotnet publish -c Release -r win-x64 /p:PublishSingleFile=true
move bin\Release\net7.0\win-x64\publish\SHX8800.exe ..\pkg\filetools\SHX8800_64
rmdir /S /Q bin obj