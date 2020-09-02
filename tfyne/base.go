package tfyne

import (
	"fmt"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
	"github.com/arlsclu7/golab/db"
	"github.com/sirupsen/logrus"
)

type f1 func(i int) string

type s1 struct {
	age int
}

func (s *s1) Change() {
	s.age += 1
}

type as func(s *s1) string

func Runbase() {
	myApp := app.New()
	myWindow := myApp.NewWindow("TabContainer")
	find := db.Search(7)
	if find == nil {
		logrus.New().Info(find)
	}
	lname := fmt.Sprintf("%s", find.Name)
	nameLabel := widget.NewLabel("Name: " + lname)
	sexLabel := widget.NewLabel("Sex: male")
	ageLabel := widget.NewLabel("Age: 18")
	addressLabel := widget.NewLabel("Province: shanghai")
	addressLabel.Hide()
	profile := widget.NewVBox(nameLabel, sexLabel, ageLabel, addressLabel)

	musicRadio := widget.NewRadio([]string{"on", "off"}, func(string) {})
	showAddressCheck := widget.NewCheck("show address?", func(value bool) {
		if !value {
			addressLabel.Hide()
		} else {
			addressLabel.Show()
		}
	})
	next := widget.NewCheck("show next?", func(value bool) {
		if !value {
			widget.ShowPopUp(content fyne.CanvasObject, canvas fyne.Canvas)
		} else {
			logrus.Info(111)
		}
	})
	memberTypeSelect := widget.NewSelect([]string{"junior", "senior", "admin"}, func(string) {})

	setting := widget.NewForm(
		&widget.FormItem{Text: "music", Widget: musicRadio},
		&widget.FormItem{Text: "check", Widget: showAddressCheck},
		&widget.FormItem{Text: "member type", Widget: memberTypeSelect},
		&widget.FormItem{Text: "next", Widget: next},
	)

	tabs := widget.NewTabContainer(
		widget.NewTabItem("Profile", profile),
		widget.NewTabItem("Setting", setting),
	)

	myWindow.SetContent(tabs)
	myWindow.Resize(fyne.NewSize(200, 200))
	myWindow.ShowAndRun()

}
