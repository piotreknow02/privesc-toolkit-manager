package types

const (
	Linux   = 0
	Windows = 1
	Mac     = 2
)

const (
	NoArch = 0
	Amd64  = 1
	Arm64  = 2
	I386   = 3
)

type PEScript struct {
	Name   string
	Url    string
	System int
	Arch   int
}

func GetNames(p []PEScript) []string {
	var names []string
	for _, v := range p {
		names = append(names, v.Name)
	}
	return names
}

func (p PEScript) CheckSystem(systemId int, ArchId int) bool {
	return p.System == systemId && (p.Arch == ArchId || p.Arch == NoArch)
}

func FilterBySystem(list []PEScript, currentSystem int, currentArch int) []PEScript {
	var filtered []PEScript
	for _, v := range list {
		if v.CheckSystem(currentSystem, currentArch) {
			filtered = append(filtered, v)
		}
	}
	return filtered
}

func GetSystemName(id int) string {
	switch id {
	case Linux:
		return "Linux"
	case Windows:
		return "Windows"
	case Mac:
		return "Mac"
	}
	panic("Unknown OS")
}
