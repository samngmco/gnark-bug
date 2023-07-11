# gnark-bug

As of this writing (July 11th, 2023), signatures generated inside gnark-crypto library can't be verified by other libraries. Seems a flaw in the implementation. 

```
> go run .
signature verification ok
hashedMsg: 918e47e7d1b7b4892fd746e691f8f44f3c9a3cc9a129de06301701895ea9676d
After HashToInt(): 918e47e7
Sig.R: 39da17ee7835a1949ec0e925e77ff42afdea28d67173a74cec4b52fda0cc8ffc
Sig.S: 327c8a6bb44e855e8ff1f11e1f6d98f8d07340059399652198a69979d0d283b7
Pub.X: 3bd2defa93e3a69cefa9765496ac537cc1280990df61e6bf18166863a09e8f39
Pub.Y: 571bff72b6dc52bd7f7bb4dd390feba74460b3ca12eeda2a19d9d185740853d1
```

Port the input to `index.js` verification failed

