GoZBar
---
Fork of https://github.com/PeterCxy/gozbar due to API breaking changes.

ZBar bindings for golang. Only scanner supported.

Read the ZBar documentations for explanations on constants and arguments.

The ZBar library must be installed for this to compile.

Example JPEG Scan.
```go
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
```
