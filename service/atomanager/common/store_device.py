import MySQLdb
import base64
import datetime
import random
import sys
import traceback
import json
import re
import redis



from distutils.version import StrictVersion
                
FREE_DEVICE_LIST = "FREE_DEVICE_LIST"

hardWareMaps = {
        "iPhone6,1": "s5l8960x",
        "iPhone6,2": "s5l8960x",
        "iPhone7,1": "t7000",
        "iPhone7,2": "t7000",
        "iPhone8,1": "s8000",
        "iPhone8,1": "s8003",
        "iPhone8,2": "s8000",
        "iPhone8,2": "s8003",
        "iPhone8,4": "s8003",
        "iPhone8,4": "s8000",
        "iPhone9,1": "s8010",
        "iPhone9,2": "s8010",
        "iPhone9,3": "s8010",
        "iPhone9,4": "s8010"
        }

buildNums = {
        "10.1"  : "14B72",
        "10.0.3": "14A551",
        "10.0.2": "14A456",
        "10.0.1": "14A403",
        "9.3.5" : "13G36",
        "9.3.4" : "13G35",
        "9.3.3" : "13G34",
        "9.3.2" : "13F69",
        "9.3.1" : "13E238",
        "9.3"   : ["13E237", "13E233"],
        "9.2.1" : "13D15",
        "9.2"   : "13C75",
        "9.1"   : "13B143",
        "9.0.2" : "13A452",
        "9.0.1" : "13A404",
        "9.0  " : "13A344",
        "8.4.1" : "12H321",
        "8.4"   : "12H143",
        "8.3"   : "12F70",
        "8.2"   : "12D508",
        "8.1.3" : "12B466",
        "8.1.2" : "12B440",
        "8.1.1" : "12B435",
        "8.1"   : "12B411",
        "8.0.2" : "12A405",
        "8.0.1" : "12A402",
        "8.0"   : "12A365",
        "7.1.2" : "11D257",
        "7.1.1" : "11D201",
        "7.1"   : "11D167",
        "7.0.6" : "11B651",
        "7.0.5" : "11B601",
        "7.0.4" : "11B554a",
        "7.0.2" : "11A501"
        }


