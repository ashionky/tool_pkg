/**
 * @Author pibing
 * @create 2022/1/25 9:38 AM
 */

package datetime

import (
	"strconv"
	"time"
)

/**
 * 获取时间的描述
 * @param  str  {string} 时间字符串
 * @return str2 {string} 日期的描述 如今天 15:04:05、昨天 15:04:05
 */
func FormatDateDesc(str string) (str2 string) {
	layout := "2006-01-02 15:04:05"
	t, _ := time.Parse(layout, str)
	nTime := time.Now()
	yesTime := nTime.AddDate(0, 0, -1)
	if nTime.Format("20060102") == t.Format("20060102") {
		return "今天 " + t.Format("15:04:05")
	} else if yesTime.Format("20060102") == t.Format("20060102") {
		return "昨天 " + t.Format("15:04:05")
	} else if nTime.Format("2006") == t.Format("2006") {
		return t.Format("01-02 15:04:05")
	} else {
		return t.Format("2006-01-02 15:04:05")
	}
}

/**
 * 获取时间的描述
 * @param  str  {string} 时间字符串
 * @return str2 {string} 日期的描述  如30分钟前、今天 15:04:05、昨天 15:04:05
 */
func FormatDateDescSuper(str string) (str2 string) {
	layout := "2006-01-02 15:04:05"
	loc := time.Local
	t, _ := time.ParseInLocation(layout, str, loc)
	nTime := time.Now()
	yesTime := nTime.AddDate(0, 0, -1)
	if nTime.Format("20060102") == t.Format("20060102") {
		mins := nTime.Sub(t)
		min := mins.Minutes()
		if min < 1 {
			return "刚刚"
		} else if min < 60 {
			return strconv.FormatFloat(min, 'f', 0, 64) + "分钟前"
		}
		return "今天 " + t.Format("15:04:05")
	} else if yesTime.Format("20060102") == t.Format("20060102") {
		return "昨天 " + t.Format("15:04:05")
	} else if nTime.Format("2006") == t.Format("2006") {
		return t.Format("01-02 15:04:05")
	} else {
		return t.Format("2006-01-02 15:04:05")
	}
}

/**
 * 获取时间的年月日
 * @param  t  {time} 时间
 * @return year、month、day {int} 年 月 日  2021  10  12
 */
func GetYMDFromDate(t *time.Time) (year int, month int, day int) {
	year = t.Year()
	month = int(t.Month())
	day = t.Day()
	return
}

/**
 * 获取时间的时分秒
 * @param  t  {time} 时间
 * @return hour、min、sec {int}  时 分 秒  12  10  12
 */
func GetHMSFromDate(t *time.Time) (hour int, min int, sec int) {
	hour = t.Hour()
	min = t.Minute()
	sec = t.Second()
	return
}

func GetNowDateFormat() string {
	t := time.Now()
	return t.Format("2006-01-02 15:04:05")
}
func GetNowDayFormat() string {
	t := time.Now()
	return t.Format("2006-01-02")
}
func GetNowMonthFormat() string {
	t := time.Now()
	return t.Format("2006-01")
}

func GetNowYearFormat() string {
	t := time.Now()
	return t.Format("2006")
}

// 获取当前时间戳，单位毫秒
func GetNowTsByMillisecond() int64 {
	return time.Now().UnixNano() / 1e6
}

// 获取当前时间戳，单位秒
func GetNowTsBySecond() int64 {
	return time.Now().Unix()
}

// 把日期字符串转为时间戳 单位秒
func GetNowTimestampByString(dataStr string, format string) int64 {
	t, _ := time.Parse(format, dataStr)
	return t.Unix()
}

func GetNowDateTimeFormatByFormat(format string) string {
	t := time.Now()
	return t.Format(format)
}

/*字符串时间转time
  2003-01-02 15:04:05 转为 2003-01-02 15:04:05 +0800 CST
*/
func GetTimeFromDefaultString(str string) time.Time {
	layout := "2006-01-02 15:04:05"
	local, _ := time.LoadLocation("Local")
	t, _ := time.ParseInLocation(layout, str, local)
	return t
}

// 今日最后时间点
func GetDayLatestTime() time.Time {
	layout := "2006-01-02 15:04:05"
	local, _ := time.LoadLocation("Local")
	t, _ := time.ParseInLocation(layout, GetNowDayFormat()+" 23:59:59", local)
	return t
}

// 格式化时间戳为日期字符串
func GetDateStrFormatTs(ts int64) string {
	t := time.Unix(ts, 0)
	return t.Format("2006-01-02 15:04:05")
}

// 格式化时间戳为指定格式的字符串
func GetTimeWithLayout(ts int64, layout string) string {
	t := time.Unix(ts, 0)
	return t.Format(layout)
}

