package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type timePoint struct {
	TheTimeStamp int  `gorm:"column:time_stamp"`
	Flag         bool `gorm:"column:is_begin"`
	IsDelete     bool `gorm:"column:is_delete"`
}

func (u timePoint) TableName() string {
	//绑定MYSQL表名为users
	return "holiday"
}

// 开始和结束时间 计算 准确的时间diff
// 原理 每个假期包括 开始结束时间 都计算一个 timestamp
// 使用开始和结束时间划一条线段， 累加 该计算的数值
// like
func diff(beginTimeStamp, endTimeStamp int) int {

	if endTimeStamp < beginTimeStamp {
		return 0
	}
	timeDiff := 0
	beginTime := beginTimeStamp

	holidayList := getHoliday(beginTimeStamp, endTimeStamp)

	fmt.Println(" in diff function holidays: ", holidayList)

	flag := isHoliday(beginTimeStamp) // 是否是 假期
	isEnd := false

	if len(holidayList) == 0 && !flag {
		return endTimeStamp - beginTimeStamp
	}

	for _, timePoint := range holidayList {
		if timePoint.Flag { // 第一个时间开始节点
			timeDiff += timePoint.TheTimeStamp - beginTime
		}
		beginTime = timePoint.TheTimeStamp
		isEnd = timePoint.Flag

	}
	if !isEnd {
		timeDiff += endTimeStamp - beginTime
	}

	return timeDiff
}

func getHoliday(beginTimeStamp, endTimeStamp int) []timePoint {
	db, err := gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/newtice?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	var holidays []timePoint
	db.Where("time_stamp >= ? and time_stamp <= ? and is_delete = 0", beginTimeStamp, endTimeStamp).Order("time_stamp, is_begin desc", true).Find(&holidays)

	return holidays
}

func isHoliday(beginTimeStamp int) bool {
	db, err := gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/newtice?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	var holidays []timePoint // 有个结束时间
	db.Where("time_stamp >= ? and time_stamp <= ? and is_delete = 0 and is_begin = 1", beginTimeStamp, beginTimeStamp+24*60*60).Find(&holidays)

	if len(holidays) > 0 {
		return true
	}

	return false
}

func main() {
	// holidays := getHoliday(1546272000, 1577635200)

	timeDiff := diff(1546271000, 1546272010) // 1546272000

	timeDiff = diff(1546272000, 1546273000) // 1546272000

	timeDiff = diff(1546272000, 1546617600) // 1546272000

	timeDiff = diff(1546617600, 1546272000)

	timeDiff = diff(1556035200, 1556380800)

	fmt.Println(timeDiff)
}