modelMaps = {
       "G07R": "iPhone5,3",
       "G07P": "iPhone5,3",
       "G07T": "iPhone5,3",
       "G07V": "iPhone5,3",
       "G07Q": "iPhone5,3",
       "G5MP": "iPhone7,2",
       "G5MK": "iPhone7,2",
       "G5MH": "iPhone7,2",
       "G5MJ": "iPhone7,2",
       "G5MG": "iPhone7,2",
       "G5ML": "iPhone7,2",
       "G5MM": "iPhone7,2",
       "G5MN": "iPhone7,2",
       "G5MD": "iPhone7,2",
       "G5MC": "iPhone7,2",
       "G5MF": "iPhone7,2",
       "G5MQ": "iPhone7,2",
       "G5MR": "iPhone7,2",
       "G5MT": "iPhone7,2",
       "G5MY": "iPhone7,2",
       "G5MW": "iPhone7,2",
       "G5N0": "iPhone7,2",
       "G5MV": "iPhone7,2",
       "G5QP": "iPhone7,1",
       "G5QR": "iPhone7,1",
       "G5QH": "iPhone7,1",
       "G5QG": "iPhone7,1",
       "G5QT": "iPhone7,1",
       "G5QY": "iPhone7,1",
       "G5QK": "iPhone7,1",
       "G5QV": "iPhone7,1",
       "G5QM": "iPhone7,1",
       "G5QL": "iPhone7,1",
       "G5QJ": "iPhone7,1",
       "G5QN": "iPhone7,1",
       "G5QQ": "iPhone7,1",
       "G5QW": "iPhone7,1",
       "G5QF": "iPhone7,1",
       "G5R0": "iPhone7,1",
       "G5R2": "iPhone7,1",
       "G5R1": "iPhone7,1",
       "GRWH": "iPhone8,2",
       "GRWD": "iPhone8,2",
       "GRWJ": "iPhone8,2",
       "GRWF": "iPhone8,2",
       "GRWL": "iPhone8,2",
       "GRWM": "iPhone8,2",
       "GRX8": "iPhone8,2",
       "GRXL": "iPhone8,2",
       "GRWX": "iPhone8,2",
       "GRWQ": "iPhone8,2",
       "GRWV": "iPhone8,2",
       "GRX5": "iPhone8,2",
       "GRXG": "iPhone8,2",
       "GRX1": "iPhone8,2",
       "GRWT": "iPhone8,2",
       "GRXC": "iPhone8,2",
       "GRXH": "iPhone8,2",
       "GRX4": "iPhone8,2",
       "GRXV": "iPhone8,1",
       "GRX2": "iPhone8,2",
       "GRWY": "iPhone8,2",
       "GRXY": "iPhone8,1",
       "GRXX": "iPhone8,1",
       "GRXW": "iPhone8,1",
       "GRXQ": "iPhone8,1",
       "GRXR": "iPhone8,1",
       "GRWP": "iPhone8,2",
       "GRXD": "iPhone8,2",
       "GRX7": "iPhone8,2",
       "GRXT": "iPhone8,1",
       "GRXK": "iPhone8,2",
       "GRY0": "iPhone8,1",
       "GRY3": "iPhone8,1",
       "GRY5": "iPhone8,1",
       "GRY4": "iPhone8,1",
       "GRY1": "iPhone8,1",
       "GRY7": "iPhone8,1",
       "GRYG": "iPhone8,1",
       "GRY2": "iPhone8,1",
       "GRYF": "iPhone8,1",
       "GRYD": "iPhone8,1",
       "GRYK": "iPhone8,1",
       "GRY8": "iPhone8,1",
       "GRYJ": "iPhone8,1",
       "GRY6": "iPhone8,1",
       "GRYH": "iPhone8,1",
       "GRYC": "iPhone8,1",
       "GRY9": "iPhone8,1",
       "F38Y": "iPhone5,1",
       "F39C": "iPhone5,1",
       "F38W": "iPhone5,1",
       "F39D": "iPhone5,1",
       "F8GL": "iPhone5,1",
       "F8GK": "iPhone5,1",
       "F8GN": "iPhone5,1",
       "F8GH": "iPhone5,1",
       "F8GJ": "iPhone5,1",
       "F8GM": "iPhone5,1",
       "F8H8": "iPhone5,1",
       "F8H4": "iPhone5,1",
       "F8H7": "iPhone5,1",
       "F8H5": "iPhone5,1",
       "F8H2": "iPhone5,1",
       "F8H6": "iPhone5,1",
       "FF9R": "iPhone6,1",
       "FF9Y": "iPhone6,1",
       "FF9V": "iPhone6,1",
       "FFDN": "iPhone6,1",
       "FFDP": "iPhone6,1",
       "FFDR": "iPhone6,1",
       "FFDQ": "iPhone6,1",
       "FFFL": "iPhone6,1",
       "FFFN": "iPhone6,1",
       "FFFK": "iPhone6,1",
       "FFFJ": "iPhone6,1",
       "FFFM": "iPhone6,1",
       "FFFR": "iPhone6,1",
       "FFFP": "iPhone6,1",
       "FFFQ": "iPhone6,1",
       "FFGC": "iPhone6,1",
       "FFG8": "iPhone6,1",
       "FFGJ": "iPhone6,1",
       "FFGD": "iPhone6,1",
       "FFGK": "iPhone6,1",
       "FFGH": "iPhone6,1",
       "FFFT": "iPhone6,1",
       "FFG9": "iPhone6,1",
       "FFFV": "iPhone6,1",
       "FFFW": "iPhone6,1",
       "FFGF": "iPhone6,1",
       "FFGG": "iPhone6,1",
       "FFHP": "iPhone5,3",
       "FFHQ": "iPhone5,3",
       "FFHN": "iPhone5,3",
       "FFHG": "iPhone5,3",
       "FFHR": "iPhone5,3",
       "FFHL": "iPhone5,3",
       "FFHM": "iPhone5,3",
       "FFHJ": "iPhone5,3",
       "FFHK": "iPhone5,3",
       "FFHH": "iPhone5,3",
       "FFT7": "iPhone5,3",
       "FFTM": "iPhone5,3",
       "FFT6": "iPhone5,3",
       "FFT5": "iPhone5,3",
       "FFTN": "iPhone5,3",
       "FH19": "iPhone5,1",
       "FH1H": "iPhone5,1",
       "FH1D": "iPhone5,1",
       "FH1G": "iPhone5,1",
       "FH1F": "iPhone5,1",
       "FH1C": "iPhone5,1",
       "FL05": "iPhone5,3",
       "FL01": "iPhone5,3",
       "FL04": "iPhone5,3",
       "FL02": "iPhone5,3",
       "FL03": "iPhone5,3",
       "FLFW": "iPhone5,3",
       "FLFY": "iPhone5,3",
       "FLFV": "iPhone5,3",
       "FLFP": "iPhone5,3",
       "FLFM": "iPhone5,3",
       "FLFT": "iPhone5,3",
       "FLG0": "iPhone5,3",
       "FLFN": "iPhone5,3",
       "FLFL": "iPhone5,3",
       "FLG2": "iPhone5,3",
       "FM1V": "iPhone5,3",
       "FM1Y": "iPhone5,3",
       "FM21": "iPhone5,3",
       "FM20": "iPhone5,3",
       "FM1R": "iPhone5,3",
       "FM1T": "iPhone5,3",
       "FM1P": "iPhone5,3",
       "FM1W": "iPhone5,3",
       "FM1N": "iPhone5,3",
       "FM1Q": "iPhone5,3",
       "FNDF": "iPhone5,3",
       "FNDJ": "iPhone5,3",
       "FNDG": "iPhone5,3",
       "FNDH": "iPhone5,3",
       "FNDK": "iPhone5,3",
       "FNDD": "iPhone5,3",
       "FNDN": "iPhone5,3",
       "FNDL": "iPhone5,3",
       "FNDM": "iPhone5,3",
       "FNDP": "iPhone5,3",
       "FNJR": "iPhone6,1",
       "FNJJ": "iPhone6,1",
       "FNJQ": "iPhone6,1",
       "FNJN": "iPhone6,1",
       "FNJP": "iPhone6,1",
       "FNJL": "iPhone6,1",
       "FNJT": "iPhone6,1",
       "FNJM": "iPhone6,1",
       "FNJK": "iPhone6,1",
       "FNLW": "iPhone5,3",
       "FNLQ": "iPhone5,3",
       "FNM1": "iPhone5,3",
       "FNM2": "iPhone5,3",
       "FNLV": "iPhone5,3",
       "FNLR": "iPhone5,3",
       "FNM0": "iPhone5,3",
       "FNLT": "iPhone5,3",
       "FNLY": "iPhone5,3",
       "FNM3": "iPhone5,3",
       "FNNM": "iPhone6,1",
       "FNNK": "iPhone6,1",
       "FNNP": "iPhone6,1",
       "FNNL": "iPhone6,1",
       "FNNN": "iPhone6,1",
       "FNNQ": "iPhone6,1",
       "FNNR": "iPhone6,1",
       "FNNV": "iPhone6,1",
       "FNNT": "iPhone6,1",
       "FP6P": "iPhone6,1",
       "FP6Q": "iPhone6,1",
       "FP6H": "iPhone6,1",
       "FP6J": "iPhone6,1",
       "FP6K": "iPhone6,1",
       "FP6M": "iPhone6,1",
       "FP6N": "iPhone6,1",
       "FP6L": "iPhone6,1",
       "FP6R": "iPhone6,1",
       "FQ0Y": "iPhone5,3",
       "FQ13": "iPhone5,3",
       "FQ10": "iPhone5,3",
       "FQ12": "iPhone5,3",
       "FQ15": "iPhone5,3",
       "FQ17": "iPhone5,3",
       "FQ18": "iPhone5,3",
       "FQ14": "iPhone5,3",
       "FQ11": "iPhone5,3",
       "FQ16": "iPhone5,3",
       "FR8G": "iPhone5,3",
       "FR8F": "iPhone5,3",
       "FR9C": "iPhone6,1",
       "FR8H": "iPhone5,3",
       "FR9D": "iPhone6,1",
       "FR97": "iPhone6,1",
       "FR9N": "iPhone6,1",
       "FR92": "iPhone5,3",
       "FR93": "iPhone5,3",
       "FR9M": "iPhone6,1",
       "FR9F": "iPhone6,1",
       "FR8W": "iPhone5,3",
       "FR94": "iPhone5,3",
       "FR98": "iPhone6,1",
       "FR95": "iPhone5,3",
       "FR8J": "iPhone5,3",
       "FR9H": "iPhone6,1",
       "FR9K": "iPhone6,1",
       "FR8Q": "iPhone5,3",
       "FR8Y": "iPhone5,3",
       "FR8M": "iPhone5,3",
       "FR9G": "iPhone6,1",
       "FR9J": "iPhone6,1",
       "FR99": "iPhone6,1",
       "FR8N": "iPhone5,3",
       "FR8P": "iPhone5,3",
       "FR96": "iPhone5,3",
       "FR9L": "iPhone6,1",
       "FR90": "iPhone5,3",
       "FR8T": "iPhone5,3",
       "FR8V": "iPhone5,3",
       "FR9P": "iPhone6,1",
       "FR9Q": "iPhone6,1",
       "FR9R": "iPhone6,1",
       "FR91": "iPhone5,3",
       "FR8R": "iPhone5,3",
       "FR9V": "iPhone6,1",
       "FR9T": "iPhone6,1",
       "FRC4": "iPhone6,1",
       "FRC7": "iPhone6,1",
       "FRCD": "iPhone6,1",
       "FRCC": "iPhone6,1",
       "FRC8": "iPhone6,1",
       "FRC6": "iPhone6,1",
       "FRC5": "iPhone6,1",
       "FRC9": "iPhone6,1",
       "FRCF": "iPhone6,1",
       "H2XG": "iPhone8,4", 
       "H2XH": "iPhone8,4", 
       "H2XJ": "iPhone8,4", 
       "H2XK": "iPhone8,4", 
       "H2XL": "iPhone8,4", 
       "H2XM": "iPhone8,4", 
       "H2XN": "iPhone8,4", 
       "H2XP": "iPhone8,4", 
       "H2XQ": "iPhone8,4", 
       "H2XR": "iPhone8,4", 
       "H2XT": "iPhone8,4", 
       "H2XV": "iPhone8,4", 
       "H2XW": "iPhone8,4", 
       "H2XX": "iPhone8,4", 
       "H2Y7": "iPhone8,4", 
       "H2Y8": "iPhone8,4"
        }

