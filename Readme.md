# ARCHIVED

> **This project has been archived.** This experiment has been replaced by [glu](https://github.com/boxesandglue/glu), which provides a more complete command-line interface for boxes and glue.

---

# Command line for boxes and glue (archived)

Experimental and incomplete

## Installation

When you have a Go installation: `go install github.com/boxesandglue/cli/bag@latest`

You can also download the latest binaries at https://github.com/boxesandglue/cli/releases/latest

## Sample document

Start with `bag myfile.rsr`, where the file `myfile.rsr` looks like this:

```go
now := time.now()

func setup_fonts(f) {
    ff := f.new_fontfamily("text")

    fs := frontend.new_fontsource({
        location: filepath.join("fonts","CrimsonPro-Regular.ttf"),
        features: ["kern","liga"],
        })

    ff.add_member({source: fs, weight: 400, style: "normal"})
    return ff
}

str := `The quick brown fox jumps
over the lazy dog with a very long line that should be wrapped
at some point. This is a test to see how the text is formatted
when it is too long to fit on one line. The quick brown fox jumps
over the lazy dog with a very long line that should be wrapped`

str = strings.join(strings.fields(str)," ")

f := frontend.new("out.pdf")
backend_doc := f.doc
backend_doc.language = frontend.get_language("en")

backend_doc.title = "A test document"

ff := setup_fonts(f)

p := f.doc.new_page()
para := frontend.new_text()
para.items = [str]

vlist := f.format_paragraph({
	text: para,
	width: bag.sp("225pt"),
	leading: bag.sp("14pt"),
	font_size: bag.sp("12pt"),
	family: ff,
})
p.output_at(bag.sp("1cm"), bag.sp("10cm"), vlist)
p.shipout()

f.doc.finish()

printf("finished in %.2fms\n",time.since(now) * 1000)
```

## Other

Contact: gundlach@speedata.de<br>
License: BSD-3 clause<br>
Status: **archived** – superseded by [glu](https://github.com/boxesandglue/glu)<br>
Legal info: boxes and glue is a registered trademark
