whoami
======

Simple HTTP docker service that prints it's container ID in a colored webpage:

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
        </body>
    </html>

Enviroment variables:
* PORT: port to listen (default 8000)
* COLOR: backgroung color (default red)

