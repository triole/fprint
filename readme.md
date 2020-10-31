# Aprint

<!--- mdtoc: toc begin -->

1.	[Synopsis](#synopsis)
2.	[Config files](#config-files)
3.	[Build](#build)<!--- mdtoc: toc end -->

## Synopsis

A simple shell tool to print text strings to ascii art. Basically a wrapper of my fork of [go-figure](https://github.com/common-nighthawk/go-figure).

Run `aprint -h` to find out what it can do.

## Config files

Settings can be passed by command line arguments or by config. Configs are in toml format and used with the `-k` parameter. This is how they look.

```toml
text = "main text"
font = "univers"
colour = "yellow"
pretext = "text before"
posttext = "text after"
```

## Build

Either build by hand or use [task](https://github.com/go-task/task).

```shell
# this
go build

# or that
task

# remember what you can do if you have 'task'
task -l
```
