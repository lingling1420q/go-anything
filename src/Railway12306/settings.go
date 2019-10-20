package Railway12306

import (
	"fmt"
	"strings"
)

var STATION_URL = "https://kyfw.12306.cn/otn/resources/js/framework/station_name.js?station_version=1.9114"
var trainUrl = "https://kyfw.12306.cn/otn/leftTicket/query?leftTicketDTO.train_date=%s&leftTicketDTO.from_station=%s&leftTicketDTO.to_station=%s&purpose_codes=%s"

func RailWayURL(date string, fromStation string, toStation string, codes string) string {
	return fmt.Sprintf(trainUrl, date, strings.ToUpper(fromStation), strings.ToUpper(toStation), strings.ToUpper(codes))
}

var TickType map[int]string

func init() {
	TickType = make(map[int]string)
	TickType[1] = "adult"      // 成人票
	TickType[2] = "child"      // 孩票
	TickType[3] = "student"    // 学生票
	TickType[4] = "disability" // 伤残军人票
}

// 参考：https://github.com/sunhailin-Leo/12306-Go
const (
	HXHBed       int = 33 // 动卧
	SeatBusiness int = 32 // 商务座
	SeatFirst    int = 31 // 一等座
	SeatSecond   int = 30 // 二等座
	HardSeat     int = 29 // 硬座
	HardBed      int = 28 // 硬卧
	NoSeat       int = 26 // 无座
	SeatSpecial  int = 25 // 特等座
	SoftSeat     int = 23 // 软座
)
