package utils

//String constants
const (
	DateFormat = "20060102"
	FilesPath  = "\\recipes\\input\\"
	FilesName  = "recipes"
	OutputPath = "\\recipes\\output\\"
	XlsExt     = ".xls"
	XlsxExt    = ".xlsx"

	SaveError         = "No se pudo guardar el fichero %s\r\n"
	CreateError       = "No se pudo crear el fichero %s\r\n"
	FileNotFoundError = "No se pudo encontrar el fichero %s\r\n"
	NotZipFileError   = "%s, no es un fichero compatible\r\n"
	UnknownError      = "Error desconocido"

	ZipPrefix  = "OpenFile: zip:"
	OpenPrefix = "OpenFile: open"

	IntroMessage          = "\r\nIntroduzca comando (create nombre/read nombre/exit):"
	FilePathMessage       = "Introduzca el nombre del fichero de recetas:"
	NoFileNameMessage     = "Debe introducir el nombre del fichero"
	UnknownCommandMessage = "Comando no reconocido, introduzca nuevo comando:"
	ExitMessage           = "Cerrando con error %d\r\n"
)
