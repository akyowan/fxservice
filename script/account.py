#!/usr/bin/python

import requests
import json
import traceback
import sys


disables = []
enables = []
fails = []

def GetAccounts(params):
    api = "http://apo.kdzs123.com:8060/accounts"
    try:
        resp = requests.get(api, params=params)
        respJson = json.loads(resp.text)
        if respJson["errno"] != "OK":
            return False
        return respJson["data"]
    except Exception, e:
        traceback.print_exc()
        return False

def CheckAccounts(accounts):
    global disables
    global enables
    global fails
    patchs = []
    try:
        for account in accounts:
            momoAccount = account["momo_account"]
            check = CheckAccount(momoAccount)
            if check == "OK":
                print "%s   OK" % momoAccount
                enables.append(account)
            elif check == "DISABLE":
                print "%s   DISABLE" % momoAccount
                patchs.append(account)
                disables.append(account)
            else:
                print "%s   CHECK_FAIL" % momoAccount
                fails.append(account)
        DisableAccount(patchs)
    except Exception, e:
        traceback.print_exc()
        return False

def CheckAccount(account):
    headers = {
            "Cache-Control": "no-cache",
            "Connection": "Keep-Alive",
            "Content-Type": "application/x-www-form-urlencoded",
            "Accept": "text/html, application/xhtml+xml, */*",
            "Accept-Encoding": "gbk, GB2312",
            "Accept-Language": "zh-cn",
            "Cookie": "SESSIONID=A74B0572-4B1A-7720-BC84-9F4A675CBC06",
            "User-Agent": "Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; WOW64; Trident/5.0)",
            "Content-Length": "0",
            "Host": "api.immomo.com"
            }
    api = "http://api.immomo.com/api/profile/%s" % account
    flag = "0EF05695-FF1A-C903-EEFD-A1CEF78F1160"
    try:
        resp = requests.get(api, headers=headers)
        if resp.status_code != 200:
            print resp.text
            return "FAILED"

        respJson = json.loads(resp.text)
        if not respJson.has_key("photos"):
            print resp.text
            return "FAILED"
        photos = respJson["photos"]
        for key in photos:
            if key == flag:
                return "DISABLE"
        return "OK"
    except Exception, e:
        traceback.print_exc()
        return False

def DisableAccount(accounts):
    api = "http://api.kdzs123.com:8060/accounts"
    try:
        params = []
        for account in accounts:
            params.append({
                    "account":account["account"],
                    "status":4
                    })
        req = requests.patch(api, data=json.dumps(params))
        if req.status_code == 200:
            print "PATCH %d OK" % len(params)
        else:
            print req.text
        return True
    except Exception, e:
        traceback.print_exc()
        return False

def main():
    global enables
    global fails
    global disables
    limit = 20
    params = {
            "status":   2,
            "limit":    limit,
            }
    offset = 0
    resultFile = "./result.json"
    if len(sys.argv) > 1:
        resultFile = sys.argv[1]
    f = open(resultFile, "w")
    while(True):
        params["offset"] = offset
        accounts = GetAccounts(params)
        CheckAccounts(accounts)
        if len(accounts) != limit:
            break
        offset += limit
    result = {
            "disable":disables,
            "enables":enables,
            "check_fails":fails
            }
    json.dump(result, f)
    f.close()

if __name__ == "__main__":
    main()
