#!/usr/bin/env python
# coding: UTF-8

'''
获取hdfs的namenode
非高可用集群中，暂只支持dpp
'''

import xml.etree.ElementTree as ET
import subprocess as SP
import socket
import os

hdfsSiteConfigFile = "/etc/hadoop/conf/hdfs-site.xml"


def execCmd(cmd):
    r = os.popen(cmd)
    text = r.read().strip()
    r.close()
    return text

if __name__ == '__main__':
        namenodehostname = execCmd("hdfs getconf -namenodes")
        namenode_list = filter(None,namenodehostname.split(" "))
        for namenode in namenode_list:
           print(execCmd("cat /etc/hosts|grep %s|awk '{print $1}'" % namenode))
