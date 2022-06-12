package timer

import "time"

//获取当前时间
func GetNowTime() time.Time {

	return time.Now()
}

//根据持续时间，计算时间
func GetCalculateTime(currentTime time.Time, d string) (time.Time, error) {
	//计算持续时间
	duration, err := time.ParseDuration(d)

	if err != nil {
		return time.Time{}, err
	}

	return currentTime.Add(duration), nil
}
