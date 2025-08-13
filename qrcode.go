/*
generate qr codes linking to submission form
*/

/*
GenerateQRCode() -> returns png image bytes for qr 

sends data to:
	admin tool
*/

package main 

import (
	"github.com/skip2/go-qrcode"
)

func GenerateQRCode(url string) ( error){ 
	err := qrcode.WriteFile(url, qrcode.Medium, 256, "qrcode.png")
	return err 
}
