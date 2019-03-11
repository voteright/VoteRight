// Create an object on the window for storing functions
window.voteright = {};

window.voteright.populateCandidates = () => {
    fetch("/candidates").then((ret) => ret.json()).then((candidates) => {
        let list = document.getElementById("candidate-list");
        list.innerHTML = "";
        for(let i = 0; i < candidates.length; i ++){
            console.log(candidates[i])
            let item = document.createElement("li")
            item.innerText = "ID: " + String(candidates[i].ID) + " | Name: " + candidates[i].Name + " | Cohort: " + String(candidates[i].Cohort)
            list.appendChild(item);
        }
    })
}

window.onload = () => {
    window.voteright.populateCandidates();
};