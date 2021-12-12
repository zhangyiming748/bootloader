package calendar

import (
	"fmt"
	"github.com/nosixtools/solarlunar"
	"strconv"
	"strings"
	"time"
)

const (
	NewYear                  = "01-01" //新年
	ValentinesDay            = "02-14" //情人节
	WomensDay                = "03-08" //妇女节
	ArborDay                 = "03-12" //植树节
	AprilFoolsDay            = "04-01" //愚人节
	LaborDay                 = "05-01" //劳动节
	YouthDay                 = "05-04" //青年节
	ChildrensDay             = "06-01" //儿童节
	NationalDay              = "10-01" //国庆节
	NationalPovertyReliefDay = "10-17" //国家扶贫日
	Halloween                = "10-31" //万圣夜
	Double11                 = "11-11"
	NationalConstitutionDay  = "12-04" //国家宪法日
	ThanksgivingDay          = "11-25" //感恩节
	Double12                 = "12-12"
	NationalMemorialDay      = "12-13" //国家公祭日
	ChristmasEve             = "12-24" //平安夜
	Christmas                = "12-25" //圣诞节
)
const (
	ChineseNewYear      = "01-01" //春节
	LanternFestival     = "01-15" //元宵节
	ChingMingFestival   = "03-05" //清明节
	MothersDay          = "04-08" //母亲节
	DragonBoatFestival  = "05-05" //端午节
	SingleDogDay        = "07-07" //七夕节
	MidyearFestival     = "07-15" //中元节
	MidAutumnFestival   = "08-15" //中秋节
	DoubleNinthFestival = "09-09" // 重阳节
	LabaFestival        = "12-08" //腊八节

)

//获取当前年份
func thisYear() string {
	ret := fmt.Sprint(time.Now().Format("2006"))
	return ret
}
func nextYear() string {
	this := thisYear()
	i, _ := strconv.Atoi(this)
	i += 1
	s := strconv.Itoa(i)
	return s
}
func SubDay() {
	fmt.Println("早上好,摸鱼人!")
	fmt.Println(time.Now().Format("今天是2006年1月2日"))
	fmt.Println("【摸鱼办】提醒您")
	fmt.Println("工作再累,一定不要忘记摸鱼哦!")
	fmt.Println("有事没事起身去茶水间,去厕所,去廊道走走")
	fmt.Println("吃饭时间就吃饭")
	fmt.Println("午休时间就午休")
	fmt.Println("别老在工位上坐着,钱是老板的,但命是自己的")
	defer func() {
		fmt.Println("上班是帮老板赚钱,摸鱼是赚老板的钱!")
		fmt.Println("最后,祝愿天下所有摸鱼人,都能愉快的渡过每一天")
	}()
	defer nextNewYear()
	lunar()
	solar()

}
func solar() {
	subValentinesDay()
	subWomensDay()
	subArborDay()
	subAprilFoolsDay()
	subLaborDay()
	subYouthDay()
	subChildrensDay()
	subNationalDay()
	subHalloween()
	subDouble11()
	subNationalConstitutionDay()
	subThanksgivingDay()
	subDouble12()
	subNationalMemorialDay()
	subChristmasEve()
	subChristmas()
}
func lunar() {
	chineseNewYear()
	chineseLanternFestival()
	chineseChingMingFestival()
	chineseMothersDay()
	chineseDragonBoatFestival()
	chineseSingleDogDay()
	chineseMidyearFestival()
	chineseMidAutumnFestival()
	chineseDoubleNinthFestival()
	chineseLabaFestival()
}

//情人节
func subValentinesDay() {
	day := strings.Join([]string{thisYear(), ValentinesDay}, "-")
	ret, _ := time.Parse("2006-01-02", day)
	unsub := ret.Sub(time.Now())
	if unsub < 0 {
		return
	}
	fmt.Printf("距离情人节还有%v天\n", int(unsub.Hours())/24)
}

