## My step to GO!

## Git

Actually great possibility to study staff that I've missed 😕

Fancy animation at learngitbranching was pretty verbose, yeah it expalins a lot. But, to be honest, their short solutions was tricky)

But I can't get why have you missed [documentation](https://git-scm.com/) at least in addition materials?

## Unix Shell

[Modules passed](./task_unix_shell/ "Modules")

Stuff that always was postponed and get studied from practical needs rather than intentionaly.

Kinda some basic commands was really new for me like

```bash
$ rmdir
```

Finally get full meaning from file permission assigning and managing.

Also found useful commands to manipulate with expansion in shell and I/O redirection. Yeah from I/O redirection commands like
```bash
$ awk
```
```bash
$ grep
```
From my side would like to add that bash have util called
```bash
$ envsubst
```
It can be useful even with k8s files templating


## Git Collaboration

[GitHub & Collaboration](./task_git_collaboration/udacity_course.png "Modules")

This lesson seems a little confusing, mostly it was mentioned in lesson 1 (Git Basics).
Generally it is big impatient to start language specific lessons.

## Go basics 1

This lessons was funny, working with string in Golang pretty comfortable. But feels like a lot of topics wasn't used in this task. Want to mention that Golang feature like Reader pretty new after NodeJS or Python. Also, it can be some pain to work with them at first time in net/http module. To complete theme with Reader want to mention io/ioutil.

And build in Golang utils would be greate to mention too. At least
```bash
go fmt
```

As unified code style made Go popular too.


## Memory Management

* What's going to happen if program reaches maximum limit of stack ?
* What's going to happen if program requests a big (more then 128KB) memory allocation on heap ?
* What's the difference between Text and Data memory segments ?

