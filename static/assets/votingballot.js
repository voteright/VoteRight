window.voteright = {}

var app = new Vue({
    el: '#app',
    data: {
      message: 'Hello Vue!',
      races: [],
      voted: [],
      alreadyvoted: false,
    },
    mounted (){
        fetch("/races").then((retval) => retval.json()).then((ret) => {
            for (let i = 0; i < ret.length; i ++){
                ret[i].voted = false;
                ret[i].for = -1;
                this.races.push(ret[i]);
                this.voted.push(0);
            }
        })
    },
    methods:{ 
        vote(raceid, candidateid, raceidx){
            this.$set(this.races[raceidx], 'voted', true);
            this.$set(this.races[raceidx], 'for', candidateid);
            for (let i = 0; i < this.races[raceidx].Candidates.length; i ++){
                if(this.races[raceidx].Candidates[i].ID == candidateid){
                    this.$set(this.races[raceidx], 'forname', this.races[raceidx].Candidates[i].Name);

                }
            }

            console.log("VOTE " + raceid, + " " + candidateid + " " + raceidx, " " + this.races[raceidx].voted);
        },
        isDisabled(id){
            return this.races[id].voted;
        },
        readyToSend(){
            for(let i = 0; i < this.races.length; i ++){
                console.log(this.races[i])
                if(this.races[i].voted == false){
                    return false
                }
            }
            return true;
        },
        submit(){
            val = []
            for(let i = 0; i < this.races.length; i ++){
                val.push({"ID": this.races[i].for});
            }

            fetch("/voters/vote",{
                method: "POST",
                credentials: "include",
                body: JSON.stringify(val),
            }).then((resp) => {
                if (resp.status == 200 ){
                    this.printBallot()
                }else{
                    this.alreadyvoted = true;
                }
            })
        },
        printBallot(){
            w = window.open()

            let x = document.createElement('div')
            x.innerHTML ="<p style='font-size: 20px;'>Ballot</p>" 
            for (let i = 0; i < this.races.length; i ++){
                x.innerHTML += "<strong>" + this.races[i].Name + ":</strong> " + this.races[i].forname + "<br>";
            }
            x.innerHTML += "<p>Signature: ____________________</p>";
            w.document.body.appendChild(x)
            w.print()
            w.close()
            window.location = "/thanks"
        }
    }
  })

// window.onload = () => {
//     document.getElementById("voteFailed").hidden = true;
//     fetch("/candidates").then((data) => data.json()).then((val) => {
//         let candidateslist = document.getElementById("candidateslist")
//         let candidateSelector = candidateslist.getElementsByClassName("candidateSelector")
        
        
//         for(let i = 0; i < val.length; i ++){
//             let myCandidateBox = candidateSelector[0].cloneNode(true);
//             candidateslist.appendChild(myCandidateBox)

//             let name = myCandidateBox.getElementsByClassName("candidateName")
//             name[0].innerText = val[i].Name;

//             let button = myCandidateBox.getElementsByClassName("submitbutton")[0]
//             button.onclick = () => {
//                 fetch("/voters/vote", {
//                     method: "POST",
//                     body: JSON.stringify({ID: val[i].ID}),
//                     credentials: "include",
//                 }).then((response) => {
//                     if (!response.ok) {
//                         document.getElementById("voteFailed").hidden = false;

//                     }else{
//                         window.location = "/thanks"

//                     }
//                 })
//             }

            
//         }
//         candidateSelector[0].remove()
//     })
// }
// // http://graben-212.dynamic2.rpi.edu:8080