let socket = new WebSocket("ws://localhost:1337/ws")


socket.onopen = ((e) => {
    console.log("Nouveau client", e)
})

socket.onclose = ((e) => {
    console.log("Le client "+e+" vient de partir")
})




