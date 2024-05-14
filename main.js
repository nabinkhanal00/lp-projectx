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
    // check if the browser supports the property
    if (element.selectionStart) {
        // Others
        var startPos = element.selectionStart;
        var endPos = element.selectionEnd;
        element.value = element.value.substring(0, startPos) +
            text +
            element.value.substring(endPos, element.value.length);
        element.selectionStart = startPos + text.length;
        element.selectionEnd = startPos + text.length;
        element.scrollLeft = element.scrollWidth
    } else {
        element.value += text;
    }
};
for (let button of calculatorButtons) {
    button.addEventListener('click', () => {
        switch (button.id) {
            case "C":
                input.value = ""
                result.innerHTML = "0"
                result.classList.remove('text-red-500')
                result.classList.add('text-blue-500')
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


let histories = JSON.parse(localStorage.getItem("histories")) || []


let current_history = histories.length;

const handleDisplayDetails = (h) => {
    input.value = h.expression
    result.classList.remove("text-red-500")
    result.classList.add("text-blue-500")
    result.innerHTML = `= ${h.result}`
    query.innerHTML = `${h.expression}`

    steps.innerHTML = ""
    if (h.steps) {
        for (let step of h.steps) {
            let li = document.createElement("li")
            li.appendChild(document.createTextNode(step))
            li.classList = "bg-gray-100 p-2 rounded-md"
            steps.appendChild(li)
        }
    }
}

historyClearBtn.addEventListener('click', () => {
    histories = []
    history.innerHTML = ""
    localStorage.setItem("histories", JSON.stringify([]))
})

const handleHistoryClick = (event) => {
    const id = event.target.id;
    const history = histories.find((h) => h.id == id)
    handleDisplayDetails(history)
}
const handleDeleteSingleHistory = (event) => {
    event.stopPropagation()
    let id = event.target.id;
    id = id.substr(0, id.length - "-delete".length)
    histories = histories.filter((h) => h.id != id)
    localStorage.setItem("histories", JSON.stringify(histories))
    history.innerHTML = ""
    for (let h of histories) {
        handleAddHistory(h)
    }
}


const handleAddHistory = (h) => {
    let button = document.createElement("button")
    button.id = h.id + "-delete"
    button.classList = "text-red-500"
    button.addEventListener('click', handleDeleteSingleHistory)
    button.innerHTML = "🗑"
    // button.innerHTML = "X"
    let li = document.createElement("li")
    li.appendChild(document.createTextNode(`${h.expression} = ${h.result}`))
    li.appendChild(button)
    li.id = h.id;
    li.classList = "flex items-center justify-between bg-gray-100 p-2 rounded-md cursor-pointer relative"
    li.addEventListener('click', handleHistoryClick)
    history.appendChild(li)
}

for (let h of histories) {
    handleAddHistory(h)
}


const handleResult = (h) => {
    handleAddHistory(h)
    handleDisplayDetails(h)
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
    let query = document.getElementById("input")
    if (query.value == "") {
        handleEmpty()
        return
    }
    let queryValue = query.value;
    try {
        const go = new Go();

        WebAssembly.instantiateStreaming(fetch("main.wasm"), go.importObject).then(result => {
            go.run(result.instance);
        })
        const content = Calculate(queryValue);

        const data = JSON.parse(content)
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
            localStorage.setItem("histories", JSON.stringify(histories))
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
