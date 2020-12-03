#!/usr/bin/python
#-*- coding:utf-8 -*-
from impala.dbapi import connect
#import urllib2
#import urllib
import json
#import psycopg2
import traceback
rows = list()
import sys,os
import ConfigParser
#from sqlalchemy import *
#from sqlalchemy.engine import create_engine
#from sqlalchemy.schema import *
import pandas as pd
import json
import signal
global dict
def handle_SIGUSR1(signum, frame):
    exit(2)
def query_hive_json():
    try:
        argc = len(sys.argv) 
        if argc ==6:
            sqlfile = sys.argv[5]
           # if not os.path.isfile(sqlfile):
            #    exit("文件不存在")
           # fHandle = open(sqlfile)
           # cont = fHandle.read().strip(";\n\r\t")
           # fHandle.read()
           # fHandle.close()
            sqls = sqlfile.split(";")
            host1 = sys.argv[1]
            port1 = sys.argv[2]
            user1 = sys.argv[3]
            password1 = sys.argv[4]
            conn = connect(host=host1,port=port1,user=user1,password=password1,auth_mechanism="PLAIN")
            num = 0
            c = list()
            global dict 
            for sql in sqls:
                if '--'not in sql  : 
                    if len(sql.replace(' ',''))!=0:    
                        df = pd.read_sql(sql = sql ,con = conn)
                        df.to_json("result.json")
                        with open("result.json","r") as f:
                              a={}
                              b = f.read()
                              a[num]=b
                              num = num + 1 
                              c.append(b)
                    else:
                        continue
                else:
                    continue
            result = dict()
            for key, value in enumerate(c):
                result[key] = value
                os.system("rm -rf result.json")  
            text = json.dumps(result,ensure_ascii=False)
            print text
        elif argc != 6:
            print """Usage: python go_tdh.py host port user password sqls
Example: 
        python go_tdh.py host port user password sqls """
    except Exception as e:
        s=sys.exc_info()
        #global dict
        dict=str(s[1]).replace('DatabaseError','')
        print "[Error]  '%s' ===>  Now  the sql is : '%s' " % (eval(dict)["message"],sql)  
        signal.signal(signal.SIGUSR1, handle_SIGUSR1)  
        os.kill(os.getpid(), signal.SIGUSR1)             
query_hive_json()
