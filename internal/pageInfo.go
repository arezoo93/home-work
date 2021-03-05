package internal

import (
	"fmt"
	"home24/app"
	"io/ioutil"
	"net/http"
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
	p.getPageTitleVersion()
	p.getPageHeadings()
	p.getLinks()
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
		return fmt.Errorf("get status: %d from smapp batch routing", resp.StatusCode)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	p.content = string(data)
	return nil
}

func (p *Page) getHTMLVersion() {

}

func (p *Page) getPageTitleVersion() {
	startIndex := strings.Index(p.content, "<title>")
	endIndex := strings.Index(p.content[startIndex:], "</title>")
	fmt.Println("page title is:", p.content[startIndex+len("<title>"):startIndex+endIndex])
}

func (p *Page) getPageHeadings() {
	for i := 1; i < 5; i++ {
		heading := fmt.Sprintf("h%d", i)
		count := strings.Count(p.content, "<"+heading)
		fmt.Println("number of "+heading, ":", count)
	}
}

func (p *Page) getLinks() {
	totalLinks := strings.Count(p.content, "href=\"")
	externalLinks := strings.Count(p.content, "href=\"http")
	internalLinks := totalLinks - externalLinks
	fmt.Println("number of external links", externalLinks)
	fmt.Println("number of internal links", internalLinks)
}

func (p *Page) hasPageLoginForm(){
	startIndex := strings.Index(p.content, "<form>")
	endIndex := strings.Index(p.content[startIndex:], "</form>")
	strings.Contains(p.content[startIndex+len("<form>"):startIndex+endIndex], "login")
}
