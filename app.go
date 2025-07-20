package main

import (
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	w := myApp.NewWindow("Profit Calculator")

	revenueEntry := widget.NewEntry()
	revenueEntry.SetPlaceHolder("Enter Revenue in cents (e.g. 10000000) here...")

	expensesEntry := widget.NewEntry()
	expensesEntry.SetPlaceHolder("Enter Expenses in cents (e.g. 5000000) here...")

	taxRateEntry := widget.NewEntry()
	taxRateEntry.SetPlaceHolder("Enter Tax rate (%) (e.g. 19.5) here...")

	resultLabel := widget.NewLabel("")

	calcBtn := widget.NewButton("Calculate", func() {
		revenue, err1 := strconv.Atoi(revenueEntry.Text)
		expenses, err2 := strconv.Atoi(expensesEntry.Text)
		taxRate, err3 := strconv.ParseFloat(taxRateEntry.Text, 64)

		if err1 != nil || err2 != nil || err3 != nil {
			resultLabel.SetText("Please enter valid numbers for all fields.")
			return
		}

		result := CalculateProfit(revenue, expenses, taxRate)
		resultLabel.SetText(FormatResult(result))
	})

	resetBtn := widget.NewButton("Reset", func() {
		revenueEntry.SetText("")
		expensesEntry.SetText("")
		taxRateEntry.SetText("")
		resultLabel.SetText("")
	})

	form := container.NewVBox(
		widget.NewLabel("Profit Calculator"),
		revenueEntry,
		expensesEntry,
		taxRateEntry,
		container.NewHBox(calcBtn, resetBtn),
		resultLabel,
	)

	w.SetContent(form)
	w.Resize(fyne.NewSize(400, 300))
	w.ShowAndRun()
}
