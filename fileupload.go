// Upload a file to the server (typically on /tmp)
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func handler(res http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		src, hdr, err := req.FormFile("my-file")
		if err != nil {
			http.Error(res, err.Error(), 500)
			return
		}
		defer src.Close()

		dst, err := os.Create(filepath.Join(os.TempDir(), hdr.Filename))
		if err != nil {
			http.Error(res, err.Error(), 500)
			return
		}
		fmt.Println("File: ", hdr.Filename)

		defer dst.Close()
		io.Copy(dst, src)
	}

	res.Header().Set("Content-Type", "text/html")
	hostname, err := os.Hostname()

	if err != nil {
		hostname = "localhost"
	}

	form := `
	<html> 
		<head> 
			<title> Upload file</title> 
		</head>
		<body> 
			<h1> Upload file to ` + os.TempDir() + ` on ` + hostname + `</h1>
			<form method="POST" enctype="multipart/form-data"><input type="file" name="my-file">
			<input type="submit">
			</form> 
		</body> 
	</html>`
	io.WriteString(res, form)
}

func main() {
	port := "9000"
	fmt.Println("[+] Server started on port", port)
	http.HandleFunc("/", handler)
	http.ListenAndServe(":"+port, nil)
}
