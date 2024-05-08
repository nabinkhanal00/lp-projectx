let calculatorButtons = document.querySelectorAll(".calcBtn")
let input = document.getElementById("input")
let steps = document.getElementById("steps")
let query = document.getElementById("query")
let result = document.getElementById("result")
let history = document.getElementById("history")
let calculate = document.getElementById("calculate")
let historyClearBtn = document.getElementById("historyClearBtn")
input.focus()


const insertAtCaret = (element, text) => {
    text = text || '';
    if (element.selectionStart || element.selectionStart === 0) {
        // Others
        var startPos = element.selectionStart;
        var endPos = element.selectionEnd;
        element.value = element.value.substring(0, startPos) +
            text +
            element.value.substring(endPos, element.value.length);
        element.selectionStart = startPos + text.length;
        element.selectionEnd = startPos + text.length;
    } else {
        element.value += text;
    }
};
for (let button of calculatorButtons) {
    button.addEventListener('click', () => {
        switch (button.id) {
            case "C":
                input.value = ""
                break
            default:
                insertAtCaret(input, button.id)
                break
        }
        input.focus()
    })
    button.addEventListener('mousedown', (event) => {
        event.preventDefault()
    })
}

let histories = [
    {
        id: "history-1",
        "error": false,
        "message": undefined,
        expression: "6+5",
        result: 11,
        steps: ["6 + 5 = 11"]
    }
]


let current_history = histories.length;

const handleDisplayHistory = (h) => {
    input.value = h.expression
    result.classList.remove("text-red-500")
    result.classList.add("text-blue-500")
    result.innerHTML = `Result: ${h.result}`
    query.innerHTML = `Query: ${h.expression}`

    steps.innerHTML = ""
    for (let step of h.steps) {
        let li = document.createElement("li")
        li.appendChild(document.createTextNode(step))
        li.classList = "bg-gray-100 p-2 rounded-md"
        steps.appendChild(li)
    }
}

historyClearBtn.addEventListener('click', () => {
    histories = []
    history.innerHTML = ""
})

const handleHistoryClick = (event) => {
    const id = event.target.id;
    const history = histories.find((h) => h.id == id)
    handleDisplayHistory(history)
}

const handleAddHistory = (h) => {
    let li = document.createElement("li")
    li.appendChild(document.createTextNode(`${h.expression} = ${h.result}`))
    li.id = h.id;
    li.classList = "bg-gray-100 p-2 rounded-md cursor-pointer"
    li.addEventListener('click', handleHistoryClick)
    history.appendChild(li)
}

for (let h of histories) {
    handleAddHistory(h)
}


const handleResult = (h) => {
    handleAddHistory(h)
    handleDisplayHistory(h)
}


const handleError = (data) => {
    result.classList.remove("text-blue-500")
    result.classList.add("text-red-500")
    query.innerHTML = "QUERY: " + data.expression
    result.innerHTML = "ERROR: " + data.message
    steps.innerHTML = ""
}

const handleEmpty = () => {
}

const handleCalculate = async () => {
    const query = document.getElementById("input")
    if (query.value == "") {
        handleEmpty()
        return
    }
    const body = {
        query: query.value
    }
    try {
        const response = await fetch("/calculate", {
            "method": "POST",
            "body": JSON.stringify(body),
            "headers": {
                "Content-Type": "application/json"
            }
        })
        const data = await response.json()
        if (data.error) {
            handleError(data)
        } else {
            current_history++;
            let h = {
                expression: query.value,
                id: "history-" + current_history,
                error: false,
                result: data.answer,
                steps: data.steps
            }
            histories.push(h)
            handleResult(h)
        }
    }
    catch (e) {
        console.log(e)
    }
}

calculate.addEventListener("click", handleCalculate)
input.addEventListener("keyup", function(event) {
    if (event.key == "Enter") {
        handleCalculate()
    }
})

input.addEventListener("input", function() {
    var input = this.value;
    var sanitizedInput = input.replace(/[^0-9+\-*\/.\(\)]/g, '');
    if (sanitizedInput !== input) {
        this.value = sanitizedInput;
    }
});