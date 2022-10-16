let wind = document.getElementById("wind");
let water = document.getElementById("water");
let windStatus = document.getElementById("windStatus");
let waterStatus = document.getElementById("waterStatus");

setTimeout(function(){
    fetch("http://localhost:8080/status").then(response => response.json())
    .then(data => {
        console.log(data)
        wind.innerHTML += `Wind: ${data.wind}`
        water.innerHTML += `Water: ${data.water}`
        windStatus.innerHTML += `Wind Status: ${data.windStatus}`
        waterStatus.innerHTML += `Water Status: ${data.waterStatus}`
    })
}, 15000);