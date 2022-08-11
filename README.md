# Share videos
## Required Features:
- Register: As a new user, I want to create an account and sign in by entering the username and password for the first time. (Picture 1)
- Login: As a registered user, I want to sign in by entering the username and password. (Picture 1)
- Share movie: As a signed in user, I want to share a youtube movie by clicking “Share a movie” button (Picture 2) and filling the share form (Picture 3).
- See movie list: As a visitor, I want to see a list of all shared movies. (Picture 1) (no need to display the number of up/down votes)

## Getting Started
This project is a starting point for a Beego application.
## Project structure
- MVC

## Features
- Register: create a user with username and password
- Login: login with username and password to receive token
- Share video: share video with authentication
- List video: get list of videos shared

## Tech stack
- Framework: beego
- Programming language: golang
- Database: postgres
- Cache: memory cache (using beego lib)

## Deployment
- Test: go run main.go
- Deploy by docker: 
  - docker build --tag remitano-share-video .
  - docker run -d -p 8080:8080 remitano-share-video
### References
https://beego.vip/docs/intro/