enableVersions = [
        "9.3.4",
        "9.3.3",
        "9.3.5",
        "9.3.2",
        "9.3.1",
        "9.2",
        "9.2.1",
        "9.0.2",
        "9.0.1",
        "9.0"
        ]

enableModels = [
        "iPhone7,2",
        "iPhone8,1",
        "iPhone8,2",
        "iPhone7,1",
        "iPhone6,2",
        "iPhone6,1"
        ]

def connectMysql(host, user, passwd, db, port):
    try:
        conn = MySQLdb.connect(host=host, user=user, passwd=passwd, db=db, port=int(port))
        conn.set_character_set('utf8')
        cur = conn.cursor()
        cur.execute('SET NAMES utf8;') 
        cur.execute('SET CHARACTER SET utf8;')
        cur.execute('SET character_set_connection=utf8;')
        return conn
    except Exception, e:
        print "Connect %s %s %s %s %s error[%r]" % (host, port, user, passwd, db, e)
        traceback.print_exc()
        return False;

def ConnectRedis():
    try:
        r = redis.Redis(host='localhost', port=6379, db=9)
        if r:
            return r
        return True
    except Exception, e:
        print "    CONNECT REDIS FAILED:%r" % e
        traceback.print_exc()
        return False

def getUploadDevice(date, conn):
    try:
        cur = conn.cursor()
        tableName = "upload_%s" % date
        cur.execute("SELECT content FROM %s WHERE action='device_connect' AND content is NOT NULL" % tableName)
        res = cur.fetchall()
        deviceInfos = []
        for ins in res:
            content = base64.decodestring(ins[0])
            deviceInfos.append(content.split("||"))
        return deviceInfos
    except Exception, e:
        print "Get device information failed[%r]" % e
        traceback.print_exc()
        return False

