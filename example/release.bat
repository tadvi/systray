rsrc -manifest app.manifest -ico=app.ico -o rsrc.syso
go build
REM use this if you do not want command prompt window: go build -ldflags="-H windowsgui"

