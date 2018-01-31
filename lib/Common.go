package lib

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

const DATE_TIME_FORMAT = "2006-01-02 15:04:05"
const DATE_FORMAT = "2006-01-02"

var cst *time.Location

func init() {
	cst = time.FixedZone("CST", 28800)
	rand.Seed(time.Now().UnixNano())
}

// 包含头和尾
func Mrand(min, max int) int {
	return rand.Intn(max-min+1) + min
}

func MakeSalt(length int) string {
	var str = []rune("0123456789abcdefghijklmnopqrstuvwxyz")
	var salt = make([]rune, length)
	var max = len(str)
	for i := 0; i < length; i++ {
		salt[i] = str[rand.Intn(max)]
	}
	return string(salt)
}

func Makeid(length int) string {
	return MakeSalt(length)
}

func GetPayForm(ordID string, money int) string {
	var v_mid = "20272562"  //商户号，不可改
	var v_oid = ordID       //订单号
	var v_amount = money    //金额
	var v_moneytype = "CNY" //币种
	var v_url = "http://www.xx.com"
	var v_key = "%()#QOKFDLS:1*&U" //密匙，不可改

	//MD5校验串生成方法：当消费者在商户端生成最终订单的时候，
	// 将订单中的v_amount v_moneytype v_oid v_mid v_url key六个参数的value值
	// 拼成一个无间隔的字符串(顺序不要改变)。参数key是商户的MD5密钥（该密匙可在登陆商户管理界面后自行更改。）
	var form = `<form method=post action="https://pay3.chinabank.com.cn/PayGate">
            <input type="hidden" name=v_mid value="%s">
            <input type="hidden" name=v_oid value="%s">
            <input type="hidden" name=v_amount value="%d">
            <input type="hidden" name=v_moneytype value="%s">
            <input type="hidden" name=v_url value="%s">
            <input type="hidden" name=v_md5info value="%s">
            <input type=submit value="网银在线支付">
            </form>`
	var sign = md5.Sum([]byte(strconv.Itoa(v_amount) + v_moneytype + v_oid + v_mid + v_url + v_key))
	var signStr = strings.ToUpper(hex.EncodeToString(sign[:]))
	form = fmt.Sprintf(form, v_mid, v_oid, v_amount, v_moneytype, v_url, signStr)
	return form
}

// func toChinaDate(date) {
//     var localTime = date.getTime();//毫秒
//     var localOffset = date.getTimezoneOffset() * 60000; //获得当地时区偏移的毫秒数
//     var utc = localTime + localOffset; //还原成utc时间

//     var offset = 8; //中国，东+8区
//     var time = utc + (3600000 * offset);//时间毫秒
//     return new Date(time);
// }

func GetPadStr(number int) string {
	if number >= 10 {
		return strconv.Itoa(number)
	}
	return "0" + strconv.Itoa(number)
}

// func getDateString(date) {
//     return date.getFullYear()
//         + '-' + getPadStr(date.getMonth() + 1)
//         + '-' + getPadStr(date.getDate())
//         + ' ' + getPadStr(date.getHours())
//         + ':' + getPadStr(date.getMinutes())
//         + ':' + getPadStr(date.getSeconds());
// }

// func toDateStr(date) {
//     return getDateString(toChinaDate(date));
// }

func GetCnDate() time.Time {
	return time.Now().In(cst)
}