def getLdsDevice(date, conn):
    try:
        cur = conn.cursor()
        tableName = "upload_%s" % date
        cur.execute("SELECT DISTINCT(sn) FROM %s WHERE action='pub_authen'" % tableName)
        res = cur.fetchall()
        return res
    except Exception, e:
        print "Get lds device information[%r] failed" % e
        traceback.print_exc()
        return False

def getAccessModel(date, conn):
    try:
        cur = conn.cursor()
        tableName = "access_%s" % date
        cur.execute("SELECT device_id, model FROM %s WHERE model>1 AND device_id!='UN_KOWN' GROUP BY device_id" % tableName)
        res = cur.fetchall()
        models = []
        for ins in res:
            models.append(ins)
        return models
    except Exception, e:
        print "Get models failed[%r]" % e
        traceback.print_exc()
        return False

def updateModel(sn, model, conn):
    try:
        cur = conn.cursor()
        cur.execute("UPDATE apo_device_info SET model=%s WHERE sn=%s" , (model, sn))
        count = cur.rowcount
        cur.close()
        conn.commit()
        return count 
    except Exception, e:
        print "UPDATE sn[%s] model[%s] failed %r" % (sn, model, e)
        traceback.print_exc()
        return False


def getModel(model):
    maps = {
            0x00010001:"iPhone3,1",
            0x00010002:"iPhone4,1",
            0x00010003:"iPhone5,1",
            0x00010004:"iPhone6,1",
            0x00010005:"iPhone5,3",
            0x00010006:"iPhone7,2",
            0x00010007:"iPhone8,1",
            0x00010008:"iPhone7,1",
            0x00010009:"iPhone8,2",
            0x0001000A:"iPhone8,4"
    }
    if model in maps:
        return maps[model]
    return None


