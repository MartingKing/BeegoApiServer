package crawer

import (
	"fmt"
	"strconv"
	"net/http"
	"regexp"
	"strings"
)

var DuanziMap map[string]string

func DuanziSpiderStart(start int, end int) {
	page := make(chan int)
	for i := start; i <= end; i++ {
		go spider(i, page)
	}

	for i := start; i <= end; i++ {
		<-page
	}
}
func spider(i int, page chan int) {
	url := "https://www.pengfu.com/xiaohua_" + strconv.Itoa(i) + ".html"
	//开始爬取页面内容
	result, err := duanziSpiderStartdoWork(url)
	if err != nil {
		fmt.Println("dowork err", err)
		return
	}
	//解析表达式  正则筛选关键信息
	//<h1 class="dp-b"><a href="https://www.pengfu.com/content_1832263_1.html" target="_blank">出来打工的第二个月</a>
	re := regexp.MustCompile(`<h1 class="dp-b"><a href="(?s:(.*?))"`)
	if re == nil {
		fmt.Println("compile err")
		return
	}
	joyurl := re.FindAllStringSubmatch(result, -1)

	sliceTitle := make([]string, 0)
	sliceContent := make([]string, 0)
	sliceKey := make([]int, 0)
	for key, data := range joyurl {
		//爬取每一个段子
		title, content, err := crowerDuanzi(data[1])
		if err != nil {
			fmt.Println("crower duanzi err")
			continue
		}
		//保存数据到数据库
		sliceTitle = append(sliceTitle, title)
		sliceContent = append(sliceContent, content)
		sliceKey = append(sliceKey, key)

	}
	//将数据保存到map
	DuanziMap = make(map[string]string)
	for i := 0; i < len(sliceKey); i++ {
		DuanziMap[sliceTitle[i]] = sliceContent[i]
	}
	page <- i
}

func duanziSpiderStartdoWork(url string) (result string, err error) {
	resp, errdo := http.Get(url)
	if err != nil {
		err = errdo
		return
	}
	defer resp.Body.Close()
	buf := make([]byte, 1024*4)
	for {
		readbuf, _ := resp.Body.Read(buf)
		if readbuf == 0 {
			break
		}
		result += string(buf[:readbuf])
	}
	return
}

func crowerDuanzi(url string) (title, content string, err error) {
	result, errw := duanziSpiderStartdoWork(url)
	if errw != nil {
		err = errw
		return
	}
	//正则匹配标题 标题只取一个
	//<h1 class="dp-b"><a href="https://www.pengfu.com/content_1832259_1.html" target="_blank">段子标题</a>
	reTitle := regexp.MustCompile(`<h1>(?s:(.*?))</h1>`)
	if reTitle == nil {
		fmt.Println("crowerDuanzi err")
		return
	}

	//获取标题  只过滤第一个  参数设置为1 -1表示过滤所有
	temTitle := reTitle.FindAllStringSubmatch(result, 1)
	for _, data := range temTitle {
		title = data[1]
		title = strings.Replace(title, "\t", "", -1)
		break
	}

	//正则匹配内容
	//<div class="content-img clearfix pt10 relative"> 段子内容</div>

	reContent := regexp.MustCompile(`div class="content-txt pt10">(?s:(.*?))<a id="prev" href="`)
	if reContent == nil {
		fmt.Println("crowerDuanzi err")
		return
	}
	//获取标题  只过滤第一个  参数设置为1 -1表示过滤所有
	temContent := reContent.FindAllStringSubmatch(result, -1)
	for _, data := range temContent {
		content = data[1]
		content = strings.Replace(content, "\t", "", -1)
		content = strings.Replace(content, "\r", "", -1)
		content = strings.Replace(content, "\n", "", -1)
		content = strings.Replace(content, "<br />", "", -1)
		break
	}
	return
}
