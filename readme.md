# VoteRight

An electronic voting system with distributed trust

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