//妇女节
func subWomensDay() {
	day := strings.Join([]string{thisYear(), WomensDay}, "-")
	ret, _ := time.Parse("2006-01-02", day)
	unsub := ret.Sub(time.Now())
	if unsub < 0 {
		return
	}
	fmt.Printf("距离妇女节还有%v天\n", int(unsub.Hours())/24)
}

//植树节
func subArborDay() {
	day := strings.Join([]string{thisYear(), ArborDay}, "-")
	ret, _ := time.Parse("2006-01-02", day)
	unsub := ret.Sub(time.Now())
	if unsub < 0 {
		return
	}
	fmt.Printf("距离植树节还有%v天\n", int(unsub.Hours())/24)
}

//愚人节
func subAprilFoolsDay() {
	day := strings.Join([]string{thisYear(), AprilFoolsDay}, "-")
	ret, _ := time.Parse("2006-01-02", day)
	unsub := ret.Sub(time.Now())
	if unsub < 0 {
		return
	}
	fmt.Printf("距离愚人节还有%v天\n", int(unsub.Hours())/24)
}

//劳动节
func subLaborDay() {
	day := strings.Join([]string{thisYear(), LaborDay}, "-")
	ret, _ := time.Parse("2006-01-02", day)
	unsub := ret.Sub(time.Now())
	if unsub < 0 {
		return
	}
	fmt.Printf("距离劳动节还有%v天\n", int(unsub.Hours())/24)
}

//青年节
func subYouthDay() {
	day := strings.Join([]string{thisYear(), YouthDay}, "-")
	ret, _ := time.Parse("2006-01-02", day)
	unsub := ret.Sub(time.Now())
	if unsub < 0 {
		return
	}
	fmt.Printf("距离青年节还有%v天\n", int(unsub.Hours())/24)
}

//儿童节
func subChildrensDay() {
	day := strings.Join([]string{thisYear(), ChildrensDay}, "-")
	ret, _ := time.Parse("2006-01-02", day)
	unsub := ret.Sub(time.Now())
	if unsub < 0 {
		return
	}
	fmt.Printf("距离儿童节还有%v天\n", int(unsub.Hours())/24)
}

//国庆节
func subNationalDay() {
	day := strings.Join([]string{thisYear(), NationalDay}, "-")
	ret, _ := time.Parse("2006-01-02", day)
	unsub := ret.Sub(time.Now())
	if unsub < 0 {
		return
	}
	fmt.Printf("距离国庆节还有%v天\n", int(unsub.Hours())/24)
}

//国家扶贫日
func subNationalPovertyReliefDay() {
	day := strings.Join([]string{thisYear(), NationalPovertyReliefDay}, "-")
	ret, _ := time.Parse("2006-01-02", day)
	unsub := ret.Sub(time.Now())
	if unsub < 0 {
		return
	}
	fmt.Printf("距离国家扶贫日还有%v天\n", int(unsub.Hours())/24)
}

//感恩节
func subThanksgivingDay() {
	day := strings.Join([]string{thisYear(), ThanksgivingDay}, "-")
	ret, _ := time.Parse("2006-01-02", day)
	unsub := ret.Sub(time.Now())
	if unsub < 0 {
		return
	}
	fmt.Printf("距离感恩节还有%v天\n", int(unsub.Hours())/24)
}

//万圣夜
func subHalloween() {
	day := strings.Join([]string{thisYear(), Halloween}, "-")
	ret, _ := time.Parse("2006-01-02", day)
	unsub := ret.Sub(time.Now())
	if unsub < 0 {
		return
	}
	fmt.Printf("距离万圣夜还有%v天\n", int(unsub.Hours())/24)
}
func subDouble11() {
	day := strings.Join([]string{thisYear(), Double11}, "-")
	ret, _ := time.Parse("2006-01-02", day)
	unsub := ret.Sub(time.Now())
	if unsub < 0 {
		return
	}
	fmt.Printf("距离双十一还有%v天\n", int(unsub.Hours())/24)
}

