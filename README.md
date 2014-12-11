# google

`google` is a command line utility to query Google easily and quickly from the terminal.
The `google` command opens a browser window with the google results webpage of the given search query.

## Installation

Assuming you have a working Go environment and `GOPATH/bin` is in your `PATH`, you can install `google` easily: 

~~~ shell
go get github.com/hariharan-uno/google
~~~

You can verify that `google` was installed correctly:

~~~ shell
google -h
~~~

## Usage 

`google` is simple to use, no flags, no options. Just `google querystring`. 

For example,

~~~ shell
google who is the president of united states?
~~~

## Contribute

I have used and tested this only on linux. So, please report the bugs by trying it out and send pull requests!

## Motivation

I use Google for searching a lot of things while coding and I felt the two step process of opening google.com and then typing the query as too slow. This is just to scratch my itch :)

## Trademark

Google is a registered trademark of Google Inc.