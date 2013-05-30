package handlers

import (
	"html/template"
	"net/http"
	"os"
)

var (
	browseTemplates = template.Must(template.ParseFiles(
		"templates/browse_downloads.html",
		"templates/browse_uploads.html",
		"templates/footer.html",
		"templates/header.html",
	))
)

func BrowseDownloads(w http.ResponseWriter, r *http.Request) {
	files := make([][]string, 0)

	d, err := os.Open("downloads/")
	if err != nil {
		panic(err)
	}

	fi, err := d.Readdir(-1)
	if err != nil {
		panic(err)
	}

	for i, fi := range fi {
		row := i / 6
		col := i % 6

		if 0 == col {
			files = append(files, make([]string, 0))
		}

		fn := fi.Name()
		fl := len(fn)

		if 44 > fl {
			continue
		} // sha1 (40 chars) + .png + x
		if ".png" != fn[fl-4:] {
			continue
		}

		files[row] = append(files[row], fn[:fl-4])
	}

	page := new(BrowsePage)
	page.NavDownloads = true
	page.PageTitle = "Browse Downloads"

	if 0 < len(files) {
		page.Files = files
	}

	browseTemplates.ExecuteTemplate(w, "browse_downloads.html", page)
}

func BrowseUploads(w http.ResponseWriter, r *http.Request) {
	files := make([][]string, 0)

	d, err := os.Open("uploads/")
	if err != nil {
		panic(err)
	}

	fi, err := d.Readdir(-1)
	if err != nil {
		panic(err)
	}

	for i, fi := range fi {
		row := i / 6
		col := i % 6

		if 0 == col {
			files = append(files, make([]string, 0))
		}

		fn := fi.Name()
		fl := len(fn)

		if 44 > fl {
			continue
		} // sha1 (40 chars) + .png + x
		if ".png" != fn[fl-4:] {
			continue
		}

		files[row] = append(files[row], fn[:fl-4])
	}

	page := new(BrowsePage)
	page.NavUploads = true
	page.PageTitle = "Browse Uploads"

	if 0 < len(files) {
		page.Files = files
	}

	browseTemplates.ExecuteTemplate(w, "browse_uploads.html", page)
}