def storeDeviceInfo(deviceInfo, conn):
    global modelMaps
    global hardWareMaps
    global buildNums
    global enableVersions
    global enableModels
    try:
        conn.autocommit(False)
        cur = conn.cursor()
        for info in deviceInfo:
            imei = info[0]
            sn = info[1]
            seq = info[2]
            version = info[3]
            mac = info[4]
            wifi = info[5]
            model = 0
            hardWare = 0
            buildNum = 0
            updateFlag = True
            if (sn == ""):
                continue
            if (imei == ""):
                imei = " "
                updateFlag = False
            if (seq == ""):
                seq = " "
                updateFlag = False
            if (version == "" or StrictVersion(version) < StrictVersion('9.0')):
                version = enableVersions[random.randint(0, len(enableVersions)-1)]
            if (mac == ""):
                mac = " "
                updateFlag = False
            if (wifi == ""):
                wifi = " "
                updateFlag = False
            if (seq[-4:] in modelMaps):
                model = modelMaps[seq[-4:]]
            else:
                model = enableModels[random.randint(0, len(enableModels)-1)]

            if (StrictVersion(model[6:].replace(',', '.')) < StrictVersion(model[6:].replace(',', '.'))):
                model = enableModels[random.randint(0, len(enableModels)-1)]

            if (model in hardWareMaps):
                hardWare = hardWareMaps[model]

            if (version in buildNums):
                if (isinstance(buildNums[version], list)):
                    buildNum = buildNums[version][random.randint(0, len(buildNums[version])-1)]
                else:
                    buildNum = buildNums[version]
            if len(sn) != 40:
                break

            dType = 0
            if len(imei) > 3  and len(seq) > 3:
                dType = 1
            if updateFlag:
                cur.execute("INSERT INTO apo_device_info(imei, sn, seq, version, mac, wifi, model, build_num, hard_ware, type) \
                        VALUES(%s, %s, %s, %s, %s, %s, %s, %s, %s, %s) \
                        ON DUPLICATE KEY UPDATE \
                        wifi=%s, imei=%s, seq=%s, version=%s, mac=%s, model=%s, hard_ware=%s, build_num=%s, type=%s", \
                        (imei, sn, seq, version, mac, wifi, model, buildNum, hardWare, dType, 
                         wifi, imei, seq, version, mac, model, hardWare, buildNum, dType))
            else:
                cur.execute("INSERT INTO apo_device_info(imei, sn, seq, version, mac, wifi, model, build_num, hard_ware, type) \
                        VALUES(%s, %s, %s, %s, %s, %s, %s, %s, %s, %s) ON DUPLICATE KEY UPDATE wifi=%s", \
                        (imei, sn, seq, version, mac, wifi, model, buildNum, hardWare, dType, wifi))
        cur.close()
        conn.commit()
    except Exception, e:
        print "Store device info faield[%r]" % e
        traceback.print_exc()
        return False

def storeLdsDeviceInfo(devices, conn):
    try:
        conn.autocommit(False)
        cur = conn.cursor()
        for info in devices:
            sn = info[0]
            if not sn:
                continue
            cur.execute("INSERT INTO apo_device_info(sn) VALUES (%s) ON DUPLICATE KEY UPDATE sn=%s" , (sn, sn))
        cur.close()
        conn.commit()
    except Exception, e:
        print "Store device info faield[%r]" % e
        traceback.print_exc()
        return False


