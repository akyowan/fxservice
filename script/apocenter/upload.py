#!/usr/bin/python
#-*- coding:UTF-8 -*-

import requests
import csv
import sys
import json
import argparse
import re

class Uploader:
    def __init__(self, domain, appKey):
        self.device_api = "http://%s.71741.com/upload/devices" % domain
        self.account_api = "http://%s.71741.com/upload/accounts" % domain
        self.appKey = appKey

    def post(self, url, data):
        headers = {"AppKey":self.appKey}
        resp = requests.post(url, json=data, headers=headers)
        return resp

    def uploadAccount(self, brief, accounts):
        url = "%s/%s" % (self.account_api, brief)
        return self.post(url, accounts)

    def uploadDevice(self, devices):
        return self.post(self.device_api, devices)
    
def UploadAccount(uploader, brief, path):
    fp = open(path, "r")
    accounts = []
    for line in fp:
        line = line.replace('\r','').replace('\n','')
        info = line.split(" ")
        if len(info) < 2:
            continue
        accounts.append({
                "account":info[0],
                "passwd":info[(len(info)-1)]
                })
        if len(accounts) >= 1000:
            res = uploader.uploadAccount(brief, accounts)
            if res.status_code != 200:
                print res.text
                return False
            else:
                r = res.json()
                print "upload account success:%d errors:%d exist:%d" % (r["data"]["success"], len(r["data"]["errors"]), len(r["data"]["exists"]))
            accounts = []
    if len(accounts) > 0:
        res = uploader.uploadAccount(brief, accounts)
        if res.status_code != 200:
            print res.text
            return False
        else:
            r = res.json()
            print "upload account success:%d errors:%d exist:%d" % (r["data"]["success"], len(r["data"]["errors"]), len(r["data"]["exists"]))

def UploadDevice(uploader, group, path):
    fp = open(path, "r")
    reader = csv.DictReader(fp)
    devices = []
    for row in reader:
        device = {
                "sn"    :row["sn"],
                "group" :group
                }
        if row.has_key("imei"):
            device["imei"] = row["imei"]
        if row.has_key("seq"):
            device["seq"] = row["seq"]
        if row.has_key("model"):
            device["model"] = row["model"]
        if row.has_key("model_num"):
            device["model_num"] = row["model_num"]
        if row.has_key("wifi"):
            device["wifi"] = row["wifi"]
        if row.has_key("build_num"):
            device["build_num"] = row["build_num"]
        if row.has_key("hard_ware"):
            device["hard_ware"] = row["hard_ware"]
        if row.has_key("hardware_model"):
            device["hardware_model"] = row["hardware_model"]
        if row.has_key("ecid"):
            device["ecid"] = row["ecid"]
        if row.has_key("region"):
            device["region"] = row["region"]
        if row.has_key("firmware"):
            device["firmware"] = row["firmware"]
        if row.has_key("mlb_seq"):
            device["mlb_seq"] = row["mlb_seq"]
        if row.has_key("baseband_version"):
            device["baseband_version"] = row["baseband_version"]
        devices.append(device)
        if len(devices) >= 1000:
            res = uploader.uploadDevice(devices)
            if res.status_code != 200:
                print res.text
                return False
            else:
                r = res.json()
                print "upload device success:%d errors:%d exists:%d" % (r["data"]["success"], len(r["data"]["errors"]), len(r["data"]["exists"]))
            devices = []

    if len(devices) > 0:
        res = uploader.uploadDevice(devices)
        if res.status_code != 200:
            print res.text
            return False
        else:
            r = res.json()
            print "upload device success:%d errors:%d exists:%d" % (r["data"]["success"], len(r["data"]["errors"]), len(r["data"]["exists"]))

def getGroup(path):
    path = "/%s" % path
    rex = re.compile(r'.*\/(.*?)\..*?$')
    matches = re.match(rex, path)
    if not matches:
        return None
    brief = matches.group(1)
    return brief

def main():
    parser = argparse.ArgumentParser()
    parser.add_argument("domain", choices=["aso", "ato"], help="service domain ato/aso")
    parser.add_argument("type", choices=["account", "device"], help="upload type account or device")
    parser.add_argument("input", help="upload data file          account[txt split by space character] device[csv]")
    parser.add_argument("-g", "--group", help="account group, default it set use input file name")
    args = parser.parse_args()
    if args.group == None:
        args.group = getGroup(args.input)
    if args.group == None:
        print "invalid input file"
        return 1
    if args.domain == "aso":
        appKey = "d51429b872ad07b936bd85ea605d13a4"
    else:
        appKey = "2b525950fdc56f96c66ec919fad15d39"
    uploader = Uploader(args.domain, appKey)
    if args.type == "account":
        UploadAccount(uploader, args.group, args.input)
    else:
        UploadDevice(uploader, args.group, args.input)
main()
