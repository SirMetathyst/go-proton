@echo off
pushd %~dp0
go-proton.exe --WatchFileEnable=true --WatchFile=./proton.proton --File=./proton.proton --PrintFilePostProcessor_C_1_4_2
popd
pause