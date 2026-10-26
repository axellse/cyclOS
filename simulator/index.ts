//First time using typescript, not sure how good this code is
declare class Go {
    constructor();
    importObject: any;
    run(a: any): void;
}

declare class ColorRGBA {
    R: number
    G: number
    B: number
    A: number
}
declare class InputState {
    Scheme: string
    Data: {[Key: string]: any;}
}

function ConfigureSetupForm(form: HTMLFormElement) {
    form.addEventListener('submit', (e) => {
        e.preventDefault()
        const data = new FormData(form)

        RunSimulator(data)
    })
}

const sForm = document.querySelector('#setup')
if (sForm == null) {
    alert("Form not found!")
} else {
    ConfigureSetupForm(sForm as HTMLFormElement)
}

let canvas: HTMLCanvasElement
let canvasCtx: CanvasRenderingContext2D
let inputState: InputState = {
    Scheme: "",
    Data: {},
}

function iGetScheme(): string {
    return inputState.Scheme
}

function iBoolStatus(key: string): boolean {
    if (typeof inputState.Data[key] != "boolean") {
        alert("cyclOS tried to get status of " + typeof inputState.Data[key] + " input parameter as boolean")
        throw new Error("iBoolStatus type error")
    }
    return inputState.Data[key]
}

function dSetPixel(x: number, y: number, colorJSON: string) {
    const color : ColorRGBA = JSON.parse(colorJSON)
    
    canvasCtx.fillStyle = "rgb(" + color.R + ", " + color.G + ", " + color.B + ", " + color.A / 255 + ")";
    canvasCtx.fillRect(x, y, 1, 1);
}

function dSize(): { width: number, height: number } {
    return { width: canvas.width, height: canvas.height }
}

async function RunSimulator(formData: FormData) {
    const sim = document.querySelector("#simulator-interface")
    if (!sim) {
        alert("cant find simulator interface!")
        return
    }

    sim.firstElementChild?.remove()
    canvas = document.createElement("canvas")
    canvas.style = "image-rendering: pixelated; image-rendering: crisp-edges;"
    canvas.width = Number(formData.get('display-width'))
    canvas.height = Number(formData.get('display-height'))
    sim.appendChild(canvas)

    let _ctx = canvas.getContext("2d");
    if (_ctx == null) {
        alert("Canvas context 2d not suppported!")
        return
    }
    canvasCtx = _ctx

    canvasCtx.fillStyle = "black";
    canvasCtx.fillRect(0, 0, canvas.width, canvas.height);

    const inputScheme = formData.get('input-scheme')
    if (!inputScheme) {
        alert("no input scheme!")
        return
    }

    inputState.Scheme = inputScheme.toString()
    if (inputScheme == '3button') {
        let inputPanel = document.createElement('div')
        inputPanel.classList.add('vertical-button-panel')

        for (const button of ["up", "confirm", "down"]) {
            let buttonElement = document.createElement('button')
            buttonElement.innerText = button

            inputState.Data[button] = false
            let makeListener = (set: boolean) => {
                return () => {
                    inputState.Data[button] = set
                }
            }
            buttonElement.addEventListener('mousedown', makeListener(true))
            buttonElement.addEventListener('mouseup', makeListener(false))

            inputPanel.appendChild(buttonElement)
        }

        sim.appendChild(inputPanel)
    }

    const go = new Go();

    const file = formData.get('wasm-binary') as File
    const wasm = await WebAssembly.instantiate(await file.arrayBuffer(), go.importObject)

    go.run(wasm.instance);
}


