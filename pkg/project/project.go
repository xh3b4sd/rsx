package project

var (
	description = "Command line tool for simulating different backing scenarios of the resistance protocol."
	gitSHA      = "n/a"
	name        = "rsx"
	source      = "https://github.com/xh3b4sd/rsx"
	version     = "n/a"
)

func Description() string {
	return description
}

func GitSHA() string {
	return gitSHA
}

func Name() string {
	return name
}

func Source() string {
	return source
}

func Version() string {
	return version
}
