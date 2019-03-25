// Create an object on the window for storing functions
window.voteright = {};

window.voteright.populateCandidates = () => {
    fetch("/candidates/votes").then((ret) => ret.json()).then((candidates) => {
        let list = document.getElementById("candidate-list");
        list.innerHTML = "";
        for(let i = 0; i < candidates.length; i ++){
            console.log(candidates[i])
            let item = document.createElement("li")
            item.innerText = "ID: " + String(candidates[i].Candidate.ID) + " | Name: " + candidates[i].Candidate.Name + " | Cohort: " + String(candidates[i].Candidate.Cohort) + " | Votes: " + String(candidates[i].Votes)
            list.appendChild(item);
        }
    })
}

window.onload = () => {
    window.voteright.populateCandidates();
};