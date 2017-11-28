#!/bin/python

# -*- coding:utf-8 -*-

import redis

self.redis = redis.Redis(host="storage", port=6379, db=10)


install = redis.hgetall("pdf_stats_install")
print install
