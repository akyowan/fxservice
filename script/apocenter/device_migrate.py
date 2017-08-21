#!/usr/bin/python
# -*- coding: UTF-8 -*-

from pymongo import MongoClient
import json

client = MomgoClient('localhost', 27017)
db = client.apo_storage
col = db.devices

source_file = open("devices.json", 'r')
devices = json.load(source_file)

for device in devices:
    print device

