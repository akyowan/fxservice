#!/usr/bin/python2.7

import os
import sys
import os.path
import json
files = []
for parent,dirnames,filenames in os.walk(sys.argv[1]):
    for d in dirnames:
        for p,dd,fs in os.walk(os.path.join(parent,d)):
            photos = []
            seq = 1
            for filename in fs:
                fname = "http://archive.kdzs123.com/momo/%s" % (os.path.join(p,filename)[2:])
                photos.append({"seq":seq, "url":fname})
                seq += 1
            files.append(photos)
print json.dumps(files)
