package calendar

import (
	"fmt"
	"github.com/nosixtools/solarlunar"
	"sort"
	"strconv"
	"strings"
	"time"
)

type countSlice []Festival

func (c countSlice) Len() int { // 重写 Len() 方法
	return len(c)
}
func (c countSlice) Swap(i, j int) { // 重写 Swap() 方法
	c[i], c[j] = c[j], c[i]
}
func (c countSlice) Less(i, j int) bool { // 重写 Less() 方法， 从大到小排序
	return c[i].subDay < c[j].subDay
}

const (
	Year = 365 //一年多少天
)

type Festival struct {
	chineseName string //节日中文名
	subDay      int    //节日剩余日期
	date        string //节日日期
}

func (f *Festival) SetChineseName(s string) {
	f.chineseName = s
	return
}
func (f Festival) GetChineseName() string {
	return f.chineseName
}
func (f *Festival) SetSubDay(i int) {
	f.subDay = i
	return
}
func (f Festival) GetSubDay() int {
	return f.subDay
}
func (f *Festival) SetDate(s string) {
	f.date = s
	return
}
func (f Festival) GetDate() string {
	return f.date
}

var (
	Countdown []Festival
	//公历节日
	NewYear                  Festival // 元旦
	ValentinesDay            Festival //情人节
	WomensDay                Festival //妇女节
	ArborDay                 Festival //植树节
	AprilFoolsDay            Festival //愚人节
	LaborDay                 Festival //劳动节
	YouthDay                 Festival //青年节
	ChildrensDay             Festival //儿童节
	NationalDay              Festival //国庆节
	NationalPovertyReliefDay Festival //国家扶贫日
	Halloween                Festival //万圣夜
	Double11                 Festival //
	NationalConstitutionDay  Festival ////国家宪法日
	ThanksgivingDay          Festival //感恩节
	Double12                 Festival //
	NationalMemorialDay      Festival //国家公祭日
	ChristmasEve             Festival //平安夜
	Christmas                Festival //圣诞节
	//农历节日
	ChineseNewYear      Festival //春节
	LanternFestival     Festival //元宵节
	ChingMingFestival   Festival //清明节
	MothersDay          Festival //母亲节
	DragonBoatFestival  Festival //端午节
	SingleDogDay        Festival //七夕节
	MidyearFestival     Festival //中元节
	MidAutumnFestival   Festival //中秋节
	DoubleNinthFestival Festival //重阳节
	LabaFestival        Festival //腊八节
)

