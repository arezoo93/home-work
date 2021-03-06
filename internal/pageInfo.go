package internal

import (
	"fmt"
	"home24/app"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

type Page struct {
	Url     string
	Client  http.Client
	content string
}

func NewPageInfo(url string) *Page {
	page := new(Page)
	page.Client.Timeout = app.Configs.TimeOut
	page.Url = url
	return page
}

func (p *Page) GetInfo() error {
	err := p.getContent()
	if err != nil {
		return err
	}
	p.getHTMLVersion()
	p.getPageTitle()
	p.getPageHeadings()
	p.getLinks()
	p.hasLoginForm()
	return nil
}

func (p *Page) getContent() error {
	req, err := http.NewRequest("GET", p.Url, nil)
	if err != nil {
		return err
	}

	resp, err := p.Client.Do(req)
	if err != nil {
		return err
	}

	defer func() {
		if err = resp.Body.Close(); err != nil {
			panic(err)
		}
	}()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Get status: %d", resp.StatusCode)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	p.content = string(data)
	return nil
}

func (p *Page) getHTMLVersion() {
	pattern := regexp.MustCompile("(?i)<!doctype html>")
	if pattern.MatchString(p.content) {
		fmt.Println("HTML version is 5")
	} else {
		startIndex := strings.Index(p.content, "# HTML ")
		if startIndex >= 0 {
			startIndex = startIndex + len("# HTML ")
		}
		endIndex := strings.Index(p.content[startIndex:], " ")
		fmt.Println("HTML version is", p.content[startIndex:startIndex+endIndex])
	}
}

func (p *Page) getPageTitle() {

	startIndex := strings.Index(p.content, "<title>")
	if startIndex >= 0 {
		endIndex := strings.Index(p.content[startIndex:], "</title>")
		fmt.Println("Page title is:", p.content[startIndex+len("<title>"):startIndex+endIndex])
	}
}

func (p *Page) getPageHeadings() {
	for i := 1; i <= app.Configs.MaxHeadingLevel; i++ {
		heading := fmt.Sprintf("h%d", i)
		count := strings.Count(p.content, "<"+heading)
		fmt.Println("Number of "+heading, ":", count)
	}
}

func (p *Page) getLinks() {
	totalLinks := strings.Count(p.content, "href=\"")
	externalLinks := strings.Count(p.content, "href=\"http")
	internalLinks := totalLinks - externalLinks
	fmt.Println("Number of external links", externalLinks)
	fmt.Println("Number of internal links", internalLinks)
}

func (p *Page) hasLoginForm() {
	pattern := regexp.MustCompile("(?is)<form.*?>.*login.*</form>")
	if pattern.MatchString(p.content) {
		fmt.Println("This page has login form")
	}
}
