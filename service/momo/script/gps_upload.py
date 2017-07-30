#!/usr/bin/python
# -*- coding:utf-8 -*-

#coding
import json
import sys
import requests



def UploadGpss():
    if len(sys.argv) == 3:
        gpsType = int(sys.argv[2])
    elif len(sys.argv) == 2:
        gpsType = 1
    else:
        print "USEAGE ./gps_upload.py [type](default 1)"
        exit(1)
    gpss = []
    filename = sys.argv[1]
    f = open(filename, "r")
    for line in f.readlines():
        line = line.strip()
        arr = line.split(",")
        province = arr[0].strip()
        city = arr[1].strip()
        longitude = float(arr[2].strip())
        latitude = float(arr[3].strip())
        gpss.append({"longitude":longitude, "latitude":latitude, "province":province, "city":city, "type":gpsType})
    #api = "http://apo.kdzs123.com:8060/gpss"
    api = "http://192.168.1.103:19802/gpss"
    r = requests.post(api, data=json.dumps(gpss))
    print r.text
UploadGpss()

