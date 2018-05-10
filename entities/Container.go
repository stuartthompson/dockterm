package entities

// Container ...
// Represents a Docker container.
type Container struct {
	ID      string
	Created string
	Path    string
	Args    []string
	State   ContainerState
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
