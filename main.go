package main

import (
	"github.com/signintech/gopdf"
	"time"
)

const(
	light = "Light"
	regular = "Regular"
	medium = "Medium"
	bold = "Bold"
	black = "Black"
)

const (
	short= "02-01-2006"
	long= "2 January, 2021"
)

type invoice struct {
	InvoiceId string
	Name string
	StreetAndNumber string
	Zip string
	City string
	Country string
	Date string
	DueDate string
}

func checkErr(err error){
	if err != nil	{
		panic(err)
	}
}

func main(){
	getInvoice() // should be called from rest api
}

func getInvoice(){
	// get invoice details in here and pass them
	date := time.Now().Format(short)

	// setup data
	inv := invoice{
		InvoiceId: "0001",
		Name:      "Jari Dewulf",
		StreetAndNumber:    "Deinzestraat 40",
		Zip:       "9800",
		City:      "Deinze",
		Country:   "Belgium",
		Date:       date,
		DueDate: 	date,
	}
	createInvoice(inv)

	// send back invoice in stream.
}

func createInvoice(inv invoice) {

	mm6ToPx := 80.6
	pdf := gopdf.GoPdf{}

	pdf.Start(gopdf.Config{
		PageSize: *gopdf.PageSizeA4,
		TrimBox: gopdf.Box{Left: mm6ToPx, Top: mm6ToPx, Right: 595-mm6ToPx, Bottom: 842-mm6ToPx},
	})
	opt := gopdf.PageOption{
		PageSize: gopdf.PageSizeA4,
		TrimBox: &gopdf.Box{Left: mm6ToPx, Top: mm6ToPx, Right: 595-mm6ToPx, Bottom: 842-mm6ToPx},
	}
	pdf.AddPageWithOption(opt)

	// add fonts
	fonts := []string{light,medium, regular, bold, black}
	for _, font := range fonts {
		err := pdf.AddTTFFont(font, "./assets/fonts/Poppins-"+font+".ttf")
		checkErr(err)
	}

	setDefaultFont(&pdf)
	pdf.Rotate(60,350, 400)
	pdf.Image("./assets/images/icon/csmm-light.png", 0,0, nil)
	pdf.RotateReset()

	printLogo(&pdf,70,95)
	printSellerAddress(&pdf)
	printBuyerAddress(&pdf, inv)

	// invoice title
	pdf.SetFont(medium, "", 22)
	pdf.SetTextColor(58,71,99)
	pdf.SetX(70)
	pdf.SetY(300)
	pdf.Text("Invoice")

	//########################
	// invoice details		##
	//########################
	// titles
	pdf.SetFont(medium, "", 10)
	pdf.SetX(70); pdf.SetY(325); pdf.Text("INVOICE ID")
	pdf.SetX(200); pdf.Text("INVOICE DATE")
	pdf.SetX(330); pdf.Text("DUE DATE")

	setDefaultFont(&pdf)

	pdf.SetX(70); pdf.SetY(348); pdf.Text("#"+ inv.InvoiceId)
	pdf.SetX(200); pdf.Text(inv.Date)
	pdf.SetX(330); pdf.Text(inv.DueDate)

	/*
	##########################
	# 	   DRAW TABLE 	 	 #
	##########################
	*/

	// header
	pdf.SetFont(medium,"",10)
	pdf.SetStrokeColor(0,205,106); pdf.SetFillColor(0,205,106); pdf.SetTextColor(255,255,255)
	pdf.Polygon([]gopdf.Point{{70,405}, {530,405}, { 530, 425}, {70, 425}}, "DF") // background
	pdf.SetX(80); pdf.SetY(417.5); pdf.Text("Description of product or service")
	pdf.SetX(320); pdf.Text("VAT excl.")
	pdf.SetX(480); pdf.Text("Total")

	// body
	setDefaultFont(&pdf)
	pdf.SetFont(light, "",10)
	pdf.SetX(80); pdf.SetY(442.5); pdf.Text("CSMM 10")
	pdf.SetX(320); pdf.SetY(442.5); pdf.Text("5 EUR")
	pdf.SetX(470); pdf.SetY(442.5); pdf.Text("6.05 EUR")
	// this could possibly loop a few times.

	// grid
	pdf.SetLineWidth(1)

	for i:=455.;i<(455+5*30);i+=30 {
		pdf.Line(70,i, 530,i)
	}
	// vertical lines
	pdf.Line(70,425,70, 575)
	pdf.Line(450,425,450,575)
	pdf.Line(530,425,530,575)

	pdf.SetFont(light,"",8)
	pdf.SetX(70);pdf.SetY(650); pdf.Text("Unless otherwise agreed in writing, all actions by or with this company are subject to the general terms and ")
	pdf.SetX(70); pdf.SetY(665); pdf.Text("conditions. The customer declares that he is aware of these provisions and accepts them without reserveration. ")
	pdf.SetX(70); pdf.SetY(680);pdf.Text("The general terms and conditions can be found on www.csmm.app."); pdf.AddExternalLink("https://csmm.app/", 275,670,70,15)
	pdf.SetLineWidth(1.5); pdf.SetStrokeColor(0,205, 106); pdf.Line(275,683,340,683)

	printFooter(&pdf)
	// write article to file / writeStream
	pdf.WritePdf("result.pdf")

}

// helper func
func setDefaultFont(pdf *gopdf.GoPdf){
	pdf.SetFont("Regular", "", 10)
	pdf.SetTextColor(150,150,150)
}

func printLogo(pdf *gopdf.GoPdf, x,y float64){
	pdf.Image("./assets/images/icon/normal.png", x, y, nil)
	pdf.SetFont(bold,"",30)
	pdf.SetTextColor(58,71,99)
	pdf.SetX(x+80); pdf.SetY(y+50); pdf.Text("CSMM")
	setDefaultFont(pdf)
}

func printBuyerAddress(pdf *gopdf.GoPdf, inv invoice){
	pdf.SetFont(regular,"",9)
	pdf.SetX(450); pdf.SetY(200); pdf.Text(inv.Name)
	pdf.SetX(450); pdf.SetY(213); pdf.Text(inv.StreetAndNumber)
	pdf.SetX(450); pdf.SetY(226); pdf.Text(inv.Zip + " " + inv.City)
	pdf.SetX(450); pdf.SetY(239); pdf.Text(inv.Country)
	setDefaultFont(pdf)
}

func printSellerAddress(pdf *gopdf.GoPdf){
	pdf.SetFont(regular,"",9); setDefaultFont(pdf)
}

func printFooter(pdf *gopdf.GoPdf){
	pdf.Image("./assets/images/icon/mini.png", 70, 760, nil)

	// smaller font
	pdf.SetFont(light, "", 7)
	pdf.SetX(70); pdf.SetY(800); pdf.Text("www.csmm.app")
	pdf.SetX(470); pdf.SetY(800); pdf.Text("BE 0439.200.175")

	pdf.SetX(70); pdf.SetY(810); pdf.Text("CSMM Billing")
	pdf.SetX(470); pdf.SetY(810); pdf.Text("RPR Gent")

	setDefaultFont(pdf)
}
