package utils

/*type OS int

var GlobalOS int

const (
	WINDOWS OS = 1 + iota
	DARWIN
	LINUX
	DEFAULT
)

var osDef = [...]string {
	"windows",
	"darwin",
	"linux",
	"default",
}

func setOS(os string){
	switch os {
	case "windows":
		GlobalOS = 0
		fmt.Println("Code not implemented")
	case "darwin":
		GlobalOS = 1
		fmt.Println("Code not implemented")
	case "linux":
		GlobalOS = 2
		fmt.Println("Code not implemented")
	default:
		GlobalOS = 4
		fmt.Println("Code not implemented")
	}
}

func (os OS) String() string {
	return osDef[os - 1]
}*/

func FormatStatSlice(rawStatSlice []string) []string {
	var statSlice []string
	for _, stat := range rawStatSlice {
		if stat != "" {
			statSlice = append(statSlice, stat)
		}
	}
	return statSlice
}
