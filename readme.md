# VoteRight

An electronic voting system with distributed trust

## Installation (primary and verification servers on one instance)

1) Install Docker and Docker Compose https://docs.docker.com/compose/install/
2) Clone the repository
3) cd into the directory `cd voteright/`
4) run `docker-compose up`

## Dev installation

1) Install golang 1.12 https://blog.golang.org/go1.12
2) clone the repo into your gopath
3) go get .
4) go run main.go

## Deploy primary vote server

1) Install Docker and Docker Compose https://docs.docker.com/compose/install/
2) Clone the repository
3) cd into the directory `cd voteright/`
4) Configure verification server locations, and port in voteright.json
4) Configure the election in dump.json
5) run `docker build . -t voteright`
6) prepare the dump.json file to be used for the creation of verification servers
7) docker run --net="host" -d voteright
    - Note port flags and other docker flags can be used if necessary

## Deploy verification server

1) Install Docker and Docker Compose https://docs.docker.com/compose/install/
2) Clone the repository
3) cd into the directory `cd voteright/`
4) Configure set the verification field to true in voteright.json
4) copy the provided dump.json into the directory
5) run `docker build . -t voteright verification -f Dockerfile.verification`
6) docker run --net="host" -d voteright
    - Note port flags and other docker flags can be used if necessary



## Voting

*Primary Vote server* - Serves voting booth, ballot
*Voting Booth* - Frontend where user casts their vote
*Verification Cluster* - Cluster to verify votes

*How to run an election*

1) Primary Vote server is online, and given a list of addresses for a verification cluster, if no cluster exists, election may be run without
2) Voting booth opened on *All* Polling stations to be used in the election, more cannot be added later
3) Vote server begins attempting to establish a connection with each server in the verification cluster
4) Upon establishing a connection, the Primary voting server sends the addresses of the other servers in the cluster to each server it connects to
5) Verification servers connect to eachother, election may be __secured__ once two or more verification servers are established
6) Upon securing the election, the addresses of the verification servers are sent to the frontend

Election now may be conducted. Verification servers must regularly request verification on votes throughout the duration of the election, if any problem occurs, it is displayed on the voting booth.

If a verification server has an error, or a dissagreement, that server's votes are ignored and it is invalidated, this is displayed to the voter

In the event that less than two verification servers remain online, the election is invalidated.

[Systems diagram](verificationdiagram.pdf)

## API and Routes

**Primary voting server**

`GET /` - Serves main voting booth page

`GET /voters` - Get all voters in the election (admin only)

`POST /voters/validate` - Check is a voter is valid and has not voted. Post body should contain voter id. Returns the json blob representing the voter if valid, null if not

**Request**

```json
{"ID": 1}
```

**Response**

```json
{
 "StudentID": 1,
 "Cohort": 1,
 "Name": "Joey Lyon"
}
```

`POST /voters/verifyself` - Uses the id stored in the session cookie to verify the voter

**Response**
```json
{"Voted": true}
```

`POST /voters/vote` - Uses the id stored in the session cookie, and json body to cast the user's vote. Body should be a json array containing the ids of the candidates to vote for. Response contains random ids

**Request**

```json
[{"ID":2},{"ID":3}]
```

**Response**

```json
{
 "RandomID": 557700679194777,
 "Candidates": [
  {
   "Name": "Stan Marsh",
   "Cohort": 1,
   "ID": 2
  },
  {
   "Name": "Randy Marsh",
   "Cohort": 1,
   "ID": 3
  }
 ]
}
```

`GET /candidates` - Get all candidates in the election

```json
[
 {
  "Name": "Eric Cartman",
  "Cohort": 1,
  "ID": 1
 },
 {
  "Name": "Stan Marsh",
  "Cohort": 1,
  "ID": 2
 },
 {
  "Name": "Randy Marsh",
  "Cohort": 1,
  "ID": 3
 },
 {
  "Name": "Kenny",
  "Cohort": 1,
  "ID": 4
 }
]
```

`GET /candidates/votes` - Get all candidates and vote totals (admin only)

```json
[
 {
  "Candidate": {
   "Name": "Eric Cartman",
   "Cohort": 1,
   "ID": 1
  },
  "Votes": 3
 },
 {
  "Candidate": {
   "Name": "Stan Marsh",
   "Cohort": 1,
   "ID": 2
  },
  "Votes": 1
 },
 {
  "Candidate": {
   "Name": "Randy Marsh",
   "Cohort": 1,
   "ID": 3
  },
  "Votes": 1
 },
 {
  "Candidate": {
   "Name": "Kenny",
   "Cohort": 1,
   "ID": 4
  },
  "Votes": 3
 }
]
```