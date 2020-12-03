from pyhive.hive import connect
con = connect(host='192.168.90.105',port=10000,auth='KERBEROS',kerberos_service_name="hive")
cursor = con.cursor()
cursor.execute('select * from chbtest.shijian')
datas = cursor.fetchall()
print(datas)
cursor.close()
con.close()