// Copyright 2018 Stuart Thompson.

// This file is part of Dockterm.

// Dockterm is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// Dockterm is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with Dockterm. If not, see <http://www.gnu.org/licenses/>.

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
