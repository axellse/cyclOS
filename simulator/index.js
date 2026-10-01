"use strict";
function ConfigureSetupForm(form) {
    form.addEventListener('submit', (e) => {
        e.preventDefault();
        const data = new FormData(form);
        RunSimulator(data);
    });
}
const sForm = document.querySelector('#setup');
if (sForm == null) {
    alert("Form not found!");
}
else {
    ConfigureSetupForm(sForm);
}
let canvas;
let canvasCtx;
function dSetPixel(x, y, colorJSON) {
    const color = JSON.parse(colorJSON);
    canvasCtx.fillStyle = "rgb(" + color.R + ", " + color.G + ", " + color.B + ", " + color.A / 255 + ")";
    canvasCtx.fillRect(x, y, 1, 1);
}
function dSize() {
    return { width: canvas.width, height: canvas.height };
}
async function RunSimulator(formData) {
    document.querySelector("#simulator-interface")?.firstElementChild?.remove();
    canvas = document.createElement("canvas");
    canvas.style = "image-rendering: pixelated; image-rendering: crisp-edges;";
    canvas.width = Number(formData.get('display-width'));
    canvas.height = Number(formData.get('display-height'));
    document.querySelector("#simulator-interface")?.appendChild(canvas);
    let _ctx = canvas.getContext("2d");
    if (_ctx == null) {
        throw new Error("Canvas context 2d not suppported!");
    }
    canvasCtx = _ctx;
    canvasCtx.fillStyle = "black";
    canvasCtx.fillRect(0, 0, canvas.width, canvas.height);
    const go = new Go();
    const file = formData.get('wasm-binary');
    const wasm = await WebAssembly.instantiate(await file.arrayBuffer(), go.importObject);
    go.run(wasm.instance);
}