func init() {
	NewYear.SetChineseName("元旦")
	NewYear.SetDate("01-01")
	NewYear.SetSubDay(allInSolar(NewYear.GetDate()))
	Countdown = append(Countdown, NewYear)

	ValentinesDay.SetChineseName("情人节")
	ValentinesDay.SetDate("02-14")
	ValentinesDay.SetSubDay(allInSolar(ValentinesDay.GetDate()))
	Countdown = append(Countdown, ValentinesDay)

	WomensDay.SetChineseName("妇女节")
	WomensDay.SetDate("03-08")
	WomensDay.SetSubDay(allInSolar(WomensDay.GetDate()))
	Countdown = append(Countdown, WomensDay)

	ArborDay.SetChineseName("植树节")
	ArborDay.SetDate("03-12")
	ArborDay.SetSubDay(allInSolar(ArborDay.GetDate()))
	Countdown = append(Countdown, ArborDay)

	AprilFoolsDay.SetChineseName("愚人节")
	AprilFoolsDay.SetDate("04-01")
	AprilFoolsDay.SetSubDay(allInSolar(AprilFoolsDay.GetDate()))
	Countdown = append(Countdown, AprilFoolsDay)

	LaborDay.SetChineseName("劳动节")
	LaborDay.SetDate("05-01")
	LaborDay.SetSubDay(allInSolar(LaborDay.GetDate()))
	Countdown = append(Countdown, LaborDay)

	YouthDay.SetChineseName("青年节")
	YouthDay.SetDate("05-04")
	YouthDay.SetSubDay(allInSolar(YouthDay.GetDate()))
	Countdown = append(Countdown, YouthDay)

	ChildrensDay.SetChineseName("儿童节")
	ChildrensDay.SetDate("06-01")
	ChildrensDay.SetSubDay(allInSolar(ChildrensDay.GetDate()))
	Countdown = append(Countdown, ChildrensDay)

	NationalDay.SetChineseName("国庆节")
	NationalDay.SetDate("10-01")
	NationalDay.SetSubDay(allInSolar(NationalDay.GetDate()))
	Countdown = append(Countdown, NationalDay)

	NationalPovertyReliefDay.SetChineseName("国家扶贫日")
	NationalPovertyReliefDay.SetDate("10-17")
	NationalPovertyReliefDay.SetSubDay(allInSolar(NationalPovertyReliefDay.GetDate()))
	Countdown = append(Countdown, NationalPovertyReliefDay)

	Halloween.SetChineseName("万圣夜")
	Halloween.SetDate("10-31")
	Halloween.SetSubDay(allInSolar(Halloween.GetDate()))
	Countdown = append(Countdown, Halloween)

	Double11.SetChineseName("双十一")
	Double11.SetDate("11-11")
	Double11.SetSubDay(allInSolar(Double11.GetDate()))
	Countdown = append(Countdown, Double11)

	NationalConstitutionDay.SetChineseName("国家宪法日")
	NationalConstitutionDay.SetDate("12-04")
	NationalConstitutionDay.SetSubDay(allInSolar(NationalConstitutionDay.GetDate()))
	Countdown = append(Countdown, NationalConstitutionDay)

	ThanksgivingDay.SetChineseName("感恩节")
	ThanksgivingDay.SetDate("11-25")
	ThanksgivingDay.SetSubDay(allInSolar(ThanksgivingDay.GetDate()))
	Countdown = append(Countdown, ThanksgivingDay)

	Double12.SetChineseName("双十二")
	Double12.SetDate("12-12")
	Double12.SetSubDay(allInSolar(Double12.GetDate()))
	Countdown = append(Countdown, Double12)

	NationalMemorialDay.SetChineseName("国家公祭日")
	NationalMemorialDay.SetDate("12-13")
	NationalMemorialDay.SetSubDay(allInSolar(NationalMemorialDay.GetDate()))
	Countdown = append(Countdown, NationalMemorialDay)

	ChristmasEve.SetChineseName("平安夜")
	ChristmasEve.SetDate("12-24")
	ChristmasEve.SetSubDay(allInSolar(ChristmasEve.GetDate()))
	Countdown = append(Countdown, ChristmasEve)

	Christmas.SetChineseName("圣诞节")
	Christmas.SetDate("12-25")
	Christmas.SetSubDay(allInSolar(Christmas.GetDate()))
	Countdown = append(Countdown, Christmas)

	ChineseNewYear.SetChineseName("春节")
	ChineseNewYear.SetDate("12-29")
	ChineseNewYear.SetSubDay(allInLuna(ChineseNewYear.GetDate()) + 2)
	Countdown = append(Countdown, ChineseNewYear)

	LanternFestival.SetChineseName("元宵节")
	LanternFestival.SetDate("01-15")
	LanternFestival.SetSubDay(allInLuna(LanternFestival.GetDate()))
	Countdown = append(Countdown, LanternFestival)

	ChingMingFestival.SetChineseName("清明节")
	ChingMingFestival.SetDate("03-05")
	ChingMingFestival.SetSubDay(allInLuna(ChingMingFestival.GetDate()))
	Countdown = append(Countdown, ChingMingFestival)

	MothersDay.SetChineseName("母亲节")
	MothersDay.SetDate("04-08")
	MothersDay.SetSubDay(allInLuna(MothersDay.GetDate()))
	Countdown = append(Countdown, MothersDay)

	DragonBoatFestival.SetChineseName("端午节")
	DragonBoatFestival.SetDate("05-05")
	DragonBoatFestival.SetSubDay(allInLuna(DragonBoatFestival.GetDate()))
	Countdown = append(Countdown, DragonBoatFestival)

	SingleDogDay.SetChineseName("七夕节")
	SingleDogDay.SetDate("07-07")
	SingleDogDay.SetSubDay(allInLuna(SingleDogDay.GetDate()))
	Countdown = append(Countdown, SingleDogDay)

	MidyearFestival.SetChineseName("中元节")
	MidyearFestival.SetDate("07-15")
	MidyearFestival.SetSubDay(allInLuna(MidyearFestival.GetDate()))
	Countdown = append(Countdown, MidyearFestival)

	MidAutumnFestival.SetChineseName("中秋节")
	MidAutumnFestival.SetDate("08-15")
	MidAutumnFestival.SetSubDay(allInLuna(MidAutumnFestival.GetDate()))
	Countdown = append(Countdown, MidAutumnFestival)

	DoubleNinthFestival.SetChineseName("重阳节")
	DoubleNinthFestival.SetDate("09-09")
	DoubleNinthFestival.SetSubDay(allInLuna(DoubleNinthFestival.GetDate()))
	Countdown = append(Countdown, DoubleNinthFestival)

	LabaFestival.SetChineseName("腊八节")
	LabaFestival.SetDate("12-08")
	LabaFestival.SetSubDay(allInLuna(LabaFestival.GetDate()))
	Countdown = append(Countdown, LabaFestival)

	sort.Sort(countSlice(Countdown))
}
func allInSolar(date string) int {
	day := strings.Join([]string{thisYear(), date}, "-")
	ret, _ := time.Parse("2006-01-02", day)
	unsub := ret.Sub(time.Now())
	if unsub < 0 {
		return int(unsub.Hours())/24 + Year
	}
	return int(unsub.Hours()) / 24
}
func allInLuna(date string) int {
	day := strings.Join([]string{thisYear(), date}, "-")
	convert := solarlunar.LunarToSolar(day, false)
	ret, _ := time.Parse("2006-01-02", convert)
	unsub := ret.Sub(time.Now())
	return int(unsub.Hours()) / 24
}

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
	for _, v := range Countdown {
		if v.GetSubDay() < 0 {
			continue
		}
		if v.GetSubDay() == 0 {
			fmt.Printf("明天是%v\n", v.GetChineseName())
			continue
		}
		fmt.Printf("距离%v还有%v天\n", v.GetChineseName(), v.GetSubDay())
	}
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
}

//计算和元旦的差值
func nextNewYear() {
	thisYearInt, _ := strconv.Atoi(thisYear())
	nextYearInt := thisYearInt + 1
	nextYearStr := strconv.Itoa(nextYearInt)
	day := strings.Join([]string{nextYearStr, "01-01"}, "-") //2021-01-01
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
