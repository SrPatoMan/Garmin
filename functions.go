package main

import (
	"net/url"
	"time"

	"fyne.io/fyne/v2"
)

func Startup(a fyne.App) {
	urls := []string{
		"https://teams.microsoft.com/v2/",
		"https://teams.microsoft.com/v2/",
		"https://chatgpt.com/",
	}

	for _, s := range urls {
		miurl, _ := url.Parse(s)
		_ = a.OpenURL(miurl)
		time.Sleep(100 * time.Millisecond)
	}
}
