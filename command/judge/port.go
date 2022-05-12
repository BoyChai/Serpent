package judge

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

//代码变量直接摆烂了到时候在优化吧

// PortRange 直接字符串转int类型不建议使用这里函数目前只用于port函数的if判断
func portrange(port string) int {
	value, _ := strconv.Atoi(port)
	return value
}

// horizontalLine 判断-类型端口的字符串，进行处理返回int类型
func horizontalLine(port string) ([]int, error) {
	a := strings.Split(port, "-")
	for _, a2 := range a {
		aa, _ := regexp.MatchString("^\\d+$|^\\d+[.]?\\d+$;", a2)
		if aa == false && 0 < portrange(a2) && portrange(a2) <= 65535 {
			return nil, errors.New("Please pay attention to the syntax specificatio")
		}
	}
	first, _ := strconv.Atoi(a[0])
	last, _ := strconv.Atoi(a[1])
	var vulues []int
	for i := first; i <= last; i++ {
		vulues = append(vulues, i)
	}
	return vulues, nil
}

//判断,类型的端口进行处理返回int端口
func comma(ports string) ([]int, error) {
	port := strings.Split(ports, ",")
	for _, a2 := range port {
		aa, _ := regexp.MatchString("^\\d+$|^\\d+[.]?\\d+$;", a2)
		if aa == false && 0 < portrange(a2) && portrange(a2) <= 65535 {
			return nil, errors.New("Please pay attention to the syntax specificatio")
		}
	}
	var values []int
	for _, value := range port {
		v, _ := strconv.Atoi(value)
		values = append(values, v)
	}
	return values, nil

}

// PortConversion 提供端口判断返回一个int切片
func PortConversion(port string) ([]int, error) {
	//正则
	horizontalbar, _ := regexp.MatchString("-", port)
	commajudge, _ := regexp.MatchString(",", port)
	whole, _ := regexp.MatchString("^\\d+$|^\\d+[.]?\\d+$;", port)
	if horizontalbar {
		ports, err := horizontalLine(port)
		return ports, err
	} else if commajudge {
		ports, err := comma(port)
		return ports, err
	} else if whole {
		var ports []int
		ports = append(ports, portrange(port))
		return ports, nil
	} else {
		return nil, errors.New("Please pay attention to the syntax specificatio")
	}
}
