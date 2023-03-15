# Groupie Tracker - Absolut Music
(A french version of the readme is disponible [here](readme-fr.md))
Absolut Music is a web application developed as part of the Groupie Tracker project at Ynov School of Computer Science. <br> The application retrieves information from an API to allow users to track the concerts of their favorite music groups.

## Prerequisiteicis

Before being able to use this application, you must have the following:

* Docker (version 23.0.1 or later) (if you want to run the application using Docker)
* Go (version 1.19.6 or later) (if you want to run the application using `go build`)

## Installation

1. Clone this code repository to your computer using the following command:
``` bash
git clone https://github.com/TomJegou/Groupie-tracker.git
```

2. Navigate to the application directory using the following command:
``` bash
cd Groupie-tracker
```

## Running with Docker

1. Launch the application using the following command:
``` bash
docker-compose up
```

2. Open your web browser and go to the following URL:
``` bash
http://localhost:80
```
3. Use the application to search for music groups and track their concerts.

## Running with Go

1. Compile the application using the following command:
``` bash
go build -o bin/absolut-music
```

2. Launch the application using the following command:
```bash
./bin/absolut-music
```

3. Open your web browser and go to the following URL:
``` bash
http://localhost:8080
```

4. Use the application to search for music groups and track their concerts.

### Organization

For the organization we use trello were you can find the task and the schedule [here](https://trello.com/invite/b/2xhTB18x/ATTI020dcaac9ef35348614f97bfc580ed476613003E/absolut-music).