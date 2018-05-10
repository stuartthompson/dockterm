package entities

// Container ...
// Represents a Docker container.
type Container struct {
	ID              string
	Created         string
	Path            string
	Args            []string
	State           ContainerState
	Image           string
	ResolvConfPath  string
	HostnamePath    string
	HostsPath       string
	LogPath         string
	Name            string
	RestartCount    int
	Driver          string
	Platform        string
	MountLabel      string
	ProcessLabel    string
	AppArmorProfile string
	ExecIDs         []string
}

// ContainerState ...
// Represents the state of a Docker container.
type ContainerState struct {
	Status     string
	Running    bool
	Paused     bool
	Restarting bool
	OOMKilled  bool
	Dead       bool
	Pid        int
	ExitCode   int
	Error      string
	StartedAt  string
	EndedAt    string
}

// Config ...
// Represents the config for a Docker container.
type Config struct {
	Hostname     string
	Domainname   string
	User         string
	AttachStdin  bool
	AttachStdout bool
	AttachStderr bool
	Tty          bool
	OpenStdin    bool
	StdinOnce    bool
	Env          []string
	Cmd          []string
	ArgsEscaped  bool
	Image        string
	//Volumes    ??
	WorkingDir string
	//Entrypoint ??
	//OnBuild    ??
	Labels []string
}
