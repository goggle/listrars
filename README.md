Listrars
========

Listrars lists the rar files in a given directory, but only shows the first part in case of multipart rar files.

Usage
-----

```
Usage:
listrars [PATH]

Options:
-h --help     Show this screen.
--version     Show version.
```

Example
-------
Let the directory `/home/user/rfiles` contain the following files:
```
hello_world.c
photos_paris_2017.rar
big_data.part001.rar
big_data.part002.rar
big_data.part003.rar
big_data.part004.rar
```
By running `listrars` in this directory or `listrars /home/user/rfiles` in an arbitrary directory, we get this output:
```
big_data.part001.rar
photos_paris_2017.rar
```
This can now be used to extract all the rar files in `/home/user/rfiles` by using `xargs` and `unrar`:
```
listrars | xargs -I{} unrar e {}
```
