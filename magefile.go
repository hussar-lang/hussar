// +build mage

package main

var Default = Build

// Build compiles the source files into usable binaries.
func Build() {
	// get git version and semver version - how to do without git repo
	// compile with correct version numbers injected
}

// Install helps to install the binaries correctly.
func Install() {
	// move binary to /usr/bin, /usr/local/bin or /bin, otherwise print a note that it should be put in $PATH
	// maybe install to GoBin instead, but still compile correctly with versions
	// init preferences in $HOME/.prefs
}

// Release prepares the built files for bundling and release - this should ONLY be used in CI and testing. Normal users will not have to use this ever.
func Release() {
	// generate changelog
	// upload binary to github releases
	// generate (install.sh) makeself.sh
	// upload install.sh via sftp to get.hussar.io
}
