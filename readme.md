# Gorter version 2.0.1
Gorter is CLI utility for sorting files. It solves a basic need of mine - keeping my home directory (and it's sub directories) organized.  
Gorter aims to be customizable - users can edit configuration file and define directory names and file types,
that will be placed in those directories.

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
gorter -i /nameof/input/directory -o /nameof/output/directory (this is optional)
```
If output directory is omitted, sorted files will be placed in directories within input directory.  
If you omit output directory, you can use even shorter command
```
gorter /nameof/input/directory 
```