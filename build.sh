#! /bin/bash

# Build web UI
cd ~/work/src/github.com/interesting1113/video_server/web
go install
cp ~/work/bin/web ~/work/bin/video_server_web_ui/web
cp -R ~/work/src/github.com/interesting1113/video_server/templates ~/work/bin/video_server_web_ui/