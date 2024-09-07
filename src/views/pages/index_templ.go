// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.771
package pages

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import (
	"github.com/VoltealProductions/TheArchiveReader/views/layout"
)

func Index() templ.Component {
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
		templ_7745c5c3_Var2 := templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
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
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"bg-skin-theme\"><h1 class=\"text-xl text-skin-base text-center uppercase font-bold p-1 mb-2 bg-black/45\">New Newspaper Articles</h1></div><div class=\"flex justify-between flex-row p-3 gap-4\"><article class=\"text-xs text-center basis-1/4\"><h1 class=\"text-center font-bold text-base\">Example Article <span class=\"text-xs\">by <a class=\"text-skin-accent hover:text-skin-accent-light\" href=\"#\">John Doe</a></span></h1><h2 class=\"text-center text-sm font-bold mb-1 italic\"><a class=\"text-skin-accent hover:text-skin-accent-light\" href=\"\">The Stormwind Tribune</a></h2><img src=\"/public/img/tww-bg-2.jpeg\" alt=\"test image\" class=\"h-24 w-full object-cover object-top mb-1\"><p class=\"mb-4 text-left\">Lorem ipsum dolor sit amet consectetur adipisicing elit. Nulla asperiores consequuntur a deleniti dignissimos nihil sint officiis quas Lorem ipsum dolor sit amet consectetur adipisicing elit. Nulla asperiores consequuntur a deleniti dignissimos nihil sint officiis quas...</p><a href=\"#\" class=\"bg-skin-theme text-skin-base px-4 p-2 hover:text-skin-accent-light hover:bg-skin-theme-dark\">read more</a></article><article class=\"text-xs basis-1/4 text-center\"><h1 class=\"text-center font-bold text-base\">Example Article <span class=\"text-xs\">by <a class=\"text-skin-accent hover:text-skin-accent-light\" href=\"#\">John Doe</a></span></h1><h2 class=\"text-center text-sm font-bold mb-1 italic\"><a class=\"text-skin-accent hover:text-skin-accent-light\" href=\"\">The Stormwind Tribune</a></h2><img src=\"/public/img/tww-bg-2.jpeg\" alt=\"test image\" class=\"h-24 w-full object-cover object-top mb-1\"><p class=\"mb-4 text-left\">Lorem ipsum dolor sit amet consectetur adipisicing elit. Nulla asperiores consequuntur a deleniti dignissimos nihil sint officiis quas Lorem ipsum dolor sit amet consectetur adipisicing elit. Nulla asperiores consequuntur a deleniti dignissimos nihil sint officiis quas...</p><a href=\"#\" class=\"bg-skin-theme text-skin-base px-4 p-2 hover:text-skin-accent-light hover:bg-skin-theme-dark\">read more</a></article><article class=\"text-xs text-center basis-1/4\"><h1 class=\"text-center font-bold text-base\">Example Article <span class=\"text-xs\">by <a class=\"text-skin-accent hover:text-skin-accent-light\" href=\"#\">John Doe</a></span></h1><h2 class=\"text-center text-sm font-bold mb-1 italic\"><a class=\"text-skin-accent hover:text-skin-accent-light\" href=\"\">The Stormwind Tribune</a></h2><img src=\"/public/img/tww-bg-2.jpeg\" alt=\"test image\" class=\"h-24 w-full object-cover object-top mb-1\"><p class=\"mb-4 text-left\">Lorem ipsum dolor sit amet consectetur adipisicing elit. Nulla asperiores consequuntur a deleniti dignissimos nihil sint officiis quas Lorem ipsum dolor sit amet consectetur adipisicing elit. Nulla asperiores consequuntur a deleniti dignissimos nihil sint officiis quas...</p><a href=\"#\" class=\"bg-skin-theme text-skin-base px-4 p-2 hover:text-skin-accent-light hover:bg-skin-theme-dark\">read more</a></article><article class=\"text-xs text-center basis-1/4\"><h1 class=\"text-center font-bold text-base\">Example Article <span class=\"text-xs\">by <a class=\"text-skin-accent hover:text-skin-accent-light\" href=\"#\">John Doe</a></span></h1><h2 class=\"text-center text-sm font-bold mb-1 italic\"><a class=\"text-skin-accent hover:text-skin-accent-light\" href=\"\">The Stormwind Tribune</a></h2><img src=\"/public/img/tww-bg-2.jpeg\" alt=\"test image\" class=\"h-24 w-full object-cover object-top mb-1\"><p class=\"mb-4 text-left\">Lorem ipsum dolor sit amet consectetur adipisicing elit. Nulla asperiores consequuntur a deleniti dignissimos nihil sint officiis quas Lorem ipsum dolor sit amet consectetur adipisicing elit. Nulla asperiores consequuntur a deleniti dignissimos nihil sint officiis quas...</p><a href=\"#\" class=\"bg-skin-theme text-skin-base px-4 p-2 hover:text-skin-accent-light hover:bg-skin-theme-dark\">read more</a></article></div><div class=\"pb-1 my-1\"><div class=\"bg-skin-theme\"><h1 class=\"text-xl text-skin-base text-center uppercase font-bold p-1 mb-2 bg-black/45\">New Adventures</h1></div></div><div class=\"relative overflow-x-auto rounded-lg\"><table class=\"w-full table-fixed text-left font-semibold text-black/70\"><tbody class=\"text-xs\"><tr class=\"border-b border-gray-200\"><td colspan=\"2\" class=\"px-3 py-1\"><a class=\"text-skin-accent hover:text-skin-accent-light\" href=\"#\">The Dark Heart</a></td><td colspan=\"2\" class=\"px-3 py-3\"><a class=\"text-skin-accent hover:text-skin-accent-light\" href=\"#\">Celestine Eveningshine</a></td><td class=\"px-3 py-1\">Val'sharah</td><td class=\"px-3 py-1\">3 hours ago</td></tr><tr class=\"border-b border-gray-200\"><td colspan=\"2\" class=\"px-3 py-1\"><a class=\"text-skin-accent hover:text-skin-accent-light\" href=\"#\">From Dusk to Dawn</a></td><td colspan=\"2\" colspan=\"2\" class=\"px-3 py-3\"><a class=\"text-skin-accent hover:text-skin-accent-light\" href=\"#\">Celestine Eveningshine</a></td><td class=\"px-3 py-1\">Suramar</td><td class=\"px-3 py-1\">6 hours ago</td></tr><tr class=\"border-b border-gray-200\"><td colspan=\"2\" class=\"px-3 py-1\"><a class=\"text-skin-accent hover:text-skin-accent-light\" href=\"#\">Lies, Secrets, and poisons</a></td><td colspan=\"2\" class=\"px-3 py-3\"><a class=\"text-skin-accent hover:text-skin-accent-light\" href=\"#\">Celestine Eveningshine</a></td><td class=\"px-3 py-1\">Orgrimmar</td><td class=\"px-3 py-1\">12 hours ago</td></tr></tbody></table></div><div class=\"flex flex-row p-3 gap-4\"><!-- Characters / Guilds --><div class=\"basis-1/2\"><div class=\"col-span-full pb-1 my-1\"><div class=\"bg-skin-theme\"><h1 class=\"text-xl text-skin-base text-center uppercase font-bold p-1 mb-2 bg-black/45\">New Characters</h1></div><table class=\"w-full table-fixed text-left font-semibold text-black/70\"><tbody class=\"text-xs\"><tr class=\"border-b border-gray-300\"><td colspan=\"2\" class=\"px-3 py-3\"><a class=\"text-skin-accent hover:text-skin-accent-light\" href=\"#\">Celestine Eveningshine</a></td><td class=\"px-3 py-3\"><span class=\"text-red-600\">Horde</span></td><td class=\"px-3 py-3\">1 hour ago</td></tr><tr class=\"border-b border-gray-300\"><td colspan=\"2\" class=\"px-3 py-3\"><a class=\"text-skin-accent hover:text-skin-accent-light\" href=\"#\">Augustina Seabreeze</a></td><td class=\"px-3 py-3\"><span class=\"text-yellow-600\">Neutral</span></td><td class=\"px-3 py-3\">1 hour ago</td></tr><tr class=\"border-b border-gray-300\"><td colspan=\"2\" class=\"px-3 py-3\"><a class=\"text-skin-accent hover:text-skin-accent-light\" href=\"#\">Ishala the Reverent</a></td><td class=\"px-3 py-3\"><span class=\"text-blue-600\">Alliance</span></td><td class=\"px-3 py-3\">1 hour ago</td></tr></tbody><!-- Characters --></table></div></div><div class=\"basis-1/2\"><div class=\"col-span-full pb-1 my-1\"><div class=\"bg-skin-theme\"><h1 class=\"text-xl text-skin-base text-center uppercase font-bold p-1 mb-2 bg-black/45\">New Guilds</h1></div></div><table class=\"w-full table-fixed text-left font-semibold text-black/70\"><tbody class=\"text-xs\"><tr class=\"border-b border-gray-300\"><td colspan=\"2\" class=\"px-3 py-3\"><a class=\"text-skin-accent hover:text-skin-accent-light\" href=\"#\">The Crimson Batallion</a></td><td class=\"px-3 py-3\"><span class=\"text-red-600\">Horde</span></td><td class=\"px-3 py-3\">10 hours ago</td></tr><tr class=\"border-b border-gray-300\"><td colspan=\"2\" class=\"px-3 py-3\"><a class=\"text-skin-accent hover:text-skin-accent-light\" href=\"#\">The Silver Convocation</a></td><td class=\"px-3 py-3\"><span class=\"text-yellow-600\">Neutral</span></td><td class=\"px-3 py-3\">20 hours ago</td></tr><tr class=\"border-b border-gray-300\"><td colspan=\"2\" class=\"px-3 py-3\"><a class=\"text-skin-accent hover:text-skin-accent-light\" href=\"#\">The Sisterhood</a></td><td class=\"px-3 py-3\"><span class=\"text-blue-600\">Alliance</span></td><td class=\"px-3 py-3\">2 days ago</td></tr></tbody><!-- Guilds --></table></div></div><div class=\"pb-1 my-1\"><div class=\"bg-skin-theme\"><h1 class=\"text-xl text-skin-base text-center font-bold p-1 mb-2 uppercase bg-black/45\">Noticeboard</h1></div></div><div class=\"relative overflow-x-auto rounded-lg\"><table class=\"w-full table-fixed text-left font-semibold text-black/70\"><tbody class=\"text-xs\"><tr class=\"border-b border-gray-200\"><td colspan=\"2\" class=\"px-3 py-1\"><a class=\"text-skin-accent hover:text-skin-accent-light\" href=\"#\">Seeking Troll Expert</a></td><td colspan=\"2\" class=\"px-3 py-3\"><a class=\"text-skin-accent hover:text-skin-accent-light\" href=\"#\">Eladrin Moonshine</a></td><td colspan=\"2\" class=\"px-3 py-1\">Posted in <span class=\"font-bold\">Stormwind City</span></td><td class=\"px-3 py-1\">5 hours ago</td></tr><tr class=\"border-b border-gray-200\"><td colspan=\"2\" class=\"px-3 py-1\"><a class=\"text-skin-accent hover:text-skin-accent-light\" href=\"#\">A Ritual Hunt</a></td><td colspan=\"2\" class=\"px-3 py-3\"><a class=\"text-skin-accent hover:text-skin-accent-light\" href=\"#\">Eas'tien Volteal</a></td><td colspan=\"2\" class=\"px-3 py-1\">Posted in <span class=\"font-bold\">Bel'ameth</span></td><td class=\"px-3 py-1\">10 hours ago</td></tr><tr class=\"border-b border-gray-200\"><td colspan=\"2\" class=\"px-3 py-1\"><a class=\"text-skin-accent hover:text-skin-accent-light\" href=\"#\">Notis of konstrukting</a></td><td colspan=\"2\" class=\"px-3 py-3\"><a class=\"text-skin-accent hover:text-skin-accent-light\" href=\"#\">Grug the Peon</a></td><td colspan=\"2\" class=\"px-3 py-1\">Posted in <span class=\"font-bold\">Orgrimmar</span></td><td class=\"px-3 py-1\">12 hours ago</td></tr></tbody></table></div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			return templ_7745c5c3_Err
		})
		templ_7745c5c3_Err = layout.Base().Render(templ.WithChildren(ctx, templ_7745c5c3_Var2), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