def updateBuildNum(conn):

    global hardWareMaps
    global buildNums
    try:
        cur = conn.cursor()
        for k in hardWareMaps:
            cur.execute("UPDATE apo_device_info SET hard_ware='%s' WHERE version='%s'" % (hardWareMaps[k], k)) 
            count = cur.rowcount
            print "UPDATE version[%s] hardWare[%s] count[%d]\n" % (k, hardWareMaps[k], count)
            conn.commit()

        for k in buildNums:
            if (isinstance(buildNums[k], list)):
                v = buildNums[k][random.randint(0, len(buildNums[k])-1)]
            else:
                v = buildNums[k]
            cur.execute("UPDATE apo_device_info SET build_num='%s' WHERE version='%s'" % (v, k))
            count = cur.rowcount
            conn.commit()
            print "UPDATE version[%s] buildNum[%s] count[%d]\n" % (k, v, count)
        cur.close()
        return True
    except Exception, e:
        print "UPDATE hardWare and buildNum failed %r" % e
        traceback.print_exc()
        return False


def getDeviceInfo(sn, conn):
    try:
        cur = conn.cursor(cursorclass=MySQLdb.cursors.DictCursor)
        cur.execute("SELECT * FROM apo_device_info WHERE sn='%s' " % sn);
        if cur.rowcount <= 0:
            return False
        res = cur.fetchall()[0]
        return res
    except Exception, e:
        print "GET DEVICE INFO[%s] FAILED:%r" % (sn, e)
        traceback.print_exc()
        return False

def bindDeviceInfo(db, redis):
    try:
        device = redis.lpop(FREE_DEVICE_LIST)
        if device:
            return json.loads(device)
        cur = db.cursor(cursorclass=MySQLdb.cursors.DictCursor)
        cur.execute("SELECT * FROM apo_device_info WHERE signin_count=0 AND bind_count=0 ORDER BY signin_count ASC LIMIT 5000")
        db.autocommit(False)
        for device in cur.fetchall():
            tid = device["id"]
            cur.execute("UPDATE apo_device_info SET bind_count=bind_count+1 WHERE id=%d" % tid)
            value = json.dumps(device)
            redis.lpush(FREE_DEVICE_LIST, value)
        db.commit()
        print "GET FREE DEVICE OK"
        device = redis.lpop(FREE_DEVICE_LIST)
        if not device:
            return False
        return json.loads(device)
    except Exception,e:
        print "GET FREE DEVICE FAILED"
        traceback.print_exc()
        return False

def storeAccount(conn, redis, fileName, weight):
    try:
        cur = conn.cursor()
        fileObj = open(fileName, "r")

        rex = re.compile(r'.*\/(.*?)\..*?$')
        matches = re.match(rex, fileName)
        if not matches:
            print "INVALID FILENAME FORMAT [%s]" % fileName
            return False
        brief = matches.group(1)
        print brief
        startId = getAccountGroup(brief, conn)
        if not startId:
            return False
        count = 0
        conn.autocommit(False)
        for line in fileObj:
            line = line.replace('\r','').replace('\n','')
            startId = startId+1
            info = line.split("\t")
            account = info[0]
            passwd = info[1]
            status = 1
            ip = getAccountRandIp()
            dinfo= bindDeviceInfo(conn, redis)
            if not dinfo:
                print "BIND DEVICE INFO FAILED"
                continue
            sn = dinfo['sn']
            cur.execute("INSERT INTO apo_account_info(\
                    id, account, passwd, brief, ip, imei, sn, seq, version, mac, wifi, model, build_num, hard_ware, status) \
                    VALUES(%d, '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', %d) \
                    ON DUPLICATE KEY UPDATE passwd='%s', brief='%s'" % (
                        startId, account, passwd, brief, ip, dinfo['imei'], sn, dinfo['seq'], dinfo['version'], dinfo['mac'], \
                                dinfo['wifi'],dinfo['model'], dinfo['build_num'], dinfo['hard_ware'], status, passwd, brief))
            #cur.execute("UPDATE apo_device_info set bind_count=100 WHERE sn='%s'" % sn)
            count += 1
        conn.commit()
        storeAccountGroup(conn, brief, count, int(weight))
        flushRedisBriefsCache()
        print "STORE %s %s ACCOUNT SUCCESS" % (fileName, count)
        return True
    except Exception, e:
        print "Store account failed %r" % e
        traceback.print_exc()
        return False;

