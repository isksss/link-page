const go = new Go();

// console.log("debug");

WebAssembly.instantiateStreaming(
    fetch("wasm/main.wasm"), go.importObject)
    .then((result) => {
        console.log(result);
        go.run(result.instance);
    }
);