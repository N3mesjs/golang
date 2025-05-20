package main 

import (
	"fmt"
	"github.com/go-vgo/robotgo"
	"github.com/vcaesar/imgo"
	"os/exec"
)

func main() {
	//for {
	//	x, y := robotgo.Location()
	//	fmt.Println("x: ", x, " y: ", y)
	//}
	purn()
}


func purn() {
	robotgo.Move(100, 100)
	fmt.Println(robotgo.Location())
	robotgo.Move(100, 200) // multi screen supported
	robotgo.MoveSmooth(120, 150)
	fmt.Println(robotgo.Location())

	color := robotgo.GetPixelColor(100, 200)
  	fmt.Println("color---- ", color)

	sx, sy := robotgo.GetScreenSize()
	fmt.Println("get screen size: ", sx, sy)

	bit := robotgo.CaptureScreen(0, 0, 2560, 1440)
  	defer robotgo.FreeBitmap(bit)
	img := robotgo.ToImage(bit)
	imgo.Save("test.png", img)

	exec.Command("cmd", "/C", "start", "https://").Run()
	robotgo.Sleep(1)
	robotgo.TypeStr("pornhub")
	robotgo.KeyTap("enter")
	robotgo.Move(720, 333)
	robotgo.Sleep(1)
	robotgo.Click("left")
	robotgo.Click("left")
}