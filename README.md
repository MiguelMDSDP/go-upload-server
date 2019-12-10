# Upload Server

File upload system using GO and live-server.


## Contents

### Testing Form

The testing-form.html script contains a html script that makes a post form to upload files.


### Main Application

The main.go script executes an post http route simulating an upload server, saving the uploaded files on my-files directiory.


## Installation

This application uses go modules, so doesn't needs a specific installation for scripts.

Neverthless, you must install liver-server through the following commands:
> sudo apt install npm
>
> npm install -g live-server


## Usage

To use this application you only have to create the my-files folder inside the repository using the following command:

> mkdir my-files

Now, you just have to run the main application and live-server in different terminals:

> go run main.go
>
> live-server


## References

- https://github.com/gorilla/mux
- https://gist.github.com/donmccurdy/20fb112949324c92c5e8