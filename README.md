# slendercards-go-server

## Table of contents
- [Introduction](#Introduction)
- [Getting Started](#Getting-Started)
- [Features](#Features)
- [Future Features](#Future-Features)
- [Technologies](#Technologies)

## Introduction
slendercards is a web app built to help people learn foreign languages. Users can choose from multiple language pairs and automatically generate context-based flashcards based on personal search terms.

slendercards-go-server is the companion server to [slendercards-go-svelte](https://github.com/jakewmiles/slendercards-go-svelte). 
slendercards-go-server makes use of web scraping using Colly, and then stores the flashcards in a PostgreSQL database.

See slendercards in action [here](https://www.youtube.com/watch?v=B_asB_UGEgM)

<code>
<img width="250" alt="slendercards home" src="https://user-images.githubusercontent.com/52141045/125161290-6400d280-e179-11eb-9d1e-a057acdaed82.png">
<img width="250" alt="slendercards create" src="https://user-images.githubusercontent.com/52141045/125161303-7418b200-e179-11eb-81ed-21248b29c245.png">
<img width="250" alt="Screenshot 2021-07-10 at 12 13 43" src="https://user-images.githubusercontent.com/52141045/125161315-84c92800-e179-11eb-9732-22c39b84f65f.png">
</code>

## Getting Started

- Ensure that Go (Golang) is installed on your machine by running `go version` in your terminal. If you need to install Go, please see instructions [here](https://golang.org/doc/install)

- Ensure that PostgreSQL is installed on your machine by running `psql --version` in your terminal. If you need to install PostgreSQL, please see instructions [here](https://www.postgresql.org/download/). The db config details on line 26 of `flashcard/flashcard.model.go` file should be filled out correctly to match your local environment.

1. Clone this repo and enter!

   ```bash
   git clone https://github.com/jakewmiles/slendercards-go-server.git
   cd slendercards-go-server
   ```

2. The command `go run main.go` should get the server up and running!
   

## Features

- Scrapes data from reverso.net using a Go package created by me: [reverso-scraper](https://pkg.go.dev/github.com/jakewmiles/reverso-scraper?tab=versions)
- Posts new flashcards to the SQL database, using GORM.

## Future Features

- Enable user authentication

## Technologies

- Go (Golang)
- Fiber
- GORM
- Colly
