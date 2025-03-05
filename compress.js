// This program comresses a .wasm into a .wasm.br
// type node compress.js in the terminal

const fs = require('fs');
const brotli = require('brotli');

const inputFile = 'wasm/TTCompanion.wasm';
const wasmBuffer = fs.readFileSync(inputFile);

// Compress with Brotli
const compressed = brotli.compress(wasmBuffer, {
    mode: 0, // 0 generic compression (default), 1 for text
    quality: 11, // compression level (0 to 11, 11 is the max)
    lgwin: 22 // window size (default 22)
});

// Verify that the compression succeed
if (compressed === null) {
    console.error('Failure');
    process.exit(1);
}

const outputFile = 'wasm/TTCompanion.wasm.br';
fs.writeFileSync(outputFile, compressed);

console.log(`File compressed with Brotli : ${outputFile}`);