//国家宪法日
func subNationalConstitutionDay() {
	day := strings.Join([]string{thisYear(), NationalConstitutionDay}, "-")
	ret, _ := time.Parse("2006-01-02", day)
	unsub := ret.Sub(time.Now())
	if unsub < 0 {
		return
	}
	fmt.Printf("距离国家宪法日还有%v天\n", int(unsub.Hours())/24)
}

func subDouble12() {
	day := strings.Join([]string{thisYear(), Double12}, "-")
	ret, _ := time.Parse("2006-01-02", day)
	unsub := ret.Sub(time.Now())
	if unsub < 0 {
		return
	}
	fmt.Printf("距离双十二还有%v天\n", int(unsub.Hours())/24)
}

//国家公祭日
func subNationalMemorialDay() {
	day := strings.Join([]string{thisYear(), NationalMemorialDay}, "-")
	ret, _ := time.Parse("2006-01-02", day)
	unsub := ret.Sub(time.Now())
	if unsub < 0 {
		return
	}
	fmt.Printf("距离国家公祭日还有%v天\n", int(unsub.Hours())/24)
}

//平安夜
func subChristmasEve() {
	day := strings.Join([]string{thisYear(), ChristmasEve}, "-")
	ret, _ := time.Parse("2006-01-02", day)
	unsub := ret.Sub(time.Now())
	if unsub < 0 {
		return
	}
	fmt.Printf("距离平安夜还有%v天\n", int(unsub.Hours())/24)
}

//圣诞节
func subChristmas() {
	day := strings.Join([]string{thisYear(), Christmas}, "-")
	ret, _ := time.Parse("2006-01-02", day)
	unsub := ret.Sub(time.Now())
	if unsub < 0 {
		return
	}
	fmt.Printf("距离圣诞节还有%v天\n", int(unsub.Hours())/24)
}

//计算和元旦的差值
func nextNewYear() {
	thisYearInt, _ := strconv.Atoi(thisYear())
	nextYearInt := thisYearInt + 1
	nextYearStr := strconv.Itoa(nextYearInt)
	day := strings.Join([]string{nextYearStr, NewYear}, "-") //2021-01-01
	ret, _ := time.Parse("2006-01-02", day)
	unsub := ret.Sub(time.Now())
	if unsub < 0 {
		return
	}
	fmt.Printf("距离明年还有%v天\n", int(unsub.Hours())/24)
	if int(unsub.Hours())/24 <= 31 {
		fmt.Println("过几天又会有人发\"新的一年,新的自己\"这种自欺欺人的话")

	}
}

//func Lunar() {
//	solarDate := "2022-01-01"
//	fmt.Println(solarlunar.SolarToChineseLuanr(solarDate))
//	fmt.Println(solarlunar.SolarToSimpleLuanr(solarDate))
//
//	lunarDate := "2022-01-01"
//	fmt.Println(solarlunar.LunarToSolar(lunarDate, false))
//}
//春节
func chineseNewYear() {
	y := nextYear()
	CHY := strings.Join([]string{y, ChineseNewYear}, "-") //农历新年
	convert := solarlunar.LunarToSolar(CHY, false)
	ret, _ := time.Parse("2006-01-02", convert)
	unsub := ret.Sub(time.Now())
	if unsub < 0 {
		return
	}
	fmt.Printf("距离明年春节还有%v天\n", int(unsub.Hours())/24)
	if int(unsub.Hours())/24 < 7 {
		fmt.Println("新的一年就要到了,谢谢大家长久以来对我的忽视,我很喜欢这种状态,大家都各忙各的,没人鸟我,我也不想鸟你们,新的一年希望大家继续加油,我会一直和你们耗下去")
	}
}

