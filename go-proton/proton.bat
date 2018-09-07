@echo off
pushd %~dp0
proton.exe --WatchFileEnable=true --WatchFile=proton.proton
popd
pause