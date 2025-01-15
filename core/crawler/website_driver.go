package crawler

import (
	"context"
	"fmt"
	"github.com/chromedp/chromedp"
	"github.com/gocolly/colly/v2"
	"github.com/google/uuid"
	"os"
	"pacyuribot/global"
	"pacyuribot/logger"
	"time"
)

type VisitWebsiteCallback func()

func VisitWebsite(url string, c *colly.Collector, callback VisitWebsiteCallback) {
	// 保证cache文件夹存在
	_, err := os.Stat("cache")
	if os.IsNotExist(err) {
		// 如果不存在，创建文件夹
		err = os.Mkdir("cache", os.ModePerm)
		if err != nil {
			panic(err)
		}
	}

	filePath := "cache/" + uuid.New().String() + ".html"
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// 设置超时时间
	ctx, cancel = context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	// 定义结果变量
	var htmlContent string

	// 启动任务
	err = chromedp.Run(ctx,
		chromedp.Navigate(url),                         // 替换为目标URL
		chromedp.WaitVisible(`body`, chromedp.ByQuery), // 等待页面加载完成
		chromedp.OuterHTML("html", &htmlContent),       // 获取完整HTML内容
	)
	if err != nil {
		logger.Error("Failed to run chromedp: %s", err.Error())
		return
	}
	// 打印获取的内容
	err = os.WriteFile(filePath, []byte(htmlContent), 0644)
	if err != nil {
		logger.Error("Failed to Write Cache File: %s", err.Error())
		return
	}
	localURL := fmt.Sprintf("http://localhost:%d/%s", global.Config.Server.Port, filePath)
	err = c.Visit(localURL)
	if err != nil {
		logger.Error("Failed to visit local resource: %s, localURL is %s", err.Error(), localURL)
		return
	}
	callback()

	err = os.Remove(filePath)
	if err != nil {
		logger.Error("Failed to delete file: %s", err.Error())
	}
}
