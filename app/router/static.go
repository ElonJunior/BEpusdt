package router

import (
	"html/template"
	"io/fs"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/v03413/bepusdt/app/log"
	"github.com/v03413/bepusdt/app/model"
	"github.com/v03413/bepusdt/app/utils"
	"github.com/v03413/bepusdt/static"
)

func staticInit(e *gin.Engine) {
	customPath := model.GetK(model.PaymentStaticPath)
	if customPath != "" && utils.IsExist(customPath) {
		initCustomPayment(e, customPath)

		return
	}

	initDefaultPayment(e)
}

func initCustomPayment(e *gin.Engine, path string) {
	tmpl := template.New("customer")

	template.Must(tmpl.ParseGlob(filepath.Join(path, "views", "*.html")))
	parseSecureTemplate(tmpl)
	e.SetHTMLTemplate(tmpl)

	e.StaticFS("/payment/assets", http.Dir(filepath.Join(path, "assets")))
	if fsContainsPath(static.Secure, "secure/assets") {
		e.StaticFS("/secure/assets", http.FS(subFS(static.Secure, "secure/assets")))
	}

	log.Info("成功注册自定义静态资源路径：", path)
}

func initDefaultPayment(e *gin.Engine) {
	tmpl := template.New("default")

	parseSecureTemplate(tmpl)
	template.Must(tmpl.ParseFS(static.Payment, "payment/views/*.html"))
	e.SetHTMLTemplate(tmpl)

	e.StaticFS("/payment/assets", http.FS(subFS(static.Payment, "payment/assets")))
	if fsContainsPath(static.Secure, "secure/assets") {
		e.StaticFS("/secure/assets", http.FS(subFS(static.Secure, "secure/assets")))
	}
}

func subFS(src fs.FS, dir string) fs.FS {
	sub, _ := fs.Sub(src, dir)

	return sub
}

func parseSecureTemplate(tmpl *template.Template) {
	if fsContainsPath(static.Secure, "secure/secure.html") {
		template.Must(tmpl.ParseFS(static.Secure, "secure/secure.html"))

		return
	}

	template.Must(tmpl.New("secure.html").Parse(`<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <title>Admin UI Missing</title>
  <style>
    body {
      margin: 0;
      padding: 24px;
      font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", sans-serif;
      background: #f7f8fa;
      color: #1f2329;
    }
    .card {
      max-width: 760px;
      margin: 40px auto;
      background: #fff;
      border: 1px solid #e5e6eb;
      border-radius: 10px;
      padding: 20px;
      line-height: 1.6;
    }
    code {
      display: inline-block;
      background: #f2f3f5;
      border: 1px solid #e5e6eb;
      border-radius: 6px;
      padding: 2px 8px;
      margin: 0 2px;
    }
  </style>
</head>
<body>
  <div class="card">
    <h2>Admin page is unavailable</h2>
    <p>The admin frontend file <code>static/secure/secure.html</code> is missing, so the login SPA cannot be rendered.</p>
    <p>Build the frontend and copy dist files to <code>static/secure</code>, then restart service.</p>
  </div>
</body>
</html>`))
}

func fsContainsPath(src fs.FS, path string) bool {
	_, err := fs.Stat(src, path)

	return err == nil
}
