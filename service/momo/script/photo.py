#!/usr/bin/python2.7

import os
import sys
import os.path
import json
import requests
files = []
for parent,dirnames,filenames in os.walk(sys.argv[1]):
    for filename in filenames:
        ff = "http://archive.kdzs123.com/momo/%s%s" % (parent, filename)
        photos = []
        photos.append({"seq":1, "url":ff})
        files.append(photos)

api = "http://apo.kdzs123.com:8060/photos"
#api = "http://localhost:19802/photos"
r = requests.post(api, json=files)
print r.text
