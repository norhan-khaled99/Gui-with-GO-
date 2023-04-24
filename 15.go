package main

import (
	"database/sql"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	_ "github.com/go-sql-driver/mysql"
	"image/color"
)

type database struct {
	db string
}

var result string
var arr []string

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Connection")
	myWindow.Resize(fyne.NewSize(400, 300))
	myWindow.CenterOnScreen()

	lab4 := widget.NewLabel("please enter query to run")
	xx := myApp.NewWindow("asds")
	xx.Resize(fyne.NewSize(200, 200))
	//btn 1
	text := canvas.NewText("connected successfully", color.White)
	text.Alignment = fyne.TextAlignCenter
	text.TextStyle = fyne.TextStyle{Italic: true}
	xx.SetContent(text)

	yy := myApp.NewWindow("nbh")
	yy.Resize(fyne.NewSize(300, 200))

	aa := myApp.NewWindow("cxz")
	aa.Resize(fyne.NewSize(300, 200))
	lab1 := widget.NewLabel("11111")
	input := widget.NewEntry()
	aa.SetContent(container.New(layout.NewVBoxLayout(), lab1, input, layout.NewSpacer()))

	bb := myApp.NewWindow("ds")
	bb.Resize(fyne.NewSize(300, 200))
	text = canvas.NewText("updated successfully", color.White)
	text.Alignment = fyne.TextAlignCenter
	text.TextStyle = fyne.TextStyle{Italic: true}
	bb.SetContent(text)

	cc := myApp.NewWindow("ds")
	cc.Resize(fyne.NewSize(300, 200))
	text = canvas.NewText("Deleleted successfully", color.White)
	text.Alignment = fyne.TextAlignCenter
	text.TextStyle = fyne.TextStyle{Italic: true}
	cc.SetContent(text)

	lab3 := widget.NewLabel(result)
	lab3.SetText(result)
	yy.SetContent(lab3)

	text2 := widget.NewButton("Show Tabels", func() {
		db, err := sql.Open("mysql", "root:"+"123456789"+"@tcp(127.0.0.1:3306)/hospital")

		if err != nil {
			panic(err.Error())
		}
		dd, err := db.Query("SHOW TABLES")
		for dd.Next() {
			var d1 database
			err = dd.Scan(&d1.db)
			arr = append(arr, d1.db)
		}
		for i := 0; i < len(arr); i++ {
			result += arr[i] + "\n "
		}
		lab3.SetText(result)
		lab3.Hide()
		defer db.Close()
		defer dd.Close()
		yy.Show()

	})
	lab3.Alignment = fyne.TextAlignCenter
	lab4.Alignment = fyne.TextAlignCenter
	text1 := widget.NewButton("open connection", func() {
		db, err := sql.Open("mysql", "root:"+"123456789"+"@tcp(127.0.0.1:3306)/hospital")
		if err != nil {
			panic(err.Error())
		}
		fmt.Println("connected")
		defer db.Close()
		xx.Show()
	})
	input = widget.NewEntry()

	text3 := widget.NewButton("INSERT", func() {
		//	aa.Show()
		db, err := sql.Open("mysql", "root:"+"123456789"+"@tcp(127.0.0.1:3306)/hospital")
		if err != nil {
			panic(err.Error())
		}
		insert, err := db.Query(input.Text)
		if err != nil {
			panic(err.Error())
		}
		defer insert.Close()
	})

	text4 := widget.NewButton("UPDATE", func() {
		db, err := sql.Open("mysql", "root:"+"123456789"+"@tcp(127.0.0.1:3306)/hospital")
		if err != nil {
			panic(err.Error())
		}
		stmt, err := db.Prepare(input.Text)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(stmt)
		defer db.Close()
		bb.Show()
	})
	text5 := widget.NewButton("DELETE", func() {
		db, err := sql.Open("mysql", "root:"+"123456789"+"@tcp(127.0.0.1:3306)/hospital")

		if err != nil {
			panic(err.Error())
		}
		del, err := db.Prepare(input.Text)
		fmt.Println(del)
		fmt.Println("delet done ")
		defer db.Close()
		cc.Show()
	})

	myWindow.SetContent(container.New(layout.NewCenterLayout()))
	myWindow.SetContent(container.New(layout.NewVBoxLayout(), text1, text2, lab4, input, text3, text4, text5, layout.NewSpacer()))
	myWindow.ShowAndRun()

}
