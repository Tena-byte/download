package handler

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
)

func Home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	fmt.Fprint(w, `
	<!DOCTYPE html>
	<html>
	<body>
		<form id="form" action="/download">
        	<input type="text" name="imageUrl" id="imageId" placeholder="enter image link">
        	<button type="submit" >Download Image</button>
    	</form>

		<script>
			document.getElementById('form').addEventListener('submit', function() {
				setTimeout(() => {
					document.getElementById('imageId').value = '';
				}, 100);
			});
		</script>
	</body>
	</html>
	`)
}

func ImageDownload(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Form not valid", 405)
		return
	}
	imageUrl := r.Form.Get("imageUrl")

	if imageUrl == "" {
		http.Error(w, "Image Url can not be empty", 400)
		return
	}

	parsedURL, err := url.Parse(imageUrl)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	imageName := path.Base(parsedURL.Path)

	if imageName == "" || imageName == "/" || imageName == "." {
		imageName = "downloadedImage.jpg"
	}

	resp, err := http.Get(imageUrl)
	ctype := resp.Header.Get("Content-Type")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	w.Header().Set("Content-Type", ctype)
	w.Header().Set("Content-Disposition", "attachment; filename=\""+imageName+"\"")

	io.Copy(w, resp.Body)

}
