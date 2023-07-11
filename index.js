var EC = require('elliptic').ec;

var ec = new EC('secp256k1');

var pub = { 
    x: "3bd2defa93e3a69cefa9765496ac537cc1280990df61e6bf18166863a09e8f39",
    y: "571bff72b6dc52bd7f7bb4dd390feba74460b3ca12eeda2a19d9d185740853d1"
};  

var key = ec.keyFromPublic(pub, 'hex');

var signature = { 
    r: '39da17ee7835a1949ec0e925e77ff42afdea28d67173a74cec4b52fda0cc8ffc', 
    s: '327c8a6bb44e855e8ff1f11e1f6d98f8d07340059399652198a69979d0d283b7' 
}


msgHash = "918e47e7d1b7b4892fd746e691f8f44f3c9a3cc9a129de06301701895ea9676d"
console.log(key.verify(msgHash, signature))

msgHash = "918e47e7"
console.log(key.verify(msgHash, signature))