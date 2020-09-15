# whoami
Simple HTTP docker service that prints its container ID and IPs in a colored webpage:

    $ docker run -d -p 8000:8000 --name whoami -e COLOR=green camandel/whoami
    736ab83847bb12dddd8b09969433f3a02d64d5b0be48f7a5c59a594e3a6a3541
    
    $ curl $(hostname --all-ip-addresses | awk '{print $1}'):8000
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

Enviroment variables:
* PORT: port to listen (default 8000)
* COLOR: backgroung color (default red)

## Examples
Run a container with different port and color:
```
docker run -d -p 8888:8888 -e PORT=8888 -e COLOR=green camandel/whoami
```
Create a new image with a different default color:
```
docker build --build-arg color=blue -t camandel/whoami:blue .
docker run -d -p 8000:8000 --name whoami camandel/whoami:blue
```