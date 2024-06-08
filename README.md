# Na hora API
<p>
  <img src='http://img.shields.io/static/v1?label=STATUS&message=ACTIVE&color=GREEN&style=for-the-badge'>
</p>

## :hammer: Project functionalities

- `#1`: Be the core API for the scheduling service called "Na hora".

<p align="right">(<a href="#top">back to top</a>)</p>
<hr>

## :computer: The technologies used were:

- [Go](https://go.dev/)
- [Docker](https://www.docker.com/)
- [PostgreSQL](https://www.postgresql.org/)

<p align="right">(<a href="#top">back to top</a>)</p>
<hr>

### :rocket: You can run this project following the steps below:

#### # *1* Clone the project from GitHub
#### # *2* Run ```go mod tidy```
#### # *3* Run ```cp .env.model .env```
#### # *4* Run ```docker-compose up -d```
#### # *5* Run ```make m-apply```
#### # *5* Run ```make dev```

<p align="right">(<a href="#top">back to top</a>)</p>

### :bulb: Project tips:

#### You can build the project running:
#### ```make build```

#### You can generate new dependency injection creating the set in wire.go and running:
#### ```make wire```

#### You can modify the DB structure using migrations changing the go entities and running:
#### ```make m-generate --name=name-for-the-migration```
#### And to aplly it to the database use:
#### ```make m-apply```