#!/bin/bash
rm -rf ../beego_myblog_release
mkdir ../beego_myblog_release
cp beego-myblog ../beego_myblog_release
cp -fr views ../beego_myblog_release
cp -fr static ../beego_myblog_release
cp -fr conf ../beego_myblog_release
tar -zcvf ../beego_myblog_release.tgz ../beego_myblog_release
