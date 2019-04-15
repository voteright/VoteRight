// // Create an object on the window for storing functions
// window.voteright = {};

// window.voteright.populateCandidates = () => {
//     fetch("/candidates/votes").then((ret) => ret.json()).then((candidates) => {
//         let list = document.getElementById("candidate-list");
//         list.innerHTML = "";
//         for(let i = 0; i < candidates.length; i ++){
//             console.log(candidates[i])
//             let item = document.createElement("li")
//             item.innerText = "ID: " + String(candidates[i].Candidate.ID) + " | Name: " + candidates[i].Candidate.Name + " | Cohort: " + String(candidates[i].Candidate.Cohort) + " | Votes: " + String(candidates[i].Votes)
//             list.appendChild(item);
//         }
//     })
// }

// window.onload = () => {
//     window.voteright.populateCandidates();
// };


var app = new Vue({
    el: '#app',
    data: {
        message: "hello world",
        verificationServers: [],
        verificationData: [],
        integrityViolations: [],
        voteTotals: [],
        serversMatch: false
    },
    mounted (){
        this.fetchVerificationServers();
        this.fetchIntegrityViolations();
        this.fetchIfServersMatch();
        this.fetchVoteTotals();
        this.fetchServerData();
    },
    methods:{ 
        fetchVerificationServers(){
            fetch("/integrity/servers").then((data) => data.json()).then((val) => { 
                this.verificationServers = val;
            })
        },
        fetchIntegrityViolations(){
            fetch("/integrity/").then((data) => data.json()).then((val) => { 
                this.integrityViolations = val;
            })
        },
        fetchIfServersMatch(){
            fetch("/integrity/match").then((data) => data.json()).then((val) => { 
                this.serversMatch = val.AllVerificaitonServersMatch;
            })
        },
        fetchServerData(){
            fetch("/integrity/data").then((data) => data.json()).then((val) => {
                this.verificationData = val;
            })
        },
        fetchVoteTotals(){
            fetch("/candidates/votes").then((data) => data.json()).then((val) => {
                this.voteTotals = val;
            })
        }
    }
  })