def storeAccountGroup(conn, brief, total, weight):
    try:
        cur = conn.cursor()
        cur.execute("INSERT INTO apo_account_groups(brief, weight, total) VALUES('%s', %d, %d)" % (brief, weight, total))
        conn.commit()
        return True
    except Exception, e:
        print "store %s accout brief failed %r" % (brief, e)
        traceback.print_exc()
        return False;

def getAccountRandIp():
    ips = ""
    for i in range(1, 5):
        ip = "192.168.1.%d|" % random.randint(2, 254)
        ips += ip
    ip = "10.%d.%d.%d" % (random.randint(2, 254), random.randint(2, 254),random.randint(2, 254))
    ips += ip
    return ips

def getAccountGroup(brief, conn):
    try:
        cur = conn.cursor(cursorclass=MySQLdb.cursors.DictCursor)
        cur.execute("SELECT MAX(id) as max FROM apo_account_info WHERE brief='%s'" % brief)
        res = cur.fetchall()[0]
        print res
        if res['max']:
            return res['max']
        cur.execute("SELECT MAX(id) as max FROM apo_account_info")
        res = cur.fetchall()[0]
        if not res['max']:
            return 1000100000
        id = res['max']
        id = ((id/100000) + 1) * 100000
        return id
    except Exception, e:
        print "Get %s group failed %r" % (brief, e)
        traceback.print_exc()
        return False
        
def range_date(start, end):
    for n in range(int((end-start).days)):
        yield start+datetime.timedelta(n)

def flushRedisBriefsCache():
    try:
        redis = ConnectRedis()
        if not redis:
            print "CONNECT REDIS FAILED"
            return False
        redis.delete('ACCOUNT_ALL_BRIEFS')
        return True
    except Exception, e:
        print "FLUSH REDIS BRIEF CACHE FAILED"
        print traceback.print_exc()
        return False

def bindAccount(db):
    try:
        cur = db.cursor()
        count = 0
        while(1):
            cur.execute("SELECT id, sn FROM apo_account_info WHERE type=0 AND status=1 LIMIT 1000")
            res = cur.fetchall()
            if cur.rowcount == 0:
                print "BIND ACCOUNT END"
                break
            for info in res:
                accountID = info[0]
                sn = info[1]
                cur.execute("SELECT bind_count FROM apo_device_info WHERE sn='%s'" % sn)
                if (cur.rowcount == 1) and (cur.fetchall()[0][0] == 0):
                    cur.execute("UPDATE apo_device_info SET bind_count=1 WHERE sn='%s'" % sn)
                    cur.execute("UPDATE apo_account_info SET type=1 WHERE id=%d" % accountID)
                    db.commit()
                else:
                    cur.execute("SELECT id, imei, sn, seq, version, mac, wifi, model, build_num, hard_ware FROM apo_device_info WHERE bind_count=0 ORDER BY id DESC LIMIT 1")
                    deviceInfo = cur.fetchall()[0]
                    deviceID = deviceInfo[0]
                    newSN = deviceInfo[2]
                    imei = deviceInfo[1]
                    if not imei:
                        imei = " "
                    seq = deviceInfo[3]
                    if not seq:
                        seq = " "
                    version = deviceInfo[4]
                    if not version:
                        seq = " "
                    mac = deviceInfo[5]
                    if not mac:
                        mac = " "
                    wifi = deviceInfo[6]
                    if not wifi:
                        wifi = " "
                    model = deviceInfo[7]
                    if not model:
                        model = " "
                    build_num = deviceInfo[8]
                    if not build_num:
                        build_num = " "
                    hard_ware = deviceInfo[9]
                    if not hard_ware:
                        hard_ware = " "
                    cur.execute("UPDATE apo_device_info SET bind_count=1 WHERE id=%d" % deviceID)
                    cur.execute("UPDATE apo_account_info SET \
                            sn='%s',\
                            imei='%s',\
                            seq='%s',\
                            version='%s',\
                            mac='%s',\
                            wifi='%s',\
                            model='%s',\
                            build_num='%s',\
                            hard_ware='%s',\
                            type=1\
                            WHERE id=%d" % (\
                            newSN,\
                            imei,\
                            seq,\
                            version,\
                            mac,\
                            wifi,\
                            model,\
                            build_num,\
                            hard_ware,\
                            accountID))
                    db.commit()
            print "BIND ACCOUNT SUCCESS %d" % count 
            sys.stdout.flush()
            count = count + 1000
    except Exception, e:
        print "Bind account failed %r" % e
        traceback.print_exc()