//判断是否为闰年
func IsLeapYear(year int) bool { //y == 2000, 2004
	//判断是否为闰年
	if year%4 == 0 && year%100 != 0 || year%400 == 0 {
		return true
	}

	return false
}

/*获取某月有多少天
  如2000 ， 2  有29天
  如2001 ， 2  有28天
*/
func GetMonthDays(year int, month int) int {
	days := 0
	if month != 2 {
		if month == 4 || month == 6 || month == 9 || month == 11 {
			days = 30
		} else {
			days = 31
		}
	} else {
		if IsLeapYear(year) {
			days = 29
		} else {
			days = 28
		}
	}
	return days
}

//时间字符串 如 2022-01-25 10:44:15 得到  20220125104415
func GetNowTime() string {
	t := time.Now()
	year := strconv.Itoa(t.Year())
	month := strconv.Itoa(int(t.Month()))
	day := strconv.Itoa(t.Day())
	hour := strconv.Itoa(t.Hour())
	min := strconv.Itoa(t.Minute())
	sec := strconv.Itoa(t.Second())

	if len(month) == 1 {
		month = "0" + month
	}

	if len(day) == 1 {
		day = "0" + day
	}
	return year + month + day + hour + min + sec
}

//  GetTimeString 获取报表时间：年 月 日
func GetYMDString() (y, m, d string) {
	t := time.Now()
	y = t.Format("2006")
	m = t.Format("2006-01")
	d = t.Format("2006-01-02")
	return
}

/*两个时间内的日期数组
  如  "2021-01-02 15:04:05","2021-01-06 15:04:05"之间的日期
  有  [20210102 20210103 20210104 20210105 20210106]
*/
func GetTimeArray(startTime, endTime string) []string {
	layout := "2006-01-02 15:04:05"
	start, err := time.Parse(layout, startTime)
	resArr := make([]string, 0)
	if err != nil {
		return resArr
	}
	str1 := start.Format("20060102")
	resArr = append(resArr, str1)
	end, err := time.Parse(layout, endTime)
	if err != nil {
		return resArr
	}
	timeBuild(start, end, &resArr)
	str2 := end.Format("20060102")

	if str2 != str1 {
		resArr = append(resArr, str2)
	}
	return resArr
}

func timeBuild(start time.Time, end time.Time, between *[]string) {
	d, err := time.ParseDuration("24h")
	if err != nil {
		return
	}
	nextDay := start.Add(d)
	if end.After(nextDay) {
		dayStr := nextDay.Format("20060102")
		*between = append(*between, dayStr)
		timeBuild(nextDay, end, between)
	}
}

// 获取某时刻多少天前的时间
// 如 5天之前 2022-01-20 10:40:44   2022-01-25 10:40:44
func GetTimeBeforeDay(start string, num int) (string, string) {
	t := time.Now()
	if start != "" {
		layout := "2006-01-02 15:04:05"
		startTime, err := time.Parse(layout, start)
		if err != nil {
			t = startTime
		}
	} else {
		start = t.Format("2006-01-02 15:04:05")
	}

	allHour := num * 24
	allHourStr := "-" + strconv.Itoa(allHour) + "h"

	d, err := time.ParseDuration(allHourStr)
	if err != nil {
		return "", start
	}
	nextDay := t.Add(d)
	day := nextDay.Format("2006-01-02 15:04:05")
	return day, start
}

// GetNowYearMoth 获得当前年份的年月数组
// 如[202201 202202 202203 202204 202205 202206 202207 202208 202209 202210 202211 202212]
func GetNowYearMoth() []string {
	t := time.Now()
	day := t.Format("2006")
	resArr := make([]string, 0)
	fMonth, err := strconv.Atoi(day + "01")
	if err != nil {
		return resArr
	}
	for i := 0; i < 12; i++ {
		mStr := strconv.Itoa(fMonth + i)
		resArr = append(resArr, mStr)
	}
	return resArr
}

func UtcToLocal(utcTime string) string {
	layout := "2006-01-02T15:04:05Z"
	utc, _ := time.LoadLocation("UTC")
	newTime, _ := time.ParseInLocation(layout, utcTime, utc)

	return newTime.Local().Format("2006-01-02 15:04:05")
}

func LocalToUtc(localTime string) string {
	layout := "2006-01-02 15:04:05"
	local, _ := time.LoadLocation("Local")
	newTime, _ := time.ParseInLocation(layout, localTime, local)

	return newTime.UTC().Format("2006-01-02T15:04:05Z")
}

//比较时间大小 true表示time2>time1
func CompareTimeString(time1, time2 string) bool {

	//先把时间字符串格式化成相同的时间类型
	t1, err := time.Parse("2006-01-02 15:04:05", time1)
	t2, err := time.Parse("2006-01-02 15:04:05", time2)
	if err == nil && t1.Before(t2) {
		//处理逻辑
		return true
	}
	return false
}