```
5623b298d000-5623b2991000 r-xp 00000000 08:02 1187433                    /usr/lib/gnome-settings-daemon/gsd-mouse
5623b2b91000-5623b2b92000 r--p 00004000 08:02 1187433                    /usr/lib/gnome-settings-daemon/gsd-mouse
5623b2b92000-5623b2b93000 rw-p 00005000 08:02 1187433                    /usr/lib/gnome-settings-daemon/gsd-mouse
5623b3e7e000-5623b3ec0000 rw-p 00000000 00:00 0                          [heap]
7f370c000000-7f370c021000 rw-p 00000000 00:00 0
7f370c021000-7f3710000000 ---p 00000000 00:00 0
7f3714000000-7f3714021000 rw-p 00000000 00:00 0
7f3714021000-7f3718000000 ---p 00000000 00:00 0
7f37183b8000-7f37183b9000 ---p 00000000 00:00 0
7f37183b9000-7f3718bb9000 rw-p 00000000 00:00 0
7f3718bb9000-7f3718bba000 ---p 00000000 00:00 0
7f3718bba000-7f37193ba000 rw-p 00000000 00:00 0
7f37193ba000-7f3719f01000 r--p 00000000 08:02 1186502                    /usr/lib/locale/locale-archive
7f3719f01000-7f3719f07000 r-xp 00000000 08:02 262413                     /lib/x86_64-linux-gnu/libuuid.so.1.3.0
7f3719f07000-7f371a106000 ---p 00006000 08:02 262413                     /lib/x86_64-linux-gnu/libuuid.so.1.3.0
7f371a106000-7f371a107000 r--p 00005000 08:02 262413                     /lib/x86_64-linux-gnu/libuuid.so.1.3.0
7f371a107000-7f371a108000 rw-p 00006000 08:02 262413                     /lib/x86_64-linux-gnu/libuuid.so.1.3.0
7f371a108000-7f371a14f000 r-xp 00000000 08:02 262425                     /lib/x86_64-linux-gnu/libblkid.so.1.1.0
7f371a14f000-7f371a34f000 ---p 00047000 08:02 262425                     /lib/x86_64-linux-gnu/libblkid.so.1.1.0
7f371a34f000-7f371a353000 r--p 00047000 08:02 262425                     /lib/x86_64-linux-gnu/libblkid.so.1.1.0
7f371a353000-7f371a354000 rw-p 0004b000 08:02 262425                     /lib/x86_64-linux-gnu/libblkid.so.1.1.0
7f371a354000-7f371a355000 rw-p 00000000 00:00 0
7f371a355000-7f371a358000 r-xp 00000000 08:02 267506                     /lib/x86_64-linux-gnu/libdl-2.27.so
7f371a358000-7f371a557000 ---p 00003000 08:02 267506                     /lib/x86_64-linux-gnu/libdl-2.27.so
7f371a557000-7f371a558000 r--p 00002000 08:02 267506                     /lib/x86_64-linux-gnu/libdl-2.27.so
7f371a558000-7f371a559000 rw-p 00003000 08:02 267506                     /lib/x86_64-linux-gnu/libdl-2.27.so
7f371a559000-7f371a560000 r-xp 00000000 08:02 267624                     /lib/x86_64-linux-gnu/librt-2.27.so
7f371a560000-7f371a75f000 ---p 00007000 08:02 267624                     /lib/x86_64-linux-gnu/librt-2.27.so
7f371a75f000-7f371a760000 r--p 00006000 08:02 267624                     /lib/x86_64-linux-gnu/librt-2.27.so
7f371a760000-7f371a761000 rw-p 00007000 08:02 267624                     /lib/x86_64-linux-gnu/librt-2.27.so
7f371a761000-7f371a77b000 r-xp 00000000 08:02 267616                     /lib/x86_64-linux-gnu/libpthread-2.27.so
7f371a77b000-7f371a97a000 ---p 0001a000 08:02 267616                     /lib/x86_64-linux-gnu/libpthread-2.27.so
7f371a97a000-7f371a97b000 r--p 00019000 08:02 267616                     /lib/x86_64-linux-gnu/libpthread-2.27.so
7f371a97b000-7f371a97c000 rw-p 0001a000 08:02 267616                     /lib/x86_64-linux-gnu/libpthread-2.27.so
7f371a97c000-7f371a980000 rw-p 00000000 00:00 0
7f371a980000-7f371a9f0000 r-xp 00000000 08:02 267605                     /lib/x86_64-linux-gnu/libpcre.so.3.13.3
7f371a9f0000-7f371abf0000 ---p 00070000 08:02 267605                     /lib/x86_64-linux-gnu/libpcre.so.3.13.3
7f371abf0000-7f371abf1000 r--p 00070000 08:02 267605                     /lib/x86_64-linux-gnu/libpcre.so.3.13.3
7f371abf1000-7f371abf2000 rw-p 00071000 08:02 267605                     /lib/x86_64-linux-gnu/libpcre.so.3.13.3
7f371abf2000-7f371abf9000 r-xp 00000000 08:02 1189269                    /usr/lib/x86_64-linux-gnu/libffi.so.6.0.4
7f371abf9000-7f371adf8000 ---p 00007000 08:02 1189269                    /usr/lib/x86_64-linux-gnu/libffi.so.6.0.4
7f371adf8000-7f371adf9000 r--p 00006000 08:02 1189269                    /usr/lib/x86_64-linux-gnu/libffi.so.6.0.4
7f371adf9000-7f371adfa000 rw-p 00007000 08:02 1189269                    /usr/lib/x86_64-linux-gnu/libffi.so.6.0.4
7f371adfa000-7f371ae4b000 r-xp 00000000 08:02 262708                     /lib/x86_64-linux-gnu/libmount.so.1.1.0
7f371ae4b000-7f371b04a000 ---p 00051000 08:02 262708                     /lib/x86_64-linux-gnu/libmount.so.1.1.0
7f371b04a000-7f371b04c000 r--p 00050000 08:02 262708                     /lib/x86_64-linux-gnu/libmount.so.1.1.0
7f371b04c000-7f371b04d000 rw-p 00052000 08:02 262708                     /lib/x86_64-linux-gnu/libmount.so.1.1.0
7f371b04d000-7f371b04e000 rw-p 00000000 00:00 0
7f371b04e000-7f371b065000 r-xp 00000000 08:02 267622                     /lib/x86_64-linux-gnu/libresolv-2.27.so
7f371b065000-7f371b265000 ---p 00017000 08:02 267622                     /lib/x86_64-linux-gnu/libresolv-2.27.so
7f371b265000-7f371b266000 r--p 00017000 08:02 267622                     /lib/x86_64-linux-gnu/libresolv-2.27.so
7f371b266000-7f371b267000 rw-p 00018000 08:02 267622                     /lib/x86_64-linux-gnu/libresolv-2.27.so
7f371b267000-7f371b269000 rw-p 00000000 00:00 0
7f371b269000-7f371b28e000 r-xp 00000000 08:02 267628                     /lib/x86_64-linux-gnu/libselinux.so.1
7f371b28e000-7f371b48d000 ---p 00025000 08:02 267628                     /lib/x86_64-linux-gnu/libselinux.so.1
7f371b48d000-7f371b48e000 r--p 00024000 08:02 267628                     /lib/x86_64-linux-gnu/libselinux.so.1
7f371b48e000-7f371b48f000 rw-p 00025000 08:02 267628                     /lib/x86_64-linux-gnu/libselinux.so.1
7f371b48f000-7f371b491000 rw-p 00000000 00:00 0
7f371b491000-7f371b4ad000 r-xp 00000000 08:02 267655                     /lib/x86_64-linux-gnu/libz.so.1.2.11
7f371b4ad000-7f371b6ac000 ---p 0001c000 08:02 267655                     /lib/x86_64-linux-gnu/libz.so.1.2.11
7f371b6ac000-7f371b6ad000 r--p 0001b000 08:02 267655                     /lib/x86_64-linux-gnu/libz.so.1.2.11
7f371b6ad000-7f371b6ae000 rw-p 0001c000 08:02 267655                     /lib/x86_64-linux-gnu/libz.so.1.2.11
7f371b6ae000-7f371b6b1000 r-xp 00000000 08:02 1180015                    /usr/lib/x86_64-linux-gnu/libgmodule-2.0.so.0.5600.4
7f371b6b1000-7f371b8b0000 ---p 00003000 08:02 1180015                    /usr/lib/x86_64-linux-gnu/libgmodule-2.0.so.0.5600.4
7f371b8b0000-7f371b8b1000 r--p 00002000 08:02 1180015                    /usr/lib/x86_64-linux-gnu/libgmodule-2.0.so.0.5600.4
7f371b8b1000-7f371b8b2000 rw-p 00003000 08:02 1180015                    /usr/lib/x86_64-linux-gnu/libgmodule-2.0.so.0.5600.4
7f371b8b2000-7f371b8bf000 r-xp 00000000 08:02 1190051                    /usr/lib/x86_64-linux-gnu/libwayland-client.so.0.3.0
7f371b8bf000-7f371babe000 ---p 0000d000 08:02 1190051                    /usr/lib/x86_64-linux-gnu/libwayland-client.so.0.3.0
7f371babe000-7f371bac0000 r--p 0000c000 08:02 1190051                    /usr/lib/x86_64-linux-gnu/libwayland-client.so.0.3.0
7f371bac0000-7f371bac1000 rw-p 0000e000 08:02 1190051                    /usr/lib/x86_64-linux-gnu/libwayland-client.so.0.3.0
7f371bac1000-7f371bca8000 r-xp 00000000 08:02 267483                     /lib/x86_64-linux-gnu/libc-2.27.so
7f371bca8000-7f371bea8000 ---p 001e7000 08:02 267483                     /lib/x86_64-linux-gnu/libc-2.27.so
7f371bea8000-7f371beac000 r--p 001e7000 08:02 267483                     /lib/x86_64-linux-gnu/libc-2.27.so
7f371beac000-7f371beae000 rw-p 001eb000 08:02 267483                     /lib/x86_64-linux-gnu/libc-2.27.so
7f371beae000-7f371beb2000 rw-p 00000000 00:00 0
7f371beb2000-7f371bfc6000 r-xp 00000000 08:02 1179992                    /usr/lib/x86_64-linux-gnu/libglib-2.0.so.0.5600.4
7f371bfc6000-7f371c1c6000 ---p 00114000 08:02 1179992                    /usr/lib/x86_64-linux-gnu/libglib-2.0.so.0.5600.4
7f371c1c6000-7f371c1c7000 r--p 00114000 08:02 1179992                    /usr/lib/x86_64-linux-gnu/libglib-2.0.so.0.5600.4
7f371c1c7000-7f371c1c8000 rw-p 00115000 08:02 1179992                    /usr/lib/x86_64-linux-gnu/libglib-2.0.so.0.5600.4
7f371c1c8000-7f371c1c9000 rw-p 00000000 00:00 0
7f371c1c9000-7f371c21b000 r-xp 00000000 08:02 1180016                    /usr/lib/x86_64-linux-gnu/libgobject-2.0.so.0.5600.4
7f371c21b000-7f371c41b000 ---p 00052000 08:02 1180016                    /usr/lib/x86_64-linux-gnu/libgobject-2.0.so.0.5600.4
7f371c41b000-7f371c41c000 r--p 00052000 08:02 1180016                    /usr/lib/x86_64-linux-gnu/libgobject-2.0.so.0.5600.4
7f371c41c000-7f371c41d000 rw-p 00053000 08:02 1180016                    /usr/lib/x86_64-linux-gnu/libgobject-2.0.so.0.5600.4
7f371c41d000-7f371c5b2000 r-xp 00000000 08:02 1179984                    /usr/lib/x86_64-linux-gnu/libgio-2.0.so.0.5600.4
7f371c5b2000-7f371c7b2000 ---p 00195000 08:02 1179984                    /usr/lib/x86_64-linux-gnu/libgio-2.0.so.0.5600.4
7f371c7b2000-7f371c7b9000 r--p 00195000 08:02 1179984                    /usr/lib/x86_64-linux-gnu/libgio-2.0.so.0.5600.4
7f371c7b9000-7f371c7ba000 rw-p 0019c000 08:02 1179984                    /usr/lib/x86_64-linux-gnu/libgio-2.0.so.0.5600.4
7f371c7ba000-7f371c7bc000 rw-p 00000000 00:00 0
7f371c7bc000-7f371c7d0000 r-xp 00000000 08:02 1187450                    /usr/lib/gnome-settings-daemon-3.0/libgsd.so
7f371c7d0000-7f371c9cf000 ---p 00014000 08:02 1187450                    /usr/lib/gnome-settings-daemon-3.0/libgsd.so
7f371c9cf000-7f371c9d1000 r--p 00013000 08:02 1187450                    /usr/lib/gnome-settings-daemon-3.0/libgsd.so
7f371c9d1000-7f371c9d2000 rw-p 00015000 08:02 1187450                    /usr/lib/gnome-settings-daemon-3.0/libgsd.so
7f371c9d2000-7f371c9f9000 r-xp 00000000 08:02 267455                     /lib/x86_64-linux-gnu/ld-2.27.so
7f371cbd4000-7f371cbdc000 rw-p 00000000 00:00 0
7f371cbf7000-7f371cbf9000 rw-p 00000000 00:00 0
7f371cbf9000-7f371cbfa000 r--p 00027000 08:02 267455                     /lib/x86_64-linux-gnu/ld-2.27.so
7f371cbfa000-7f371cbfb000 rw-p 00028000 08:02 267455                     /lib/x86_64-linux-gnu/ld-2.27.so
7f371cbfb000-7f371cbfc000 rw-p 00000000 00:00 0
7ffc9d413000-7ffc9d434000 rw-p 00000000 00:00 0                          [stack]
7ffc9d498000-7ffc9d49b000 r--p 00000000 00:00 0                          [vvar]
7ffc9d49b000-7ffc9d49c000 r-xp 00000000 00:00 0                          [vdso]
ffffffffff600000-ffffffffff601000 r-xp 00000000 00:00 0                  [vsyscall]

```



## TCP. UDP. Network