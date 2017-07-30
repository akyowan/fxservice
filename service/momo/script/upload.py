#!/usr/bin/python

import json
import sys
import requests

def UploadAccounts(filename):
    accounts = []
    f = open(filename, "r")
    for line in f.readlines():
        line = line.strip()
        arr = line.split("----")
        accounts.append({"account":arr[0], "password":arr[1]})
    api = "http://apo.kdzs123.com:8060/accounts"
    r = requests.post(api, data=json.dumps(accounts))
    print r.text

UploadAccounts(sys.argv[1])
