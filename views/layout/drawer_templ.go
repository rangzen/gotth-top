// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.793
package layout

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import "github.com/rangzen/gotth-top/views/home"

func Drawer() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div x-data=\"{ drawerOpen: false }\" class=\"drawer\"><input id=\"my-drawer-1\" type=\"checkbox\" class=\"drawer-toggle\" x-model=\"drawerOpen\"><div class=\"drawer-content flex flex-col\"><!-- Navbar --><div class=\"navbar bg-base-300 w-full\"><div class=\"flex-none lg:hidden\"><label for=\"my-drawer-1\" aria-label=\"open sidebar\" class=\"btn btn-square btn-ghost\"><svg xmlns=\"http://www.w3.org/2000/svg\" fill=\"none\" viewBox=\"0 0 24 24\" class=\"inline-block h-6 w-6 stroke-current\"><path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M4 6h16M4 12h16M4 18h16\"></path></svg></label></div><div class=\"mx-2 flex-1 px-2\">gotth top</div><div class=\"hidden flex-none lg:block\"><ul class=\"menu menu-horizontal\"><!-- Navbar menu content here --><li><a hx-get=\"/home\" hx-trigger=\"click\" hx-target=\"#content\">Home</a></li><li><a hx-get=\"/proc\" hx-trigger=\"click\" hx-target=\"#content\">Processes</a></li><li><a hx-get=\"/general\" hx-trigger=\"click\" hx-target=\"#content\">General</a></li></ul></div></div><!-- Page content here --><div id=\"content\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = home.Home().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></div><div class=\"drawer-side\"><label for=\"my-drawer-1\" aria-label=\"close sidebar\" class=\"drawer-overlay\"></label><ul class=\"menu bg-base-200 min-h-full w-80 p-4\"><!-- Sidebar content here --><li><a hx-get=\"/home\" hx-trigger=\"click\" hx-target=\"#content\" x-on:click=\"drawerOpen = false\">Home</a></li><li><a hx-get=\"/proc\" hx-trigger=\"click\" hx-target=\"#content\" x-on:click=\"drawerOpen = false\">Processes</a></li><li><a hx-get=\"/general\" hx-trigger=\"click\" hx-target=\"#content\" x-on:click=\"drawerOpen = false\">General</a></li></ul></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
