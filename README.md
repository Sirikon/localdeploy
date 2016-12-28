# Molly #

[![CircleCI](https://circleci.com/gh/Sirikon/molly.svg?style=shield)](https://circleci.com/gh/Sirikon/molly)
[![Go Report Card](https://goreportcard.com/badge/github.com/sirikon/molly)](https://goreportcard.com/report/github.com/sirikon/molly)

![Molly Malone](http://i.imgur.com/vpbfqlb.jpg)

## Installation ##

### Automatic ###

```bash
curl https://molly.sirikon.me/install.sh | bash
```
### Manual ###

```bash
wget https://dl.sirikon.me/molly.zip
```

This will download a zip file with molly binary inside,
put it wherever you want (`/usr/bin` for example) be sure it's
a path inside your `$PATH` environment variable.

Whatever option you choose, once done, you should be able to run `molly`
command on your terminal.

## Usage ##

### Creating a new project ###

```bash
molly project add [project name]
```

This will create the folder structures inside `/srv/molly/[project name]`.

```
/srv/molly/[project name]
 ├─ run.sh
 ├─ deploy.sh
 └─ files/
```

Here you
can find the `run.sh` and `deploy.sh` files, which are used to define how
to run and deploy your project.

### Configuring a project ###

When molly receives a new artifact will first __deploy__ and then __run__
the project (both files are executed with `/srv/molly/[project name]/files`
as CWD)

For example, for a Node.js project, the `deploy.sh` file would look like this:

```bash
unzip $MOLLY_ARTIFACT
/path/to/npm install
```

And the `run.sh` file should look like this:

```bash
EXAMPLE_ENV_VAR=example_value
/path/to/node index.js
```

### Running molly daemon ###

To run the molly daemon just run the following command:

```bash
molly daemon
```

This will start a web server on port 8080 ready to receive the deployment requests.

__Note__: Right now molly doesn't self-install as a service, you need
to do this by yourself to be sure it's always working.

### Deploying a project ###

Once the project is properly configured and the molly daemon is running, you are able to deploy
your project.

```bash
zip -r artifact * # Compress whatever you need to deploy into artifact.zip
curl -F artifact=@./artifact.zip -F project=[project name] -F token=[project token] "http://yourserver.com:8080/deploy"
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
