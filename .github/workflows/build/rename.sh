#!/bin/bash
cd pkg/filetools/
rm -rf SHX8800
if [ "$1" = "386" ]; then
    mv SHX8800_32 SHX8800
else
    mv SHX8800_64 SHX8800
fi