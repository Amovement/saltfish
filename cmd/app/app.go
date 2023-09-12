package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"fyne.io/systray"
	"github.com/elpsyr/saltfish/internal/job"
	"github.com/elpsyr/saltfish/pkg/win"
	"log"
	"os"
	"time"
)

var rewardCount int
var fishingCount int

func main() {
	myApp := app.New()
	w := myApp.NewWindow("salt fish")
	//resource, err := fyne.LoadResourceFromPath("./images/fish.svg")
	//if err != nil {
	//	fmt.Println("LoadResourceFromPath ERROR:", err)
	//}
	//w.SetIcon(resource)
	w.Resize(fyne.Size{
		Width: 300,
	})

	// 点击关闭进行隐藏
	w.SetCloseIntercept(func() {
		w.Hide()
	})

	manager := &job.Manager{}

	workLabel := widget.NewLabel(fmt.Sprintf("Rewards : %d Fishing : %d", rewardCount, fishingCount))
	timeLabel := widget.NewLabel("Run Time : 00:00:00")
	w.SetContent(container.NewVBox(
		widget.NewButton("hide", func() {
			//hello.SetText("🦈️")
			hwnd := win.GetHwndByTitle("咸鱼之王")
			win.SetTopWindow(hwnd)
			win.HideWindow(hwnd)
		}),
		widget.NewButton("show", func() {
			//hello.SetText("🐟")
			hwnd := win.GetHwndByTitle("咸鱼之王")
			win.ShowWindow(hwnd)
			win.SetTopWindow(hwnd)
		}),
		widget.NewButton("reward", func() {
			//hello.SetText("🪙")
			hwnd := win.GetHwndByTitle("咸鱼之王")
			go manager.GetReward(hwnd)
		}),
		widget.NewButton("fishing", func() {
			hwnd := win.GetHwndByTitle("咸鱼之王")
			go manager.GetFish(hwnd)
		}),
		container.New(layout.NewHBoxLayout(), layout.NewSpacer(), workLabel, layout.NewSpacer()),
		container.New(layout.NewHBoxLayout(), layout.NewSpacer(), timeLabel, layout.NewSpacer()),
	))

	menu := fyne.NewMenu("MyApp",
		fyne.NewMenuItem("Show", func() {
			log.Println("Tapped show")
			w.Show()
		}))

	if desk, ok := myApp.(desktop.App); ok {
		//resourceIco, err := fyne.LoadResourceFromPath("./images/fish.ico")
		//if err != nil {
		//	fmt.Println("LoadResourceFromPath ERROR:", err)
		//}
		//desk.SetSystemTrayIcon(resourceIco)
		desk.SetSystemTrayMenu(menu)
	}

	go updateTimeLabel(timeLabel)
	GetReward2Hour(manager, workLabel) // 注册
	GetFish8Hour(manager, workLabel)   // 注册

	//systray.Run(onReady, onExit)
	w.ShowAndRun()
}

func onReady() {

	// 使用 ioutil.ReadFile 读取图片文件内容
	imageBytes, err := os.ReadFile("./images/fish.ico")
	if err != nil {
		fmt.Println("无法读取图片文件:", err)
		return
	}
	systray.SetIcon(imageBytes)
	systray.SetTitle("Awesome App")
	systray.SetTooltip("Pretty awesome超级棒")
	mQuit := systray.AddMenuItem("Quit", "Quit the whole app")

	// Sets the icon of a menu item.
	mQuit.SetIcon(imageBytes)
}

func onExit() {
	// clean up here
}

func elapsedTime(startTime time.Time) string {
	elapsedTime := time.Since(startTime)
	hours := int(elapsedTime.Hours())
	minutes := int(elapsedTime.Minutes()) % 60
	seconds := int(elapsedTime.Seconds()) % 60
	return fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds)
}

func updateTimeLabel(label *widget.Label) {
	startTime := time.Now()

	for {
		currentElapsedTime := elapsedTime(startTime)
		time.Sleep(time.Second)
		label.SetText("Run Time : " + currentElapsedTime)
	}
}

func GetReward2Hour(m *job.Manager, label *widget.Label) {
	// 创建一个每隔2小时触发一次的Ticker
	ticker := time.NewTicker(2 * time.Hour)
	//ticker := time.NewTicker(30 * time.Second)

	// 启动一个goroutine来处理Ticker触发的事件
	go func() {
		for {
			select {
			case <-ticker.C:
				// 在Ticker触发时调用方法A
				hwnd := win.GetHwndByTitle("咸鱼之王")
				m.GetReward(hwnd)
				rewardCount++
				label.SetText(fmt.Sprintf("Rewards : %d Fishing : %d", rewardCount, fishingCount))
			}
		}
	}()

}

func GetFish8Hour(m *job.Manager, label *widget.Label) {
	// 创建一个每隔8小时触发一次的Ticker
	ticker := time.NewTicker(8*time.Hour + time.Minute)

	// 启动一个goroutine来处理Ticker触发的事件
	go func() {
		for {
			select {
			case <-ticker.C:
				// 在Ticker触发时调用方法A
				hwnd := win.GetHwndByTitle("咸鱼之王")
				m.GetReward(hwnd)
				rewardCount++
				label.SetText(fmt.Sprintf("Rewards : %d Fishing : %d", rewardCount, fishingCount))
			}
		}
	}()

}
