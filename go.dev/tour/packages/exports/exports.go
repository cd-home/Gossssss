package exports

// Exported names
// a name is exported if it begins with a capital letter
// includes packages name space
// eg: variables
// eg: const
// eg: struct
// eg: function

// tips:
// struct fields also satisfies this rule

var ExportVar = "Export Var"

var UnExportVar = "UnExport Var By Getter Func To Export"

const ExportConst = "ExportConst"

// ExportStruct Can Export
type ExportStruct struct {
	Name string
	// Can Not Export When Use This Struct
	age uint8
}

// ExportFunc
// Maybe Some Name Can Not Exported, But User 'Getter' Func To Export
func ExportFunc() string {
	return UnExportVar
}
