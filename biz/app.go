package biz

import (
	"fmt"
	"regexp"
	"DTSong/util"
	"DTSong/const"
	"github.com/tidwall/gjson"
	"DTSong/model"
	"os"

	url2 "net/url"
)

type App struct{}

func (this *App) Start() {

	var url string
	//fmt.Println("Please enter your song URL")
	//fmt.Scanln(&url)
	//fmt.Println("url is : %s", url)
	url = "http://m.kugou.com/share/zlist.html?listid=2&type=0&uid=725617981&token=4d029e9bb3216ca221367035ca65203b62d4c1d96003500aa69dac8a51980d78&sign=640957d14c669d07f6bdb42bb4272b34&_t=1513488218&chl=wechat&from=singlemessage&isappinstalled=0";
	songList := this.ParseUrl(url);
	if songList != nil {
		this.DownloadRes(songList)
	}
}

//下载文件
func (this *App) DownloadRes(songList []model.Song) {

	dir, _ := os.Getwd()
	dir += "/DTSongRes"

	//下载失败列表
	failArray := []string{}

	if len(songList) > 0 {
		err := os.MkdirAll(dir, 0777)
		fmt.Println("解析完成开始下载,下载目录为-->",dir)
		if err != nil {
			fmt.Println(err)
		}
		for i := 0; i < len(songList); i++ {
			song := songList[i]
			fileName := dir + "/" + song.Name + ".mp3"
			netError := util.CurlToFile(song.Url, fileName)
			if netError != nil {
				fmt.Println("下载失败 --> " + song.Name)
				failArray = append(failArray, song.Name)
			} else {
				fmt.Printf("下载进度%v/%v\n", i+1, len(songList))
			}
		}
		fmt.Printf("下载完成,目录为:%v\n", dir)
		fmt.Printf("下载失败个数:%v\n", len(failArray))
	}

}

//解析url
func (this *App) ParseUrl(url string) ([]model.Song) {

	isUrl, err := regexp.Match("(https?|ftp|file)://[-A-Za-z0-9+&@#/%?=~_|!:,.;]+[-A-Za-z0-9+&@#/%=~_|]", []byte(url))



	if !isUrl || err != nil {
		fmt.Println("url 格式错误!")
		return nil
	}

	//从url中获取参数
	u, _ := url2.Parse(url)
	m, _ := url2.ParseQuery(u.RawQuery)
	listParams := map[string]string{}
	listParams["listid"] = m["listid"][0]
	listParams["type"] = m["type"][0]
	listParams["uid"] = m["uid"][0]
	listParams["sign"] = m["sign"][0]
	listParams["_t"] = m["_t"][0]
	listParams["pagesize"] = "30000"
	listParams["json"] = ""
	listParams["page"] = "1"
	listParams["token"] = m["token"][0]

	fmt.Println("加载用户歌曲列表....")

	//获得歌曲列表
	res, err := util.GetUrl(_const.GET_USER_MUSIC_LIST, listParams)
	if err != nil {
		fmt.Println(err)
		return nil;
	}

	list := gjson.Get(res, "list.info").Array()

	if len(list) <= 0 {
		fmt.Println("当前用户歌单为空!")
		return nil
	}

	var songList []model.Song



	for i := 0; i < len(list); i++ {
		obj := list[i]
		song := model.Song{}
		song.Name = obj.Get("name").String()
		song.Hash = obj.Get("hash").String()

		infoParams := map[string]string{}
		infoParams["hash"] = song.Hash
		infoParams["cmd"] = "playInfo"
		infoParams["is_share"] = "1"

		fmt.Printf("加载用户歌曲详情%v/%v\n",i+1,len(list))

		//加载歌曲详情
		infoRes, err := util.GetUrl(_const.GET_USER_MUSIC_INFO, infoParams)
		if err != nil {
			fmt.Println(err)
			return nil;
		}
		song.Url = gjson.Get(infoRes, "url").String()
		song.ExtName = gjson.Get(infoRes, "extName").String()
		songList = append(songList, song)
	}

	return songList;
}
