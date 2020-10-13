# whoami
Simple HTTP docker service that prints its container ID and IPs in different pages and formats (/, /html, /colored and /text):

<p align="center">
  <img src="assets/whoami.jpg?raw=true" alt="img" />
</p>

Enviroment variables:
* PORT: port to listen (default 8000)
* COLOR: backgroung color (default red)

## Examples
Run a container with default settings:
```
docker run -d -p 8000:8000 --name whoami camandel/whoami
```
Run a container with different port and color:
```
docker run -d -p 8888:8888 -e PORT=8888 -e COLOR=green camandel/whoami
```
Create a new image with a different default color:
```
docker build --build-arg color=blue -t camandel/whoami:blue .
docker run -d -p 8000:8000 --name whoami camandel/whoami:blue
```
The output format on / depends on the browser (with curl you'll get a green colored text):
```
$ curl $(hostname --all-ip-addresses | awk '{print $1}'):8000
GREEN: I'm 736ab83847bb [172.17.0.2]
```
Get html page:
```
$ curl $(hostname --all-ip-addresses | awk '{print $1}'):8000/html
<!DOCTYPE html>
<html>
    <head>
        <meta charset="UTF-8">
        <title>whoami</title>
        <style>
        body {
            background-color: green;
            color: white;
            padding-top: 50px;
            text-align: center;
            font-family: sans-serif, serif;
        }
        </style>
    </head>
    <body>
        <h1>I'm 736ab83847bb</h1>
        <h1>172.17.0.2</h1>
    </body>
</html>
```
Get plain text:
```
$ curl $(hostname --all-ip-addresses | awk '{print $1}'):8000/text
GREEN: I'm 736ab83847bb [172.17.0.2]
```
Get colored text output:
```
$ curl $(hostname --all-ip-addresses | awk '{print $1}'):8000/colored
GREEN: I'm 736ab83847bb [172.17.0.2]
```