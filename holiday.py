#!/usr/bin/env python
# -*- coding: UTF-8* -*-
"""
@author panshuangqing
@create 2017.9.17
"""
import requests
import datetime
import time
import MySQLdb.cursors

def mysql_init():
    connect = MySQLdb.connect(
        host = "127.0.0.1",
        port = 3306,
        user = "root",
        passwd = "root",
        db = "test",
        charset = 'utf8',
        cursorclass = MySQLdb.cursors.DictCursor
    )
    #cursor = connect.cursor()
    return connect

def add_mysql(time_format="2020-01-01"):
    """
     如果追求性能的话 可以批量写入 暂时没有必要
    """
    connect = mysql_init()
    cursor = connect.cursor()
    time_array = time.strptime(time_format, "%Y-%m-%d")
    begin_timestamp = time.mktime(time_array)
    end_timestamp = begin_timestamp + 24 * 60 * 60
    insert_sql = "insert into holiday(time_stamp, is_begin) values(%s, 1)" % begin_timestamp

    cursor.execute(insert_sql)
    connect.commit()

    insert_sql =  "insert into holiday(time_stamp, is_begin) values(%s, 0)" % end_timestamp

    cursor.execute(insert_sql)

    connect.commit()

    cursor.close()
    connect.close()

def add(year=2020):
    """
     先从接口获取 所有的假期包括 周六和周日
     将数据写到数据库中
    """
    url = 'http://timor.tech/api/holiday/year/2020'
    result = requests.get(url)
    db_json = result.json()
    holiday_set = set()
    no_holiday_set = set()
    for aa in db_json["holiday"]:
        bb = db_json["holiday"][aa]
        time_format = "2020-" + aa
        if bb["holiday"] == False:
            no_holiday_set.add(time_format)
            continue
        holiday_set.add(time_format)
    begin_date_time = datetime.datetime(year, 1, 1)
    end_date_time = datetime.datetime(year + 1, 1, 1)
    while begin_date_time.strftime("%Y-%m-%d") < end_date_time.strftime("%Y-%m-%d"):
        if begin_date_time.weekday() in [6, 5] and begin_date_time.strftime("%Y-%m-%d") not in no_holiday_set:
            holiday_set.add(begin_date_time.strftime("%Y-%m-%d"))
        begin_date_time = datetime.timedelta(days=1) + begin_date_time
    bb = sorted(list(holiday_set))

    for holiday_time_format in bb:
        add_mysql(int(holiday_time_format))
        time.sleep(1)

if __name__ == "__main__":
    add()