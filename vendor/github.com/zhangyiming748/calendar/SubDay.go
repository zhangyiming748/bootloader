package calendar

import (
	"fmt"
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

//获取当前年份
func thisYear() string {
	ret := fmt.Sprint(time.Now().Format("2006"))
	return ret
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
