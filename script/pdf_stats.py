#!/bin/python

# -*- coding:utf-8 -*-

import time
import redis
import traceback
from pymongo import MongoClient
from datetime import datetime

class PDFStats:
    def __init__(self):
        conn = MongoClient("storage", 27017)
        self.db = conn.reports
        now = time.time()
        self.start = time.localtime(now)
        self.end = time.localtime(now - (3600*24))
        self.redis = redis.Redis(host="storage", port=6379, db=12)

    def InstallStart(self):
        try:
            date = time.strftime('%Y-%m-%d',self.start)
            query =  {
                "action":"install",
                "status":"start",
                "log_time":{"$gte":datetime(self.start.tm_year, self.start.tm_mon, self.start.tm_mday)},
                "log_time":{"$lt":datetime(self.end.tm_year, self.end.tm_mon, self.end.tm_mday)}
            }
            report = self.db.pdf_report
            ids = report.distinct("device_id", query)
            count = len(ids)
            self.redis.hset("install_start", date, count)
            print "%s %d install" % (date, count)

            ids = report.distinct("device_id", {"action":"install"})
            count = len(ids)
            self.redis.hset("install_start", "total", count)

        except Exception, e:
            traceback.print_exc()
            return False

    def InstallEnd(self):
        try:
            date = time.strftime('%Y-%m-%d',self.start)
            query =  {
                "action":"install",
                "status":"end",
                "log_time":{"$gte":datetime(self.start.tm_year, self.start.tm_mon, self.start.tm_mday)},
                "log_time":{"$lt":datetime(self.end.tm_year, self.end.tm_mon, self.end.tm_mday)}
            }
            report = self.db.pdf_report
            ids = report.distinct("device_id", query)
            count = len(ids)
            self.redis.hset("install_end", date, count)
            print "%s %d install" % (date, count)

            ids = report.distinct("device_id", {"action":"install"})
            count = len(ids)
            self.redis.hset("install_end", "total", count)

        except Exception, e:
            traceback.print_exc()
            return False


    def UninstallStart(self):
        try:
            date = time.strftime('%Y-%m-%d',self.start)
            query =  {
                "action":"uninstall",
                "status":"start",
                "log_time":{"$gte":datetime(self.start.tm_year, self.start.tm_mon, self.start.tm_mday)},
                "log_time":{"$lt":datetime(self.end.tm_year, self.end.tm_mon, self.end.tm_mday)}
            }
            report = self.db.pdf_report
            ids = report.distinct("device_id", query)
            count = len(ids)
            self.redis.hset("uninstall_start", date, count)
            print "%s %d uninstall" % (date, count)

            ids = report.distinct("device_id", {"action":"uninstall"})
            count = len(ids)
            self.redis.hset("uninstall_start", "total", count)
        except Exception, e:
            traceback.print_exc()
            return False

    def UninstallEnd(self):
        try:
            date = time.strftime('%Y-%m-%d',self.start)
            query =  {
                "action":"uninstall",
                "status":"end",
                "log_time":{"$gte":datetime(self.start.tm_year, self.start.tm_mon, self.start.tm_mday)},
                "log_time":{"$lt":datetime(self.end.tm_year, self.end.tm_mon, self.end.tm_mday)}
            }
            report = self.db.pdf_report
            ids = report.distinct("device_id", query)
            count = len(ids)
            self.redis.hset("uninstall_end", date, count)
            print "%s %d uninstall" % (date, count)

            ids = report.distinct("device_id", {"action":"uninstall"})
            count = len(ids)
            self.redis.hset("uninstall_end", "total", count)
        except Exception, e:
            traceback.print_exc()
            return False

    def MFShow(self):
        try:
            date = time.strftime('%Y-%m-%d',self.start)
            query =  {
                "action":"mf_show",
                "log_time":{"$gte":datetime(self.start.tm_year, self.start.tm_mon, self.start.tm_mday)},
                "log_time":{"$lt":datetime(self.end.tm_year, self.end.tm_mon, self.end.tm_mday)}
            }
            report = self.db.pdf_report
            ids = report.distinct("device_id", query)
            count = len(ids)
            self.redis.hset("mf_show", date, count)
            print "%s %d mf_show" % (date, count)

            ids = report.distinct("device_id", {"action":"mf_show"})
            count = len(ids)
            self.redis.hset("mf_show", "total", count)
        except Exception, e:
            traceback.print_exc()
            return False

    def ServerRun(self):
        try:
            date = time.strftime('%Y-%m-%d',self.start)
            query =  {
                "action":"server_run",
                "log_time":{"$gte":datetime(self.start.tm_year, self.start.tm_mon, self.start.tm_mday)},
                "log_time":{"$lt":datetime(self.end.tm_year, self.end.tm_mon, self.end.tm_mday)}
            }
            report = self.db.pdf_report
            ids = report.distinct("device_id", query)
            count = len(ids)
            self.redis.hset("server_run", date, count)
            print "%s %d mf_show" % (date, count)

            ids = report.distinct("device_id", {"action":"mf_show"})
            count = len(ids)
            self.redis.hset("server_run", "total", count)
        except Exception, e:
            traceback.print_exc()
            return False

    def Stats(self):
        stats.InstallStart()
        stats.InstallEnd()
        stats.UninstallStart()
        stats.UninstallEnd()
        stats.MFShow()
        stats.ServerRun()

stats = PDFStats()
stats.Stats()
