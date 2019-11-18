# Gorter
Gorter is CLI utility for sorting files. It solves a basic need of mine - keep downloads folder tidy. I think many can relate. 
Gorter aims to be user customizable with usage of configuration file, in which user can define directory names and file types,
that will be placed in said directory.

# Instalation

Build the executable
```
go build gorter.go config.go
```

Place the executable in your PATH, for example /usr/local/bin.
Make sure gorter is executable!
```
chmod +x gorter
```

Place the config file in $HOME/.config/ for individual user(s) or in /etc/ for all users.
Gorter also can read config file from the same directory it's placed in.

# Using Gorter

```
gorter input_dir output_dir(optional)
```
If output directory is omitted, sorted files will be placed in directories within input directory.