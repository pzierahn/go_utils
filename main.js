Main = {};

Main.createElem = function(elem, parent)
{
    let ele = document.createElement(elem);
    parent.appendChild(ele);
    return ele;
};

Main.main = function()
{
    let ws = new WebSocket("{{.}}");
    let button = Main.createElem("button", document.body);
    button.innerText = "Button";
    button.onclick = function()
    {
        ws.send("Hallo");
    };

    let div = Main.createElem("div", document.body);

    let print = function(message)
    {
        let tag = Main.createElem("div", div);
        tag.innerHTML = message;
    };

    ws.onopen = function(evt)
    {
        print("OPEN");
    };

    ws.onclose = function(evt)
    {
        print("CLOSE");
        ws = null;
    };

    ws.onmessage = function(evt)
    {
        print("RESPONSE: " + evt.data);
    };

    ws.onerror = function(evt)
    {
        print("ERROR: " + evt.data);
    };

    // print("SEND: " + input.value);
    // ws.send(input.value);
    // ws.close();
};

window.addEventListener("load", Main.main);
