{{$tAlias := .Aliases.Table .Table.Name -}}

func New{{$tAlias.UpSingular}}(mods ...{{$tAlias.UpSingular}}Mod) *{{$tAlias.UpSingular}}Template {
	return defaultFactory.New{{$tAlias.UpSingular}}(mods...)
}

func (f *factory) New{{$tAlias.UpSingular}}(mods ...{{$tAlias.UpSingular}}Mod) *{{$tAlias.UpSingular}}Template {
	o := &{{$tAlias.UpSingular}}Template{f: f}

	f.base{{$tAlias.UpSingular}}Mods.Apply(o)
 {{$tAlias.UpSingular}}ModSlice(mods).Apply(o)

	return o
}

func New{{$tAlias.UpPlural}}(number int, mods ...{{$tAlias.UpSingular}}Mod) {{$tAlias.UpSingular}}TemplateSlice {
	return defaultFactory.New{{$tAlias.UpPlural}}(number, mods...)
}

func (f *factory) New{{$tAlias.UpPlural}}(number int, mods ...{{$tAlias.UpSingular}}Mod) {{$tAlias.UpSingular}}TemplateSlice {
  var templates = make({{$tAlias.UpSingular}}TemplateSlice, number)

  for i := 0; i < number; i++ {
		templates[i] = f.New{{$tAlias.UpSingular}}(mods...)
	}

	return templates
}

