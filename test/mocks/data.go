package mocks

func getMockContent(url string) string {
	var content string
	switch url {
	case "http://adjust.com":
		content = `<!doctype html><html lang=en><head><meta charset=utf-8><title>The Mobile Measurement Company | Adjust</title></head><body><p>Mobile Measurement made easy: Adjust unifies all your marketing activities into one powerful platform, giving you the insights you need to scale your business.</p></body></html>`
	case "http://facebook.com":
		content = `<!doctype html><html lang=en><head><meta charset=utf-8><title>Facebook</title></head><body><p>Hello from Facebook</p></body></html>`
	case "http://google.com":
		content = `<!doctype html><html lang=en><head><meta charset=utf-8><title>Google</title></head><body><p>Hello from Google</p></body></html>`
	case "http://twitter.com":
		content = `<!doctype html><html lang=en><head><meta charset=utf-8><title>Twitter</title></head><body><p>Hello from Twitter</p></body></html>`
	}
	return content
}

func GetMockMD5(url string) string {
	var hash string
	switch url {
	case "http://adjust.com":
		hash = "fa26298f5deef2ef2d23af8c57fd1866"
	case "http://facebook.com":
		hash = "56a2c0d167d15ca53f484d50ad2fd73b"
	case "http://google.com":
		hash = "afdaaeba6c02ae7485789f71ba9e0ab6"
	case "http://twitter.com":
		hash = "558333cb84dc818084f0c4eef270ec8c"
	}
	return hash
}
