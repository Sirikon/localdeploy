# Molly #

[![CircleCI](https://circleci.com/gh/Sirikon/molly.svg?style=svg)](https://circleci.com/gh/Sirikon/molly)

![Molly Malone](http://i.imgur.com/vpbfqlb.jpg)

## Installation ##

### Automatic ###

```
curl https://molly.sirikon.me/install.sh | bash
```
### Manual ###

```
wget https://dl.sirikon.me/molly.zip
```

This will download a zip file with molly binary inside,
put it wherever you want (`/usr/bin` for example) be sure it's
a path inside your `$PATH` environment variable.

Whatever option you choose, once done, you should be able to run `molly`
command on your terminal.

__Note__: Right now molly doesn't self-install as a service, you need
to do this by yourself to be sure it's always working. This service
should run the command `molly daemon`.

## Usage ##

### Creating a new project ###

Run:

```
molly project add [project name]
```

This will create the folder structures inside `/srv/molly/[project name]`, here you
can find the `run.sh` and `deploy.sh` files, which are used to define how
to run and deploy your project. (both files are executed with
`/srv/molly/[project name]/files` as CWD)

For example, for a Node.js project, the deploy.sh file would look like this:

```
unzip $MOLLY_ARTIFACT
/path/to/npm install
```

And the run.sh file should look like this:

```
EXAMPLE_ENV_VAR=example_value
/path/to/node index.js
```

## License ##

```
MIT License

Copyright (c) 2016 Carlos Fernández

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```