uploadDb = connectMysql("123.57.191.144", "koudai", "GoodLuck1956", "koudai_upload", 3306)
if not uploadDb:
    print "CONNECT DB[koudai_upload] FAILED"
    exit(1)

storeDb = connectMysql("rdshqyyhbl530971fle8.mysql.rds.aliyuncs.com", "wanliping", "wans198059", "koudai_apo", 3306)
#storeDb = connectMysql("192.168.2.152", "wans", "koudai123456", "koudai_apo", 3306)
#storeDb = connectMysql("192.168.1.103", "wanliping", "wans198059", "koudai_apo", 3306)
if not storeDb:
    print "CONNECT DB[koudai_apo] FAILED"
    exit(1)

r = ConnectRedis()
if not r:
    print "CONNECT REDIS ERROR"
    exit(0)

if (sys.argv[1] == "DEVICE"):
    start = sys.argv[2]
    end = sys.argv[3]
    if (not start):
        print "PLEASE INPUT START DATE LIKE[0000-00-00]"
        exit(1)
    if (not end):
        end = start
    start = start.split('-')
    end = end.split('-')

    start = datetime.datetime(int(start[0]), int(start[1]), int(start[2]), 0, 0, 0)
    end = datetime.datetime(int(end[0]), int(end[1]), int(end[2]), 0, 0, 0)
    for i in range_date(start, end):
        date = i.strftime("%Y%m%d")
        print "START    %s" % date
        deviceInfo = getUploadDevice(date, uploadDb)
        if not deviceInfo:
            continue
        print "GET      %s DEVICE SUCCESS %d" % (date, len(deviceInfo))
        storeDeviceInfo(deviceInfo, storeDb)
        print "STORE    %s DEVICE SUCCESS" % date
elif (sys.argv[1] == "MODEL"):
    start = sys.argv[2]
    end = sys.argv[3]
    if (not start):
        print "PLEASE INPUT START DATE LIKE[0000-00-00]"
        exit(1)
    if (not end):
        end = start
    start = start.split('-')
    end = end.split('-')
    start = datetime.datetime(start[0], start[1], start[2], 0, 0, 0)
    end = datetime.datetime(end[0], end[0], end[0], 0, 0, 0)
    for i in range_date(start, end):
        date = i.strftime("%Y%m%d")
        modelsInfo = getAccessModel(date, uploadDb)
        if not modelsInfo:
            continue
        print "GET %s MODEL SUCCESS" % date
        for info in modelsInfo:
            model = getModel(info[1])
            if model:
                count = updateModel(info[0], info[1], storeDb)
                print info[0], model, count
elif (sys.argv[1] == "HARD"):
    updateBuildNum(storeDb)
elif (sys.argv[1] == "ACCOUNT"):
    storeAccount(storeDb, r, sys.argv[2], sys.argv[3])
elif (sys.argv[1] == "LDS_DEVICE"):
    start = sys.argv[2]
    end = sys.argv[3]
    if (not start):
        print "PLEASE INPUT START DATE LIKE[0000-00-00]"
        exit(1)
    if (not end):
        end = start
    start = start.split('-')
    end = end.split('-')
   
    start = datetime.datetime(int(start[0]), int(start[1]), int(start[2]), 0, 0, 0)
    end = datetime.datetime(int(end[0]), int(end[1]), int(end[2]), 0, 0, 0)
    for i in range_date(start, end):
        date = i.strftime("%Y%m%d")
        print "STARTLDS %s" % date
        devices = getLdsDevice(date, uploadDb)
        if not devices:
            continue
        print "GET      %s DEVICE SUCCESS %d" % (date, len(devices))
        storeLdsDeviceInfo(devices, storeDb)
        print "STORE    %s DEVICE SUCCESS" % date
elif (sys.argv[1] == "BIND_ACCOUNT"):
    bindAccount(storeDb)
else:
    print "NO ACTION"
