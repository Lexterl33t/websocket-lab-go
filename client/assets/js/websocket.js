let socket = new WebSocket("ws://localhost:1337/ws")

// event when new client is connected
let ev_new_user = ((...d) => {
    console.log(d[0].Msg)
})

// event when client quit the websocket
let ev_exit_user = ((...d) => {
    console.log(d[0].Msg)
})

let ev_list_user = ((...d) => {
    console.log(d[0])
})

// Commands object
const commands = {
    1: ev_new_user,
    2: ev_exit_user,
    3: ev_list_user
}

socket.onopen = ((e) => {
    console.log("New client")
})

socket.onmessage = ((e) => {

    console.log(e)
    data = JSON.parse(e.data)

    commands[data.ID](data.CMD)
})

socket.onclose = ((e) => {
    data = JSON.parse(e.data)

    commands[data.ID](data.CMD)
})




