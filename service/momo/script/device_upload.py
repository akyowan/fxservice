#!/usr/bin/python

import json
import sys
import requests

def uploadDevices(filename):
    devices = json.load(open(filename))
    api = "http://apo.kdzs123.com:8060/devices"
    r = requests.post(api, json=devices)
    print r.text

uploadDevices(sys.argv[1])
