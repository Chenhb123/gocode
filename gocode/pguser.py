#!/usr/bin/python
#-*- coding:utf-8 -*-

import sys
import urllib2
import urllib
import json
import requests
import psycopg2
import os
import pyhs2
import traceback

if __name__ == '__main__':
    if len(sys.argv) < 4:
        exit("参数admin,adminPasswd, 不得为空")
    sqlid = sys.argv[1]
    sqldb = sys.argv[2]
    sqlfile = sys.argv[3]
    #res = requests.post(url = urlParam, params=dataParam, headers=headersParam)
    res = requests.post(url = urlParam, data = json.dumps(dataParam), headers=headersParam)
    print(res.text)
