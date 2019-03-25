window.voteright = {}

window.onload = () => {
    document.getElementById("voteFailed").hidden = true;
    fetch("/candidates").then((data) => data.json()).then((val) => {
        let candidateslist = document.getElementById("candidateslist")
        let candidateSelector = candidateslist.getElementsByClassName("candidateSelector")
        
        
        for(let i = 0; i < val.length; i ++){
            let myCandidateBox = candidateSelector[0].cloneNode(true);
            candidateslist.appendChild(myCandidateBox)

            let name = myCandidateBox.getElementsByClassName("candidateName")
            name[0].innerText = val[i].Name;

            let button = myCandidateBox.getElementsByClassName("submitbutton")[0]
            button.onclick = () => {
                fetch("/voters/vote", {
                    method: "POST",
                    body: JSON.stringify({ID: val[i].ID}),
                    credentials: "include",
                }).then((response) => {
                    if (!response.ok) {
                        document.getElementById("voteFailed").hidden = false;

                    }else{
                        window.location = "/thanks"

                    }
                })
            }

            
        }
        candidateSelector[0].remove()
    })
}
// http://graben-212.dynamic2.rpi.edu:8080