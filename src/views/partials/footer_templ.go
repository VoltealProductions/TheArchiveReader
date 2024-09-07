// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.771
package partials

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func Footer() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\" basis-1/4 px-5 py-1\"><div class=\"\"><div class=\"flex flex-row gap-5 items-center text-center\"><ul class=\"basis-1/2\"><li><a class=\"text-base text-skin-base hover:text-skin-accent-light\" href=\"#\">Characters</a></li><li><a class=\"text-base text-skin-base hover:text-skin-accent-light\" href=\"#\">Guilds</a></li><li><a class=\"text-base text-skin-base hover:text-skin-accent-light\" href=\"#\">Adventures</a></li><li><a class=\"text-base text-skin-base hover:text-skin-accent-light\" href=\"#\">Noticeboard</a></li><li><a class=\"text-base text-skin-base hover:text-skin-accent-light\" href=\"#\">News Archive</a></li></ul><ul class=\"basis-1/2\"><li><a class=\"text-base text-skin-base hover:text-skin-accent-light\" href=\"#\">Characters</a></li><li><a class=\"text-base text-skin-base hover:text-skin-accent-light\" href=\"#\">Guilds</a></li><li><a class=\"text-base text-skin-base hover:text-skin-accent-light\" href=\"#\">Adventures</a></li><li><a class=\"text-base text-skin-base hover:text-skin-accent-light\" href=\"#\">Noticeboard</a></li><li><a class=\"text-base text-skin-base hover:text-skin-accent-light\" href=\"#\">News Archive</a></li></ul></div></div></div><div class=\" basis-2/4 text-center* px-5 py-1\"><div class=\"\"><p class=\"text-xs mb text-center md:text-base\">Application copyrighted &copy; (2024) by <a class=\"text-base text-skin-base hover:text-skin-accent-light\" target=\"_blank\" href=\"https://github.com/Volteal\" class=\"text-link visited:text-visitedlink italic hover:underline\">Volteal Productions</a><p class=\"text-xs text-center md:text-base\">World of Warcraft is copyrighted &copy; (2004) by <a class=\"text-base text-skin-base hover:text-skin-accent-light\" target=\"_blank\" href=\"https://www.blizzard.com\" class=\"text-link visited:text-visitedlink italic hover:underline\">Blizzard Entertainment, Inc</a> All Rights Reserved.</p><p class=\"text text-center md:text-base\">All content on this app is &copy; (2024) by the owners.</p></p></div></div><div class=\" basis-1/4 px-5 py-1\"><div class=\"\"><ul class=\"basis-1 text-center\"><li><a class=\"text-base text-skin-base hover:text-skin-accent-light\" href=\"#\">Warcraft Roleplay Guidebook</a></li><li><a class=\"text-base text-skin-base hover:text-skin-accent-light\" href=\"#\">Copyright Notice</a></li><li><a class=\"text-base text-skin-base hover:text-skin-accent-light\" href=\"#\">Terms & Conditions</a></li><li><a class=\"text-base text-skin-base hover:text-skin-accent-light\" href=\"#\">FAQ</a></li><li><a class=\"text-base text-skin-base hover:text-skin-accent-light\" href=\"#\">Contact</a></li></ul></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
