// Create an object on the window for storing functions
window.voteright = {};

let idbox = document.getElementById("idtextbox");
let helptext = document.getElementById("idhelp");
helptext.innerHTML = ""

idbox.onchange = () => {
    if (idbox.value < 0) {
        idbox.value = 0;
    }
}

document.getElementById("submitid").onclick = () =>{
    var x = Number(document.getElementById("idtextbox").value)
    fetch("/voters/validate", {
        method: "POST", 
        headers: {
            "Content-Type": "application/json",
            // "Content-Type": "application/x-www-form-urlencoded",
        },
        body: JSON.stringify({ID: x})
    }).then((ret) => ret.json()).then((val) => {
        if (val != null){
            fetch("/voters/login", {
                method: "POST", 
                headers: {
                    "Content-Type": "application/json",
                    // "Content-Type": "application/x-www-form-urlencoded",
                },
                credentials: 'include',
                body: JSON.stringify(val)
            }).then((data) => {
                return data.text()
            }).then((val) => {
                if (val == "voted"){
                    alert("Already voted")
                }else{
                    // console.log(val)
                    window.location = "/votingbooth"
                    
                }
            })

        }else{
            helptext.innerText = "Invalid ID"
        }
    })
}