package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"regexp"
	"strings"
	"sync"
	"time"
)

var indexURL = "https://www.weibo.com" //这里就用example来替换掉它的域名，免得被人找麻烦
/*
   ptnIndexItem是一个个播放视频网页的链接
   ptnVideoItem是为了匹配视频播放网页里的视频链接
   dir 是你要下载的路径
*/
var ptnIndexItem = regexp.MustCompile(`<a[^<>]+href *\= *[\"']?(\/[\d]+)\"[^<>]*title\=\"([^\"]*)\".*name.*>`)
var dir = "./xe_video"
var ptnVideoItem = regexp.MustCompile(`<a[^<>]+href *\= *[\"']?(https\:\/\/[^\"]+)\"[^<>]*download[^<>]*>`)

//增加一个等待组
var wg sync.WaitGroup

// DownList xxx
/*
	DownList xxx
   用于记录下载进度的结构体
*/
type DownList struct {
	Data map[string][]int64
	Lock sync.Mutex
}

/*
   检查点
*/
func check(e error) {
	if e != nil {
		panic(e)
	}
}

//Get  xxx
/*
	Get  xxx
  	url: 需要获取的网页

   	return:
       content        抓取到的网页源码
       statusCode    返回的状态码

*/
func Get(url string) (content string, statusCode int) {
	resp, err1 := http.Get(url)
	if err1 != nil {
		statusCode = -100
		return
	}
	defer resp.Body.Close()
	data, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		statusCode = -200
		return
	}
	statusCode = resp.StatusCode
	content = string(data)
	return
}

/*
   param:
       filename:    文件名
       text            需要比较的字符串
   return:
       ture            没
       false        有

*/
func readOnLine(filename string, text string) bool {
	fi, err := os.OpenFile(filename, os.O_RDWR|os.O_APPEND, os.ModeAppend)
	check(err)
	defer fi.Close()
	text = text + "\n"
	br := bufio.NewReader(fi)
	for {
		a, c := br.ReadString('\n')
		if c == io.EOF {
			fmt.Println(text, "不存在，现在写入")
			fi.WriteString(text)
			return true
		}
		if string(a) == text {
			fmt.Println("存在", text)
			break
		}
	}
	return false
}

/*
   输出下载进度的方法

*/
func (downList *DownList) process() {
	for {
		downList.Lock.Lock()
		for key, arr := range downList.Data {
			fmt.Printf("%s progress: [%-50s] %d%% Done\n", key, strings.Repeat("#", int(arr[0]*50/arr[1])), arr[0]*100/arr[1])
		}
		//fmt.Println(downList)
		downList.Lock.Unlock()
		time.Sleep(time.Second * 3)
		fmt.Printf("\033[2J")
	}
}

// Down xxx
/*
   url: 视频链接
   filename: 本地文件名
   downList: 用来记录下载进度的一个结构体的指针
*/
func Down(url string, filename string, downList *DownList) bool {
	b := make([]byte, 1024)
	f, err := os.Create(filename)
	if err != nil {
		fmt.Println("创建文件失败")
		return false
	}
	defer f.Close()
	repo, err := http.Get(url)
	if err != nil {
		fmt.Println("获取资源失败")
		return false
	}
	defer repo.Body.Close()
	bufRead := bufio.NewReader(repo.Body)
	for {
		n, err := bufRead.Read(b)
		if err == io.EOF {
			break
		}
		f.Write(b[:n])
		fileInfo, err := os.Stat(filename)
		fileSize := fileInfo.Size()
		//fmt.Println(fileSize, "--", repo.ContentLength)
		downList.Lock.Lock()
		downList.Data[filename] = []int64{fileSize, repo.ContentLength}
		downList.Lock.Unlock()
	}

	wg.Done()
	return true
}

func main() {
	//初始化downList 与 map
	var downListF DownList
	downListF.Data = make(map[string][]int64)
	downList := &downListF
	fmt.Println("-==============================1")
	//首先获取index网页的内容
	context, statusCode := Get(indexURL)
	if statusCode != 200 {
		fmt.Println("error")
		return
	}
	//fmt.Println("-==============================2", context)
	//return
	/*    提取并复制到二维数组
	      htmlResult              [][]string
	      htmlResult[n]          []string    匹配到的链接
	      htmlResult[n][0]          string     全匹配数据
	      htmlResult[n][1]         url        string
	      htmlResult[n][2]         title        string
	*/

	//播放视频网页的链接
	htmlResult := ptnIndexItem.FindAllStringSubmatch(context, -1)
	fmt.Println(htmlResult)
	return
	length := len(htmlResult)
	go downList.process()
	for i := 0; i < length; i++ {
		v := htmlResult[i]
		videoHTML, videoStatus := Get(indexURL + v[1])
		if videoStatus != 200 {
			fmt.Println("error")
			continue
		}
		videoResult := ptnVideoItem.FindAllStringSubmatch(videoHTML, -1)
		ok := readOnLine("test", v[1])
		if len(videoResult) > 0 && len(videoResult[0]) > 0 && ok {
			//fmt.Println(videoResult[0][1])
			wg.Add(1)
			dirFile := path.Join(dir, htmlResult[i][2])
			go Down(videoResult[0][1], dirFile, downList)

		}
		//fmt.Println(i, videoHTML)
	}
	wg.Wait()
	return

}
