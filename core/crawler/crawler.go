package crawler

import "C"
import (
	"container/list"
	"github.com/gocolly/colly/v2"
	"net/url"
	"pacyuribot/core/pair"
	"pacyuribot/logger"
	"pacyuribot/utils"
	"regexp"
)

// Crawler 仅包含基本的功能，扩展功能使用filter实现
type Crawler interface {
	AddAllowedDomains(domains []string) Crawler
	AddDisallowedDomains(domains []string) Crawler
	AddDisallowedURLFilter(reg *regexp.Regexp) Crawler
	AddContentFilter(filter ContentFilter, priority int) Crawler // 越大越优先
	AddUrlFilter(filter UrlFilter, priority int) Crawler         // 越大越优先
	AddPageCrawledCallback(callback PageCrawledCallback, priority int) Crawler
	AddTargetUrls(urls []string) Crawler
	Run()
}

type ContentBuilder struct {
	Contents  *list.List
	SourceURL string
}

func (c *ContentBuilder) Add(content string) {
	c.Contents.PushBack(content)
}

func (c *ContentBuilder) Text() string {
	text := ""
	for e := c.Contents.Front(); e != nil; e = e.Next() {
		text = text + e.Value.(string)
	}
	return text
}

func NewContentBuilder(url string) *ContentBuilder {
	return &ContentBuilder{list.New(), url}
}

type DefaultCrawlerCTX struct {
	contentBuilder *ContentBuilder
	element        *colly.HTMLElement
	currentURL     *url.URL
	currentDepth   int
}

func newDefaultCrawlerCTX() *DefaultCrawlerCTX {
	return &DefaultCrawlerCTX{
		nil,
		nil,
		nil,
		0,
	}
}

type UrlListItem struct {
	targetURL *url.URL
	depth     int
}

type DefaultCrawler struct {
	C                   *colly.Collector
	urlList             *list.List
	contentFilterList   *list.List
	pageCrawledCallback *list.List
	urlFilterList       *list.List
	ctx                 *DefaultCrawlerCTX
	AllowDomains        []string
	DisallowDomains     []string
}

func (d *DefaultCrawler) AddTargetUrls(urls []string) Crawler {
	for _, s := range urls {
		parse, err := url.Parse(s)
		if err != nil {
			logger.Debug("Failed to parse url: %s", s)
			continue
		}
		d.urlList.PushBack(UrlListItem{targetURL: parse, depth: 0}) // 使用 PushBack 插入到 list 的尾部
	}
	return d
}

func NewDefaultCrawler() *DefaultCrawler {
	return &DefaultCrawler{
		colly.NewCollector(),
		list.New(),
		list.New(),
		list.New(),
		list.New(),
		newDefaultCrawlerCTX(),
		[]string{},
		[]string{},
	}
}

func (d *DefaultCrawler) AddUrlFilter(filter UrlFilter, priority int) Crawler {
	insertByPriority(d.urlFilterList, filter, priority)
	return d
}

func (d *DefaultCrawler) AddDisallowedURLFilter(reg *regexp.Regexp) Crawler {
	d.C.DisallowedURLFilters = append(d.C.DisallowedURLFilters, reg)
	return d
}

func (d *DefaultCrawler) AddAllowedDomains(domains []string) Crawler {
	d.AllowDomains = append(d.AllowDomains, domains...)
	return d
}

func (d *DefaultCrawler) AddDisallowedDomains(domains []string) Crawler {
	d.DisallowDomains = append(d.DisallowDomains, domains...)
	return d
}

func (d *DefaultCrawler) AddContentFilter(filter ContentFilter, priority int) Crawler {
	insertByPriority(d.contentFilterList, filter, priority)
	return d
}

func (d *DefaultCrawler) AddPageCrawledCallback(callback PageCrawledCallback, priority int) Crawler {
	insertByPriority(d.pageCrawledCallback, callback, priority)
	return d
}

// 泛型函数，根据 priority 将元素插入到双向链表中
func insertByPriority[T any](l *list.List, key T, priority int) {
	p := pair.Pair[T, int]{
		Key:   key,
		Value: priority,
	}

	// 如果链表为空，直接放入头部即可
	if l.Len() == 0 {
		l.PushFront(p)
		return
	}

	// 从前往后遍历，找到合适位置插入
	for e := l.Front(); e != nil; e = e.Next() {
		currentPair := e.Value.(pair.Pair[T, int])
		if currentPair.Value > priority {
			l.InsertBefore(p, e)
			return
		}
	}

	// 如果遍历完都没找到比他大的，插到链表尾部
	l.PushBack(p)
}

func (d *DefaultCrawler) Run() {
	// 过滤所有标签的内容
	d.C.OnHTML("p", func(element *colly.HTMLElement) {
		d.ctx.contentBuilder.Contents.PushBack(element.Text)
		blocked := false
		for f := d.contentFilterList.Front(); f != nil; f = f.Next() {
			d.ctx.contentBuilder.Contents.Back().Value, blocked = f.Value.(pair.Pair[ContentFilter, int]).Key(
				d.ctx.contentBuilder.Contents.Back().Value.(string), d,
			)
			// 被过滤，则溢出当时的内容
			if blocked {
				logger.Debug("Content Filter: %s; Reason: %s",
					element.Text, d.ctx.contentBuilder.Contents.Back().Value.(string))
				d.ctx.contentBuilder.Contents.Remove(d.ctx.contentBuilder.Contents.Back())
				break
			}
		}
	})

	// 加上外链
	d.C.OnHTML("a", func(element *colly.HTMLElement) {
		href := element.Attr("href")
		targetURL, _ := utils.CompleteURL(href, element.Request.URL)
		if targetURL != nil {
			blocked := false
			for f := d.urlFilterList.Front(); f != nil && !blocked; f = f.Next() {
				targetURL, blocked = f.Value.(pair.Pair[UrlFilter, int]).Key(
					targetURL, d,
				)
			}

			if !blocked {
				logger.Debug("added url: " + targetURL.String())
				d.urlList.PushBack(UrlListItem{targetURL: targetURL, depth: d.ctx.currentDepth + 1})
			}
		}
	})

	for e := d.urlList.Front(); e != nil; e = e.Next() {
		urlItem := e.Value.(UrlListItem)
		d.ctx.currentDepth = urlItem.depth
		d.runOnSingleURL(urlItem.targetURL.String())
	}
}

func (d *DefaultCrawler) runOnSingleURL(targetURL string) {
	d.ctx.contentBuilder = NewContentBuilder(targetURL)
	d.ctx.currentURL, _ = url.Parse(targetURL)
	logger.Debug("visiting site: %s", targetURL)
	VisitWebsite(targetURL, d.C, func() {
		// build完后，使用pageCallback，优先级最大的callback会存数据
		for f := d.pageCrawledCallback.Front(); f != nil; f = f.Next() {
			if f.Value.(pair.Pair[PageCrawledCallback, int]).Key(d) {
				break
			}
		}
	})
}