//元宵节
func chineseLanternFestival() {
	y := thisYear()
	CHY := strings.Join([]string{y, LanternFestival}, "-")
	convert := solarlunar.LunarToSolar(CHY, false)
	ret, _ := time.Parse("2006-01-02", convert)
	unsub := ret.Sub(time.Now())
	if unsub < 0 {
		return
	}
	fmt.Printf("距离元宵节还有%v天\n", int(unsub.Hours())/24)
}

//清明节
func chineseChingMingFestival() {
	y := thisYear()
	CHY := strings.Join([]string{y, ChingMingFestival}, "-")
	convert := solarlunar.LunarToSolar(CHY, false)
	ret, _ := time.Parse("2006-01-02", convert)
	unsub := ret.Sub(time.Now())
	if unsub < 0 {
		return
	}
	fmt.Printf("距离清明节还有%v天\n", int(unsub.Hours())/24)
}

//母亲节
func chineseMothersDay() {
	y := thisYear()
	CHY := strings.Join([]string{y, MothersDay}, "-")
	convert := solarlunar.LunarToSolar(CHY, false)
	ret, _ := time.Parse("2006-01-02", convert)
	unsub := ret.Sub(time.Now())
	if unsub < 0 {
		return
	}
	fmt.Printf("距离母亲节还有%v天\n", int(unsub.Hours())/24)
}

//端午节
func chineseDragonBoatFestival() {
	y := thisYear()
	CHY := strings.Join([]string{y, DragonBoatFestival}, "-")
	convert := solarlunar.LunarToSolar(CHY, false)
	ret, _ := time.Parse("2006-01-02", convert)
	unsub := ret.Sub(time.Now())
	if unsub < 0 {
		return
	}
	fmt.Printf("距离端午节还有%v天\n", int(unsub.Hours())/24)
}

//七夕节
func chineseSingleDogDay() {
	y := thisYear()
	CHY := strings.Join([]string{y, SingleDogDay}, "-")
	convert := solarlunar.LunarToSolar(CHY, false)
	ret, _ := time.Parse("2006-01-02", convert)
	unsub := ret.Sub(time.Now())
	if unsub < 0 {
		return
	}
	fmt.Printf("距离七夕节还有%v天\n", int(unsub.Hours())/24)
}

//中元节
func chineseMidyearFestival() {
	y := thisYear()
	CHY := strings.Join([]string{y, MidyearFestival}, "-")
	convert := solarlunar.LunarToSolar(CHY, false)
	ret, _ := time.Parse("2006-01-02", convert)
	unsub := ret.Sub(time.Now())
	if unsub < 0 {
		return
	}
	fmt.Printf("距离中元节还有%v天\n", int(unsub.Hours())/24)
}

//中秋节
func chineseMidAutumnFestival() {
	y := thisYear()
	CHY := strings.Join([]string{y, MidAutumnFestival}, "-")
	convert := solarlunar.LunarToSolar(CHY, false)
	ret, _ := time.Parse("2006-01-02", convert)
	unsub := ret.Sub(time.Now())
	if unsub < 0 {
		return
	}
	fmt.Printf("距离中秋节还有%v天\n", int(unsub.Hours())/24)
}

//重阳节
func chineseDoubleNinthFestival() {
	y := thisYear()
	CHY := strings.Join([]string{y, DoubleNinthFestival}, "-")
	convert := solarlunar.LunarToSolar(CHY, false)
	ret, _ := time.Parse("2006-01-02", convert)
	unsub := ret.Sub(time.Now())
	if unsub < 0 {
		return
	}
	fmt.Printf("距离重阳节还有%v天\n", int(unsub.Hours())/24)
}

//腊八节
func chineseLabaFestival() {
	y := thisYear()
	CHY := strings.Join([]string{y, LabaFestival}, "-")
	convert := solarlunar.LunarToSolar(CHY, false)
	ret, _ := time.Parse("2006-01-02", convert)
	unsub := ret.Sub(time.Now())
	if unsub < 0 {
		return
	}
	fmt.Printf("距离腊八节还有%v天\n", int(unsub.Hours())/24)
}
