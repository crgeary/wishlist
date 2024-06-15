// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.707
package views

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import (
	"github.com/crgeary/wishlist/components"
)

func Register() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<!doctype html><html lang=\"en-gb\" class=\"h-full\"><head>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = components.Head(components.HeadProps{
			Title: "Register",
		}).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</head><body class=\"flex items-center justify-center h-full\"><section class=\"w-1/3 border p-8\"><h1 class=\"text-2xl\">Register</h1><form class=\"mt-4\"><div><label for=\"name\">Name</label><div class=\"mt-1.5\"><input type=\"text\" name=\"name\" id=\"name\" class=\"border w-full p-1\"></div></div><div><label for=\"email\">E-mail address</label><div class=\"mt-1.5\"><input type=\"email\" name=\"email\" id=\"email\" class=\"border w-full p-1\"></div></div><div class=\"mt-2\"><label for=\"password\">Password</label><div class=\"mt-1.5\"><input type=\"password\" name=\"password\" id=\"password\" class=\"border w-full p-1\"></div></div><div class=\"mt-2\"><button class=\"bg-purple-400 px-2 py-1\">Register</button></div><p class=\"mt-2\">Or <a href=\"/signin\" class=\"underline\">sign in</a></p></form></section></body></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
