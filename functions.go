package main

import (
	"image/color"
	"net/url"
	"os"
	"os/exec"
	"sort"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
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

func OsInfo(osWin fyne.Window) {
	hostname, _ := os.Hostname()

	ip, ipErr := exec.Command("powershell", "-Command", "Get-NetIPAddress -AddressFamily IPv4 | Select -ExpandProperty IPAddress").Output()

	ipLabel := widget.NewLabel("Direcciones IP")
	var ipInfo *canvas.Text
	if ipErr != nil {
		ipInfo = canvas.NewText("Error al mostrar IP", color.Black)
	} else {
		actualIPs := strings.TrimSpace(string(ip))
		actualIPs = strings.ReplaceAll(actualIPs, "\r\n", " | ")
		ipInfo = canvas.NewText(actualIPs, color.Black)
	}

	envs := os.Environ()
	sort.Strings(envs)
	rt := widget.NewRichTextWithText(strings.Join(envs, "\n"))
	rt.Wrapping = fyne.TextWrapWord

	envsLabel := widget.NewLabel("Variables de entorno")
	envsInfo := container.NewScroll(rt)
	envsInfo.SetMinSize(fyne.NewSize(600, 300))

	gap := canvas.NewText("", color.Transparent)
	gap.TextSize = 10

	hostnameLabel := widget.NewLabel("Hostname")
	hostnameInfo := canvas.NewText(hostname, color.Black)

	box1 := container.NewVBox(hostnameLabel, hostnameInfo)
	box2 := container.NewVBox(ipLabel, ipInfo)
	box3 := container.NewVBox(envsLabel, envsInfo)
	content := container.NewVBox(box1, gap, box2, gap, box3)

	osWin.SetContent(content)
	osWin.Show()
}
