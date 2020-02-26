package extractors

import (
        //"fmt"
        // "encoding/json"

        "annie/downloader"
        "annie/extractors/utils"
)

type DouyinVideoUrlData struct {
        Url_list []string
}

type DouyinVideoData struct {
        Play_addr   DouyinVideoUrlData
        Real_play_addr string
}

type DouyinData struct {
        Video DouyinVideoData
        Desc    string
}

func Douyin(url string) utils.VideoData {
        html := downloader.Get(url)
       // fmt.Println("html is:", html)
        vDataUrl := downloader.Match1(`playAddr: \"(.*?)\",`, html)[1]
        //fmt.Println("vDataUrl is:", vDataUrl)

        var dataDict DouyinData
        dataDict.Desc = downloader.Match1(`<p class="desc">(.*?)</p>`, html)[1]
        //fmt.Println("dataDict.Desc:", dataDict.Desc)
        //json.Unmarshal([]byte(vData), &dataDict)

        data := utils.VideoData {
                Site: "Douyin douyin.com",
                Title:  dataDict.Desc,
                Url:  vDataUrl,// dataDict.Video.Real_play_addr,
                Ext:    "mp4",
        }
        data.Size = downloader.UrlSize(data.Url)
        downloader.UlrSave(data)
        return data
}
