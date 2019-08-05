package action

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

func Crawler(c echo.Context) error {


}

//下面是个爬虫，暂时与业务代码没有关系
var (
	// 默认爬取网页 url
	TargetUrl = "https://baidu.com"
	// 默认存储图片地址
	DirPath = "D:/test"
	// 最大爬取图片数量
	MaxNum = 100
	CurNum = 0

	waitGroup sync.WaitGroup

	// 图片url的通道
	chanImgUrl = make(chan string)
	// 网页url的通道(供图片解析)
	chanUrl = make(chan string, 10)
	// 所有网页
	urls []string
)

func Init() {
	var url, dirPath string

	fmt.Printf("请输入链接：(默认%s)\n", TargetUrl)
inputUrl:
	fmt.Scanln(&url)
	if url != "" {
		_, err := http.Get(url)
		if err != nil {
			fmt.Println("输入链接无法访问，请重新输入！")
			goto inputUrl
		} else {
			TargetUrl = url
		}
	}

	fmt.Printf("请输入存储图片地址：（默认%s）\n", DirPath)
	fmt.Scanln(&dirPath)
	if dirPath != "" {
		_, err := os.Open(dirPath)
		if err != nil {
			err = os.MkdirAll(dirPath, os.ModePerm)
			if err != nil {
				fmt.Println("存储地址错误，即将使用默认配置！")
			} else {
				DirPath = dirPath
				fmt.Println("创建文件夹成功：" + dirPath + " 成功！")
			}
		}
	}

	var num int
	fmt.Printf("请输入需要获取图片数：（默认%d）\n", MaxNum)
	fmt.Scanln(&num)
	if num > 0 {
		MaxNum = num
	}

	_, err := os.Open(DirPath)
	if err != nil {
		err = os.MkdirAll(DirPath, os.ModePerm)
		if err == nil {
			fmt.Println("创建文件夹：" + DirPath + " 成功！")
		} else {
			fmt.Println("创建文件夹失败：" + err.Error())
			os.Exit(1)
		}
	}

	// 解析目标url
	chanUrl <- TargetUrl

	waitGroup.Add(1)
	go func() {
		urls = append(urls, TargetUrl)
		for url := urls[0]; len(urls) > 0; {
			getUrl(url)
		}
		fmt.Println("解析所有url结束！")

		// 解析完成，释放chanUrl
		close(chanUrl)
		waitGroup.Done()
	}()

	waitGroup.Add(1)
	go func() {
		for url := range chanUrl {
			getPictureUrl(url)
		}

		fmt.Println("解析所有图片url结束！")
		// 解析完成，释放chanUrl
		close(chanImgUrl)
		waitGroup.Done()
	}()

	// 开启多个下载线程
	for i := 0; i < 5; i++ {
		waitGroup.Add(1)
		go readChanUrl(chanImgUrl)
	}

	waitGroup.Wait()

	fmt.Println("输入任意键退出！")
	fmt.Scanln()
}

/*
 根据传入的url 地址，解析出该网页上所有的图片链接
*/
func getPictureUrl(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("URL ERROR")
		return
	}
	html, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	length := len(html)

	for index := 0; index < length; index++ {
		if string(html[index]) == "<" && index < length-4 {
			// 解析img，截取<img >，并获取其中的url
			if string(html[index+1]) == "i" && string(html[index+2]) == "m" && string(html[index+3]) == "g" {
				for j := index + 3; j < length; j++ {
					if string(html[j]) == ">" {
						imgTag := string(html[index : j+1])
						imgUrl := parseImgUrlFromImgTag(imgTag)
						if imgUrl != "" {
							chanImgUrl <- imgUrl
							index += j - index
						}
						break
					}
				}
			}
		}
	}
}

/*
 根据传入的url 地址，解析出该网页上所有的图片链接
*/
func getUrl(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("URL ERROR")
		return
	}
	html, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	length := len(html)

	for index := 0; index < length; index++ {
		if string(html[index]) == "<" && index < length-4 {
			if string(html[index+1]) == "a" { // 解析<a>
				for j := index + 1; j < length; j++ {
					if string(html[j]) == ">" {
						aTag := string(html[index : j+1])
						aUrl := parseUrlFromATag(aTag)
						if aUrl != "" {
							chanUrl <- aUrl
							urls = append(urls, aUrl)
							index += j - index
						}
						break
					}
				}
			}
		}
	}
}

/*
 解析<img /> 标签中的图片链接
*/
func parseImgUrlFromImgTag(imgTag string) string {
	var result string

	slice := strings.Split(imgTag, "\"")
	for _, v := range slice {
		if strings.HasPrefix(v, "http") {
			result = v
			break
		}
	}

	return result
}

/*
 解析<a /> 标签中的图片链接
*/
func parseUrlFromATag(aTag string) string {
	var result string

	slice := strings.Split(aTag, "\"")
	for _, v := range slice {
		if strings.HasPrefix(v, "http") {
			result = v
			break
		}
	}

	return result
}

/*
 从chan 中不断读取并下载
*/
func readChanUrl(chanUrl chan string) {
	for url := range chanImgUrl {
		if url == "" {
			fmt.Println("【CLOSE】")
			break
		}

		var prefix string
		/*slice :=strings.Split(url, "/")
		prefix = slice[len(slice) - 1]
		prefix = strings.Split(prefix, ".")[0]
		if prefix == "" {
			prefix = string(rand.Int())
		}*/
		prefix = strconv.FormatInt(time.Now().UnixNano(), 10)

		fileName := DirPath + "/" + prefix + ".jpg"

		downLoad(url, fileName)
	}
	fmt.Println("下载图片任务结束！")
	waitGroup.Done()
}

/*
 下载
*/
func downLoad(url, fileName string) {
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("【ERROR】" + url)
		return
	}
	defer res.Body.Close()

	bytes, _ := ioutil.ReadAll(res.Body)
	err = ioutil.WriteFile(fileName, bytes, 0644)
	if err != nil {
		fmt.Println("download file " + url + " error")
		fmt.Println(err)
	} else {
		CurNum++
		fmt.Printf("当前获取图片数：%d\n", CurNum)
		if CurNum == MaxNum {
			fmt.Println("获取所有图片完成!")
			os.Exit(1)
		}
	}
}
