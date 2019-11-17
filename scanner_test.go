package gozbar

import (
	"image/jpeg"
	"image/png"
	"os"
	"testing"
)

func TestBarcode(t *testing.T) {
	const expectedVal = "9876543210128"

	f, err := os.Open("testdata/barcode.png")
	if err != nil {
		t.Fatal(err)
	}

	i, err := png.Decode(f)
	if err != nil {
		t.Fatal(err)
	}

	img := FromImage(i)

	s := NewScanner()
	err = s.SetConfig(0, CFG_ENABLE, 1)
	if err != nil {
		t.Fatal("error setting config", err)
	}

	err = s.Scan(img)
	if err != nil {
		t.Fatal("error scanning", err)
	}
	defer s.Destroy()

	img.First().Each(func(str string) {
		if str != expectedVal {
			t.Fatalf("expected [%s] got [%s]", expectedVal, str)
		}
	})
}

func TestQRCode(t *testing.T) {
	const expectedVal = "ZBar big law good! ZBar螟ｧ豕募･ｽ!"

	f, err := os.Open("testdata/qr.jpg")
	if err != nil {
		t.Fail()
		return
	}

	i, _ := jpeg.Decode(f)
	img := FromImage(i)

	s := NewScanner()
	err = s.SetConfig(0, CFG_ENABLE, 1)
	if err != nil {
		t.Fatal("error setting config", err)
	}

	err = s.Scan(img)
	if err != nil {
		t.Fatal("error scanning", err)
	}
	defer s.Destroy()

	img.First().Each(func(str string) {
		if str != expectedVal {
			t.Fatalf("expected [%s] got [%s]", expectedVal, str)
		}
	})
}

func TestPhoto(t *testing.T) {
	const expectedStr = "http://www.searchenginestrategies.com/sanfrancisco/share.html"

	f, err := os.Open("./testdata/photo.jpg")
	if err != nil {
		t.Fatal(err)
	}

	i, err := jpeg.Decode(f)
	if err != nil {
		t.Fatal(err)
	}

	img := FromImage(i)

	s := NewScanner()
	err = s.SetConfig(0, CFG_ENABLE, 1)
	if err != nil {
		t.Fatal("error setting config", err)
	}

	err = s.Scan(img)
	if err != nil {
		t.Fatal("error scanning", err)
	}

	defer s.Destroy()

	img.First().Each(func(str string) {
		if str != expectedStr {
			t.Fatalf("expected [%s] got [%s]", expectedStr, str)
		}
	})
}
