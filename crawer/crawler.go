package crawer

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"golang.org/x/text/transform"
	"io"
	"golang.org/x/text/encoding"
	"bufio"
	"golang.org/x/net/html/charset"
	"regexp"
)

func Start() [][][]byte{
	resp, err := http.Get("http://www.zhenai.com/zhenghun")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Println("errpage:", resp.StatusCode)
		return nil
	}
	//将GBK转为UTF-8，否则出现乱码
	e := determingEncoding(resp.Body)
	utf8reader := transform.NewReader(resp.Body, e.NewDecoder())
	//将转化后的编码器传入
	all, err := ioutil.ReadAll(utf8reader)
	if err != nil {
		panic(err)
	}
	return printCityList(all)
}

func determingEncoding(r io.Reader) encoding.Encoding {
	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		panic(err)
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}

func printCityList(contents []byte) [][][]byte{
	/**
	<a href="http://www.zhenai.com/zhenghun/zigong"class="">自贡</a>
	分析：这一部分是相同的不需要多做处理 <a href="http://www.zhenai.com/zhenghun/
		  [0-9a-z]这里是指城市的拼音部分为数字和字母组成
		  [^>]*  这里表示通配除了右括号以外的所有东西，可以是class或者别的
		  [^<]+  表示城市的汉字通配部分
	  	  ()括号包围的表示正则的提取功能，表示我们需要的部分
	 */
	re := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`)
	matchesitem := re.FindAllSubmatch(contents, -1)
	//for _, item := range matchesitem {
	//	for _, submatch := range item {
	//		fmt.Printf("%s ", submatch)
	//	}
	//	fmt.Println()
	//}
	return matchesitem
}
