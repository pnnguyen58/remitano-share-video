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
- Login: login with username and password to receive token.
  Token will be cached in memory and expired after 1 hour
- Share video: share video with authentication
- List video: get list of videos shared

## Tech stack
- Framework: beego
- Programming language: golang
- Database: postgres
- Cache: memory cache (using beego lib)

## Database

- docker run --name remitano-share-video-postgres --rm -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -e PGDATA=/var/lib/postgresql/data/pgdata -v /tmp:/var/lib/postgresql/data -p 5432:5432 -it postgres:14.1-alpine 
- docker exec -it nkeeping-staging-postgres /bin/sh
- CREATE USER worker WITH ENCRYPTED PASSWORD '1qazxsw23edc';
- CREATE DATABASE videos WITH OWNER = worker;

- CREATE TABLE public.users (
id bigserial NOT NULL,
username text NOT NULL,
"password" text NOT NULL,
created_at timestamptz NULL DEFAULT now(),
CONSTRAINT users_pk PRIMARY KEY (id),
CONSTRAINT users_un UNIQUE (username)
);
  
- CREATE TABLE public.videos (
id bigserial NOT NULL,
user_id int8 NOT NULL,
title text NOT NULL,
description text NULL,
link text NOT NULL,
created_at timestamptz NULL DEFAULT now(),
CONSTRAINT videos_pk PRIMARY KEY (id)
);
  
## Deployment
- Test env: go run main.go
- Deploy by docker: 
  - docker build --tag remitano-share-video .
  - docker run -d -p 8080:8080 remitano-share-video

## Test
- Unit test: go test unit_test.go
- Integration test: 
  - To implement integration_test, go to conf/app.conf
  and comment "SessionOn = true"
  execute: go test integration_test.go

## Scalability
- The system must handle many users
  - In case of using SQL database, we should using partition by time,
  could be 3 months, to increase loading video list
  - In case of using noSQL, because users only need add and get by
  filter or search keywords. With using noSQL, we can optimaize speed of
  inserting or loading videos.
  - To increase sharing with concurrent users, we should use 
  queuing technique (like kafka). Or we can implement by ourselves with using 
  goroutine combine channel
### References
https://beego.vip/docs/intro/