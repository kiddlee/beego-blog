#!/bin/bash
rm -rf ../beego_myblog_release
mkdir ../beego_myblog_release
cp beegomyblog ../beego_myblog_release
cp -fr views ../beego_myblog_release
cp -fr static ../beego_myblog_release
cp -fr conf ../beego_myblog_release
cp -fr content ../beego_myblog_release
cd ..
tar -zcvf beego_myblog_release.tgz beego_myblog_release
