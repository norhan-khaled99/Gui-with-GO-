/*

package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello")

	hello := widget.NewLabel("Hello Fyne!")
	w.SetContent(container.NewVBox(
		hello,
		widget.NewButton("Hi!", func() {
			hello.SetText("Welcome :)")
		}),
	))

	w.ShowAndRun()
}
*/
/*
package main

import (
	"fyne.io/fyne/v2/theme"
	"log"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
	//"fyne.io/fyne/v2/theme"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Button Widget")

	content := widget.NewButton("click me", func() {
		log.Println("tapped")
	})

	content = widget.NewButtonWithIcon("Home", theme.HomeIcon(), func() {
		log.Println("tapped home")
	})
	content = widget.NewButtonWithIcon("5465", theme.HomeIcon(), func() {
		log.Println("tapped home")
	})
	content = widget.NewButtonWithIcon("1233", theme.HomeIcon(), func() {
		log.Println("tapped home")
	})

	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}
*/
/*
package main

// import fyne
import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	// New app
	a := app.New()
	// New Window & title
	w := a.NewWindow("Resize any widget")
	//Resize main/parent window
	w.Resize(fyne.NewSize(400, 400))
	// 1st Widget Entry
	entry := widget.NewEntry()
	entry.SetPlaceHolder("Enter your name")
	entry.Resize(fyne.NewSize(250, 30)) // my widget size
	entry.Move(fyne.NewPos(40, 50))     // position of widget
	entry_email := widget.NewEntry()
	entry_email.SetPlaceHolder("Enter your email")
	entry_email.Resize(fyne.NewSize(250, 30)) // my widget size
	entry_email.Move(fyne.NewPos(40, 100))    // position of widget
	entry_address := widget.NewEntry()
	entry_address.SetPlaceHolder("Enter your name")
	entry_address.Resize(fyne.NewSize(250, 30)) // my widget size
	entry_address.Move(fyne.NewPos(40, 150))    // position of widget
	// button
	btn_submit := widget.NewButton("Submit", nil)
	btn_submit.Resize(fyne.NewSize(150, 30)) // my widget size
	btn_submit.Move(fyne.NewPos(40, 200))    // position of widget
	w.SetContent(
		container.NewWithoutLayout(
			entry,
			entry_email,
			entry_address,
			btn_submit,
		),
	)
	//show and run
	w.ShowAndRun()
}
*/
/*
package main

// import fyne
import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	// New app
	a := app.New()
	// new title and window
	w := a.NewWindow("Table widget")
	// resize
	w.Resize(fyne.NewSize(400, 400))
	// create table widget
	// table is like list , just 2 values, instead of one
	table := widget.NewTable(
		// row and col . I want to create 3 row , 3 col
		func() (int, int) { return 3, 3 },
		// Now I want to specify widget. like label, checkbox
		func() fyne.CanvasObject { return widget.NewLabel("....") },
		// update/details data in widget
		func(i widget.TableCellID, obj fyne.CanvasObject) {
			// remember it is label and not newlabel
			// i is for index
			// i.col will set col value
			// i.row will present row value
			obj.(*widget.Label).SetText(fmt.Sprintf("%d %d", i.Col, i.Row))
		},
	)
	w.SetContent(table)
	w.ShowAndRun()
}
*/
/*
package main

// import fyne
import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	// window and title
	w := a.NewWindow("List View")
	// resize
	w.Resize(fyne.NewSize(400, 400))
	// list View
	// 3 arguments
	// item count, 3 itmes in list
	// widget, I want to use label widget
	// update data of widget
	list := widget.NewList(
		// lets change item count from 3 to 30
		func() int { return 30 }, // my list contain 3 items
		func() fyne.CanvasObject { return widget.NewLabel("I want to use label") },
		// last one
		func(lii widget.ListItemID, co fyne.CanvasObject) {
			// update data of widget
			co.(*widget.Label).SetText("Here is my text")
		},
	)
	// setup data on screen
	w.SetContent(list)
	w.ShowAndRun()
}
*/
/*
package main

// import fyne
import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	// new app
	a := app.New()
	// new title and window
	w := a.NewWindow("Select Entry - DropDown advance")
	// resize window
	w.Resize(fyne.NewSize(400, 400))
	// lets create a label
	label1 := widget.NewLabel("...")
	// newselectentry widget
	// []string{} it take only slice of option
	//& you can your options also in run time. Not hardcoded
	select_entry := widget.NewSelectEntry([]string{"peshawar", "multan", "gujrat", "kabul", "dehli"})
	// what to do with the selected entry ?
	// here is what we are going to define
	select_entry.OnSubmitted = func(s string) {
		fmt.Printf("my city %s is awesome", s)
		// update label with our values
		label1.Text = s
		label1.Refresh()
	}
	// container .. we have more than one widgets
	c := container.NewVBox(
		label1,
		select_entry,
	)
	w.SetContent(c)
	w.ShowAndRun()
}
*/
/*
package main

// import fyne
import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	// new app
	a := app.New()
	// new title and window
	w := a.NewWindow("Select Entry - DropDown advance")
	// resize window
	w.Resize(fyne.NewSize(400, 400))
	// lets create a label
	label1 := widget.NewLabel("...")
	// newselectentry widget
	// []string{} it take only slice of option
	//& you can your options also in run time. Not hardcoded
	select_entry := widget.NewSelectEntry([]string{"peshawar", "multan", "gujrat", "kabul", "dehli"})
	// what to do with the selected entry ?
	// here is what we are going to define
	select_entry.OnSubmitted = func(s string) {
		fmt.Printf("my city %s is awesome", s)
		// update label with our values
		label1.Text = s
		label1.Refresh()
	}
	// container .. we have more than one widgets
	c := container.NewVBox(
		label1,
		select_entry,
	)
	w.SetContent(c)
	w.ShowAndRun()
}
*/
/*
package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"image/color"
)

func main() {
	a := app.New()
	w := a.NewWindow("Colored Button")
	//resize window
	w.Resize(fyne.NewSize(400, 400))
	// Let show out button on screen
	w.SetContent(
		container.NewVBox(
			red_button(),
			red_button(),
			red_button(),
			red_button(),
			// let show / display buttons
			blue_button(),
			green_button(),
			img_button(),
		),
	)
	w.ShowAndRun() // show and run app
}

// first colored button
func red_button() *fyne.Container { // return type
	btn := widget.NewButton("Visit", nil) // button widget
	// button color
	btn_color := canvas.NewRectangle(
		color.NRGBA{R: 255, G: 0, B: 0, A: 255})
	// container for colored button
	container1 := container.New(
		// layout of container
		layout.NewMaxLayout(),
		// first use btn color
		btn_color,
		// 2nd btn widget
		btn,
	)
	// our button is ready
	return container1
}

// Blue button
// copy previous button code and change it
func blue_button() *fyne.Container { // return type
	btn := widget.NewButton("Visit", nil) // button widget
	// button color
	btn_color := canvas.NewRectangle(
		color.NRGBA{R: 05, G: 0, B: 255, A: 255})
	// container for colored button
	container1 := container.New(
		// layout of container
		layout.NewMaxLayout(),
		// first use btn color
		btn_color,
		// 2nd btn widget
		btn,
	)
	// our button is ready
	return container1
}

// Green button
// copy and paste previous button
func green_button() *fyne.Container { // return type
	btn := widget.NewButton("Visit", nil) // button widget
	// button color
	btn_color := canvas.NewRectangle(
		color.NRGBA{R: 0, G: 255, B: 0, A: 255})
	// container for colored button
	container1 := container.New(
		// layout of container
		layout.NewMaxLayout(),
		// first use btn color
		btn_color,
		// 2nd btn widget
		btn,
	)
	// our button is ready
	return container1
}

// button with image as background
func img_button() *fyne.Container { // return type
	btn := widget.NewButton("Visit", nil) // button widget
	// img button color
	//img := canvas.NewImageFromFile("c:/cartoon/blue-bg.jpg")
	// container for colored button
	container1 := container.New(
		// layout of container
		layout.NewMaxLayout(),
		// first use btn color
		//	img,
		// 2nd btn widget
		btn,
	)
	// our button is ready
	return container1
}
*/
/*
package main

// import fyne
import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"image/color"
)

func main() {
	// New app
	a := app.New()
	// New title and window
	w := a.NewWindow("Hide and show objects")
	w.Resize(fyne.NewSize(400, 400))
	// new rectangle
	rect := canvas.NewRectangle(color.Black)
	rect.Resize(fyne.NewSize(100, 100)) // 100x100 size
	rect.Move(fyne.NewPos(50, 50))      // new position of rectangle
	// check box to show and hide rectangle or any other object
	// first argument is title of check box
	// 2nd is function
	chkbox := widget.NewCheck("hide", func(b bool) {
		if b != true {
			rect.Hide()
		} else {
			rect.Show() // show object/rectangle
		}
	})
	// size of chkbox and position
	chkbox.Move(fyne.NewPos(50, 200))
	chkbox.Resize(fyne.NewSize(100, 50))
	// setup content
	w.SetContent(
		// container without layout
		container.NewWithoutLayout(
			rect,
			chkbox,
		),
	)
	w.ShowAndRun()
}
*/
/*
package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Hello")
	myWindow.Resize(fyne.NewSize(500, 600))
	label1 := widget.NewLabel("You are in content 1")
	button1 := widget.NewButton("Button 1",
		func() {
			fmt.Println("tapped button1")
			myWindow.SetContent(content)
			myWindow.Show()
		})

	label2 := widget.NewLabel("You are in content 2")
	button2 := widget.NewButton("Button 2",
		func() {
			fmt.Println("tapped button2")
			myWindow.SetContent(content1)
			myWindow.Show()
		})
	content1 := container.NewVBox(
		label1,
		button1,
	)
	content2 := container.NewVBox(
		label2,
		button2,
	)
	myWindow.SetContent(content1)
	myWindow.ShowAndRun()
}
*/
/*
package main

import (
	// importing fyne
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	// New app
	a := app.New()
	//New Window
	w := a.NewWindow("Here is my title for 3 tutorial")
	// Resize Windows
	w.Resize(fyne.NewSize(300, 300))
	// Our First widget
	labelX := widget.NewLabel("I can write anything")
	w.SetContent(labelX)
	w.ShowAndRun() // Running app
}
*/
/*
package main

// importing fyne
import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	// creating new app
	a := app.New()
	// New window
	w := a.NewWindow("Here is my title for 4th tutorial")
	// resizing my window
	w.Resize(fyne.NewSize(400, 400))
	// Now Its Time to use BUTTONs
	// First value is button name
	// 2nd value is any function
	btn := widget.NewButton("button Name", func() {
		// Our is ready
		fmt.Println("Here is Go Button")
	})
	// using our button on
	w.SetContent(btn)
	// Running app
	w.ShowAndRun()
}
*/
/*
package main

// import fyne
import (
	"encoding/json"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"io/ioutil"
)

func main() {
	// student struct, you can use any name
	type Student struct {
		Name  string // name N is capital
		Phone string
	}
	// Now creat a slice/ array to store data
	var myStudentData []Student
	// welcome again
	//lets read data from file and display it in list
	data_from_file, _ := ioutil.ReadFile("myFile_name.txt")
	// file name with extension .txt or .json
	// unmarshall/parse data received from file and save/push in slice
	// 2 argument 1. data source, 2. data slice to store data
	json.Unmarshal(data_from_file, &myStudentData)
	// now we can use this data in our list
	// lets create our list
	// new app
	a := app.New()
	// new title and window
	w := a.NewWindow("CRUD APP")
	// resize window
	w.Resize(fyne.NewSize(400, 400))
	// New label to dispaly name and phone number details
	l_name := widget.NewLabel("...")
	l_name.TextStyle = fyne.TextStyle{Bold: true}
	// for phone number
	l_phone := widget.NewLabel("...")
	// entry widget for name
	e_name := widget.NewEntry()
	e_name.SetPlaceHolder("Enter name here...")
	// entry widget for phone
	e_phone := widget.NewEntry()
	e_phone.SetPlaceHolder("Enter phone here...")
	// submit button
	submit_btn := widget.NewButton("Submit", func() {
		// logic part- store our data in json format
		// let create a struct for our data
		// Get data from entry widget and push to slice
		myData1 := &Student{
			Name:  e_name.Text, // data from name entry widget
			Phone: e_phone.Text,
		}
		// append / push data to our slice
		myStudentData = append(myStudentData, *myData1)
		// * star is very important
		// convert / parse data to json format
		// 3 arguments
		// 1st is our slice data
		// ignore 2nd
		// 3rd is identation, we are using space to indent our data
		final_data, _ := json.MarshalIndent(myStudentData, "", " ")
		// writing data to file
		// it take 3 things
		//file name .txt or .json or anything you want to use
		// data source, final_data in our case
		// writing/reading/execute permission
		ioutil.WriteFile("myFile_name.txt", final_data, 0644)
		/// we are done 🙂
		// lets test
		// empty input field after data insertion
		e_name.Text = ""
		e_phone.Text = ""
		// refresh name & phone entry object
		e_name.Refresh()
		e_phone.Refresh()
	})
	// list widget
	list := widget.NewList(
		// first argument is item count
		// len() function to get myStudentData slice len
		func() int { return len(myStudentData) },
		// 2nd argument is for widget choice. I want to use label
		func() fyne.CanvasObject { return widget.NewLabel("") },
		// 3rd argument is to update widget with our new data
		func(lii widget.ListItemID, co fyne.CanvasObject) {
			co.(*widget.Label).SetText(myStudentData[lii].Name)
		},
	)
	// update on clicked/ selected
	list.OnSelected = func(id widget.ListItemID) {
		l_name.Text = myStudentData[id].Name
		l_name.Refresh()
		// for phone number
		l_phone.Text = myStudentData[id].Phone
		l_phone.Refresh()
	}
	// show and run
	w.SetContent(
		// lets create Hsplite container
		container.NewHSplit(
			// first argument is list of data
			list,
			// 2nd is
			// vbox container
			// show both label
			container.NewVBox(l_name, l_phone, e_name, e_phone, submit_btn),
		),
	)
	w.ShowAndRun()
}
*/

package main

import (
	"log"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Entry Widget")

	input := widget.NewEntry()
	input.SetPlaceHolder("Enter text...")

	content := container.NewVBox(input, widget.NewButton("Save", func() {
		log.Println("Content was:", input.Text)
	}))

	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}
