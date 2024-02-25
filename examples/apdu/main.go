package main

import (
	"log"
	"log/slog"

	"github.com/spilikin/go-card/smartcard"
)

func main() {
	ctx, err := smartcard.EstablishContext()
	if err != nil {
		log.Fatal(err)
	}
	defer ctx.Release()

	reader, err := ctx.WaitForCardPresent()
	if err != nil {
		log.Fatal(err)
	}
	slog.Info("Reader created", "reader", reader.Name())

	card, err := reader.Connect()
	if err != nil {
		log.Fatal(err)
	}

	defer card.Disconnect()

	slog.Info("Card connected", "atr", card.ATR().String())

	efv, err := card.EFVersion2()
	if err != nil {
		log.Fatal(err)
	}
	slog.Info("EF Version 2", "efv", efv)

	efdir, err := card.EFDIR()
	if err != nil {
		log.Fatal(err)
	}
	slog.Info("EF.DIR", "EF.DIR", efdir)

	err = card.SelectDF(smartcard.DF_ESIGN)
	if err != nil {
		log.Fatal(err)
	}

	cert, err := card.ReadCertificate(smartcard.EF_C_CH_AUT_E256)
	if err != nil {
		log.Fatal(err)
	}
	slog.Info("EF.C.CH.AUT.E256", "subject", cert.Subject.CommonName, "notAfter", cert.NotAfter)
}
