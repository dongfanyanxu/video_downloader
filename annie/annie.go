package main

import (
        "flag"
        "fmt"
        "net/url"

        "annie/downloader"
        "annie/extractors"
)

var (
        debug   bool
        version bool
)

func init() {
        flag.BoolVar(&debug, "d", false, "Debug mode")
        flag.BoolVar(&version, "v", false, "Show version")
}

func main() {
        flag.Parse()
        args := flag.Args()
        //fmt.Println("args is:", args)
        //fmt.Println("len(args) is:", len(args))
        if len(args) <1 {
                fmt.Println("error")
                return
        }
        videoUrl := args[0]
        //fmt.Println("videoUrl is:", videoUrl)
        u, err := url.ParseRequestURI(videoUrl)
        //fmt.Println("u, err is:", *u, err)
        if err != nil {
                fmt.Println(err)
                return
        }
        domainPattern := `([a-z0-9][-a-z0-9]{0,62})\.` +
                `(com\.cn|com\.hk|` +
                `cn|com|net|edu|gov|biz|org|info|pro|name|xxx|xyz|` +
                `me|top|cc|tv|tt)`
        domain := downloader.Match1(domainPattern, u.Host)[1]
        //fmt.Println("domain is:", domain)
        switch domain {
            case "iesdouyin":
                extractors.Douyin(videoUrl)
            default:
                fmt.Println("unsupported URL")
        }
}
